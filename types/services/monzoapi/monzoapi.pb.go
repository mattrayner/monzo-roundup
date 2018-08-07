// Code generated by protoc-gen-go. DO NOT EDIT.
// source: types/services/monzoapi.proto

package monzoapi // import "github.com/mattrayner/monzo-roundup/types/services/monzoapi"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/mattrayner/monzo-roundup/types/services/dynamodb"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RefreshTokenInput struct {
	RefreshToken         string   `protobuf:"bytes,1,opt,name=refreshToken,proto3" json:"refreshToken,omitempty"`
	MonzoClient          *Client  `protobuf:"bytes,2,opt,name=monzoClient,proto3" json:"monzoClient,omitempty"`
	AuthKey              string   `protobuf:"bytes,3,opt,name=authKey,proto3" json:"authKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RefreshTokenInput) Reset()         { *m = RefreshTokenInput{} }
func (m *RefreshTokenInput) String() string { return proto.CompactTextString(m) }
func (*RefreshTokenInput) ProtoMessage()    {}
func (*RefreshTokenInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_monzoapi_8777985c2df56e76, []int{0}
}
func (m *RefreshTokenInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RefreshTokenInput.Unmarshal(m, b)
}
func (m *RefreshTokenInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RefreshTokenInput.Marshal(b, m, deterministic)
}
func (dst *RefreshTokenInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RefreshTokenInput.Merge(dst, src)
}
func (m *RefreshTokenInput) XXX_Size() int {
	return xxx_messageInfo_RefreshTokenInput.Size(m)
}
func (m *RefreshTokenInput) XXX_DiscardUnknown() {
	xxx_messageInfo_RefreshTokenInput.DiscardUnknown(m)
}

var xxx_messageInfo_RefreshTokenInput proto.InternalMessageInfo

func (m *RefreshTokenInput) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

func (m *RefreshTokenInput) GetMonzoClient() *Client {
	if m != nil {
		return m.MonzoClient
	}
	return nil
}

func (m *RefreshTokenInput) GetAuthKey() string {
	if m != nil {
		return m.AuthKey
	}
	return ""
}

type RefreshTokenOutput struct {
	RefreshToken         string   `protobuf:"bytes,1,opt,name=refreshToken,proto3" json:"refreshToken,omitempty"`
	AuthToken            string   `protobuf:"bytes,2,opt,name=authToken,proto3" json:"authToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RefreshTokenOutput) Reset()         { *m = RefreshTokenOutput{} }
func (m *RefreshTokenOutput) String() string { return proto.CompactTextString(m) }
func (*RefreshTokenOutput) ProtoMessage()    {}
func (*RefreshTokenOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_monzoapi_8777985c2df56e76, []int{1}
}
func (m *RefreshTokenOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RefreshTokenOutput.Unmarshal(m, b)
}
func (m *RefreshTokenOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RefreshTokenOutput.Marshal(b, m, deterministic)
}
func (dst *RefreshTokenOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RefreshTokenOutput.Merge(dst, src)
}
func (m *RefreshTokenOutput) XXX_Size() int {
	return xxx_messageInfo_RefreshTokenOutput.Size(m)
}
func (m *RefreshTokenOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_RefreshTokenOutput.DiscardUnknown(m)
}

var xxx_messageInfo_RefreshTokenOutput proto.InternalMessageInfo

func (m *RefreshTokenOutput) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

func (m *RefreshTokenOutput) GetAuthToken() string {
	if m != nil {
		return m.AuthToken
	}
	return ""
}

