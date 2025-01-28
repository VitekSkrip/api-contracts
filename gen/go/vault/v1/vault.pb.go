// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.29.1
// source: vault/v1/vault.proto

package vaultv1

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

type CreateCollectionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description   *string                `protobuf:"bytes,2,opt,name=description,proto3,oneof" json:"description,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateCollectionRequest) Reset() {
	*x = CreateCollectionRequest{}
	mi := &file_vault_v1_vault_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCollectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCollectionRequest) ProtoMessage() {}

func (x *CreateCollectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vault_v1_vault_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCollectionRequest.ProtoReflect.Descriptor instead.
func (*CreateCollectionRequest) Descriptor() ([]byte, []int) {
	return file_vault_v1_vault_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCollectionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateCollectionRequest) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

type CreateCollectionResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateCollectionResponse) Reset() {
	*x = CreateCollectionResponse{}
	mi := &file_vault_v1_vault_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCollectionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCollectionResponse) ProtoMessage() {}

func (x *CreateCollectionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vault_v1_vault_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCollectionResponse.ProtoReflect.Descriptor instead.
func (*CreateCollectionResponse) Descriptor() ([]byte, []int) {
	return file_vault_v1_vault_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCollectionResponse) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetCollectionsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCollectionsRequest) Reset() {
	*x = GetCollectionsRequest{}
	mi := &file_vault_v1_vault_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCollectionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCollectionsRequest) ProtoMessage() {}

func (x *GetCollectionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vault_v1_vault_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCollectionsRequest.ProtoReflect.Descriptor instead.
func (*GetCollectionsRequest) Descriptor() ([]byte, []int) {
	return file_vault_v1_vault_proto_rawDescGZIP(), []int{2}
}

type GetCollectionsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Collections   []*Collection          `protobuf:"bytes,1,rep,name=collections,proto3" json:"collections,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCollectionsResponse) Reset() {
	*x = GetCollectionsResponse{}
	mi := &file_vault_v1_vault_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCollectionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCollectionsResponse) ProtoMessage() {}

func (x *GetCollectionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vault_v1_vault_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCollectionsResponse.ProtoReflect.Descriptor instead.
func (*GetCollectionsResponse) Descriptor() ([]byte, []int) {
	return file_vault_v1_vault_proto_rawDescGZIP(), []int{3}
}

func (x *GetCollectionsResponse) GetCollections() []*Collection {
	if x != nil {
		return x.Collections
	}
	return nil
}

type Collection struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Sort          uint32                 `protobuf:"varint,4,opt,name=sort,proto3" json:"sort,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Collection) Reset() {
	*x = Collection{}
	mi := &file_vault_v1_vault_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Collection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Collection) ProtoMessage() {}

func (x *Collection) ProtoReflect() protoreflect.Message {
	mi := &file_vault_v1_vault_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Collection.ProtoReflect.Descriptor instead.
func (*Collection) Descriptor() ([]byte, []int) {
	return file_vault_v1_vault_proto_rawDescGZIP(), []int{4}
}

func (x *Collection) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Collection) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Collection) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Collection) GetSort() uint32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

type GetCollectionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCollectionRequest) Reset() {
	*x = GetCollectionRequest{}
	mi := &file_vault_v1_vault_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCollectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCollectionRequest) ProtoMessage() {}

func (x *GetCollectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vault_v1_vault_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCollectionRequest.ProtoReflect.Descriptor instead.
func (*GetCollectionRequest) Descriptor() ([]byte, []int) {
	return file_vault_v1_vault_proto_rawDescGZIP(), []int{5}
}

func (x *GetCollectionRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetCollectionResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Saves         []*Save                `protobuf:"bytes,3,rep,name=saves,proto3" json:"saves,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCollectionResponse) Reset() {
	*x = GetCollectionResponse{}
	mi := &file_vault_v1_vault_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCollectionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCollectionResponse) ProtoMessage() {}

func (x *GetCollectionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vault_v1_vault_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCollectionResponse.ProtoReflect.Descriptor instead.
func (*GetCollectionResponse) Descriptor() ([]byte, []int) {
	return file_vault_v1_vault_proto_rawDescGZIP(), []int{6}
}

func (x *GetCollectionResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetCollectionResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *GetCollectionResponse) GetSaves() []*Save {
	if x != nil {
		return x.Saves
	}
	return nil
}

type CreateSavesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Title         string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Url           string                 `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Collection    uint64                 `protobuf:"varint,3,opt,name=collection,proto3" json:"collection,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateSavesRequest) Reset() {
	*x = CreateSavesRequest{}
	mi := &file_vault_v1_vault_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateSavesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSavesRequest) ProtoMessage() {}

