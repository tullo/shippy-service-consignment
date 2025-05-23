// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.19.4
// source: proto/consignment.proto

package consignment

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Consignment struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description   string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Weight        int32                  `protobuf:"varint,3,opt,name=weight,proto3" json:"weight,omitempty"`
	Containers    []*Container           `protobuf:"bytes,4,rep,name=containers,proto3" json:"containers,omitempty"`
	VesselId      string                 `protobuf:"bytes,5,opt,name=vessel_id,json=vesselId,proto3" json:"vessel_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Consignment) Reset() {
	*x = Consignment{}
	mi := &file_proto_consignment_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Consignment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Consignment) ProtoMessage() {}

func (x *Consignment) ProtoReflect() protoreflect.Message {
	mi := &file_proto_consignment_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Consignment.ProtoReflect.Descriptor instead.
func (*Consignment) Descriptor() ([]byte, []int) {
	return file_proto_consignment_proto_rawDescGZIP(), []int{0}
}

func (x *Consignment) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Consignment) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Consignment) GetWeight() int32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *Consignment) GetContainers() []*Container {
	if x != nil {
		return x.Containers
	}
	return nil
}

func (x *Consignment) GetVesselId() string {
	if x != nil {
		return x.VesselId
	}
	return ""
}

type Container struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomerId    string                 `protobuf:"bytes,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Origin        string                 `protobuf:"bytes,3,opt,name=origin,proto3" json:"origin,omitempty"`
	UserId        string                 `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Container) Reset() {
	*x = Container{}
	mi := &file_proto_consignment_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Container) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Container) ProtoMessage() {}

func (x *Container) ProtoReflect() protoreflect.Message {
	mi := &file_proto_consignment_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Container.ProtoReflect.Descriptor instead.
func (*Container) Descriptor() ([]byte, []int) {
	return file_proto_consignment_proto_rawDescGZIP(), []int{1}
}

func (x *Container) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Container) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *Container) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

func (x *Container) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	mi := &file_proto_consignment_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_consignment_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_proto_consignment_proto_rawDescGZIP(), []int{2}
}

type Response struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Created       bool                   `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
	Consignment   *Consignment           `protobuf:"bytes,2,opt,name=consignment,proto3" json:"consignment,omitempty"`
	Consignments  []*Consignment         `protobuf:"bytes,3,rep,name=consignments,proto3" json:"consignments,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Response) Reset() {
	*x = Response{}
	mi := &file_proto_consignment_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_consignment_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_proto_consignment_proto_rawDescGZIP(), []int{3}
}

func (x *Response) GetCreated() bool {
	if x != nil {
		return x.Created
	}
	return false
}

func (x *Response) GetConsignment() *Consignment {
	if x != nil {
		return x.Consignment
	}
	return nil
}

func (x *Response) GetConsignments() []*Consignment {
	if x != nil {
		return x.Consignments
	}
	return nil
}

var File_proto_consignment_proto protoreflect.FileDescriptor

const file_proto_consignment_proto_rawDesc = "" +
	"\n" +
	"\x17proto/consignment.proto\x12\vconsignment\"\xac\x01\n" +
	"\vConsignment\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12 \n" +
	"\vdescription\x18\x02 \x01(\tR\vdescription\x12\x16\n" +
	"\x06weight\x18\x03 \x01(\x05R\x06weight\x126\n" +
	"\n" +
	"containers\x18\x04 \x03(\v2\x16.consignment.ContainerR\n" +
	"containers\x12\x1b\n" +
	"\tvessel_id\x18\x05 \x01(\tR\bvesselId\"m\n" +
	"\tContainer\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x1f\n" +
	"\vcustomer_id\x18\x02 \x01(\tR\n" +
	"customerId\x12\x16\n" +
	"\x06origin\x18\x03 \x01(\tR\x06origin\x12\x17\n" +
	"\auser_id\x18\x04 \x01(\tR\x06userId\"\f\n" +
	"\n" +
	"GetRequest\"\x9e\x01\n" +
	"\bResponse\x12\x18\n" +
	"\acreated\x18\x01 \x01(\bR\acreated\x12:\n" +
	"\vconsignment\x18\x02 \x01(\v2\x18.consignment.ConsignmentR\vconsignment\x12<\n" +
	"\fconsignments\x18\x03 \x03(\v2\x18.consignment.ConsignmentR\fconsignments2\x9e\x01\n" +
	"\x0fShippingService\x12F\n" +
	"\x11CreateConsignment\x12\x18.consignment.Consignment\x1a\x15.consignment.Response\"\x00\x12C\n" +
	"\x0fGetConsignments\x12\x17.consignment.GetRequest\x1a\x15.consignment.Response\"\x00B\x15Z\x13./proto;consignmentb\x06proto3"

var (
	file_proto_consignment_proto_rawDescOnce sync.Once
	file_proto_consignment_proto_rawDescData []byte
)

func file_proto_consignment_proto_rawDescGZIP() []byte {
	file_proto_consignment_proto_rawDescOnce.Do(func() {
		file_proto_consignment_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_consignment_proto_rawDesc), len(file_proto_consignment_proto_rawDesc)))
	})
	return file_proto_consignment_proto_rawDescData
}

var file_proto_consignment_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_consignment_proto_goTypes = []any{
	(*Consignment)(nil), // 0: consignment.Consignment
	(*Container)(nil),   // 1: consignment.Container
	(*GetRequest)(nil),  // 2: consignment.GetRequest
	(*Response)(nil),    // 3: consignment.Response
}
var file_proto_consignment_proto_depIdxs = []int32{
	1, // 0: consignment.Consignment.containers:type_name -> consignment.Container
	0, // 1: consignment.Response.consignment:type_name -> consignment.Consignment
	0, // 2: consignment.Response.consignments:type_name -> consignment.Consignment
	0, // 3: consignment.ShippingService.CreateConsignment:input_type -> consignment.Consignment
	2, // 4: consignment.ShippingService.GetConsignments:input_type -> consignment.GetRequest
	3, // 5: consignment.ShippingService.CreateConsignment:output_type -> consignment.Response
	3, // 6: consignment.ShippingService.GetConsignments:output_type -> consignment.Response
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_consignment_proto_init() }
func file_proto_consignment_proto_init() {
	if File_proto_consignment_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_consignment_proto_rawDesc), len(file_proto_consignment_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_consignment_proto_goTypes,
		DependencyIndexes: file_proto_consignment_proto_depIdxs,
		MessageInfos:      file_proto_consignment_proto_msgTypes,
	}.Build()
	File_proto_consignment_proto = out.File
	file_proto_consignment_proto_goTypes = nil
	file_proto_consignment_proto_depIdxs = nil
}
