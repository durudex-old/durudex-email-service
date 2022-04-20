// Copyright © 2021-2022 Durudex

// This file is part of Durudex: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.

// Durudex is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.

// You should have received a copy of the GNU Affero General Public License
// along with Durudex. If not, see <https://www.gnu.org/licenses/>.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: internal/delivery/grpc/pb/email.proto

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

type SendEmailUserCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Code     uint64 `protobuf:"varint,3,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *SendEmailUserCodeRequest) Reset() {
	*x = SendEmailUserCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_delivery_grpc_pb_email_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendEmailUserCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendEmailUserCodeRequest) ProtoMessage() {}

func (x *SendEmailUserCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_delivery_grpc_pb_email_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendEmailUserCodeRequest.ProtoReflect.Descriptor instead.
func (*SendEmailUserCodeRequest) Descriptor() ([]byte, []int) {
	return file_internal_delivery_grpc_pb_email_proto_rawDescGZIP(), []int{0}
}

func (x *SendEmailUserCodeRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SendEmailUserCodeRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *SendEmailUserCodeRequest) GetCode() uint64 {
	if x != nil {
		return x.Code
	}
	return 0
}

type SendEmailUserCodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendEmailUserCodeResponse) Reset() {
	*x = SendEmailUserCodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_delivery_grpc_pb_email_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendEmailUserCodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendEmailUserCodeResponse) ProtoMessage() {}

func (x *SendEmailUserCodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_delivery_grpc_pb_email_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendEmailUserCodeResponse.ProtoReflect.Descriptor instead.
func (*SendEmailUserCodeResponse) Descriptor() ([]byte, []int) {
	return file_internal_delivery_grpc_pb_email_proto_rawDescGZIP(), []int{1}
}

type SendEmailUserLoggedInRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Ip    string `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
}

func (x *SendEmailUserLoggedInRequest) Reset() {
	*x = SendEmailUserLoggedInRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_delivery_grpc_pb_email_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendEmailUserLoggedInRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendEmailUserLoggedInRequest) ProtoMessage() {}

func (x *SendEmailUserLoggedInRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_delivery_grpc_pb_email_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendEmailUserLoggedInRequest.ProtoReflect.Descriptor instead.
func (*SendEmailUserLoggedInRequest) Descriptor() ([]byte, []int) {
	return file_internal_delivery_grpc_pb_email_proto_rawDescGZIP(), []int{2}
}

func (x *SendEmailUserLoggedInRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SendEmailUserLoggedInRequest) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

type SendEmailUserLoggedInResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendEmailUserLoggedInResponse) Reset() {
	*x = SendEmailUserLoggedInResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_delivery_grpc_pb_email_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendEmailUserLoggedInResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendEmailUserLoggedInResponse) ProtoMessage() {}

func (x *SendEmailUserLoggedInResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_delivery_grpc_pb_email_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendEmailUserLoggedInResponse.ProtoReflect.Descriptor instead.
func (*SendEmailUserLoggedInResponse) Descriptor() ([]byte, []int) {
	return file_internal_delivery_grpc_pb_email_proto_rawDescGZIP(), []int{3}
}

type SendEmailUserRegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *SendEmailUserRegisterRequest) Reset() {
	*x = SendEmailUserRegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_delivery_grpc_pb_email_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendEmailUserRegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendEmailUserRegisterRequest) ProtoMessage() {}

