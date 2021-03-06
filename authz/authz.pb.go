// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.1
// source: authz/authz.proto

package authz

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

type CheckPermissionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ShopId int64  `protobuf:"varint,2,opt,name=shop_id,json=shopId,proto3" json:"shop_id,omitempty"`
	Act    string `protobuf:"bytes,3,opt,name=act,proto3" json:"act,omitempty"`
}

func (x *CheckPermissionRequest) Reset() {
	*x = CheckPermissionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authz_authz_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckPermissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckPermissionRequest) ProtoMessage() {}

func (x *CheckPermissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authz_authz_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckPermissionRequest.ProtoReflect.Descriptor instead.
func (*CheckPermissionRequest) Descriptor() ([]byte, []int) {
	return file_authz_authz_proto_rawDescGZIP(), []int{0}
}

func (x *CheckPermissionRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CheckPermissionRequest) GetShopId() int64 {
	if x != nil {
		return x.ShopId
	}
	return 0
}

func (x *CheckPermissionRequest) GetAct() string {
	if x != nil {
		return x.Act
	}
	return ""
}

type CheckPermissionResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool  `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	Code   int32 `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *CheckPermissionResult) Reset() {
	*x = CheckPermissionResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authz_authz_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckPermissionResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckPermissionResult) ProtoMessage() {}

func (x *CheckPermissionResult) ProtoReflect() protoreflect.Message {
	mi := &file_authz_authz_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckPermissionResult.ProtoReflect.Descriptor instead.
func (*CheckPermissionResult) Descriptor() ([]byte, []int) {
	return file_authz_authz_proto_rawDescGZIP(), []int{1}
}

func (x *CheckPermissionResult) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

func (x *CheckPermissionResult) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

type AddRoleToDomainRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ShopId int64  `protobuf:"varint,2,opt,name=shop_id,json=shopId,proto3" json:"shop_id,omitempty"`
	Act    string `protobuf:"bytes,3,opt,name=act,proto3" json:"act,omitempty"`
}

func (x *AddRoleToDomainRequest) Reset() {
	*x = AddRoleToDomainRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authz_authz_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRoleToDomainRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRoleToDomainRequest) ProtoMessage() {}

func (x *AddRoleToDomainRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authz_authz_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRoleToDomainRequest.ProtoReflect.Descriptor instead.
func (*AddRoleToDomainRequest) Descriptor() ([]byte, []int) {
	return file_authz_authz_proto_rawDescGZIP(), []int{2}
}

func (x *AddRoleToDomainRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *AddRoleToDomainRequest) GetShopId() int64 {
	if x != nil {
		return x.ShopId
	}
	return 0
}

func (x *AddRoleToDomainRequest) GetAct() string {
	if x != nil {
		return x.Act
	}
	return ""
}

type AddRoleToDomainResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool  `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	Code   int32 `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *AddRoleToDomainResult) Reset() {
	*x = AddRoleToDomainResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authz_authz_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRoleToDomainResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRoleToDomainResult) ProtoMessage() {}

func (x *AddRoleToDomainResult) ProtoReflect() protoreflect.Message {
	mi := &file_authz_authz_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRoleToDomainResult.ProtoReflect.Descriptor instead.
func (*AddRoleToDomainResult) Descriptor() ([]byte, []int) {
	return file_authz_authz_proto_rawDescGZIP(), []int{3}
}

func (x *AddRoleToDomainResult) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

func (x *AddRoleToDomainResult) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

type GetRolesInDomainRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ShopId int64 `protobuf:"varint,2,opt,name=shop_id,json=shopId,proto3" json:"shop_id,omitempty"`
}

func (x *GetRolesInDomainRequest) Reset() {
	*x = GetRolesInDomainRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authz_authz_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRolesInDomainRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRolesInDomainRequest) ProtoMessage() {}