type GetTransactionInput struct {
	TransactionID        string   `protobuf:"bytes,1,opt,name=transactionID,proto3" json:"transactionID,omitempty"`
	MonzoClient          *Client  `protobuf:"bytes,2,opt,name=monzoClient,proto3" json:"monzoClient,omitempty"`
	AuthKey              string   `protobuf:"bytes,3,opt,name=authKey,proto3" json:"authKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTransactionInput) Reset()         { *m = GetTransactionInput{} }
func (m *GetTransactionInput) String() string { return proto.CompactTextString(m) }
func (*GetTransactionInput) ProtoMessage()    {}
func (*GetTransactionInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_monzoapi_8777985c2df56e76, []int{2}
}
func (m *GetTransactionInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTransactionInput.Unmarshal(m, b)
}
func (m *GetTransactionInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTransactionInput.Marshal(b, m, deterministic)
}
func (dst *GetTransactionInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTransactionInput.Merge(dst, src)
}
func (m *GetTransactionInput) XXX_Size() int {
	return xxx_messageInfo_GetTransactionInput.Size(m)
}
func (m *GetTransactionInput) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTransactionInput.DiscardUnknown(m)
}

var xxx_messageInfo_GetTransactionInput proto.InternalMessageInfo

func (m *GetTransactionInput) GetTransactionID() string {
	if m != nil {
		return m.TransactionID
	}
	return ""
}

func (m *GetTransactionInput) GetMonzoClient() *Client {
	if m != nil {
		return m.MonzoClient
	}
	return nil
}

func (m *GetTransactionInput) GetAuthKey() string {
	if m != nil {
		return m.AuthKey
	}
	return ""
}

type GetTransactionOutput struct {
	TransactionID        string   `protobuf:"bytes,1,opt,name=transactionID,proto3" json:"transactionID,omitempty"`
	Created              string   `protobuf:"bytes,2,opt,name=created,proto3" json:"created,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Amount               int32    `protobuf:"zigzag32,4,opt,name=amount,proto3" json:"amount,omitempty"`
	LocalAmount          int32    `protobuf:"zigzag32,5,opt,name=localAmount,proto3" json:"localAmount,omitempty"`
	Originator           bool     `protobuf:"varint,6,opt,name=originator,proto3" json:"originator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTransactionOutput) Reset()         { *m = GetTransactionOutput{} }
func (m *GetTransactionOutput) String() string { return proto.CompactTextString(m) }
func (*GetTransactionOutput) ProtoMessage()    {}
func (*GetTransactionOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_monzoapi_8777985c2df56e76, []int{3}
}
func (m *GetTransactionOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTransactionOutput.Unmarshal(m, b)
}
func (m *GetTransactionOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTransactionOutput.Marshal(b, m, deterministic)
}
func (dst *GetTransactionOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTransactionOutput.Merge(dst, src)
}
func (m *GetTransactionOutput) XXX_Size() int {
	return xxx_messageInfo_GetTransactionOutput.Size(m)
}
func (m *GetTransactionOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTransactionOutput.DiscardUnknown(m)
}

var xxx_messageInfo_GetTransactionOutput proto.InternalMessageInfo

func (m *GetTransactionOutput) GetTransactionID() string {
	if m != nil {
		return m.TransactionID
	}
	return ""
}

func (m *GetTransactionOutput) GetCreated() string {
	if m != nil {
		return m.Created
	}
	return ""
}

func (m *GetTransactionOutput) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *GetTransactionOutput) GetAmount() int32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *GetTransactionOutput) GetLocalAmount() int32 {
	if m != nil {
		return m.LocalAmount
	}
	return 0
}

func (m *GetTransactionOutput) GetOriginator() bool {
	if m != nil {
		return m.Originator
	}
	return false
}

type GetCoinJarInput struct {
	MonzoClient          *Client  `protobuf:"bytes,1,opt,name=monzoClient,proto3" json:"monzoClient,omitempty"`
	AuthKey              string   `protobuf:"bytes,2,opt,name=authKey,proto3" json:"authKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCoinJarInput) Reset()         { *m = GetCoinJarInput{} }
func (m *GetCoinJarInput) String() string { return proto.CompactTextString(m) }
func (*GetCoinJarInput) ProtoMessage()    {}
func (*GetCoinJarInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_monzoapi_8777985c2df56e76, []int{4}
}
func (m *GetCoinJarInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCoinJarInput.Unmarshal(m, b)
}
func (m *GetCoinJarInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCoinJarInput.Marshal(b, m, deterministic)
}
func (dst *GetCoinJarInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCoinJarInput.Merge(dst, src)
}
func (m *GetCoinJarInput) XXX_Size() int {
	return xxx_messageInfo_GetCoinJarInput.Size(m)
}
func (m *GetCoinJarInput) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCoinJarInput.DiscardUnknown(m)
}

var xxx_messageInfo_GetCoinJarInput proto.InternalMessageInfo

func (m *GetCoinJarInput) GetMonzoClient() *Client {
	if m != nil {
		return m.MonzoClient
	}
	return nil
}

