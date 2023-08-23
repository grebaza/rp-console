// Copyright 2023 Redpanda Data, Inc.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.md
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0

package serde

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/redpanda-data/console/backend/pkg/schema"
	"github.com/twmb/franz-go/pkg/kgo"

	"github.com/santhosh-tekuri/jsonschema/v5"
)

var _ Serde = (*JsonSchemaSerde)(nil)

type JsonSchemaSerde struct {
	SchemaSvc *schema.Service
}

func (JsonSchemaSerde) Name() PayloadEncoding {
	return PayloadEncodingJSON
}

func (JsonSchemaSerde) DeserializePayload(record *kgo.Record, payloadType PayloadType) (RecordPayload, error) {
	payload := payloadFromRecord(record, payloadType)

	if len(payload) <= 5 {
		return RecordPayload{}, fmt.Errorf("payload size is < 5 for json schema")
	}

	if payload[0] != byte(0) {
		return RecordPayload{}, fmt.Errorf("incorrect magic byte for json schema")
	}

	// TODO: For more confidence we could just ask the schema service for the given
	// schema and based on the response we can check the schema type (avro, json, ..)

	return jsonDeserializePayload(payload[5:])
}

func (s JsonSchemaSerde) SerializeObject(obj any, payloadType PayloadType, opts ...SerdeOpt) ([]byte, error) {
	so := serdeCfg{}
	for _, o := range opts {
		o.apply(&so)
	}

	if !so.schemaIDSet {
		return nil, errors.New("no schema id specified")
	}

	schemaRes, err := s.SchemaSvc.GetSchemaByID(so.schemaId)
	if err != nil {
		return nil, fmt.Errorf("getting json schema from registry '%+v': %w", so.schemaId, err)
	}

	var byteData []byte
	switch v := obj.(type) {
	case string:
		byteData = []byte(v)
	case []byte:
		byteData = v
	default:
		encoded, err := json.Marshal(v)
		if err != nil {
			return nil, fmt.Errorf("error serializing to JSON: %w", err)
		}
		byteData = encoded
	}

	trimmed := bytes.TrimLeft(byteData, " \t\r\n")

	if len(trimmed) == 0 {
		return nil, fmt.Errorf("after trimming whitespaces there were no characters left")
	}

	startsWithJSON := trimmed[0] == '[' || trimmed[0] == '{'
	if !startsWithJSON {
		return nil, fmt.Errorf("first byte indicates this it not valid JSON, expected brackets")
	}

	// validate
	sch, err := s.compileJSONSchema(schemaRes)
	if err != nil {
		return nil, fmt.Errorf("error compiling json schema: %w", err)
	}

	var vObj interface{}
	if err := json.Unmarshal(trimmed, &vObj); err != nil {
		return nil, fmt.Errorf("error validating json schema: %w", err)
	}

	if err = sch.Validate(vObj); err != nil {
		return nil, fmt.Errorf("error validating json schema: %w", err)
	}

	header, err := appendEncode(nil, int(so.schemaId), nil)
	if err != nil {
		return nil, fmt.Errorf("failed encode json schema payload: %w", err)
	}

	binData := append(header, trimmed...)

	return binData, nil
}

func (s *JsonSchemaSerde) compileJSONSchema(schemaRes *schema.SchemaResponse) (*jsonschema.Schema, error) {
	c := jsonschema.NewCompiler()
	schemaName := "redpanda_json_schema_main.json"

	err := s.buildJSONSchemaWithReferences(c, schemaName, schemaRes)
	if err != nil {
		return nil, err
	}

	return c.Compile(schemaName)
}

func (s *JsonSchemaSerde) buildJSONSchemaWithReferences(compiler *jsonschema.Compiler, name string, schemaRes *schema.SchemaResponse) error {
	if err := compiler.AddResource(name, strings.NewReader(schemaRes.Schema)); err != nil {
		return err
	}

	for _, reference := range schemaRes.References {
		schemaRef, err := s.SchemaSvc.GetSchemaBySubjectAndVersion(reference.Subject, strconv.Itoa(reference.Version))
		if err != nil {
			return err
		}
		if err := compiler.AddResource(reference.Name, strings.NewReader(schemaRef.Schema)); err != nil {
			return err
		}
		if err := s.buildJSONSchemaWithReferences(compiler, reference.Name, &schema.SchemaResponse{
			Schema:     schemaRef.Schema,
			References: schemaRef.References,
		}); err != nil {
			return err
		}
	}

	return nil
}