func (x *GetRolesInDomainRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authz_authz_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRolesInDomainRequest.ProtoReflect.Descriptor instead.
func (*GetRolesInDomainRequest) Descriptor() ([]byte, []int) {
	return file_authz_authz_proto_rawDescGZIP(), []int{4}
}

func (x *GetRolesInDomainRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetRolesInDomainRequest) GetShopId() int64 {
	if x != nil {
		return x.ShopId
	}
	return 0
}

type GetRolesInDomainResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Roles []string `protobuf:"bytes,1,rep,name=roles,proto3" json:"roles,omitempty"`
	Code  int32    `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *GetRolesInDomainResult) Reset() {
	*x = GetRolesInDomainResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authz_authz_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRolesInDomainResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRolesInDomainResult) ProtoMessage() {}

func (x *GetRolesInDomainResult) ProtoReflect() protoreflect.Message {
	mi := &file_authz_authz_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRolesInDomainResult.ProtoReflect.Descriptor instead.
func (*GetRolesInDomainResult) Descriptor() ([]byte, []int) {
	return file_authz_authz_proto_rawDescGZIP(), []int{5}
}

func (x *GetRolesInDomainResult) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *GetRolesInDomainResult) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

type GetImplicitRolesInDomainRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ShopId int64 `protobuf:"varint,2,opt,name=shop_id,json=shopId,proto3" json:"shop_id,omitempty"`
}

func (x *GetImplicitRolesInDomainRequest) Reset() {
	*x = GetImplicitRolesInDomainRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authz_authz_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetImplicitRolesInDomainRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetImplicitRolesInDomainRequest) ProtoMessage() {}

func (x *GetImplicitRolesInDomainRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authz_authz_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetImplicitRolesInDomainRequest.ProtoReflect.Descriptor instead.
func (*GetImplicitRolesInDomainRequest) Descriptor() ([]byte, []int) {
	return file_authz_authz_proto_rawDescGZIP(), []int{6}
}

func (x *GetImplicitRolesInDomainRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetImplicitRolesInDomainRequest) GetShopId() int64 {
	if x != nil {
		return x.ShopId
	}
	return 0
}

type GetImplicitRolesInDomainResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Roles []string `protobuf:"bytes,1,rep,name=roles,proto3" json:"roles,omitempty"`
	Code  int32    `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *GetImplicitRolesInDomainResult) Reset() {
	*x = GetImplicitRolesInDomainResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authz_authz_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetImplicitRolesInDomainResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetImplicitRolesInDomainResult) ProtoMessage() {}

func (x *GetImplicitRolesInDomainResult) ProtoReflect() protoreflect.Message {
	mi := &file_authz_authz_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetImplicitRolesInDomainResult.ProtoReflect.Descriptor instead.
func (*GetImplicitRolesInDomainResult) Descriptor() ([]byte, []int) {
	return file_authz_authz_proto_rawDescGZIP(), []int{7}
}

func (x *GetImplicitRolesInDomainResult) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *GetImplicitRolesInDomainResult) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

type GenerateOwnerRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ShopId int64 `protobuf:"varint,2,opt,name=shop_id,json=shopId,proto3" json:"shop_id,omitempty"`
}

func (x *GenerateOwnerRoleRequest) Reset() {
	*x = GenerateOwnerRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authz_authz_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateOwnerRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateOwnerRoleRequest) ProtoMessage() {}

func (x *GenerateOwnerRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authz_authz_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateOwnerRoleRequest.ProtoReflect.Descriptor instead.
func (*GenerateOwnerRoleRequest) Descriptor() ([]byte, []int) {
	return file_authz_authz_proto_rawDescGZIP(), []int{8}
}

func (x *GenerateOwnerRoleRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GenerateOwnerRoleRequest) GetShopId() int64 {
	if x != nil {
		return x.ShopId
	}
	return 0
}

type GenerateOwnerRoleResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *GenerateOwnerRoleResult) Reset() {
	*x = GenerateOwnerRoleResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authz_authz_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateOwnerRoleResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateOwnerRoleResult) ProtoMessage() {}

