// Code generated by protoc-gen-go. DO NOT EDIT.
// source: account.proto

package account

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

type AccountResponse struct {
	StatusCode           int32             `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	StatusMessage        string            `protobuf:"bytes,2,opt,name=statusMessage,proto3" json:"statusMessage,omitempty"`
	Header               map[string]string `protobuf:"bytes,3,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Body                 []byte            `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *AccountResponse) Reset()         { *m = AccountResponse{} }
func (m *AccountResponse) String() string { return proto.CompactTextString(m) }
func (*AccountResponse) ProtoMessage()    {}
func (*AccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{0}
}

func (m *AccountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountResponse.Unmarshal(m, b)
}
func (m *AccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountResponse.Marshal(b, m, deterministic)
}
func (m *AccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountResponse.Merge(m, src)
}
func (m *AccountResponse) XXX_Size() int {
	return xxx_messageInfo_AccountResponse.Size(m)
}
func (m *AccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AccountResponse proto.InternalMessageInfo

func (m *AccountResponse) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *AccountResponse) GetStatusMessage() string {
	if m != nil {
		return m.StatusMessage
	}
	return ""
}

func (m *AccountResponse) GetHeader() map[string]string {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *AccountResponse) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type CreateAccountRequest struct {
	SessionToken         string   `protobuf:"bytes,1,opt,name=SessionToken,proto3" json:"SessionToken,omitempty"`
	RequestBody          []byte   `protobuf:"bytes,2,opt,name=RequestBody,proto3" json:"RequestBody,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateAccountRequest) Reset()         { *m = CreateAccountRequest{} }
func (m *CreateAccountRequest) String() string { return proto.CompactTextString(m) }
func (*CreateAccountRequest) ProtoMessage()    {}
func (*CreateAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{1}
}

func (m *CreateAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateAccountRequest.Unmarshal(m, b)
}
func (m *CreateAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateAccountRequest.Marshal(b, m, deterministic)
}
func (m *CreateAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAccountRequest.Merge(m, src)
}
func (m *CreateAccountRequest) XXX_Size() int {
	return xxx_messageInfo_CreateAccountRequest.Size(m)
}
func (m *CreateAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAccountRequest proto.InternalMessageInfo

func (m *CreateAccountRequest) GetSessionToken() string {
	if m != nil {
		return m.SessionToken
	}
	return ""
}

func (m *CreateAccountRequest) GetRequestBody() []byte {
	if m != nil {
		return m.RequestBody
	}
	return nil
}

type AccountRequest struct {
	SessionToken         string   `protobuf:"bytes,1,opt,name=SessionToken,proto3" json:"SessionToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountRequest) Reset()         { *m = AccountRequest{} }
func (m *AccountRequest) String() string { return proto.CompactTextString(m) }
func (*AccountRequest) ProtoMessage()    {}
func (*AccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{2}
}

func (m *AccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountRequest.Unmarshal(m, b)
}
func (m *AccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountRequest.Marshal(b, m, deterministic)
}
func (m *AccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountRequest.Merge(m, src)
}
func (m *AccountRequest) XXX_Size() int {
	return xxx_messageInfo_AccountRequest.Size(m)
}
func (m *AccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AccountRequest proto.InternalMessageInfo

func (m *AccountRequest) GetSessionToken() string {
	if m != nil {
		return m.SessionToken
	}
	return ""
}

type GetAccountRequest struct {
	SessionToken         string   `protobuf:"bytes,1,opt,name=SessionToken,proto3" json:"SessionToken,omitempty"`
	AccountID            string   `protobuf:"bytes,2,opt,name=AccountID,proto3" json:"AccountID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAccountRequest) Reset()         { *m = GetAccountRequest{} }
func (m *GetAccountRequest) String() string { return proto.CompactTextString(m) }
func (*GetAccountRequest) ProtoMessage()    {}
func (*GetAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{3}
}

func (m *GetAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAccountRequest.Unmarshal(m, b)
}
func (m *GetAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAccountRequest.Marshal(b, m, deterministic)
}
func (m *GetAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAccountRequest.Merge(m, src)
}
func (m *GetAccountRequest) XXX_Size() int {
	return xxx_messageInfo_GetAccountRequest.Size(m)
}
func (m *GetAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAccountRequest proto.InternalMessageInfo

func (m *GetAccountRequest) GetSessionToken() string {
	if m != nil {
		return m.SessionToken
	}
	return ""
}

func (m *GetAccountRequest) GetAccountID() string {
	if m != nil {
		return m.AccountID
	}
	return ""
}

type UpdateAccountRequest struct {
	SessionToken         string   `protobuf:"bytes,1,opt,name=SessionToken,proto3" json:"SessionToken,omitempty"`
	AccountID            string   `protobuf:"bytes,2,opt,name=AccountID,proto3" json:"AccountID,omitempty"`
	RequestBody          []byte   `protobuf:"bytes,3,opt,name=RequestBody,proto3" json:"RequestBody,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateAccountRequest) Reset()         { *m = UpdateAccountRequest{} }
func (m *UpdateAccountRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateAccountRequest) ProtoMessage()    {}
func (*UpdateAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{4}
}

func (m *UpdateAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateAccountRequest.Unmarshal(m, b)
}
func (m *UpdateAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateAccountRequest.Marshal(b, m, deterministic)
}
func (m *UpdateAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateAccountRequest.Merge(m, src)
}
func (m *UpdateAccountRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateAccountRequest.Size(m)
}
func (m *UpdateAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateAccountRequest proto.InternalMessageInfo

func (m *UpdateAccountRequest) GetSessionToken() string {
	if m != nil {
		return m.SessionToken
	}
	return ""
}

func (m *UpdateAccountRequest) GetAccountID() string {
	if m != nil {
		return m.AccountID
	}
	return ""
}

func (m *UpdateAccountRequest) GetRequestBody() []byte {
	if m != nil {
		return m.RequestBody
	}
	return nil
}

type DeleteAccountRequest struct {
	SessionToken         string   `protobuf:"bytes,1,opt,name=SessionToken,proto3" json:"SessionToken,omitempty"`
	AccountID            string   `protobuf:"bytes,2,opt,name=AccountID,proto3" json:"AccountID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteAccountRequest) Reset()         { *m = DeleteAccountRequest{} }
func (m *DeleteAccountRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteAccountRequest) ProtoMessage()    {}
func (*DeleteAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{5}
}

func (m *DeleteAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteAccountRequest.Unmarshal(m, b)
}
func (m *DeleteAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteAccountRequest.Marshal(b, m, deterministic)
}
func (m *DeleteAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteAccountRequest.Merge(m, src)
}
func (m *DeleteAccountRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteAccountRequest.Size(m)
}
func (m *DeleteAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteAccountRequest proto.InternalMessageInfo

func (m *DeleteAccountRequest) GetSessionToken() string {
	if m != nil {
		return m.SessionToken
	}
	return ""
}

func (m *DeleteAccountRequest) GetAccountID() string {
	if m != nil {
		return m.AccountID
	}
	return ""
}

func init() {
	proto.RegisterType((*AccountResponse)(nil), "AccountResponse")
	proto.RegisterMapType((map[string]string)(nil), "AccountResponse.HeaderEntry")
	proto.RegisterType((*CreateAccountRequest)(nil), "CreateAccountRequest")
	proto.RegisterType((*AccountRequest)(nil), "AccountRequest")
	proto.RegisterType((*GetAccountRequest)(nil), "GetAccountRequest")
	proto.RegisterType((*UpdateAccountRequest)(nil), "UpdateAccountRequest")
	proto.RegisterType((*DeleteAccountRequest)(nil), "DeleteAccountRequest")
}

func init() { proto.RegisterFile("account.proto", fileDescriptor_8e28828dcb8d24f0) }

var fileDescriptor_8e28828dcb8d24f0 = []byte{
	// 378 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0x4d, 0x6f, 0xda, 0x40,
	0x10, 0xad, 0x6d, 0x70, 0xc5, 0xf0, 0xd9, 0x91, 0x91, 0x2c, 0x84, 0x2a, 0xcb, 0xea, 0xc1, 0xa7,
	0x3d, 0x00, 0x95, 0x4a, 0x6f, 0x2d, 0x44, 0x49, 0x0e, 0xb9, 0x98, 0x20, 0xe5, 0x90, 0x8b, 0xc1,
	0xa3, 0x04, 0x61, 0x79, 0x89, 0x77, 0x8d, 0x44, 0x7e, 0x65, 0x7e, 0x40, 0x7e, 0x4c, 0x84, 0xed,
	0x84, 0x2f, 0x27, 0x22, 0x28, 0xb7, 0x9d, 0xb7, 0xf3, 0x66, 0xde, 0xbe, 0x67, 0x43, 0xd5, 0x9b,
	0x4e, 0x79, 0x1c, 0x4a, 0xb6, 0x88, 0xb8, 0xe4, 0xf6, 0xb3, 0x02, 0xf5, 0x7f, 0x29, 0xe2, 0x92,
	0x58, 0xf0, 0x50, 0x10, 0xfe, 0x04, 0x10, 0xd2, 0x93, 0xb1, 0x18, 0x70, 0x9f, 0x4c, 0xc5, 0x52,
	0x9c, 0xa2, 0xbb, 0x85, 0xe0, 0x2f, 0xa8, 0xa6, 0xd5, 0x15, 0x09, 0xe1, 0xdd, 0x91, 0xa9, 0x5a,
	0x8a, 0x53, 0x72, 0x77, 0x41, 0xec, 0x81, 0x7e, 0x4f, 0x9e, 0x4f, 0x91, 0xa9, 0x59, 0x9a, 0x53,
	0xee, 0xb4, 0xd9, 0xde, 0x1e, 0x76, 0x91, 0x5c, 0x9f, 0x85, 0x32, 0x5a, 0xb9, 0x59, 0x2f, 0x22,
	0x14, 0x26, 0xdc, 0x5f, 0x99, 0x05, 0x4b, 0x71, 0x2a, 0x6e, 0x72, 0x6e, 0xf5, 0xa1, 0xbc, 0xd5,
	0x8a, 0x0d, 0xd0, 0xe6, 0xb4, 0x4a, 0x74, 0x95, 0xdc, 0xf5, 0x11, 0x0d, 0x28, 0x2e, 0xbd, 0x20,
	0x7e, 0x15, 0x92, 0x16, 0x7f, 0xd5, 0x3f, 0x8a, 0x7d, 0x0b, 0xc6, 0x20, 0x22, 0x4f, 0xd2, 0xdb,
	0xee, 0x87, 0x98, 0x84, 0x44, 0x1b, 0x2a, 0x23, 0x12, 0x62, 0xc6, 0xc3, 0x6b, 0x3e, 0xa7, 0x30,
	0x1b, 0xb6, 0x83, 0xa1, 0x05, 0xe5, 0xac, 0xfd, 0xff, 0x5a, 0x91, 0x9a, 0x28, 0xda, 0x86, 0xec,
	0x1e, 0xd4, 0x3e, 0x3f, 0xd7, 0x1e, 0xc3, 0x8f, 0x73, 0x92, 0x27, 0x08, 0x6a, 0x43, 0x29, 0x63,
	0x5d, 0x0e, 0xb3, 0xa7, 0x6e, 0x00, 0xfb, 0x11, 0x8c, 0xf1, 0xc2, 0x3f, 0xed, 0xa9, 0x1f, 0x4e,
	0xde, 0x37, 0x42, 0x3b, 0x34, 0xe2, 0x06, 0x8c, 0x21, 0x05, 0xf4, 0xf5, 0xbb, 0x3b, 0x4f, 0x2a,
	0x7c, 0xcf, 0x2a, 0xec, 0x82, 0x9e, 0x86, 0x89, 0x4d, 0x96, 0x97, 0x6a, 0xab, 0xb1, 0xff, 0x89,
	0xd9, 0xdf, 0xf0, 0x37, 0xd4, 0xd6, 0x6e, 0x07, 0x41, 0x76, 0x25, 0xb0, 0xce, 0x8e, 0xa0, 0xf5,
	0x00, 0x36, 0x21, 0x21, 0xb2, 0x83, 0xc4, 0x72, 0x59, 0x7d, 0xc0, 0x4d, 0xe3, 0x88, 0xa2, 0xe5,
	0x6c, 0x4a, 0x47, 0x2e, 0xec, 0x82, 0x9e, 0xc6, 0x87, 0x4d, 0x96, 0x97, 0xe3, 0x7b, 0xa4, 0xd4,
	0x77, 0x6c, 0xb2, 0xbc, 0x00, 0xf2, 0x48, 0x13, 0x3d, 0xf9, 0xf3, 0xbb, 0x2f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x27, 0xb3, 0x27, 0x8e, 0x0a, 0x04, 0x00, 0x00,
}