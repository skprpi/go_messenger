// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Credential struct {
	Login                string   `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Credential) Reset()         { *m = Credential{} }
func (m *Credential) String() string { return proto.CompactTextString(m) }
func (*Credential) ProtoMessage()    {}
func (*Credential) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{0}
}

func (m *Credential) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Credential.Unmarshal(m, b)
}
func (m *Credential) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Credential.Marshal(b, m, deterministic)
}
func (m *Credential) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Credential.Merge(m, src)
}
func (m *Credential) XXX_Size() int {
	return xxx_messageInfo_Credential.Size(m)
}
func (m *Credential) XXX_DiscardUnknown() {
	xxx_messageInfo_Credential.DiscardUnknown(m)
}

var xxx_messageInfo_Credential proto.InternalMessageInfo

func (m *Credential) GetLogin() string {
	if m != nil {
		return m.Login
	}
	return ""
}

func (m *Credential) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type TokenPayload struct {
	AccountID            string   `protobuf:"bytes,1,opt,name=accountID,proto3" json:"accountID,omitempty"`
	UserRole             string   `protobuf:"bytes,2,opt,name=userRole,proto3" json:"userRole,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TokenPayload) Reset()         { *m = TokenPayload{} }
func (m *TokenPayload) String() string { return proto.CompactTextString(m) }
func (*TokenPayload) ProtoMessage()    {}
func (*TokenPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{1}
}

func (m *TokenPayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TokenPayload.Unmarshal(m, b)
}
func (m *TokenPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TokenPayload.Marshal(b, m, deterministic)
}
func (m *TokenPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenPayload.Merge(m, src)
}
func (m *TokenPayload) XXX_Size() int {
	return xxx_messageInfo_TokenPayload.Size(m)
}
func (m *TokenPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenPayload.DiscardUnknown(m)
}

var xxx_messageInfo_TokenPayload proto.InternalMessageInfo

func (m *TokenPayload) GetAccountID() string {
	if m != nil {
		return m.AccountID
	}
	return ""
}

func (m *TokenPayload) GetUserRole() string {
	if m != nil {
		return m.UserRole
	}
	return ""
}

type Token struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{2}
}

func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

type AccountID struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountID) Reset()         { *m = AccountID{} }
func (m *AccountID) String() string { return proto.CompactTextString(m) }
func (*AccountID) ProtoMessage()    {}
func (*AccountID) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{3}
}

func (m *AccountID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountID.Unmarshal(m, b)
}
func (m *AccountID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountID.Marshal(b, m, deterministic)
}
func (m *AccountID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountID.Merge(m, src)
}
func (m *AccountID) XXX_Size() int {
	return xxx_messageInfo_AccountID.Size(m)
}
func (m *AccountID) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountID.DiscardUnknown(m)
}

var xxx_messageInfo_AccountID proto.InternalMessageInfo

func (m *AccountID) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func init() {
	proto.RegisterType((*Credential)(nil), "proto.Credential")
	proto.RegisterType((*TokenPayload)(nil), "proto.TokenPayload")
	proto.RegisterType((*Token)(nil), "proto.Token")
	proto.RegisterType((*AccountID)(nil), "proto.AccountID")
}

func init() {
	proto.RegisterFile("auth.proto", fileDescriptor_8bbd6f3875b0e874)
}

var fileDescriptor_8bbd6f3875b0e874 = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x4d, 0x4b, 0xfb, 0x40,
	0x10, 0xc6, 0x93, 0xd0, 0xfc, 0xff, 0xcd, 0x34, 0xbe, 0x8d, 0x1e, 0x4a, 0xf4, 0x50, 0xf6, 0xa4,
	0x07, 0x23, 0x28, 0xf4, 0x28, 0xd4, 0x16, 0xb4, 0xb7, 0x50, 0xc4, 0x83, 0xb7, 0x35, 0x59, 0xd2,
	0x60, 0xba, 0x53, 0x92, 0x0d, 0xe2, 0x47, 0xf4, 0x5b, 0x89, 0xfb, 0x92, 0x56, 0xf0, 0x34, 0x3c,
	0xcf, 0x3c, 0xfb, 0xdb, 0x99, 0x01, 0xe0, 0x9d, 0x5a, 0xa7, 0xdb, 0x86, 0x14, 0x61, 0xa8, 0x4b,
	0x12, 0xe7, 0xb4, 0xd9, 0x90, 0x34, 0x26, 0xbb, 0x07, 0x98, 0x37, 0xa2, 0x10, 0x52, 0x55, 0xbc,
	0xc6, 0x33, 0x08, 0x6b, 0x2a, 0x2b, 0x39, 0xf6, 0x27, 0xfe, 0x65, 0xb4, 0x32, 0x02, 0x13, 0x18,
	0x6e, 0x79, 0xdb, 0x7e, 0x50, 0x53, 0x8c, 0x03, 0xdd, 0xe8, 0x35, 0x7b, 0x82, 0xf8, 0x99, 0xde,
	0x85, 0xcc, 0xf8, 0x67, 0x4d, 0xbc, 0xc0, 0x0b, 0x88, 0x78, 0x9e, 0x53, 0x27, 0xd5, 0x72, 0x61,
	0x29, 0x3b, 0xe3, 0x87, 0xd4, 0xb5, 0xa2, 0x59, 0x51, 0x2d, 0x1c, 0xc9, 0x69, 0x76, 0x05, 0xa1,
	0x26, 0xe1, 0x04, 0x46, 0x3c, 0xcf, 0x45, 0xdb, 0x6a, 0x69, 0x21, 0xfb, 0x16, 0x3b, 0x87, 0x68,
	0xd6, 0x33, 0x0f, 0x21, 0xe8, 0xbf, 0x0a, 0x96, 0x8b, 0xdb, 0x2f, 0x1f, 0x06, 0xb3, 0x4e, 0xad,
	0x71, 0x0a, 0x07, 0xf3, 0x46, 0x70, 0x25, 0x6c, 0x16, 0x4f, 0xcc, 0xce, 0xe9, 0x6e, 0xe1, 0xe4,
	0xd8, 0x5a, 0x3d, 0x8e, 0x79, 0x78, 0x0d, 0xc3, 0x47, 0xa1, 0xcc, 0x2c, 0x7f, 0x3c, 0x89, 0xad,
	0x65, 0x46, 0xf1, 0x70, 0x0a, 0x47, 0x2e, 0xee, 0x8e, 0xf0, 0x2b, 0x92, 0x9c, 0xee, 0x2b, 0x1b,
	0x61, 0x1e, 0x32, 0x18, 0x64, 0x95, 0x2c, 0x71, 0x64, 0xdb, 0x2f, 0x54, 0x15, 0x89, 0x13, 0x19,
	0xc9, 0x92, 0x79, 0x0f, 0xd1, 0xeb, 0xff, 0xf4, 0x46, 0x3b, 0x6f, 0xff, 0x74, 0xb9, 0xfb, 0x0e,
	0x00, 0x00, 0xff, 0xff, 0xe0, 0x89, 0x4a, 0xc1, 0xd2, 0x01, 0x00, 0x00,
}