func (x *GenerateOwnerRoleResult) ProtoReflect() protoreflect.Message {
	mi := &file_authz_authz_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateOwnerRoleResult.ProtoReflect.Descriptor instead.
func (*GenerateOwnerRoleResult) Descriptor() ([]byte, []int) {
	return file_authz_authz_proto_rawDescGZIP(), []int{9}
}

func (x *GenerateOwnerRoleResult) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

var File_authz_authz_proto protoreflect.FileDescriptor

var file_authz_authz_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x22, 0x5c, 0x0a, 0x16, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x73, 0x68, 0x6f, 0x70, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x63, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x63, 0x74, 0x22, 0x43, 0x0a, 0x15, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x5c, 0x0a,
	0x16, 0x41, 0x64, 0x64, 0x52, 0x6f, 0x6c, 0x65, 0x54, 0x6f, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x73, 0x68, 0x6f, 0x70, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x63, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x63, 0x74, 0x22, 0x43, 0x0a, 0x15, 0x41,
	0x64, 0x64, 0x52, 0x6f, 0x6c, 0x65, 0x54, 0x6f, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x22, 0x4b, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x49, 0x6e, 0x44, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x68, 0x6f, 0x70, 0x49, 0x64, 0x22, 0x42, 0x0a,
	0x16, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x49, 0x6e, 0x44, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x22, 0x53, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x70, 0x6c, 0x69, 0x63, 0x69, 0x74,
	0x52, 0x6f, 0x6c, 0x65, 0x73, 0x49, 0x6e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x73, 0x68, 0x6f, 0x70, 0x49, 0x64, 0x22, 0x4a, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x70,
	0x6c, 0x69, 0x63, 0x69, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x49, 0x6e, 0x44, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x22, 0x4c, 0x0a, 0x18, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4f, 0x77,
	0x6e, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x68, 0x6f, 0x70, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x68, 0x6f, 0x70, 0x49, 0x64,
	0x22, 0x2d, 0x0a, 0x17, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4f, 0x77, 0x6e, 0x65,
	0x72, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x32,
	0xc8, 0x03, 0x0a, 0x08, 0x41, 0x75, 0x74, 0x68, 0x7a, 0x52, 0x50, 0x43, 0x12, 0x50, 0x0a, 0x0f,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x1d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x50,
	0x0a, 0x0f, 0x41, 0x64, 0x64, 0x52, 0x6f, 0x6c, 0x65, 0x54, 0x6f, 0x44, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x12, 0x1d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x6f, 0x6c,
	0x65, 0x54, 0x6f, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x6f, 0x6c, 0x65,
	0x54, 0x6f, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00,
	0x12, 0x53, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x49, 0x6e, 0x44, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x12, 0x1e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x47, 0x65, 0x74,
	0x52, 0x6f, 0x6c, 0x65, 0x73, 0x49, 0x6e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x47, 0x65, 0x74,
	0x52, 0x6f, 0x6c, 0x65, 0x73, 0x49, 0x6e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x6b, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x70, 0x6c,
	0x69, 0x63, 0x69, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x49, 0x6e, 0x44, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x12, 0x26, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x70,
	0x6c, 0x69, 0x63, 0x69, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x49, 0x6e, 0x44, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x7a, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x70, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x52, 0x6f, 0x6c,
	0x65, 0x73, 0x49, 0x6e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x22, 0x00, 0x12, 0x56, 0x0a, 0x11, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4f, 0x77,
	0x6e, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x1f, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x52, 0x6f, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a,
	0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x52, 0x6f,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x42, 0x1d, 0x5a, 0x1b, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x73, 0x2d, 0x68, 0x73, 0x2f, 0x65,
	0x72, 0x70, 0x63, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_authz_authz_proto_rawDescOnce sync.Once
	file_authz_authz_proto_rawDescData = file_authz_authz_proto_rawDesc
)

