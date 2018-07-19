// Code generated by protoc-gen-go. DO NOT EDIT.
// source: types/services/dynamodb.proto

package dynamodb // import "github.com/mattrayner/monzo-roundup/types/services/dynamodb"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetUserInput struct {
	AccountID            string   `protobuf:"bytes,1,opt,name=accountID,proto3" json:"accountID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserInput) Reset()         { *m = GetUserInput{} }
func (m *GetUserInput) String() string { return proto.CompactTextString(m) }
func (*GetUserInput) ProtoMessage()    {}
func (*GetUserInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_dynamodb_785027c5e4aef12e, []int{0}
}
func (m *GetUserInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserInput.Unmarshal(m, b)
}
func (m *GetUserInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserInput.Marshal(b, m, deterministic)
}
func (dst *GetUserInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserInput.Merge(dst, src)
}
func (m *GetUserInput) XXX_Size() int {
	return xxx_messageInfo_GetUserInput.Size(m)
}
func (m *GetUserInput) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserInput.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserInput proto.InternalMessageInfo

func (m *GetUserInput) GetAccountID() string {
	if m != nil {
		return m.AccountID
	}
	return ""
}

type GetUserOutput struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserOutput) Reset()         { *m = GetUserOutput{} }
func (m *GetUserOutput) String() string { return proto.CompactTextString(m) }
func (*GetUserOutput) ProtoMessage()    {}
func (*GetUserOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_dynamodb_785027c5e4aef12e, []int{1}
}
func (m *GetUserOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserOutput.Unmarshal(m, b)
}
func (m *GetUserOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserOutput.Marshal(b, m, deterministic)
}
func (dst *GetUserOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserOutput.Merge(dst, src)
}
func (m *GetUserOutput) XXX_Size() int {
	return xxx_messageInfo_GetUserOutput.Size(m)
}
func (m *GetUserOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserOutput.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserOutput proto.InternalMessageInfo

func (m *GetUserOutput) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type UpdateUserInput struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserInput) Reset()         { *m = UpdateUserInput{} }
func (m *UpdateUserInput) String() string { return proto.CompactTextString(m) }
func (*UpdateUserInput) ProtoMessage()    {}
func (*UpdateUserInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_dynamodb_785027c5e4aef12e, []int{2}
}
func (m *UpdateUserInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserInput.Unmarshal(m, b)
}
func (m *UpdateUserInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserInput.Marshal(b, m, deterministic)
}
func (dst *UpdateUserInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserInput.Merge(dst, src)
}
func (m *UpdateUserInput) XXX_Size() int {
	return xxx_messageInfo_UpdateUserInput.Size(m)
}
func (m *UpdateUserInput) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserInput.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserInput proto.InternalMessageInfo

func (m *UpdateUserInput) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type UpdateUserOutput struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserOutput) Reset()         { *m = UpdateUserOutput{} }
func (m *UpdateUserOutput) String() string { return proto.CompactTextString(m) }
func (*UpdateUserOutput) ProtoMessage()    {}
func (*UpdateUserOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_dynamodb_785027c5e4aef12e, []int{3}
}
func (m *UpdateUserOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserOutput.Unmarshal(m, b)
}
func (m *UpdateUserOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserOutput.Marshal(b, m, deterministic)
}
func (dst *UpdateUserOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserOutput.Merge(dst, src)
}
func (m *UpdateUserOutput) XXX_Size() int {
	return xxx_messageInfo_UpdateUserOutput.Size(m)
}
func (m *UpdateUserOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserOutput.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserOutput proto.InternalMessageInfo

func (m *UpdateUserOutput) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type User struct {
	AccountID            string   `protobuf:"bytes,1,opt,name=accountID,proto3" json:"accountID,omitempty"`
	AuthKey              string   `protobuf:"bytes,2,opt,name=authKey,proto3" json:"authKey,omitempty"`
	RefreshKey           string   `protobuf:"bytes,3,opt,name=refreshKey,proto3" json:"refreshKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_dynamodb_785027c5e4aef12e, []int{4}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetAccountID() string {
	if m != nil {
		return m.AccountID
	}
	return ""
}

func (m *User) GetAuthKey() string {
	if m != nil {
		return m.AuthKey
	}
	return ""
}

func (m *User) GetRefreshKey() string {
	if m != nil {
		return m.RefreshKey
	}
	return ""
}

