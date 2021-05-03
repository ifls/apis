// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

package auth

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

type AuthRequest struct {
	ResponseType         string   `protobuf:"bytes,1,opt,name=response_type,json=responseType,proto3" json:"response_type,omitempty"`
	ClientId             string   `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	RedirectUri          string   `protobuf:"bytes,3,opt,name=redirect_uri,json=redirectUri,proto3" json:"redirect_uri,omitempty"`
	Scope                string   `protobuf:"bytes,4,opt,name=scope,proto3" json:"scope,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthRequest) Reset()         { *m = AuthRequest{} }
func (m *AuthRequest) String() string { return proto.CompactTextString(m) }
func (*AuthRequest) ProtoMessage()    {}
func (*AuthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{0}
}

func (m *AuthRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthRequest.Unmarshal(m, b)
}
func (m *AuthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthRequest.Marshal(b, m, deterministic)
}
func (m *AuthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthRequest.Merge(m, src)
}
func (m *AuthRequest) XXX_Size() int {
	return xxx_messageInfo_AuthRequest.Size(m)
}
func (m *AuthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthRequest proto.InternalMessageInfo

func (m *AuthRequest) GetResponseType() string {
	if m != nil {
		return m.ResponseType
	}
	return ""
}

func (m *AuthRequest) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *AuthRequest) GetRedirectUri() string {
	if m != nil {
		return m.RedirectUri
	}
	return ""
}

func (m *AuthRequest) GetScope() string {
	if m != nil {
		return m.Scope
	}
	return ""
}

type AuthResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthResponse) Reset()         { *m = AuthResponse{} }
func (m *AuthResponse) String() string { return proto.CompactTextString(m) }
func (*AuthResponse) ProtoMessage()    {}
func (*AuthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{1}
}

