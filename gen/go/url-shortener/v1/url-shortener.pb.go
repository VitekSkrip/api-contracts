// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.29.1
// source: url-shortener/v1/url-shortener.proto

package urlv1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type SaveRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Url           string                 `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Alias         *string                `protobuf:"bytes,1,opt,name=alias,proto3,oneof" json:"alias,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SaveRequest) Reset() {
	*x = SaveRequest{}
	mi := &file_url_shortener_v1_url_shortener_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SaveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveRequest) ProtoMessage() {}

func (x *SaveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_url_shortener_v1_url_shortener_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveRequest.ProtoReflect.Descriptor instead.
func (*SaveRequest) Descriptor() ([]byte, []int) {
	return file_url_shortener_v1_url_shortener_proto_rawDescGZIP(), []int{0}
}

func (x *SaveRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *SaveRequest) GetAlias() string {
	if x != nil && x.Alias != nil {
		return *x.Alias
	}
	return ""
}

type SaveResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Alias         string                 `protobuf:"bytes,1,opt,name=alias,proto3" json:"alias,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SaveResponse) Reset() {
	*x = SaveResponse{}
	mi := &file_url_shortener_v1_url_shortener_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SaveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveResponse) ProtoMessage() {}

func (x *SaveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_url_shortener_v1_url_shortener_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveResponse.ProtoReflect.Descriptor instead.
func (*SaveResponse) Descriptor() ([]byte, []int) {
	return file_url_shortener_v1_url_shortener_proto_rawDescGZIP(), []int{1}
}

func (x *SaveResponse) GetAlias() string {
	if x != nil {
		return x.Alias
	}
	return ""
}

var File_url_shortener_v1_url_shortener_proto protoreflect.FileDescriptor

var file_url_shortener_v1_url_shortener_proto_rawDesc = []byte{
	0x0a, 0x24, 0x75, 0x72, 0x6c, 0x2d, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2f,
	0x76, 0x31, 0x2f, 0x75, 0x72, 0x6c, 0x2d, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x75, 0x72, 0x6c, 0x5f, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x65, 0x6e, 0x65, 0x72, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x44, 0x0a, 0x0b, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x6c, 0x12, 0x19, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x88, 0x01, 0x01, 0x42,
	0x08, 0x0a, 0x06, 0x5f, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x22, 0x24, 0x0a, 0x0c, 0x53, 0x61, 0x76,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x69,
	0x61, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x32,
	0x6b, 0x0a, 0x13, 0x55, 0x72, 0x6c, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x04, 0x53, 0x61, 0x76, 0x65, 0x12, 0x1a,
	0x2e, 0x75, 0x72, 0x6c, 0x5f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x53,
	0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x75, 0x72, 0x6c,
	0x5f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x3a,
	0x01, 0x2a, 0x22, 0x08, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x61, 0x76, 0x65, 0x42, 0x26, 0x5a, 0x24,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x75, 0x72,
	0x6c, 0x2d, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x3b, 0x75,
	0x72, 0x6c, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_url_shortener_v1_url_shortener_proto_rawDescOnce sync.Once
	file_url_shortener_v1_url_shortener_proto_rawDescData = file_url_shortener_v1_url_shortener_proto_rawDesc
)

func file_url_shortener_v1_url_shortener_proto_rawDescGZIP() []byte {
	file_url_shortener_v1_url_shortener_proto_rawDescOnce.Do(func() {
		file_url_shortener_v1_url_shortener_proto_rawDescData = protoimpl.X.CompressGZIP(file_url_shortener_v1_url_shortener_proto_rawDescData)
	})
	return file_url_shortener_v1_url_shortener_proto_rawDescData
}

var file_url_shortener_v1_url_shortener_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_url_shortener_v1_url_shortener_proto_goTypes = []any{
	(*SaveRequest)(nil),  // 0: url_shortener.SaveRequest
	(*SaveResponse)(nil), // 1: url_shortener.SaveResponse
}
var file_url_shortener_v1_url_shortener_proto_depIdxs = []int32{
	0, // 0: url_shortener.UrlShortenerService.Save:input_type -> url_shortener.SaveRequest
	1, // 1: url_shortener.UrlShortenerService.Save:output_type -> url_shortener.SaveResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_url_shortener_v1_url_shortener_proto_init() }
func file_url_shortener_v1_url_shortener_proto_init() {
	if File_url_shortener_v1_url_shortener_proto != nil {
		return
	}
	file_url_shortener_v1_url_shortener_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_url_shortener_v1_url_shortener_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_url_shortener_v1_url_shortener_proto_goTypes,
		DependencyIndexes: file_url_shortener_v1_url_shortener_proto_depIdxs,
		MessageInfos:      file_url_shortener_v1_url_shortener_proto_msgTypes,
	}.Build()
	File_url_shortener_v1_url_shortener_proto = out.File
	file_url_shortener_v1_url_shortener_proto_rawDesc = nil
	file_url_shortener_v1_url_shortener_proto_goTypes = nil
	file_url_shortener_v1_url_shortener_proto_depIdxs = nil
}
