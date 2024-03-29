// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.6
// source: common.proto

package entity

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// Entity message type for responding
type Entity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Url         string `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *Entity) Reset() {
	*x = Entity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entity) ProtoMessage() {}

func (x *Entity) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entity.ProtoReflect.Descriptor instead.
func (*Entity) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

func (x *Entity) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Entity) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Entity) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Entity) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

// Defintion of request and respone for each services
type CreateEntityReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Entity `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
}

func (x *CreateEntityReq) Reset() {
	*x = CreateEntityReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEntityReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEntityReq) ProtoMessage() {}

func (x *CreateEntityReq) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEntityReq.ProtoReflect.Descriptor instead.
func (*CreateEntityReq) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{1}
}

func (x *CreateEntityReq) GetEntity() *Entity {
	if x != nil {
		return x.Entity
	}
	return nil
}

type CreateEntityRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Entity `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
}

func (x *CreateEntityRes) Reset() {
	*x = CreateEntityRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEntityRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEntityRes) ProtoMessage() {}

func (x *CreateEntityRes) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEntityRes.ProtoReflect.Descriptor instead.
func (*CreateEntityRes) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{2}
}

func (x *CreateEntityRes) GetEntity() *Entity {
	if x != nil {
		return x.Entity
	}
	return nil
}

type UpdateEntityReqBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Url         string `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *UpdateEntityReqBody) Reset() {
	*x = UpdateEntityReqBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEntityReqBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEntityReqBody) ProtoMessage() {}

func (x *UpdateEntityReqBody) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEntityReqBody.ProtoReflect.Descriptor instead.
func (*UpdateEntityReqBody) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateEntityReqBody) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateEntityReqBody) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateEntityReqBody) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type UpdateEntityReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Entity *UpdateEntityReqBody `protobuf:"bytes,2,opt,name=entity,proto3" json:"entity,omitempty"`
}

func (x *UpdateEntityReq) Reset() {
	*x = UpdateEntityReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEntityReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEntityReq) ProtoMessage() {}

func (x *UpdateEntityReq) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEntityReq.ProtoReflect.Descriptor instead.
func (*UpdateEntityReq) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateEntityReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateEntityReq) GetEntity() *UpdateEntityReqBody {
	if x != nil {
		return x.Entity
	}
	return nil
}

type UpdateEntityRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Entity `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
}

func (x *UpdateEntityRes) Reset() {
	*x = UpdateEntityRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEntityRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEntityRes) ProtoMessage() {}

func (x *UpdateEntityRes) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEntityRes.ProtoReflect.Descriptor instead.
func (*UpdateEntityRes) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateEntityRes) GetEntity() *Entity {
	if x != nil {
		return x.Entity
	}
	return nil
}

type ReadEntityReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ReadEntityReq) Reset() {
	*x = ReadEntityReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadEntityReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadEntityReq) ProtoMessage() {}

func (x *ReadEntityReq) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadEntityReq.ProtoReflect.Descriptor instead.
func (*ReadEntityReq) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{6}
}

func (x *ReadEntityReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ReadEntityRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Entity `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
}

func (x *ReadEntityRes) Reset() {
	*x = ReadEntityRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadEntityRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadEntityRes) ProtoMessage() {}

func (x *ReadEntityRes) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadEntityRes.ProtoReflect.Descriptor instead.
func (*ReadEntityRes) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{7}
}

func (x *ReadEntityRes) GetEntity() *Entity {
	if x != nil {
		return x.Entity
	}
	return nil
}

type DeleteEntityReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteEntityReq) Reset() {
	*x = DeleteEntityReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEntityReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEntityReq) ProtoMessage() {}

func (x *DeleteEntityReq) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEntityReq.ProtoReflect.Descriptor instead.
func (*DeleteEntityReq) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteEntityReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteEntityRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteEntityRes) Reset() {
	*x = DeleteEntityRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEntityRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEntityRes) ProtoMessage() {}

func (x *DeleteEntityRes) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEntityRes.ProtoReflect.Descriptor instead.
func (*DeleteEntityRes) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteEntityRes) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type ListEntitiesRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entities []*Entity `protobuf:"bytes,1,rep,name=entities,proto3" json:"entities,omitempty"`
}

