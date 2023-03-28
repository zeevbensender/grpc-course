// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.5.0
// source: adder.proto

package proto

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

type AdderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	First  int32 `protobuf:"varint,1,opt,name=first,proto3" json:"first,omitempty"`
	Second int32 `protobuf:"varint,2,opt,name=second,proto3" json:"second,omitempty"`
}

func (x *AdderRequest) Reset() {
	*x = AdderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adder_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdderRequest) ProtoMessage() {}

func (x *AdderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_adder_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdderRequest.ProtoReflect.Descriptor instead.
func (*AdderRequest) Descriptor() ([]byte, []int) {
	return file_adder_proto_rawDescGZIP(), []int{0}
}

func (x *AdderRequest) GetFirst() int32 {
	if x != nil {
		return x.First
	}
	return 0
}

func (x *AdderRequest) GetSecond() int32 {
	if x != nil {
		return x.Second
	}
	return 0
}

type AdderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int32 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *AdderResponse) Reset() {
	*x = AdderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adder_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdderResponse) ProtoMessage() {}

func (x *AdderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_adder_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdderResponse.ProtoReflect.Descriptor instead.
func (*AdderResponse) Descriptor() ([]byte, []int) {
	return file_adder_proto_rawDescGZIP(), []int{1}
}

func (x *AdderResponse) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_adder_proto protoreflect.FileDescriptor

var file_adder_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x61, 0x64, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61,
	0x64, 0x64, 0x65, 0x72, 0x22, 0x3c, 0x0a, 0x0c, 0x41, 0x64, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x72, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x66, 0x69, 0x72, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x22, 0x27, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x42, 0x0a, 0x0c, 0x41,
	0x64, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x41,
	0x64, 0x64, 0x65, 0x72, 0x12, 0x13, 0x2e, 0x61, 0x64, 0x64, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x64, 0x64, 0x65,
	0x72, 0x2e, 0x41, 0x64, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x65,
	0x65, 0x76, 0x62, 0x65, 0x6e, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2d, 0x63, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2d, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2f,
	0x61, 0x64, 0x64, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_adder_proto_rawDescOnce sync.Once
	file_adder_proto_rawDescData = file_adder_proto_rawDesc
)

func file_adder_proto_rawDescGZIP() []byte {
	file_adder_proto_rawDescOnce.Do(func() {
		file_adder_proto_rawDescData = protoimpl.X.CompressGZIP(file_adder_proto_rawDescData)
	})
	return file_adder_proto_rawDescData
}

var file_adder_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_adder_proto_goTypes = []interface{}{
	(*AdderRequest)(nil),  // 0: adder.AdderRequest
	(*AdderResponse)(nil), // 1: adder.AdderResponse
}
var file_adder_proto_depIdxs = []int32{
	0, // 0: adder.AdderService.Adder:input_type -> adder.AdderRequest
	1, // 1: adder.AdderService.Adder:output_type -> adder.AdderResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_adder_proto_init() }
func file_adder_proto_init() {
	if File_adder_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_adder_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdderRequest); i {
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
		file_adder_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdderResponse); i {
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
			RawDescriptor: file_adder_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_adder_proto_goTypes,
		DependencyIndexes: file_adder_proto_depIdxs,
		MessageInfos:      file_adder_proto_msgTypes,
	}.Build()
	File_adder_proto = out.File
	file_adder_proto_rawDesc = nil
	file_adder_proto_goTypes = nil
	file_adder_proto_depIdxs = nil
}