func (m *GetCoinJarInput) GetAuthKey() string {
	if m != nil {
		return m.AuthKey
	}
	return ""
}

type GetCoinJarOutput struct {
	Pot                  *Pot     `protobuf:"bytes,1,opt,name=pot,proto3" json:"pot,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCoinJarOutput) Reset()         { *m = GetCoinJarOutput{} }
func (m *GetCoinJarOutput) String() string { return proto.CompactTextString(m) }
func (*GetCoinJarOutput) ProtoMessage()    {}
func (*GetCoinJarOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_monzoapi_8777985c2df56e76, []int{5}
}
func (m *GetCoinJarOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCoinJarOutput.Unmarshal(m, b)
}
func (m *GetCoinJarOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCoinJarOutput.Marshal(b, m, deterministic)
}
func (dst *GetCoinJarOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCoinJarOutput.Merge(dst, src)
}
func (m *GetCoinJarOutput) XXX_Size() int {
	return xxx_messageInfo_GetCoinJarOutput.Size(m)
}
func (m *GetCoinJarOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCoinJarOutput.DiscardUnknown(m)
}

var xxx_messageInfo_GetCoinJarOutput proto.InternalMessageInfo

func (m *GetCoinJarOutput) GetPot() *Pot {
	if m != nil {
		return m.Pot
	}
	return nil
}

type Pot struct {
	PotID                string   `protobuf:"bytes,1,opt,name=potID,proto3" json:"potID,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type                 string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Balance              int32    `protobuf:"zigzag32,4,opt,name=balance,proto3" json:"balance,omitempty"`
	Currency             string   `protobuf:"bytes,5,opt,name=currency,proto3" json:"currency,omitempty"`
	Created              string   `protobuf:"bytes,6,opt,name=created,proto3" json:"created,omitempty"`
	Updated              string   `protobuf:"bytes,7,opt,name=updated,proto3" json:"updated,omitempty"`
	Deleted              bool     `protobuf:"varint,8,opt,name=deleted,proto3" json:"deleted,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pot) Reset()         { *m = Pot{} }
func (m *Pot) String() string { return proto.CompactTextString(m) }
func (*Pot) ProtoMessage()    {}
func (*Pot) Descriptor() ([]byte, []int) {
	return fileDescriptor_monzoapi_8777985c2df56e76, []int{6}
}
func (m *Pot) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pot.Unmarshal(m, b)
}
func (m *Pot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pot.Marshal(b, m, deterministic)
}
func (dst *Pot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pot.Merge(dst, src)
}
func (m *Pot) XXX_Size() int {
	return xxx_messageInfo_Pot.Size(m)
}
func (m *Pot) XXX_DiscardUnknown() {
	xxx_messageInfo_Pot.DiscardUnknown(m)
}

var xxx_messageInfo_Pot proto.InternalMessageInfo

func (m *Pot) GetPotID() string {
	if m != nil {
		return m.PotID
	}
	return ""
}

func (m *Pot) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Pot) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Pot) GetBalance() int32 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func (m *Pot) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *Pot) GetCreated() string {
	if m != nil {
		return m.Created
	}
	return ""
}

func (m *Pot) GetUpdated() string {
	if m != nil {
		return m.Updated
	}
	return ""
}

func (m *Pot) GetDeleted() bool {
	if m != nil {
		return m.Deleted
	}
	return false
}

type DepositInput struct {
	Amount               int32    `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	Pot                  *Pot     `protobuf:"bytes,2,opt,name=pot,proto3" json:"pot,omitempty"`
	MonzoClient          *Client  `protobuf:"bytes,3,opt,name=monzoClient,proto3" json:"monzoClient,omitempty"`
	AuthKey              string   `protobuf:"bytes,4,opt,name=authKey,proto3" json:"authKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DepositInput) Reset()         { *m = DepositInput{} }
func (m *DepositInput) String() string { return proto.CompactTextString(m) }
func (*DepositInput) ProtoMessage()    {}
func (*DepositInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_monzoapi_8777985c2df56e76, []int{7}
}
func (m *DepositInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DepositInput.Unmarshal(m, b)
}
func (m *DepositInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DepositInput.Marshal(b, m, deterministic)
}
func (dst *DepositInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DepositInput.Merge(dst, src)
}
func (m *DepositInput) XXX_Size() int {
	return xxx_messageInfo_DepositInput.Size(m)
}
func (m *DepositInput) XXX_DiscardUnknown() {
	xxx_messageInfo_DepositInput.DiscardUnknown(m)
}

var xxx_messageInfo_DepositInput proto.InternalMessageInfo

func (m *DepositInput) GetAmount() int32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *DepositInput) GetPot() *Pot {
	if m != nil {
		return m.Pot
	}
	return nil
}

func (m *DepositInput) GetMonzoClient() *Client {
	if m != nil {
		return m.MonzoClient
	}
	return nil
}

func (m *DepositInput) GetAuthKey() string {
	if m != nil {
		return m.AuthKey
	}
	return ""
}

type DepositOutput struct {
	Pot                  *Pot     `protobuf:"bytes,1,opt,name=pot,proto3" json:"pot,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DepositOutput) Reset()         { *m = DepositOutput{} }
func (m *DepositOutput) String() string { return proto.CompactTextString(m) }
func (*DepositOutput) ProtoMessage()    {}
func (*DepositOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_monzoapi_8777985c2df56e76, []int{8}
}
func (m *DepositOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DepositOutput.Unmarshal(m, b)
}
func (m *DepositOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DepositOutput.Marshal(b, m, deterministic)
}
func (dst *DepositOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DepositOutput.Merge(dst, src)
}
func (m *DepositOutput) XXX_Size() int {
	return xxx_messageInfo_DepositOutput.Size(m)
}
func (m *DepositOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_DepositOutput.DiscardUnknown(m)
}

var xxx_messageInfo_DepositOutput proto.InternalMessageInfo

func (m *DepositOutput) GetPot() *Pot {
	if m != nil {
		return m.Pot
	}
	return nil
}

type Client struct {
	ClientID             string   `protobuf:"bytes,1,opt,name=clientID,proto3" json:"clientID,omitempty"`
	ClientSecret         string   `protobuf:"bytes,2,opt,name=clientSecret,proto3" json:"clientSecret,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Client) Reset()         { *m = Client{} }
func (m *Client) String() string { return proto.CompactTextString(m) }
func (*Client) ProtoMessage()    {}
func (*Client) Descriptor() ([]byte, []int) {
	return fileDescriptor_monzoapi_8777985c2df56e76, []int{9}
}
func (m *Client) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Client.Unmarshal(m, b)
}
func (m *Client) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Client.Marshal(b, m, deterministic)
}
func (dst *Client) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Client.Merge(dst, src)
}
func (m *Client) XXX_Size() int {
	return xxx_messageInfo_Client.Size(m)
}
func (m *Client) XXX_DiscardUnknown() {
	xxx_messageInfo_Client.DiscardUnknown(m)
}

var xxx_messageInfo_Client proto.InternalMessageInfo

func (m *Client) GetClientID() string {
	if m != nil {
		return m.ClientID
	}
	return ""
}

func (m *Client) GetClientSecret() string {
	if m != nil {
		return m.ClientSecret
	}
	return ""
}

func init() {
	proto.RegisterType((*RefreshTokenInput)(nil), "monzoapi.RefreshTokenInput")
	proto.RegisterType((*RefreshTokenOutput)(nil), "monzoapi.RefreshTokenOutput")
	proto.RegisterType((*GetTransactionInput)(nil), "monzoapi.GetTransactionInput")
	proto.RegisterType((*GetTransactionOutput)(nil), "monzoapi.GetTransactionOutput")
	proto.RegisterType((*GetCoinJarInput)(nil), "monzoapi.GetCoinJarInput")
	proto.RegisterType((*GetCoinJarOutput)(nil), "monzoapi.GetCoinJarOutput")
	proto.RegisterType((*Pot)(nil), "monzoapi.Pot")
	proto.RegisterType((*DepositInput)(nil), "monzoapi.DepositInput")
	proto.RegisterType((*DepositOutput)(nil), "monzoapi.DepositOutput")
	proto.RegisterType((*Client)(nil), "monzoapi.Client")
}

func init() {
	proto.RegisterFile("types/services/monzoapi.proto", fileDescriptor_monzoapi_8777985c2df56e76)
}

var fileDescriptor_monzoapi_8777985c2df56e76 = []byte{
	// 618 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x96, 0x93, 0x36, 0x49, 0x27, 0x09, 0xb4, 0x4b, 0x55, 0x8c, 0x69, 0x4b, 0x64, 0x71, 0xc8,
	0x85, 0x04, 0xa5, 0x37, 0x10, 0x87, 0x92, 0x4a, 0x25, 0x20, 0xd4, 0xc8, 0x54, 0x1c, 0xb8, 0xa0,
	0x8d, 0x3d, 0x34, 0x16, 0xf1, 0xae, 0xb5, 0x5e, 0x23, 0x85, 0x3b, 0x07, 0x1e, 0x80, 0x67, 0xe0,
	0x51, 0xb8, 0xf3, 0x44, 0x68, 0xed, 0xf5, 0x5f, 0x9a, 0x48, 0xa9, 0xc4, 0x6d, 0xe7, 0xfb, 0xc6,
	0xb3, 0x33, 0xdf, 0xe7, 0xdd, 0x85, 0x13, 0xb9, 0x0c, 0x31, 0x1a, 0x46, 0x28, 0xbe, 0xf9, 0x2e,
	0x46, 0xc3, 0x80, 0xb3, 0xef, 0x9c, 0x86, 0xfe, 0x20, 0x14, 0x5c, 0x72, 0xd2, 0xca, 0x62, 0x6b,
	0x35, 0xd1, 0x5b, 0x32, 0x1a, 0x70, 0x6f, 0x96, 0x26, 0xda, 0x3f, 0x0c, 0x38, 0x70, 0xf0, 0x8b,
	0xc0, 0x68, 0x7e, 0xcd, 0xbf, 0x22, 0x9b, 0xb0, 0x30, 0x96, 0xc4, 0x86, 0x8e, 0x28, 0x81, 0xa6,
	0xd1, 0x33, 0xfa, 0x7b, 0x4e, 0x05, 0x23, 0x23, 0x68, 0x27, 0x9b, 0x8c, 0x17, 0x3e, 0x32, 0x69,
	0xd6, 0x7a, 0x46, 0xbf, 0x3d, 0xda, 0x1f, 0xe4, 0x8d, 0xa4, 0xb8, 0x53, 0x4e, 0x22, 0x26, 0x34,
	0x69, 0x2c, 0xe7, 0xef, 0x70, 0x69, 0xd6, 0x93, 0x92, 0x59, 0x68, 0x7f, 0x04, 0x52, 0x6e, 0xe3,
	0x2a, 0x96, 0xdb, 0xf6, 0x71, 0x0c, 0x7b, 0xaa, 0x48, 0x9a, 0x50, 0x4b, 0x12, 0x0a, 0xc0, 0xfe,
	0x69, 0xc0, 0x83, 0x4b, 0x94, 0xd7, 0x82, 0xb2, 0x88, 0xba, 0xd2, 0xe7, 0x7a, 0xc2, 0xa7, 0xd0,
	0x95, 0x25, 0xec, 0x42, 0x97, 0xae, 0x82, 0xff, 0x79, 0xc6, 0xbf, 0x06, 0x1c, 0x56, 0x7b, 0xd1,
	0x63, 0x6e, 0xd7, 0x8c, 0x09, 0x4d, 0x57, 0x20, 0x95, 0xe8, 0xe9, 0x31, 0xb3, 0x90, 0xf4, 0xa0,
	0xed, 0x61, 0xe4, 0x0a, 0x3f, 0x54, 0xa9, 0x7a, 0xdb, 0x32, 0x44, 0x8e, 0xa0, 0x41, 0x03, 0x1e,
	0x33, 0x69, 0xee, 0xf4, 0x8c, 0xfe, 0x81, 0xa3, 0x23, 0xf5, 0xe5, 0x82, 0xbb, 0x74, 0x71, 0x9e,
	0x92, 0xbb, 0x09, 0x59, 0x86, 0xc8, 0x29, 0x00, 0x17, 0xfe, 0x8d, 0xcf, 0xa8, 0xe4, 0xc2, 0x6c,
	0xf4, 0x8c, 0x7e, 0xcb, 0x29, 0x21, 0xf6, 0x67, 0xb8, 0x7f, 0x89, 0x72, 0xcc, 0x7d, 0xf6, 0x96,
	0x8a, 0x54, 0xdb, 0x15, 0xd5, 0x8c, 0x3b, 0xaa, 0x56, 0xab, 0xaa, 0x76, 0x06, 0xfb, 0xc5, 0x06,
	0x5a, 0xb0, 0x27, 0x50, 0x0f, 0x79, 0x56, 0xb9, 0x5b, 0x54, 0x9e, 0x72, 0xe9, 0x28, 0xc6, 0xfe,
	0x63, 0x40, 0x7d, 0xca, 0x25, 0x39, 0x84, 0xdd, 0x90, 0xcb, 0x5c, 0xd1, 0x34, 0x20, 0x04, 0x76,
	0x18, 0x0d, 0x50, 0xef, 0x94, 0xac, 0x15, 0xa6, 0x4e, 0x8a, 0x16, 0x2f, 0x59, 0xab, 0xa6, 0x66,
	0x74, 0x41, 0x99, 0x8b, 0x5a, 0xb6, 0x2c, 0x24, 0x16, 0xb4, 0xdc, 0x58, 0x08, 0x64, 0xee, 0x32,
	0x11, 0x6d, 0xcf, 0xc9, 0xe3, 0xb2, 0x4f, 0x8d, 0xaa, 0x4f, 0x26, 0x34, 0xe3, 0xd0, 0x4b, 0x98,
	0x66, 0xca, 0xe8, 0x50, 0x31, 0x1e, 0x2e, 0x50, 0x31, 0xad, 0x44, 0xe2, 0x2c, 0xb4, 0x7f, 0x19,
	0xd0, 0xb9, 0xc0, 0x90, 0x47, 0xbe, 0x4c, 0xd5, 0x2d, 0xac, 0x54, 0x33, 0xed, 0xe6, 0x56, 0x6a,
	0x4d, 0x6a, 0x9b, 0x34, 0x59, 0xb5, 0xa5, 0x7e, 0x47, 0x5b, 0x76, 0xaa, 0xb6, 0x3c, 0x87, 0xae,
	0x6e, 0x6b, 0x5b, 0x4f, 0xde, 0x40, 0x43, 0x57, 0x55, 0xea, 0x25, 0xab, 0xdc, 0x98, 0x3c, 0x56,
	0x47, 0x3e, 0x5d, 0x7f, 0x40, 0x57, 0xa0, 0xd4, 0x1e, 0x55, 0xb0, 0xd1, 0xef, 0x1a, 0xb4, 0xde,
	0xab, 0xfa, 0xe7, 0xd3, 0x09, 0x99, 0x40, 0xa7, 0x7c, 0x73, 0x90, 0xc7, 0xc5, 0xd6, 0xb7, 0x2e,
	0x36, 0xeb, 0x78, 0x3d, 0xa9, 0x47, 0xb8, 0x82, 0x7b, 0xd5, 0xf3, 0x49, 0x4e, 0x8a, 0xfc, 0x35,
	0xb7, 0x88, 0x75, 0xba, 0x89, 0xd6, 0x05, 0xc7, 0x00, 0xc5, 0xbf, 0x4b, 0x1e, 0x55, 0xb2, 0xcb,
	0x47, 0xc6, 0xb2, 0xd6, 0x51, 0xba, 0xc8, 0x0b, 0x68, 0x6a, 0xa5, 0xc9, 0x51, 0x91, 0x56, 0xfe,
	0x27, 0xac, 0x87, 0xb7, 0xf0, 0xf4, 0xdb, 0xd7, 0xaf, 0x3e, 0xbd, 0xbc, 0xf1, 0xe5, 0x3c, 0x9e,
	0x0d, 0x5c, 0x1e, 0x0c, 0x03, 0x2a, 0xa5, 0xa0, 0x4b, 0x86, 0x22, 0x7d, 0x2f, 0x9e, 0x09, 0x1e,
	0x33, 0x2f, 0x0e, 0x87, 0x1b, 0x1e, 0x93, 0x59, 0x23, 0x79, 0x24, 0xce, 0xfe, 0x05, 0x00, 0x00,
	0xff, 0xff, 0x9e, 0x2a, 0x9a, 0xb2, 0x6e, 0x06, 0x00, 0x00,
}