func file_authz_authz_proto_rawDescGZIP() []byte {
	file_authz_authz_proto_rawDescOnce.Do(func() {
		file_authz_authz_proto_rawDescData = protoimpl.X.CompressGZIP(file_authz_authz_proto_rawDescData)
	})
	return file_authz_authz_proto_rawDescData
}

var file_authz_authz_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_authz_authz_proto_goTypes = []interface{}{
	(*CheckPermissionRequest)(nil),          // 0: authz.CheckPermissionRequest
	(*CheckPermissionResult)(nil),           // 1: authz.CheckPermissionResult
	(*AddRoleToDomainRequest)(nil),          // 2: authz.AddRoleToDomainRequest
	(*AddRoleToDomainResult)(nil),           // 3: authz.AddRoleToDomainResult
	(*GetRolesInDomainRequest)(nil),         // 4: authz.GetRolesInDomainRequest
	(*GetRolesInDomainResult)(nil),          // 5: authz.GetRolesInDomainResult
	(*GetImplicitRolesInDomainRequest)(nil), // 6: authz.GetImplicitRolesInDomainRequest
	(*GetImplicitRolesInDomainResult)(nil),  // 7: authz.GetImplicitRolesInDomainResult
	(*GenerateOwnerRoleRequest)(nil),        // 8: authz.GenerateOwnerRoleRequest
	(*GenerateOwnerRoleResult)(nil),         // 9: authz.GenerateOwnerRoleResult
}
var file_authz_authz_proto_depIdxs = []int32{
	0, // 0: authz.AuthzRPC.CheckPermission:input_type -> authz.CheckPermissionRequest
	2, // 1: authz.AuthzRPC.AddRoleToDomain:input_type -> authz.AddRoleToDomainRequest
	4, // 2: authz.AuthzRPC.GetRolesInDomain:input_type -> authz.GetRolesInDomainRequest
	6, // 3: authz.AuthzRPC.GetImplicitRolesInDomain:input_type -> authz.GetImplicitRolesInDomainRequest
	8, // 4: authz.AuthzRPC.GenerateOwnerRole:input_type -> authz.GenerateOwnerRoleRequest
	1, // 5: authz.AuthzRPC.CheckPermission:output_type -> authz.CheckPermissionResult
	3, // 6: authz.AuthzRPC.AddRoleToDomain:output_type -> authz.AddRoleToDomainResult
	5, // 7: authz.AuthzRPC.GetRolesInDomain:output_type -> authz.GetRolesInDomainResult
	7, // 8: authz.AuthzRPC.GetImplicitRolesInDomain:output_type -> authz.GetImplicitRolesInDomainResult
	9, // 9: authz.AuthzRPC.GenerateOwnerRole:output_type -> authz.GenerateOwnerRoleResult
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_authz_authz_proto_init() }
func file_authz_authz_proto_init() {
	if File_authz_authz_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_authz_authz_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckPermissionRequest); i {
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
		file_authz_authz_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckPermissionResult); i {
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
		file_authz_authz_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRoleToDomainRequest); i {
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
		file_authz_authz_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRoleToDomainResult); i {
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
		file_authz_authz_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRolesInDomainRequest); i {
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
		file_authz_authz_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRolesInDomainResult); i {
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
		file_authz_authz_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetImplicitRolesInDomainRequest); i {
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
		file_authz_authz_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetImplicitRolesInDomainResult); i {
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
		file_authz_authz_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateOwnerRoleRequest); i {
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
		file_authz_authz_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateOwnerRoleResult); i {
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
			RawDescriptor: file_authz_authz_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_authz_authz_proto_goTypes,
		DependencyIndexes: file_authz_authz_proto_depIdxs,
		MessageInfos:      file_authz_authz_proto_msgTypes,
	}.Build()
	File_authz_authz_proto = out.File
	file_authz_authz_proto_rawDesc = nil
	file_authz_authz_proto_goTypes = nil
	file_authz_authz_proto_depIdxs = nil
}
