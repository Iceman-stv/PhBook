// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: PhBook.proto

package gen

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

type RegisterUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterUserRequest) Reset() {
	*x = RegisterUserRequest{}
	mi := &file_PhBook_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterUserRequest) ProtoMessage() {}

func (x *RegisterUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_PhBook_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterUserRequest.ProtoReflect.Descriptor instead.
func (*RegisterUserRequest) Descriptor() ([]byte, []int) {
	return file_PhBook_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterUserRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *RegisterUserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type RegisterUserResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterUserResponse) Reset() {
	*x = RegisterUserResponse{}
	mi := &file_PhBook_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterUserResponse) ProtoMessage() {}

func (x *RegisterUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_PhBook_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterUserResponse.ProtoReflect.Descriptor instead.
func (*RegisterUserResponse) Descriptor() ([]byte, []int) {
	return file_PhBook_proto_rawDescGZIP(), []int{1}
}

type AuthUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuthUserRequest) Reset() {
	*x = AuthUserRequest{}
	mi := &file_PhBook_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthUserRequest) ProtoMessage() {}

func (x *AuthUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_PhBook_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthUserRequest.ProtoReflect.Descriptor instead.
func (*AuthUserRequest) Descriptor() ([]byte, []int) {
	return file_PhBook_proto_rawDescGZIP(), []int{2}
}

func (x *AuthUserRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AuthUserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type AuthUserResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int32                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Token         string                 `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuthUserResponse) Reset() {
	*x = AuthUserResponse{}
	mi := &file_PhBook_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthUserResponse) ProtoMessage() {}

func (x *AuthUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_PhBook_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthUserResponse.ProtoReflect.Descriptor instead.
func (*AuthUserResponse) Descriptor() ([]byte, []int) {
	return file_PhBook_proto_rawDescGZIP(), []int{3}
}

func (x *AuthUserResponse) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *AuthUserResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type AddContactRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int32                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Phone         string                 `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddContactRequest) Reset() {
	*x = AddContactRequest{}
	mi := &file_PhBook_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddContactRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddContactRequest) ProtoMessage() {}

func (x *AddContactRequest) ProtoReflect() protoreflect.Message {
	mi := &file_PhBook_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddContactRequest.ProtoReflect.Descriptor instead.
func (*AddContactRequest) Descriptor() ([]byte, []int) {
	return file_PhBook_proto_rawDescGZIP(), []int{4}
}

func (x *AddContactRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *AddContactRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddContactRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type AddContactResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddContactResponse) Reset() {
	*x = AddContactResponse{}
	mi := &file_PhBook_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddContactResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddContactResponse) ProtoMessage() {}

func (x *AddContactResponse) ProtoReflect() protoreflect.Message {
	mi := &file_PhBook_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddContactResponse.ProtoReflect.Descriptor instead.
func (*AddContactResponse) Descriptor() ([]byte, []int) {
	return file_PhBook_proto_rawDescGZIP(), []int{5}
}

type DelContactRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int32                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DelContactRequest) Reset() {
	*x = DelContactRequest{}
	mi := &file_PhBook_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DelContactRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelContactRequest) ProtoMessage() {}

func (x *DelContactRequest) ProtoReflect() protoreflect.Message {
	mi := &file_PhBook_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelContactRequest.ProtoReflect.Descriptor instead.
func (*DelContactRequest) Descriptor() ([]byte, []int) {
	return file_PhBook_proto_rawDescGZIP(), []int{6}
}

func (x *DelContactRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *DelContactRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DelContactResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DelContactResponse) Reset() {
	*x = DelContactResponse{}
	mi := &file_PhBook_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DelContactResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelContactResponse) ProtoMessage() {}

func (x *DelContactResponse) ProtoReflect() protoreflect.Message {
	mi := &file_PhBook_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelContactResponse.ProtoReflect.Descriptor instead.
func (*DelContactResponse) Descriptor() ([]byte, []int) {
	return file_PhBook_proto_rawDescGZIP(), []int{7}
}

type FindContactRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int32                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FindContactRequest) Reset() {
	*x = FindContactRequest{}
	mi := &file_PhBook_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FindContactRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindContactRequest) ProtoMessage() {}

func (x *FindContactRequest) ProtoReflect() protoreflect.Message {
	mi := &file_PhBook_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindContactRequest.ProtoReflect.Descriptor instead.
func (*FindContactRequest) Descriptor() ([]byte, []int) {
	return file_PhBook_proto_rawDescGZIP(), []int{8}
}

func (x *FindContactRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *FindContactRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type FindContactResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Contacts      []*Contact             `protobuf:"bytes,1,rep,name=contacts,proto3" json:"contacts,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FindContactResponse) Reset() {
	*x = FindContactResponse{}
	mi := &file_PhBook_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FindContactResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindContactResponse) ProtoMessage() {}

func (x *FindContactResponse) ProtoReflect() protoreflect.Message {
	mi := &file_PhBook_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindContactResponse.ProtoReflect.Descriptor instead.
func (*FindContactResponse) Descriptor() ([]byte, []int) {
	return file_PhBook_proto_rawDescGZIP(), []int{9}
}

func (x *FindContactResponse) GetContacts() []*Contact {
	if x != nil {
		return x.Contacts
	}
	return nil
}

type GetContactsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int32                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetContactsRequest) Reset() {
	*x = GetContactsRequest{}
	mi := &file_PhBook_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetContactsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetContactsRequest) ProtoMessage() {}

func (x *GetContactsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_PhBook_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetContactsRequest.ProtoReflect.Descriptor instead.
func (*GetContactsRequest) Descriptor() ([]byte, []int) {
	return file_PhBook_proto_rawDescGZIP(), []int{10}
}

func (x *GetContactsRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetContactsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Contacts      []*Contact             `protobuf:"bytes,1,rep,name=contacts,proto3" json:"contacts,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetContactsResponse) Reset() {
	*x = GetContactsResponse{}
	mi := &file_PhBook_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetContactsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetContactsResponse) ProtoMessage() {}

func (x *GetContactsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_PhBook_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetContactsResponse.ProtoReflect.Descriptor instead.
func (*GetContactsResponse) Descriptor() ([]byte, []int) {
	return file_PhBook_proto_rawDescGZIP(), []int{11}
}

func (x *GetContactsResponse) GetContacts() []*Contact {
	if x != nil {
		return x.Contacts
	}
	return nil
}

type Contact struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Phone         string                 `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Contact) Reset() {
	*x = Contact{}
	mi := &file_PhBook_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Contact) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Contact) ProtoMessage() {}

func (x *Contact) ProtoReflect() protoreflect.Message {
	mi := &file_PhBook_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Contact.ProtoReflect.Descriptor instead.
func (*Contact) Descriptor() ([]byte, []int) {
	return file_PhBook_proto_rawDescGZIP(), []int{12}
}

func (x *Contact) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Contact) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

var File_PhBook_proto protoreflect.FileDescriptor

var file_PhBook_proto_rawDesc = string([]byte{
	0x0a, 0x0c, 0x50, 0x68, 0x42, 0x6f, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x50, 0x68, 0x42, 0x6f, 0x6f, 0x6b, 0x22, 0x4d, 0x0a, 0x13, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x16, 0x0a, 0x14, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x49, 0x0a,
	0x0f, 0x41, 0x75, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x41, 0x0a, 0x10, 0x41, 0x75, 0x74, 0x68,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x56, 0x0a, 0x11, 0x41,
	0x64, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x40, 0x0a, 0x11, 0x44, 0x65, 0x6c,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x44,
	0x65, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x41, 0x0a, 0x12, 0x46, 0x69, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x42, 0x0a, 0x13, 0x46, 0x69, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x08, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x50, 0x68, 0x42, 0x6f, 0x6f, 0x6b, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x08,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x22, 0x2d, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x42, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b,
	0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x50, 0x68, 0x42, 0x6f, 0x6f, 0x6b, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x22, 0x33, 0x0a, 0x07, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x32, 0xb6, 0x03, 0x0a, 0x10, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x50, 0x68, 0x42, 0x6f, 0x6f, 0x6b, 0x2e, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x50, 0x68, 0x42, 0x6f, 0x6f, 0x6b, 0x2e, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3d, 0x0a, 0x08, 0x41, 0x75, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x50,
	0x68, 0x42, 0x6f, 0x6f, 0x6b, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x50, 0x68, 0x42, 0x6f, 0x6f, 0x6b, 0x2e, 0x41,
	0x75, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x43, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x19, 0x2e,
	0x50, 0x68, 0x42, 0x6f, 0x6f, 0x6b, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x50, 0x68, 0x42, 0x6f, 0x6f,
	0x6b, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x12, 0x19, 0x2e, 0x50, 0x68, 0x42, 0x6f, 0x6f, 0x6b, 0x2e, 0x44, 0x65, 0x6c, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x50, 0x68, 0x42, 0x6f, 0x6f, 0x6b, 0x2e, 0x44, 0x65, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x0b, 0x46, 0x69, 0x6e,
	0x64, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x1a, 0x2e, 0x50, 0x68, 0x42, 0x6f, 0x6f,
	0x6b, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x50, 0x68, 0x42, 0x6f, 0x6f, 0x6b, 0x2e, 0x46, 0x69,
	0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x46, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73,
	0x12, 0x1a, 0x2e, 0x50, 0x68, 0x42, 0x6f, 0x6f, 0x6b, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e,
	0x74, 0x61, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x50,
	0x68, 0x42, 0x6f, 0x6f, 0x6b, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x49, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x2d, 0x73,
	0x74, 0x76, 0x2f, 0x50, 0x68, 0x42, 0x6f, 0x6f, 0x6b, 0x2f, 0x67, 0x65, 0x6e, 0x3b, 0x67, 0x65,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_PhBook_proto_rawDescOnce sync.Once
	file_PhBook_proto_rawDescData []byte
)

func file_PhBook_proto_rawDescGZIP() []byte {
	file_PhBook_proto_rawDescOnce.Do(func() {
		file_PhBook_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_PhBook_proto_rawDesc), len(file_PhBook_proto_rawDesc)))
	})
	return file_PhBook_proto_rawDescData
}

var file_PhBook_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_PhBook_proto_goTypes = []any{
	(*RegisterUserRequest)(nil),  // 0: PhBook.RegisterUserRequest
	(*RegisterUserResponse)(nil), // 1: PhBook.RegisterUserResponse
	(*AuthUserRequest)(nil),      // 2: PhBook.AuthUserRequest
	(*AuthUserResponse)(nil),     // 3: PhBook.AuthUserResponse
	(*AddContactRequest)(nil),    // 4: PhBook.AddContactRequest
	(*AddContactResponse)(nil),   // 5: PhBook.AddContactResponse
	(*DelContactRequest)(nil),    // 6: PhBook.DelContactRequest
	(*DelContactResponse)(nil),   // 7: PhBook.DelContactResponse
	(*FindContactRequest)(nil),   // 8: PhBook.FindContactRequest
	(*FindContactResponse)(nil),  // 9: PhBook.FindContactResponse
	(*GetContactsRequest)(nil),   // 10: PhBook.GetContactsRequest
	(*GetContactsResponse)(nil),  // 11: PhBook.GetContactsResponse
	(*Contact)(nil),              // 12: PhBook.Contact
}
var file_PhBook_proto_depIdxs = []int32{
	12, // 0: PhBook.FindContactResponse.contacts:type_name -> PhBook.Contact
	12, // 1: PhBook.GetContactsResponse.contacts:type_name -> PhBook.Contact
	0,  // 2: PhBook.PhoneBookService.RegisterUser:input_type -> PhBook.RegisterUserRequest
	2,  // 3: PhBook.PhoneBookService.AuthUser:input_type -> PhBook.AuthUserRequest
	4,  // 4: PhBook.PhoneBookService.AddContact:input_type -> PhBook.AddContactRequest
	6,  // 5: PhBook.PhoneBookService.DelContact:input_type -> PhBook.DelContactRequest
	8,  // 6: PhBook.PhoneBookService.FindContact:input_type -> PhBook.FindContactRequest
	10, // 7: PhBook.PhoneBookService.GetContacts:input_type -> PhBook.GetContactsRequest
	1,  // 8: PhBook.PhoneBookService.RegisterUser:output_type -> PhBook.RegisterUserResponse
	3,  // 9: PhBook.PhoneBookService.AuthUser:output_type -> PhBook.AuthUserResponse
	5,  // 10: PhBook.PhoneBookService.AddContact:output_type -> PhBook.AddContactResponse
	7,  // 11: PhBook.PhoneBookService.DelContact:output_type -> PhBook.DelContactResponse
	9,  // 12: PhBook.PhoneBookService.FindContact:output_type -> PhBook.FindContactResponse
	11, // 13: PhBook.PhoneBookService.GetContacts:output_type -> PhBook.GetContactsResponse
	8,  // [8:14] is the sub-list for method output_type
	2,  // [2:8] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_PhBook_proto_init() }
func file_PhBook_proto_init() {
	if File_PhBook_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_PhBook_proto_rawDesc), len(file_PhBook_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_PhBook_proto_goTypes,
		DependencyIndexes: file_PhBook_proto_depIdxs,
		MessageInfos:      file_PhBook_proto_msgTypes,
	}.Build()
	File_PhBook_proto = out.File
	file_PhBook_proto_goTypes = nil
	file_PhBook_proto_depIdxs = nil
}
