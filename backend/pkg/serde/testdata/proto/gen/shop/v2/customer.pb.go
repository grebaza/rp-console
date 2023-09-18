// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: shop/v2/customer.proto

package v2

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Customer_CustomerType int32

const (
	Customer_CUSTOMER_TYPE_UNSPECIFIED Customer_CustomerType = 0
	Customer_CUSTOMER_TYPE_PERSONAL    Customer_CustomerType = 1
	Customer_CUSTOMER_TYPE_BUSINESS    Customer_CustomerType = 2
)

// Enum value maps for Customer_CustomerType.
var (
	Customer_CustomerType_name = map[int32]string{
		0: "CUSTOMER_TYPE_UNSPECIFIED",
		1: "CUSTOMER_TYPE_PERSONAL",
		2: "CUSTOMER_TYPE_BUSINESS",
	}
	Customer_CustomerType_value = map[string]int32{
		"CUSTOMER_TYPE_UNSPECIFIED": 0,
		"CUSTOMER_TYPE_PERSONAL":    1,
		"CUSTOMER_TYPE_BUSINESS":    2,
	}
)

func (x Customer_CustomerType) Enum() *Customer_CustomerType {
	p := new(Customer_CustomerType)
	*p = x
	return p
}

func (x Customer_CustomerType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Customer_CustomerType) Descriptor() protoreflect.EnumDescriptor {
	return file_shop_v2_customer_proto_enumTypes[0].Descriptor()
}

func (Customer_CustomerType) Type() protoreflect.EnumType {
	return &file_shop_v2_customer_proto_enumTypes[0]
}

func (x Customer_CustomerType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Customer_CustomerType.Descriptor instead.
func (Customer_CustomerType) EnumDescriptor() ([]byte, []int) {
	return file_shop_v2_customer_proto_rawDescGZIP(), []int{0, 0}
}

type Customer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version      int32                 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Id           string                `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	FirstName    string                `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName     string                `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Gender       string                `protobuf:"bytes,5,opt,name=gender,proto3" json:"gender,omitempty"`
	CompanyName  string                `protobuf:"bytes,6,opt,name=company_name,json=companyName,proto3" json:"company_name,omitempty"`
	Email        string                `protobuf:"bytes,7,opt,name=email,proto3" json:"email,omitempty"`
	CustomerType Customer_CustomerType `protobuf:"varint,8,opt,name=customer_type,json=customerType,proto3,enum=shop.v2.Customer_CustomerType" json:"customer_type,omitempty"`
	Revision     int32                 `protobuf:"varint,9,opt,name=revision,proto3" json:"revision,omitempty"`
}

func (x *Customer) Reset() {
	*x = Customer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shop_v2_customer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Customer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Customer) ProtoMessage() {}

func (x *Customer) ProtoReflect() protoreflect.Message {
	mi := &file_shop_v2_customer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Customer.ProtoReflect.Descriptor instead.
func (*Customer) Descriptor() ([]byte, []int) {
	return file_shop_v2_customer_proto_rawDescGZIP(), []int{0}
}

func (x *Customer) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Customer) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Customer) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Customer) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Customer) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *Customer) GetCompanyName() string {
	if x != nil {
		return x.CompanyName
	}
	return ""
}

func (x *Customer) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Customer) GetCustomerType() Customer_CustomerType {
	if x != nil {
		return x.CustomerType
	}
	return Customer_CUSTOMER_TYPE_UNSPECIFIED
}

func (x *Customer) GetRevision() int32 {
	if x != nil {
		return x.Revision
	}
	return 0
}

var File_shop_v2_customer_proto protoreflect.FileDescriptor

var file_shop_v2_customer_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x76,
	0x32, 0x22, 0x89, 0x03, 0x0a, 0x08, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x18,
	0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x43, 0x0a, 0x0d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x73,
	0x68, 0x6f, 0x70, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65,
	0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x72, 0x65,
	0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x65, 0x0a, 0x0c, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x19, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d,
	0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x45,
	0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x45, 0x52, 0x53, 0x4f, 0x4e, 0x41, 0x4c, 0x10,
	0x01, 0x12, 0x1a, 0x0a, 0x16, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x45, 0x52, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x42, 0x55, 0x53, 0x49, 0x4e, 0x45, 0x53, 0x53, 0x10, 0x02, 0x42, 0x4f, 0x5a,
	0x4d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x64, 0x70,
	0x61, 0x6e, 0x64, 0x61, 0x2d, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c,
	0x65, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x65,
	0x72, 0x64, 0x65, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x76, 0x32, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shop_v2_customer_proto_rawDescOnce sync.Once
	file_shop_v2_customer_proto_rawDescData = file_shop_v2_customer_proto_rawDesc
)

func file_shop_v2_customer_proto_rawDescGZIP() []byte {
	file_shop_v2_customer_proto_rawDescOnce.Do(func() {
		file_shop_v2_customer_proto_rawDescData = protoimpl.X.CompressGZIP(file_shop_v2_customer_proto_rawDescData)
	})
	return file_shop_v2_customer_proto_rawDescData
}

var file_shop_v2_customer_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_shop_v2_customer_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_shop_v2_customer_proto_goTypes = []interface{}{
	(Customer_CustomerType)(0), // 0: shop.v2.Customer.CustomerType
	(*Customer)(nil),           // 1: shop.v2.Customer
}
var file_shop_v2_customer_proto_depIdxs = []int32{
	0, // 0: shop.v2.Customer.customer_type:type_name -> shop.v2.Customer.CustomerType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_shop_v2_customer_proto_init() }
func file_shop_v2_customer_proto_init() {
	if File_shop_v2_customer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shop_v2_customer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Customer); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_shop_v2_customer_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_shop_v2_customer_proto_goTypes,
		DependencyIndexes: file_shop_v2_customer_proto_depIdxs,
		EnumInfos:         file_shop_v2_customer_proto_enumTypes,
		MessageInfos:      file_shop_v2_customer_proto_msgTypes,
	}.Build()
	File_shop_v2_customer_proto = out.File
	file_shop_v2_customer_proto_rawDesc = nil
	file_shop_v2_customer_proto_goTypes = nil
	file_shop_v2_customer_proto_depIdxs = nil
}