func (x *ListEntitiesRes) Reset() {
	*x = ListEntitiesRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEntitiesRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEntitiesRes) ProtoMessage() {}

func (x *ListEntitiesRes) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEntitiesRes.ProtoReflect.Descriptor instead.
func (*ListEntitiesRes) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{10}
}

func (x *ListEntitiesRes) GetEntities() []*Entity {
	if x != nil {
		return x.Entities
	}
	return nil
}

var File_common_proto protoreflect.FileDescriptor

var file_common_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x73, 0x0a, 0x06, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x02,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x90, 0x01, 0x01, 0x52,
	0x03, 0x75, 0x72, 0x6c, 0x22, 0x43, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x12, 0x30, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10,
	0x01, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x39, 0x0a, 0x0f, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x06,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x06, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x22, 0x70, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x1b, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02,
	0x10, 0x02, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x03, 0x75, 0x72,
	0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x90, 0x01,
	0x01, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x6a, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x3d, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x42, 0x6f, 0x64, 0x79,
	0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x22, 0x39, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x52, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x29, 0x0a,
	0x0d, 0x52, 0x65, 0x61, 0x64, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x12, 0x18,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72,
	0x03, 0xb0, 0x01, 0x01, 0x52, 0x02, 0x69, 0x64, 0x22, 0x37, 0x0a, 0x0d, 0x52, 0x65, 0x61, 0x64,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x06, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x22, 0x2b, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2b,
	0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x3d, 0x0a, 0x0f, 0x4c,
	0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x12, 0x2a,
	0x0a, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_proto_rawDescOnce sync.Once
	file_common_proto_rawDescData = file_common_proto_rawDesc
)

func file_common_proto_rawDescGZIP() []byte {
	file_common_proto_rawDescOnce.Do(func() {
		file_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_proto_rawDescData)
	})
	return file_common_proto_rawDescData
}

var file_common_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_common_proto_goTypes = []interface{}{
	(*Entity)(nil),              // 0: entity.Entity
	(*CreateEntityReq)(nil),     // 1: entity.CreateEntityReq
	(*CreateEntityRes)(nil),     // 2: entity.CreateEntityRes
	(*UpdateEntityReqBody)(nil), // 3: entity.UpdateEntityReqBody
	(*UpdateEntityReq)(nil),     // 4: entity.UpdateEntityReq
	(*UpdateEntityRes)(nil),     // 5: entity.UpdateEntityRes
	(*ReadEntityReq)(nil),       // 6: entity.ReadEntityReq
	(*ReadEntityRes)(nil),       // 7: entity.ReadEntityRes
	(*DeleteEntityReq)(nil),     // 8: entity.DeleteEntityReq
	(*DeleteEntityRes)(nil),     // 9: entity.DeleteEntityRes
	(*ListEntitiesRes)(nil),     // 10: entity.ListEntitiesRes
}
var file_common_proto_depIdxs = []int32{
	0, // 0: entity.CreateEntityReq.entity:type_name -> entity.Entity
	0, // 1: entity.CreateEntityRes.entity:type_name -> entity.Entity
	3, // 2: entity.UpdateEntityReq.entity:type_name -> entity.UpdateEntityReqBody
	0, // 3: entity.UpdateEntityRes.entity:type_name -> entity.Entity
	0, // 4: entity.ReadEntityRes.entity:type_name -> entity.Entity
	0, // 5: entity.ListEntitiesRes.entities:type_name -> entity.Entity
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_common_proto_init() }
func file_common_proto_init() {
	if File_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Entity); i {
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
		file_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEntityReq); i {
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
		file_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEntityRes); i {
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
		file_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEntityReqBody); i {
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
		file_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEntityReq); i {
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
		file_common_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEntityRes); i {
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
		file_common_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadEntityReq); i {
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
		file_common_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadEntityRes); i {
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
		file_common_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEntityReq); i {
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
		file_common_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEntityRes); i {
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
		file_common_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEntitiesRes); i {
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
			RawDescriptor: file_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_proto_goTypes,
		DependencyIndexes: file_common_proto_depIdxs,
		MessageInfos:      file_common_proto_msgTypes,
	}.Build()
	File_common_proto = out.File
	file_common_proto_rawDesc = nil
	file_common_proto_goTypes = nil
	file_common_proto_depIdxs = nil
}