func (x *CreateSavesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vault_v1_vault_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSavesRequest.ProtoReflect.Descriptor instead.
func (*CreateSavesRequest) Descriptor() ([]byte, []int) {
	return file_vault_v1_vault_proto_rawDescGZIP(), []int{7}
}

func (x *CreateSavesRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateSavesRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *CreateSavesRequest) GetCollection() uint64 {
	if x != nil {
		return x.Collection
	}
	return 0
}

type CreateSavesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateSavesResponse) Reset() {
	*x = CreateSavesResponse{}
	mi := &file_vault_v1_vault_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateSavesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSavesResponse) ProtoMessage() {}

func (x *CreateSavesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vault_v1_vault_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSavesResponse.ProtoReflect.Descriptor instead.
func (*CreateSavesResponse) Descriptor() ([]byte, []int) {
	return file_vault_v1_vault_proto_rawDescGZIP(), []int{8}
}

func (x *CreateSavesResponse) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Save struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Uri           string                 `protobuf:"bytes,3,opt,name=uri,proto3" json:"uri,omitempty"`
	Viewed        bool                   `protobuf:"varint,4,opt,name=viewed,proto3" json:"viewed,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Save) Reset() {
	*x = Save{}
	mi := &file_vault_v1_vault_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Save) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Save) ProtoMessage() {}

func (x *Save) ProtoReflect() protoreflect.Message {
	mi := &file_vault_v1_vault_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Save.ProtoReflect.Descriptor instead.
func (*Save) Descriptor() ([]byte, []int) {
	return file_vault_v1_vault_proto_rawDescGZIP(), []int{9}
}

func (x *Save) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Save) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Save) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *Save) GetViewed() bool {
	if x != nil {
		return x.Viewed
	}
	return false
}

type GetSaveRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetSaveRequest) Reset() {
	*x = GetSaveRequest{}
	mi := &file_vault_v1_vault_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSaveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSaveRequest) ProtoMessage() {}

func (x *GetSaveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vault_v1_vault_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSaveRequest.ProtoReflect.Descriptor instead.
func (*GetSaveRequest) Descriptor() ([]byte, []int) {
	return file_vault_v1_vault_proto_rawDescGZIP(), []int{10}
}

func (x *GetSaveRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetSaveResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Save          *Save                  `protobuf:"bytes,1,opt,name=save,proto3" json:"save,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetSaveResponse) Reset() {
	*x = GetSaveResponse{}
	mi := &file_vault_v1_vault_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSaveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSaveResponse) ProtoMessage() {}