func (x *SendEmailUserRegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_delivery_grpc_pb_email_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendEmailUserRegisterRequest.ProtoReflect.Descriptor instead.
func (*SendEmailUserRegisterRequest) Descriptor() ([]byte, []int) {
	return file_internal_delivery_grpc_pb_email_proto_rawDescGZIP(), []int{4}
}

func (x *SendEmailUserRegisterRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SendEmailUserRegisterRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type SendEmailUserRegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendEmailUserRegisterResponse) Reset() {
	*x = SendEmailUserRegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_delivery_grpc_pb_email_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendEmailUserRegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendEmailUserRegisterResponse) ProtoMessage() {}

func (x *SendEmailUserRegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_delivery_grpc_pb_email_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendEmailUserRegisterResponse.ProtoReflect.Descriptor instead.
func (*SendEmailUserRegisterResponse) Descriptor() ([]byte, []int) {
	return file_internal_delivery_grpc_pb_email_proto_rawDescGZIP(), []int{5}
}

var File_internal_delivery_grpc_pb_email_proto protoreflect.FileDescriptor

var file_internal_delivery_grpc_pb_email_proto_rawDesc = []byte{
	0x0a, 0x25, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x64, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x79, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x2f, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x64, 0x75, 0x72, 0x75, 0x64, 0x65, 0x78,
	0x22, 0x60, 0x0a, 0x18, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x55, 0x73, 0x65,
	0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x22, 0x1b, 0x0a, 0x19, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x55,
	0x73, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x44, 0x0a, 0x1c, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x55, 0x73, 0x65, 0x72,
	0x4c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x70, 0x22, 0x1f, 0x0a, 0x1d, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x49, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x50, 0x0a, 0x1c, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x1f, 0x0a, 0x1d, 0x53, 0x65, 0x6e, 0x64,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xba, 0x02, 0x0a, 0x0c, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5a, 0x0a, 0x11, 0x53, 0x65,
	0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x21, 0x2e, 0x64, 0x75, 0x72, 0x75, 0x64, 0x65, 0x78, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x22, 0x2e, 0x64, 0x75, 0x72, 0x75, 0x64, 0x65, 0x78, 0x2e, 0x53, 0x65, 0x6e,
	0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x66, 0x0a, 0x15, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x49, 0x6e, 0x12,
	0x25, 0x2e, 0x64, 0x75, 0x72, 0x75, 0x64, 0x65, 0x78, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x49, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x64, 0x75, 0x72, 0x75, 0x64, 0x65, 0x78,
	0x2e, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f,
	0x67, 0x67, 0x65, 0x64, 0x49, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x66,
	0x0a, 0x15, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x25, 0x2e, 0x64, 0x75, 0x72, 0x75, 0x64, 0x65,
	0x78, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26,
	0x2e, 0x64, 0x75, 0x72, 0x75, 0x64, 0x65, 0x78, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x47, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x75, 0x72, 0x75, 0x64, 0x65, 0x78, 0x2f, 0x64, 0x75, 0x72,
	0x75, 0x64, 0x65, 0x78, 0x2d, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x64, 0x65, 0x6c, 0x69,
	0x76, 0x65, 0x72, 0x79, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x3b, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_delivery_grpc_pb_email_proto_rawDescOnce sync.Once
	file_internal_delivery_grpc_pb_email_proto_rawDescData = file_internal_delivery_grpc_pb_email_proto_rawDesc
)

func file_internal_delivery_grpc_pb_email_proto_rawDescGZIP() []byte {
	file_internal_delivery_grpc_pb_email_proto_rawDescOnce.Do(func() {
		file_internal_delivery_grpc_pb_email_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_delivery_grpc_pb_email_proto_rawDescData)
	})
	return file_internal_delivery_grpc_pb_email_proto_rawDescData
}

var file_internal_delivery_grpc_pb_email_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_internal_delivery_grpc_pb_email_proto_goTypes = []interface{}{
	(*SendEmailUserCodeRequest)(nil),      // 0: durudex.SendEmailUserCodeRequest
	(*SendEmailUserCodeResponse)(nil),     // 1: durudex.SendEmailUserCodeResponse
	(*SendEmailUserLoggedInRequest)(nil),  // 2: durudex.SendEmailUserLoggedInRequest
	(*SendEmailUserLoggedInResponse)(nil), // 3: durudex.SendEmailUserLoggedInResponse
	(*SendEmailUserRegisterRequest)(nil),  // 4: durudex.SendEmailUserRegisterRequest
	(*SendEmailUserRegisterResponse)(nil), // 5: durudex.SendEmailUserRegisterResponse
}
var file_internal_delivery_grpc_pb_email_proto_depIdxs = []int32{
	0, // 0: durudex.EmailService.SendEmailUserCode:input_type -> durudex.SendEmailUserCodeRequest
	2, // 1: durudex.EmailService.SendEmailUserLoggedIn:input_type -> durudex.SendEmailUserLoggedInRequest
	4, // 2: durudex.EmailService.SendEmailUserRegister:input_type -> durudex.SendEmailUserRegisterRequest
	1, // 3: durudex.EmailService.SendEmailUserCode:output_type -> durudex.SendEmailUserCodeResponse
	3, // 4: durudex.EmailService.SendEmailUserLoggedIn:output_type -> durudex.SendEmailUserLoggedInResponse
	5, // 5: durudex.EmailService.SendEmailUserRegister:output_type -> durudex.SendEmailUserRegisterResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_delivery_grpc_pb_email_proto_init() }
func file_internal_delivery_grpc_pb_email_proto_init() {
	if File_internal_delivery_grpc_pb_email_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_delivery_grpc_pb_email_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendEmailUserCodeRequest); i {
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
		file_internal_delivery_grpc_pb_email_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendEmailUserCodeResponse); i {
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
		file_internal_delivery_grpc_pb_email_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendEmailUserLoggedInRequest); i {
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
		file_internal_delivery_grpc_pb_email_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendEmailUserLoggedInResponse); i {
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
		file_internal_delivery_grpc_pb_email_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendEmailUserRegisterRequest); i {
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
		file_internal_delivery_grpc_pb_email_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendEmailUserRegisterResponse); i {
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
			RawDescriptor: file_internal_delivery_grpc_pb_email_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_delivery_grpc_pb_email_proto_goTypes,
		DependencyIndexes: file_internal_delivery_grpc_pb_email_proto_depIdxs,
		MessageInfos:      file_internal_delivery_grpc_pb_email_proto_msgTypes,
	}.Build()
	File_internal_delivery_grpc_pb_email_proto = out.File
	file_internal_delivery_grpc_pb_email_proto_rawDesc = nil
	file_internal_delivery_grpc_pb_email_proto_goTypes = nil
	file_internal_delivery_grpc_pb_email_proto_depIdxs = nil
}
