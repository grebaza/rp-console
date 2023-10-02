/**
 * Copyright 2022 Redpanda Data, Inc.
 *
 * Use of this software is governed by the Business Source License
 * included in the file https://github.com/redpanda-data/redpanda/blob/dev/licenses/bsl.md
 *
 * As of the Change Date specified in that file, in accordance with
 * the Business Source License, use of this software will be governed
 * by the Apache License, Version 2.0
 */

import React, { RefObject } from 'react';
import { observer } from 'mobx-react';
import { PageComponent, PageInitHelper } from '../Page';
import { api } from '../../../state/backendApi';
import { Empty, } from 'antd';
import { appGlobal } from '../../../state/appGlobal';
import { sortField } from '../../misc/common';
import { DefaultSkeleton, InlineSkeleton } from '../../../utils/tsxUtils';
import { uiSettings } from '../../../state/ui';

import './Schema.List.scss';
import SearchBar from '../../misc/SearchBar';
import { makeObservable, observable } from 'mobx';
import { KowlTable } from '../../misc/KowlTable';
import Section from '../../misc/Section';
import PageContent from '../../misc/PageContent';
import { Alert, AlertIcon, Button, Checkbox, Divider, Flex, Skeleton } from '@redpanda-data/ui';
import { SmallStat } from '../../misc/SmallStat';

function renderRequestErrors(requestErrors?: string[]) {
    if (!requestErrors || requestErrors.length === 0) {
        return null;
    }

    return (
        <Section>
            <div className="SchemaList__error-card">
                {requestErrors.map((errorMessage, idx) => (
                    <Alert key={idx} marginTop="1em" status="error">
                        <AlertIcon />
                        <div>{errorMessage}</div>
                    </Alert>
                ))}
            </div>
        </Section>
    );
}

function renderNotConfigured() {
    return (
        <PageContent>
            <Section>
                <Empty description={null}>
                    <div style={{ marginBottom: '1.5rem' }}>
                        <h2>Not Configured</h2>

                        <p>
                            Schema Registry is not configured in Redpanda Console.
                            <br />
                            To view all registered schemas, their documentation and their versioned history simply provide the connection credentials in the Redpanda Console config.
                        </p>
                    </div>

                    {/* todo: fix link once we have a better guide */}
                    <a target="_blank" rel="noopener noreferrer" href="https://docs.redpanda.com/docs/manage/console/">
                        <Button variant="solid">Redpanda Console Config Documentation</Button>
                    </a>
                </Empty>
            </Section>
        </PageContent>
    );
}

@observer
class SchemaList extends PageComponent<{}> {
    @observable searchBar: RefObject<SearchBar<any>> = React.createRef();
    @observable filteredSchemaSubjects: { name: string }[];

    constructor(p: any) {
        super(p);
        makeObservable(this);
    }

    initPage(p: PageInitHelper): void {
        p.title = 'Schema Registry';
        p.addBreadcrumb('Schema Registry', '/schema-registry');
        this.refreshData(true);
        appGlobal.onRefresh = () => this.refreshData(true);
    }

    refreshData(force?: boolean) {
        api.refreshSchemaConfig(force);
        api.refreshSchemaMode(force);
        api.refreshSchemaSubjects(force);
        api.refreshSchemaTypes(force);
    }

    isFilterMatch(filterString: string, subject: { name: string }) {
        return subject.name.toLowerCase().includes(filterString.toLowerCase());
    }

    render() {
        if (api.schemaSubjects === undefined) return DefaultSkeleton; // request in progress
        if (api.schemaOverviewIsConfigured == false) return renderNotConfigured();

        const filteredSubjects = api.schemaSubjects
            .filter(x => uiSettings.schemaList.showSoftDeleted || (!uiSettings.schemaList.showSoftDeleted && !x.isSoftDeleted))
            .filter(x => x.name.toLowerCase().includes(uiSettings.schemaList.quickSearch.toLowerCase()));

        return (
            <PageContent key="b">
                {/* Statistics Bar */}
                <Flex gap="1rem" alignItems="center">
                    <SmallStat title="Mode">{api.schemaConfig ?? <InlineSkeleton width="100px" />}</SmallStat>
                    <Divider height="2ch" orientation="vertical" />
                    <SmallStat title="Compatability">{api.schemaMode ?? <InlineSkeleton width="100px" />}</SmallStat>
                </Flex>

                <Button variant="outline" mb="4" width="fit-content" onClick={() => appGlobal.history.push('/schema-registry/edit-compatability')}>
                    Edit Compatability
                </Button>

                {renderRequestErrors()}

                <SearchBar<{ name: string }>
                    dataSource={() => (api.schemaSubjects || []).map(str => ({ name: str.name }))}
                    isFilterMatch={this.isFilterMatch}
                    filterText={uiSettings.schemaList.quickSearch}
                    onQueryChanged={(filterText) => (uiSettings.schemaList.quickSearch = filterText)}
                    onFilteredDataChanged={data => this.filteredSchemaSubjects = data}
                />

                <Section>
                    <Flex justifyContent={'space-between'} pb={3}>
                        <Button colorScheme="brand" onClick={() => appGlobal.history.push('/schema-registry/create')}>Create new schema</Button>
                        <Checkbox
                            isChecked={uiSettings.schemaList.showSoftDeleted}
                            onChange={e => uiSettings.schemaList.showSoftDeleted = e.target.checked}
                        >
                            Show soft-deleted
                        </Checkbox>
                    </Flex>

                    <KowlTable
                        dataSource={filteredSubjects}
                        columns={[
                            { title: 'Name', dataIndex: 'name', sorter: sortField('name'), defaultSortOrder: 'ascend' },
                            { title: 'Type', render: (_, r) => <SchemaTypeColumn name={r.name} /> },
                            { title: 'Latest Version', render: (_, r) => <LatestVersionColumn name={r.name} /> },
                        ]}

                        observableSettings={uiSettings.schemaList}

                        rowClassName={() => 'hoverLink'}
                        rowKey="name"
                        onRow={({ name }) => ({
                            onClick: () => appGlobal.history.push(`/schema-registry/subjects/${encodeURIComponent(name)}?version=latest`),
                        })}
                    />
                </Section>
            </PageContent>
        );
    }
}

const SchemaTypeColumn = observer((p: { name: string }) => {
    const details = api.schemaDetails.get(p.name);
    if (!details) {
        api.refreshSchemaDetails(p.name);
        return <Skeleton height="15px" />;
    }

    return <>{details.type}</>;
});

const LatestVersionColumn = observer((p: { name: string }) => {
    const details = api.schemaDetails.get(p.name);
    if (!details) {
        api.refreshSchemaDetails(p.name);
        return <Skeleton height="15px" />;
    }

    return <>{details.latestActiveVersion}</>;
});

export default SchemaList;