func (m *AuthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthResponse.Unmarshal(m, b)
}
func (m *AuthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthResponse.Marshal(b, m, deterministic)
}
func (m *AuthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthResponse.Merge(m, src)
}
func (m *AuthResponse) XXX_Size() int {
	return xxx_messageInfo_AuthResponse.Size(m)
}
func (m *AuthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthResponse proto.InternalMessageInfo

type GetTokenRequest struct {
	ClientId             string   `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ClientSecret         string   `protobuf:"bytes,2,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret,omitempty"`
	GrantType            string   `protobuf:"bytes,3,opt,name=grant_type,json=grantType,proto3" json:"grant_type,omitempty"`
	Code                 string   `protobuf:"bytes,4,opt,name=code,proto3" json:"code,omitempty"`
	RedirectUri          string   `protobuf:"bytes,5,opt,name=redirect_uri,json=redirectUri,proto3" json:"redirect_uri,omitempty"`
	RefreshToken         string   `protobuf:"bytes,6,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTokenRequest) Reset()         { *m = GetTokenRequest{} }
func (m *GetTokenRequest) String() string { return proto.CompactTextString(m) }
func (*GetTokenRequest) ProtoMessage()    {}
func (*GetTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{2}
}

func (m *GetTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTokenRequest.Unmarshal(m, b)
}
func (m *GetTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTokenRequest.Marshal(b, m, deterministic)
}
func (m *GetTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTokenRequest.Merge(m, src)
}
func (m *GetTokenRequest) XXX_Size() int {
	return xxx_messageInfo_GetTokenRequest.Size(m)
}
func (m *GetTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTokenRequest proto.InternalMessageInfo

func (m *GetTokenRequest) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *GetTokenRequest) GetClientSecret() string {
	if m != nil {
		return m.ClientSecret
	}
	return ""
}

func (m *GetTokenRequest) GetGrantType() string {
	if m != nil {
		return m.GrantType
	}
	return ""
}

func (m *GetTokenRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *GetTokenRequest) GetRedirectUri() string {
	if m != nil {
		return m.RedirectUri
	}
	return ""
}

func (m *GetTokenRequest) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

type TokenInfo struct {
	AccessToken          string            `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	TokenType            string            `protobuf:"bytes,2,opt,name=token_type,json=tokenType,proto3" json:"token_type,omitempty"`
	ExpiresIn            int64             `protobuf:"varint,3,opt,name=expires_in,json=expiresIn,proto3" json:"expires_in,omitempty"`
	RefreshToken         string            `protobuf:"bytes,4,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	Scope                string            `protobuf:"bytes,5,opt,name=scope,proto3" json:"scope,omitempty"`
	Uid                  string            `protobuf:"bytes,6,opt,name=uid,proto3" json:"uid,omitempty"`
	Info                 map[string]string `protobuf:"bytes,7,rep,name=info,proto3" json:"info,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TokenInfo) Reset()         { *m = TokenInfo{} }
func (m *TokenInfo) String() string { return proto.CompactTextString(m) }
func (*TokenInfo) ProtoMessage()    {}
func (*TokenInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{3}
}

func (m *TokenInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TokenInfo.Unmarshal(m, b)
}
func (m *TokenInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TokenInfo.Marshal(b, m, deterministic)
}
func (m *TokenInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenInfo.Merge(m, src)
}
func (m *TokenInfo) XXX_Size() int {
	return xxx_messageInfo_TokenInfo.Size(m)
}
func (m *TokenInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TokenInfo proto.InternalMessageInfo

func (m *TokenInfo) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *TokenInfo) GetTokenType() string {
	if m != nil {
		return m.TokenType
	}
	return ""
}

func (m *TokenInfo) GetExpiresIn() int64 {
	if m != nil {
		return m.ExpiresIn
	}
	return 0
}

func (m *TokenInfo) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

func (m *TokenInfo) GetScope() string {
	if m != nil {
		return m.Scope
	}
	return ""
}

func (m *TokenInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *TokenInfo) GetInfo() map[string]string {
	if m != nil {
		return m.Info
	}
	return nil
}

type GetTokenResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTokenResponse) Reset()         { *m = GetTokenResponse{} }
func (m *GetTokenResponse) String() string { return proto.CompactTextString(m) }
func (*GetTokenResponse) ProtoMessage()    {}
func (*GetTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{4}
}

func (m *GetTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTokenResponse.Unmarshal(m, b)
}
func (m *GetTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTokenResponse.Marshal(b, m, deterministic)
}
func (m *GetTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTokenResponse.Merge(m, src)
}
func (m *GetTokenResponse) XXX_Size() int {
	return xxx_messageInfo_GetTokenResponse.Size(m)
}
func (m *GetTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTokenResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AuthRequest)(nil), "heyifeng.auth.v1.AuthRequest")
	proto.RegisterType((*AuthResponse)(nil), "heyifeng.auth.v1.AuthResponse")
	proto.RegisterType((*GetTokenRequest)(nil), "heyifeng.auth.v1.GetTokenRequest")
	proto.RegisterType((*TokenInfo)(nil), "heyifeng.auth.v1.TokenInfo")
	proto.RegisterMapType((map[string]string)(nil), "heyifeng.auth.v1.TokenInfo.InfoEntry")
	proto.RegisterType((*GetTokenResponse)(nil), "heyifeng.auth.v1.GetTokenResponse")
}

func init() { proto.RegisterFile("auth.proto", fileDescriptor_8bbd6f3875b0e874) }

var fileDescriptor_8bbd6f3875b0e874 = []byte{
	// 445 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0xcf, 0x6e, 0xd4, 0x30,
	0x10, 0xc6, 0xc9, 0x6e, 0x76, 0x69, 0x66, 0xb7, 0xb0, 0xb2, 0x38, 0x44, 0x8b, 0x16, 0xb5, 0xa9,
	0x90, 0x7a, 0x8a, 0x44, 0x39, 0xf0, 0xe7, 0x06, 0x12, 0xaa, 0xf6, 0xba, 0x2d, 0x17, 0x2e, 0x51,
	0x48, 0x26, 0x8d, 0xd5, 0xca, 0x0e, 0xb6, 0x53, 0x91, 0x37, 0xe0, 0x41, 0xb8, 0xf2, 0x2a, 0x3c,
	0x13, 0xf2, 0xd8, 0x6e, 0xcb, 0x46, 0x70, 0xd9, 0x1d, 0x7f, 0x33, 0xf6, 0xfc, 0x3e, 0x7b, 0x02,
	0x50, 0xf6, 0xa6, 0xcd, 0x3b, 0x25, 0x8d, 0x64, 0xab, 0x16, 0x07, 0xde, 0xa0, 0xb8, 0xca, 0x49,
	0xbc, 0x7d, 0x95, 0xfd, 0x88, 0x60, 0xf1, 0xa1, 0x37, 0xed, 0x0e, 0xbf, 0xf5, 0xa8, 0x0d, 0x3b,
	0x81, 0x43, 0x85, 0xba, 0x93, 0x42, 0x63, 0x61, 0x86, 0x0e, 0xd3, 0xe8, 0x28, 0x3a, 0x4d, 0x76,
	0xcb, 0x20, 0x5e, 0x0e, 0x1d, 0xb2, 0xe7, 0x90, 0x54, 0x37, 0x1c, 0x85, 0x29, 0x78, 0x9d, 0x4e,
	0xa8, 0xe0, 0xc0, 0x09, 0xdb, 0x9a, 0x1d, 0xc3, 0x52, 0x61, 0xcd, 0x15, 0x56, 0xa6, 0xe8, 0x15,
	0x4f, 0xa7, 0x94, 0x5f, 0x04, 0xed, 0xb3, 0xe2, 0xec, 0x19, 0xcc, 0x74, 0x25, 0x3b, 0x4c, 0x63,
	0xca, 0xb9, 0x45, 0xf6, 0x04, 0x96, 0x8e, 0xc4, 0x75, 0xca, 0x7e, 0x47, 0xf0, 0xf4, 0x1c, 0xcd,
	0xa5, 0xbc, 0x46, 0x11, 0xf0, 0xfe, 0xea, 0x1c, 0xed, 0x75, 0x3e, 0x81, 0x43, 0x9f, 0xd4, 0x58,
	0x29, 0x34, 0x1e, 0x6d, 0xe9, 0xc4, 0x0b, 0xd2, 0xd8, 0x06, 0xe0, 0x4a, 0x95, 0xc2, 0x38, 0x77,
	0x0e, 0x2e, 0x21, 0x85, 0xac, 0x31, 0x88, 0x2b, 0x59, 0x07, 0x32, 0x8a, 0x47, 0x8e, 0x66, 0x63,
	0x47, 0x74, 0x6d, 0x8d, 0x42, 0xdd, 0x16, 0xc6, 0xf2, 0xa6, 0xf3, 0x70, 0x6d, 0x24, 0x92, 0x87,
	0xec, 0xd7, 0x04, 0x12, 0x8a, 0xb6, 0xa2, 0x91, 0xf6, 0xd4, 0xb2, 0xaa, 0x50, 0x6b, 0xbf, 0xc3,
	0xb9, 0x59, 0x38, 0x8d, 0xca, 0x2c, 0x2b, 0xe5, 0x1c, 0xab, 0x73, 0x93, 0x90, 0x42, 0xac, 0x1b,
	0x00, 0xfc, 0xde, 0x71, 0x85, 0xba, 0xe0, 0x82, 0xac, 0x4c, 0x77, 0x89, 0x57, 0xb6, 0x62, 0xcc,
	0x14, 0x8f, 0x99, 0xee, 0x9f, 0x62, 0xf6, 0xe0, 0x29, 0xd8, 0x0a, 0xa6, 0x3d, 0xaf, 0xbd, 0x09,
	0x1b, 0xb2, 0x77, 0x10, 0x73, 0xd1, 0xc8, 0xf4, 0xf1, 0xd1, 0xf4, 0x74, 0x71, 0xf6, 0x32, 0xdf,
	0x1f, 0xa4, 0xfc, 0xce, 0x58, 0x6e, 0x7f, 0x3e, 0x09, 0xa3, 0x86, 0x1d, 0x6d, 0x59, 0xbf, 0x81,
	0xe4, 0x4e, 0xb2, 0x27, 0x5f, 0xe3, 0xe0, 0xcd, 0xda, 0xd0, 0x12, 0xdc, 0x96, 0x37, 0x7d, 0xf0,
	0xe7, 0x16, 0xef, 0x27, 0x6f, 0xa3, 0x8c, 0xc1, 0xea, 0xfe, 0xfd, 0xdd, 0x50, 0x9c, 0xfd, 0x8c,
	0x20, 0xb6, 0x53, 0xc2, 0xce, 0xfd, 0xff, 0x66, 0x8c, 0xf2, 0x60, 0x9e, 0xd7, 0x2f, 0xfe, 0x95,
	0xf6, 0x43, 0xf6, 0x88, 0x5d, 0xc0, 0x41, 0xe8, 0xc2, 0x8e, 0xc7, 0xd5, 0x7b, 0x13, 0xb8, 0xce,
	0xfe, 0x57, 0x12, 0x0e, 0xfd, 0x38, 0xff, 0x12, 0xdb, 0xec, 0xd7, 0x39, 0x7d, 0x77, 0xaf, 0xff,
	0x04, 0x00, 0x00, 0xff, 0xff, 0x32, 0x64, 0xda, 0xfc, 0x85, 0x03, 0x00, 0x00,
}