func init() {
	proto.RegisterType((*GetUserInput)(nil), "dynamodb.GetUserInput")
	proto.RegisterType((*GetUserOutput)(nil), "dynamodb.GetUserOutput")
	proto.RegisterType((*UpdateUserInput)(nil), "dynamodb.UpdateUserInput")
	proto.RegisterType((*UpdateUserOutput)(nil), "dynamodb.UpdateUserOutput")
	proto.RegisterType((*User)(nil), "dynamodb.User")
}

func init() {
	proto.RegisterFile("types/services/dynamodb.proto", fileDescriptor_dynamodb_785027c5e4aef12e)
}

var fileDescriptor_dynamodb_785027c5e4aef12e = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x25, 0x5a, 0x6c, 0x3b, 0x7e, 0xb2, 0x07, 0x8d, 0x41, 0x45, 0x72, 0xf2, 0xa0, 0x09, 0xb4,
	0xe8, 0x41, 0xf1, 0x52, 0x03, 0x52, 0x3c, 0x08, 0x85, 0x5e, 0x3c, 0x08, 0x9b, 0x64, 0xb4, 0x3d,
	0x64, 0x37, 0xec, 0xce, 0x0a, 0xf1, 0x2f, 0xf8, 0xa7, 0x25, 0x6b, 0xea, 0x46, 0xfc, 0xc0, 0x1e,
	0xe7, 0xbd, 0x37, 0xfb, 0xf6, 0x3d, 0x06, 0x0e, 0xa9, 0x2a, 0x51, 0xc7, 0x1a, 0xd5, 0xcb, 0x3c,
	0x43, 0x1d, 0xe7, 0x95, 0xe0, 0x85, 0xcc, 0xd3, 0xa8, 0x54, 0x92, 0x24, 0xeb, 0x2d, 0xe6, 0xf0,
	0x14, 0x36, 0x6e, 0x91, 0xa6, 0x1a, 0xd5, 0x58, 0x94, 0x86, 0xd8, 0x01, 0xf4, 0x79, 0x96, 0x49,
	0x23, 0x68, 0x9c, 0xf8, 0xde, 0xb1, 0x77, 0xd2, 0x9f, 0x38, 0x20, 0x1c, 0xc2, 0x66, 0xa3, 0xbe,
	0x37, 0x54, 0xcb, 0x43, 0xe8, 0x18, 0x8d, 0xca, 0x2a, 0xd7, 0x07, 0x5b, 0xd1, 0xa7, 0x4f, 0xad,
	0x99, 0x58, 0x2e, 0x3c, 0x87, 0xed, 0x69, 0x99, 0x73, 0x42, 0xe7, 0xf2, 0x9f, 0xb5, 0x0b, 0xd8,
	0x71, 0x6b, 0x4b, 0xd8, 0x3d, 0x42, 0xa7, 0x9e, 0xfe, 0x4e, 0xc2, 0x7c, 0xe8, 0x72, 0x43, 0xb3,
	0x3b, 0xac, 0xfc, 0x15, 0xcb, 0x2d, 0x46, 0x76, 0x04, 0xa0, 0xf0, 0x49, 0xa1, 0xb6, 0xe4, 0xaa,
	0x25, 0x5b, 0xc8, 0xe0, 0xcd, 0x83, 0x5e, 0x62, 0x7d, 0x93, 0x11, 0xbb, 0x84, 0x6e, 0x53, 0x08,
	0xdb, 0x75, 0xbf, 0x69, 0x37, 0x1a, 0xec, 0x7d, 0xc3, 0x9b, 0x30, 0x37, 0x00, 0x2e, 0x20, 0xdb,
	0x6f, 0x85, 0xf9, 0xda, 0x56, 0x10, 0xfc, 0x44, 0x7d, 0x3c, 0x32, 0xba, 0x7e, 0xb8, 0x7a, 0x9e,
	0xd3, 0xcc, 0xa4, 0x51, 0x26, 0x8b, 0xb8, 0xe0, 0x44, 0x8a, 0x57, 0x02, 0x55, 0x5c, 0x48, 0xf1,
	0x2a, 0xcf, 0x94, 0x34, 0x22, 0x37, 0x65, 0xfc, 0xcb, 0x39, 0xa4, 0x6b, 0xf6, 0x1e, 0x86, 0xef,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x26, 0x1d, 0x9e, 0x0d, 0x30, 0x02, 0x00, 0x00,
}
