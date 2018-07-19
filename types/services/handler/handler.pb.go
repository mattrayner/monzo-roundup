// Code generated by protoc-gen-go. DO NOT EDIT.
// source: types/services/handler.proto

package handler // import "github.com/mattrayner/monzo-roundup/types/services/handler"

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

type HandleInput struct {
	AccountID            string   `protobuf:"bytes,1,opt,name=accountID,proto3" json:"accountID,omitempty"`
	TransactionID        string   `protobuf:"bytes,2,opt,name=transactionID,proto3" json:"transactionID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HandleInput) Reset()         { *m = HandleInput{} }
func (m *HandleInput) String() string { return proto.CompactTextString(m) }
func (*HandleInput) ProtoMessage()    {}
func (*HandleInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_handler_dc972465e0dcc7ab, []int{0}
}
func (m *HandleInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HandleInput.Unmarshal(m, b)
}
func (m *HandleInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HandleInput.Marshal(b, m, deterministic)
}
func (dst *HandleInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HandleInput.Merge(dst, src)
}
func (m *HandleInput) XXX_Size() int {
	return xxx_messageInfo_HandleInput.Size(m)
}
func (m *HandleInput) XXX_DiscardUnknown() {
	xxx_messageInfo_HandleInput.DiscardUnknown(m)
}

var xxx_messageInfo_HandleInput proto.InternalMessageInfo

func (m *HandleInput) GetAccountID() string {
	if m != nil {
		return m.AccountID
	}
	return ""
}

func (m *HandleInput) GetTransactionID() string {
	if m != nil {
		return m.TransactionID
	}
	return ""
}

type HandleOutput struct {
	Remainder            int32    `protobuf:"varint,1,opt,name=remainder,proto3" json:"remainder,omitempty"`
	Balance              int32    `protobuf:"varint,2,opt,name=balance,proto3" json:"balance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HandleOutput) Reset()         { *m = HandleOutput{} }
func (m *HandleOutput) String() string { return proto.CompactTextString(m) }
func (*HandleOutput) ProtoMessage()    {}
func (*HandleOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_handler_dc972465e0dcc7ab, []int{1}
}
func (m *HandleOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HandleOutput.Unmarshal(m, b)
}
func (m *HandleOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HandleOutput.Marshal(b, m, deterministic)
}
func (dst *HandleOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HandleOutput.Merge(dst, src)
}
func (m *HandleOutput) XXX_Size() int {
	return xxx_messageInfo_HandleOutput.Size(m)
}
func (m *HandleOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_HandleOutput.DiscardUnknown(m)
}

var xxx_messageInfo_HandleOutput proto.InternalMessageInfo

func (m *HandleOutput) GetRemainder() int32 {
	if m != nil {
		return m.Remainder
	}
	return 0
}

func (m *HandleOutput) GetBalance() int32 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func init() {
	proto.RegisterType((*HandleInput)(nil), "handler.HandleInput")
	proto.RegisterType((*HandleOutput)(nil), "handler.HandleOutput")
}

func init() {
	proto.RegisterFile("types/services/handler.proto", fileDescriptor_handler_dc972465e0dcc7ab)
}

var fileDescriptor_handler_dc972465e0dcc7ab = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0xa9, 0xd0, 0x86, 0xae, 0x7a, 0x59, 0x14, 0x82, 0xf4, 0x20, 0xc5, 0x83, 0x17, 0xb3,
	0xa0, 0x78, 0x11, 0x0f, 0x22, 0x45, 0xcc, 0x49, 0xcc, 0xd1, 0xdb, 0x64, 0x33, 0xd8, 0x40, 0x33,
	0x13, 0x26, 0xb3, 0x42, 0xfd, 0xf5, 0xc2, 0x6e, 0x4b, 0x54, 0x7a, 0x7b, 0xef, 0x63, 0xf9, 0x76,
	0xf7, 0x99, 0x85, 0x6e, 0x7b, 0x1c, 0xdc, 0x80, 0xf2, 0xd5, 0x7a, 0x1c, 0xdc, 0x1a, 0xa8, 0xd9,
	0xa0, 0x14, 0xbd, 0xb0, 0xb2, 0xcd, 0x76, 0x75, 0xf9, 0x6e, 0x8e, 0x5f, 0x63, 0x2c, 0xa9, 0x0f,
	0x6a, 0x17, 0x66, 0x0e, 0xde, 0x73, 0x20, 0x2d, 0x57, 0xf9, 0xe4, 0x72, 0x72, 0x3d, 0xaf, 0x46,
	0x60, 0xaf, 0xcc, 0xa9, 0x0a, 0xd0, 0x00, 0x5e, 0x5b, 0xa6, 0x72, 0x95, 0x1f, 0xc5, 0x13, 0x7f,
	0xe1, 0xf2, 0xc5, 0x9c, 0x24, 0xe5, 0x5b, 0xd0, 0x9d, 0x53, 0xb0, 0x83, 0x96, 0x1a, 0x94, 0xe8,
	0x9c, 0x56, 0x23, 0xb0, 0xb9, 0xc9, 0x6a, 0xd8, 0x00, 0x79, 0x8c, 0xb6, 0x69, 0xb5, 0xaf, 0xb7,
	0x4f, 0x26, 0x4b, 0x1e, 0xb1, 0xf7, 0x66, 0x96, 0xa2, 0x3d, 0x2b, 0xf6, 0x1f, 0xf9, 0xf5, 0xec,
	0x8b, 0xf3, 0x7f, 0x34, 0xdd, 0xfc, 0xfc, 0xf8, 0xf1, 0xf0, 0xd9, 0xea, 0x3a, 0xd4, 0x85, 0xe7,
	0xce, 0x75, 0xa0, 0x2a, 0xb0, 0x25, 0x14, 0xd7, 0x31, 0x7d, 0xf3, 0x8d, 0x70, 0xa0, 0x26, 0xf4,
	0xee, 0xf0, 0x52, 0xf5, 0x2c, 0x4e, 0x75, 0xf7, 0x13, 0x00, 0x00, 0xff, 0xff, 0xf2, 0x06, 0xe9,
	0x35, 0x4a, 0x01, 0x00, 0x00,
}
