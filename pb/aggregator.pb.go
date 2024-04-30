// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: pb/aggregator.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AggregateDistanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value     float64 `protobuf:"fixed64,1,opt,name=Value,proto3" json:"Value,omitempty"`
	OBUId     int64   `protobuf:"varint,2,opt,name=OBUId,proto3" json:"OBUId,omitempty"`
	Timestamp int64   `protobuf:"varint,3,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
}

func (x *AggregateDistanceRequest) Reset() {
	*x = AggregateDistanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_aggregator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AggregateDistanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AggregateDistanceRequest) ProtoMessage() {}

func (x *AggregateDistanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_aggregator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AggregateDistanceRequest.ProtoReflect.Descriptor instead.
func (*AggregateDistanceRequest) Descriptor() ([]byte, []int) {
	return file_pb_aggregator_proto_rawDescGZIP(), []int{0}
}

func (x *AggregateDistanceRequest) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *AggregateDistanceRequest) GetOBUId() int64 {
	if x != nil {
		return x.OBUId
	}
	return 0
}

func (x *AggregateDistanceRequest) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type AggregateDistanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *AggregateDistanceResponse) Reset() {
	*x = AggregateDistanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_aggregator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AggregateDistanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AggregateDistanceResponse) ProtoMessage() {}

func (x *AggregateDistanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_aggregator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AggregateDistanceResponse.ProtoReflect.Descriptor instead.
func (*AggregateDistanceResponse) Descriptor() ([]byte, []int) {
	return file_pb_aggregator_proto_rawDescGZIP(), []int{1}
}

func (x *AggregateDistanceResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GenerateInvoiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OBUId int64 `protobuf:"varint,1,opt,name=OBUId,proto3" json:"OBUId,omitempty"`
}

func (x *GenerateInvoiceRequest) Reset() {
	*x = GenerateInvoiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_aggregator_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateInvoiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateInvoiceRequest) ProtoMessage() {}

func (x *GenerateInvoiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_aggregator_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateInvoiceRequest.ProtoReflect.Descriptor instead.
func (*GenerateInvoiceRequest) Descriptor() ([]byte, []int) {
	return file_pb_aggregator_proto_rawDescGZIP(), []int{2}
}

func (x *GenerateInvoiceRequest) GetOBUId() int64 {
	if x != nil {
		return x.OBUId
	}
	return 0
}

type GenerateInvoiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OBUId         int64   `protobuf:"varint,1,opt,name=OBUId,proto3" json:"OBUId,omitempty"`
	TotalDistance float64 `protobuf:"fixed64,2,opt,name=TotalDistance,proto3" json:"TotalDistance,omitempty"`
	Amount        float64 `protobuf:"fixed64,3,opt,name=Amount,proto3" json:"Amount,omitempty"`
}

func (x *GenerateInvoiceResponse) Reset() {
	*x = GenerateInvoiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_aggregator_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateInvoiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateInvoiceResponse) ProtoMessage() {}

func (x *GenerateInvoiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_aggregator_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateInvoiceResponse.ProtoReflect.Descriptor instead.
func (*GenerateInvoiceResponse) Descriptor() ([]byte, []int) {
	return file_pb_aggregator_proto_rawDescGZIP(), []int{3}
}

func (x *GenerateInvoiceResponse) GetOBUId() int64 {
	if x != nil {
		return x.OBUId
	}
	return 0
}

func (x *GenerateInvoiceResponse) GetTotalDistance() float64 {
	if x != nil {
		return x.TotalDistance
	}
	return 0
}

func (x *GenerateInvoiceResponse) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

var File_pb_aggregator_proto protoreflect.FileDescriptor

var file_pb_aggregator_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x62, 0x2f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x64, 0x0a, 0x18, 0x41, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x4f,
	0x42, 0x55, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x4f, 0x42, 0x55, 0x49,
	0x64, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22,
	0x35, 0x0a, 0x19, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x44, 0x69, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2e, 0x0a, 0x16, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x4f, 0x42, 0x55, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x4f, 0x42, 0x55, 0x49, 0x64, 0x22, 0x6d, 0x0a, 0x17, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x4f, 0x42, 0x55, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x4f, 0x42, 0x55, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x54, 0x6f, 0x74, 0x61, 0x6c,
	0x44, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d,
	0x54, 0x6f, 0x74, 0x61, 0x6c, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0xb2, 0x01, 0x0a, 0x0e, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x52, 0x0a, 0x11, 0x41, 0x67, 0x67, 0x72,
	0x65, 0x67, 0x61, 0x74, 0x65, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1c, 0x2e,
	0x70, 0x62, 0x2e, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x44, 0x69, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x62,
	0x2e, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x0f,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x12,
	0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x76,
	0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x62,
	0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x65, 0x76, 0x61, 0x6c, 0x73, 0x61,
	0x62, 0x68, 0x61, 0x6e, 0x69, 0x2f, 0x74, 0x6f, 0x6c, 0x6c, 0x2d, 0x63, 0x61, 0x6c, 0x63, 0x75,
	0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_aggregator_proto_rawDescOnce sync.Once
	file_pb_aggregator_proto_rawDescData = file_pb_aggregator_proto_rawDesc
)

func file_pb_aggregator_proto_rawDescGZIP() []byte {
	file_pb_aggregator_proto_rawDescOnce.Do(func() {
		file_pb_aggregator_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_aggregator_proto_rawDescData)
	})
	return file_pb_aggregator_proto_rawDescData
}

var file_pb_aggregator_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pb_aggregator_proto_goTypes = []interface{}{
	(*AggregateDistanceRequest)(nil),  // 0: pb.AggregateDistanceRequest
	(*AggregateDistanceResponse)(nil), // 1: pb.AggregateDistanceResponse
	(*GenerateInvoiceRequest)(nil),    // 2: pb.GenerateInvoiceRequest
	(*GenerateInvoiceResponse)(nil),   // 3: pb.GenerateInvoiceResponse
}
var file_pb_aggregator_proto_depIdxs = []int32{
	0, // 0: pb.InvoiceService.AggregateDistance:input_type -> pb.AggregateDistanceRequest
	2, // 1: pb.InvoiceService.GenerateInvoice:input_type -> pb.GenerateInvoiceRequest
	1, // 2: pb.InvoiceService.AggregateDistance:output_type -> pb.AggregateDistanceResponse
	3, // 3: pb.InvoiceService.GenerateInvoice:output_type -> pb.GenerateInvoiceResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_aggregator_proto_init() }
func file_pb_aggregator_proto_init() {
	if File_pb_aggregator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_aggregator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AggregateDistanceRequest); i {
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
		file_pb_aggregator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AggregateDistanceResponse); i {
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
		file_pb_aggregator_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateInvoiceRequest); i {
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
		file_pb_aggregator_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateInvoiceResponse); i {
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
			RawDescriptor: file_pb_aggregator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_aggregator_proto_goTypes,
		DependencyIndexes: file_pb_aggregator_proto_depIdxs,
		MessageInfos:      file_pb_aggregator_proto_msgTypes,
	}.Build()
	File_pb_aggregator_proto = out.File
	file_pb_aggregator_proto_rawDesc = nil
	file_pb_aggregator_proto_goTypes = nil
	file_pb_aggregator_proto_depIdxs = nil
}