func (x *GetSaveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vault_v1_vault_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSaveResponse.ProtoReflect.Descriptor instead.
func (*GetSaveResponse) Descriptor() ([]byte, []int) {
	return file_vault_v1_vault_proto_rawDescGZIP(), []int{11}
}

func (x *GetSaveResponse) GetSave() *Save {
	if x != nil {
		return x.Save
	}
	return nil
}

var File_vault_v1_vault_proto protoreflect.FileDescriptor

var file_vault_v1_vault_proto_rawDesc = []byte{
	0x0a, 0x14, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x61, 0x75, 0x6c, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x64, 0x0a, 0x17, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01,
	0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x2a, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x17, 0x0a,
	0x15, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x4d, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x33, 0x0a, 0x0b, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x43, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x66, 0x0a, 0x0a, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x22, 0x26, 0x0a,
	0x14, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x70, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x05, 0x73, 0x61, 0x76, 0x65, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x53, 0x61, 0x76, 0x65,
	0x52, 0x05, 0x73, 0x61, 0x76, 0x65, 0x73, 0x22, 0x5c, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x53, 0x61, 0x76, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x25, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x61, 0x76, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x56, 0x0a, 0x04,
	0x53, 0x61, 0x76, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72,
	0x69, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12, 0x16, 0x0a, 0x06,
	0x76, 0x69, 0x65, 0x77, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x76, 0x69,
	0x65, 0x77, 0x65, 0x64, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x53, 0x61, 0x76, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x32, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x61, 0x76,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x73, 0x61, 0x76,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2e,
	0x53, 0x61, 0x76, 0x65, 0x52, 0x04, 0x73, 0x61, 0x76, 0x65, 0x32, 0xff, 0x03, 0x0a, 0x0c, 0x56,
	0x61, 0x75, 0x6c, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6f, 0x0a, 0x10, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1e, 0x2e, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1f, 0x2e, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a, 0x22, 0x0f, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x66, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1c,
	0x2e, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x76,
	0x61, 0x75, 0x6c, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x68, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x2e, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x47, 0x65,
	0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x5a,
	0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x61, 0x76, 0x65, 0x73, 0x12, 0x19, 0x2e,
	0x76, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x61, 0x76, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x76, 0x61, 0x75, 0x6c, 0x74,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x61, 0x76, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x3a, 0x01, 0x2a, 0x22,
	0x09, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x61, 0x76, 0x65, 0x73, 0x12, 0x50, 0x0a, 0x07, 0x47, 0x65,
	0x74, 0x53, 0x61, 0x76, 0x65, 0x12, 0x15, 0x2e, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x47, 0x65,
	0x74, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x76,
	0x61, 0x75, 0x6c, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x61, 0x76, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x20, 0x5a, 0x1e,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x76, 0x61,
	0x75, 0x6c, 0x74, 0x2e, 0x76, 0x31, 0x3b, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vault_v1_vault_proto_rawDescOnce sync.Once
	file_vault_v1_vault_proto_rawDescData = file_vault_v1_vault_proto_rawDesc
)

func file_vault_v1_vault_proto_rawDescGZIP() []byte {
	file_vault_v1_vault_proto_rawDescOnce.Do(func() {
		file_vault_v1_vault_proto_rawDescData = protoimpl.X.CompressGZIP(file_vault_v1_vault_proto_rawDescData)
	})
	return file_vault_v1_vault_proto_rawDescData
}

var file_vault_v1_vault_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_vault_v1_vault_proto_goTypes = []any{
	(*CreateCollectionRequest)(nil),  // 0: vault.CreateCollectionRequest
	(*CreateCollectionResponse)(nil), // 1: vault.CreateCollectionResponse
	(*GetCollectionsRequest)(nil),    // 2: vault.GetCollectionsRequest
	(*GetCollectionsResponse)(nil),   // 3: vault.GetCollectionsResponse
	(*Collection)(nil),               // 4: vault.Collection
	(*GetCollectionRequest)(nil),     // 5: vault.GetCollectionRequest
	(*GetCollectionResponse)(nil),    // 6: vault.GetCollectionResponse
	(*CreateSavesRequest)(nil),       // 7: vault.CreateSavesRequest
	(*CreateSavesResponse)(nil),      // 8: vault.CreateSavesResponse
	(*Save)(nil),                     // 9: vault.Save
	(*GetSaveRequest)(nil),           // 10: vault.GetSaveRequest
	(*GetSaveResponse)(nil),          // 11: vault.GetSaveResponse
}
var file_vault_v1_vault_proto_depIdxs = []int32{
	4,  // 0: vault.GetCollectionsResponse.collections:type_name -> vault.Collection
	9,  // 1: vault.GetCollectionResponse.saves:type_name -> vault.Save
	9,  // 2: vault.GetSaveResponse.save:type_name -> vault.Save
	0,  // 3: vault.VaultService.CreateCollection:input_type -> vault.CreateCollectionRequest
	2,  // 4: vault.VaultService.GetCollections:input_type -> vault.GetCollectionsRequest
	5,  // 5: vault.VaultService.GetCollection:input_type -> vault.GetCollectionRequest
	7,  // 6: vault.VaultService.CreateSaves:input_type -> vault.CreateSavesRequest
	10, // 7: vault.VaultService.GetSave:input_type -> vault.GetSaveRequest
	1,  // 8: vault.VaultService.CreateCollection:output_type -> vault.CreateCollectionResponse
	3,  // 9: vault.VaultService.GetCollections:output_type -> vault.GetCollectionsResponse
	6,  // 10: vault.VaultService.GetCollection:output_type -> vault.GetCollectionResponse
	8,  // 11: vault.VaultService.CreateSaves:output_type -> vault.CreateSavesResponse
	11, // 12: vault.VaultService.GetSave:output_type -> vault.GetSaveResponse
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_vault_v1_vault_proto_init() }
func file_vault_v1_vault_proto_init() {
	if File_vault_v1_vault_proto != nil {
		return
	}
	file_vault_v1_vault_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_vault_v1_vault_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_vault_v1_vault_proto_goTypes,
		DependencyIndexes: file_vault_v1_vault_proto_depIdxs,
		MessageInfos:      file_vault_v1_vault_proto_msgTypes,
	}.Build()
	File_vault_v1_vault_proto = out.File
	file_vault_v1_vault_proto_rawDesc = nil
	file_vault_v1_vault_proto_goTypes = nil
	file_vault_v1_vault_proto_depIdxs = nil
}
