// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
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

type Sex int32

const (
	Sex_UNSPECTED Sex = 0
	Sex_MALE      Sex = 1
	Sex_FEMALE    Sex = 2
	Sex_OTHER     Sex = 3
)

var Sex_name = map[int32]string{
	0: "UNSPECTED",
	1: "MALE",
	2: "FEMALE",
	3: "OTHER",
}

var Sex_value = map[string]int32{
	"UNSPECTED": 0,
	"MALE":      1,
	"FEMALE":    2,
	"OTHER":     3,
}

func (x Sex) String() string {
	return proto.EnumName(Sex_name, int32(x))
}

func (Sex) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

type GetUserRequest struct {
	UserId               string                `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FieldMask            *field_mask.FieldMask `protobuf:"bytes,2,opt,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *GetUserRequest) Reset()         { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()    {}
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *GetUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserRequest.Unmarshal(m, b)
}
func (m *GetUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserRequest.Marshal(b, m, deterministic)
}
func (m *GetUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserRequest.Merge(m, src)
}
func (m *GetUserRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserRequest.Size(m)
}
func (m *GetUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserRequest proto.InternalMessageInfo

func (m *GetUserRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *GetUserRequest) GetFieldMask() *field_mask.FieldMask {
	if m != nil {
		return m.FieldMask
	}
	return nil
}

type UserInfo struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Sex                  string   `protobuf:"bytes,2,opt,name=sex,proto3" json:"sex,omitempty"`
	Age                  int32    `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	Height               float32  `protobuf:"fixed32,4,opt,name=height,proto3" json:"height,omitempty"`
	Addr                 string   `protobuf:"bytes,5,opt,name=addr,proto3" json:"addr,omitempty"`
	Email                string   `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfo) Reset()         { *m = UserInfo{} }
func (m *UserInfo) String() string { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()    {}
func (*UserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *UserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfo.Unmarshal(m, b)
}
func (m *UserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfo.Marshal(b, m, deterministic)
}
func (m *UserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfo.Merge(m, src)
}
func (m *UserInfo) XXX_Size() int {
	return xxx_messageInfo_UserInfo.Size(m)
}
func (m *UserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfo proto.InternalMessageInfo

func (m *UserInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserInfo) GetSex() string {
	if m != nil {
		return m.Sex
	}
	return ""
}

func (m *UserInfo) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *UserInfo) GetHeight() float32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *UserInfo) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *UserInfo) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type SetUserRequest struct {
	UserId               string                `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FieldMask            *field_mask.FieldMask `protobuf:"bytes,2,opt,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *SetUserRequest) Reset()         { *m = SetUserRequest{} }
func (m *SetUserRequest) String() string { return proto.CompactTextString(m) }
func (*SetUserRequest) ProtoMessage()    {}
func (*SetUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *SetUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetUserRequest.Unmarshal(m, b)
}
func (m *SetUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetUserRequest.Marshal(b, m, deterministic)
}
func (m *SetUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetUserRequest.Merge(m, src)
}
func (m *SetUserRequest) XXX_Size() int {
	return xxx_messageInfo_SetUserRequest.Size(m)
}
func (m *SetUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetUserRequest proto.InternalMessageInfo

func (m *SetUserRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *SetUserRequest) GetFieldMask() *field_mask.FieldMask {
	if m != nil {
		return m.FieldMask
	}
	return nil
}

type SetUserResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetUserResponse) Reset()         { *m = SetUserResponse{} }
func (m *SetUserResponse) String() string { return proto.CompactTextString(m) }
func (*SetUserResponse) ProtoMessage()    {}
func (*SetUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *SetUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetUserResponse.Unmarshal(m, b)
}
func (m *SetUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetUserResponse.Marshal(b, m, deterministic)
}
func (m *SetUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetUserResponse.Merge(m, src)
}
func (m *SetUserResponse) XXX_Size() int {
	return xxx_messageInfo_SetUserResponse.Size(m)
}
func (m *SetUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetUserResponse proto.InternalMessageInfo

func (m *SetUserResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterEnum("heyifeng.user.v1.Sex", Sex_name, Sex_value)
	proto.RegisterType((*GetUserRequest)(nil), "heyifeng.user.v1.GetUserRequest")
	proto.RegisterType((*UserInfo)(nil), "heyifeng.user.v1.UserInfo")
	proto.RegisterType((*SetUserRequest)(nil), "heyifeng.user.v1.SetUserRequest")
	proto.RegisterType((*SetUserResponse)(nil), "heyifeng.user.v1.SetUserResponse")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 370 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x51, 0x4d, 0xaf, 0x93, 0x50,
	0x10, 0xed, 0x2d, 0x5f, 0x65, 0x8c, 0x4a, 0x26, 0x46, 0x09, 0x2b, 0x64, 0x45, 0x34, 0xe1, 0xc5,
	0x67, 0x5c, 0xb8, 0xf4, 0x83, 0xa7, 0x4d, 0xac, 0x36, 0x97, 0x76, 0xe3, 0xa6, 0xa1, 0x65, 0xa0,
	0xa4, 0x2d, 0x54, 0x2e, 0x98, 0xba, 0xf7, 0x9f, 0xf8, 0x47, 0xcd, 0xbd, 0x14, 0x4d, 0x35, 0xba,
	0x7b, 0xbb, 0x73, 0xe6, 0xe3, 0x9c, 0xcc, 0x1c, 0x80, 0x4e, 0x50, 0x13, 0x1d, 0x9b, 0xba, 0xad,
	0xd1, 0xd9, 0xd2, 0xb7, 0x32, 0xa7, 0xaa, 0x88, 0x54, 0xf1, 0xeb, 0x33, 0xcf, 0x2f, 0xea, 0xba,
	0xd8, 0xd3, 0x95, 0xea, 0xaf, 0xbb, 0xfc, 0x2a, 0x2f, 0x69, 0x9f, 0xad, 0x0e, 0xa9, 0xd8, 0xf5,
	0x3b, 0x41, 0x06, 0xf7, 0xde, 0x51, 0xbb, 0x14, 0xd4, 0x70, 0xfa, 0xd2, 0x91, 0x68, 0xf1, 0x11,
	0x58, 0x72, 0x7d, 0x55, 0x66, 0x2e, 0xf3, 0x59, 0x68, 0x73, 0x53, 0xd2, 0x69, 0x86, 0x2f, 0x01,
	0x7e, 0xaf, 0xbb, 0x63, 0x9f, 0x85, 0x77, 0xae, 0xbd, 0xa8, 0x77, 0x88, 0x06, 0x87, 0xe8, 0x46,
	0x8e, 0xcc, 0x52, 0xb1, 0xe3, 0x76, 0x3e, 0xc0, 0xe0, 0x3b, 0x83, 0x89, 0xf4, 0x98, 0x56, 0x79,
	0x8d, 0x08, 0x7a, 0x95, 0x1e, 0xe8, 0xac, 0xae, 0x30, 0x3a, 0xa0, 0x09, 0x3a, 0x29, 0x51, 0x9b,
	0x4b, 0x28, 0x2b, 0x69, 0x41, 0xae, 0xe6, 0xb3, 0xd0, 0xe0, 0x12, 0xe2, 0x43, 0x30, 0xb7, 0x54,
	0x16, 0xdb, 0xd6, 0xd5, 0x7d, 0x16, 0x8e, 0xf9, 0x99, 0x49, 0xbd, 0x34, 0xcb, 0x1a, 0xd7, 0xe8,
	0xf5, 0x24, 0xc6, 0x07, 0x60, 0xd0, 0x21, 0x2d, 0xf7, 0xae, 0xa9, 0x8a, 0x3d, 0x91, 0xc7, 0x26,
	0xb7, 0x7f, 0xec, 0x53, 0xb8, 0xff, 0xcb, 0x45, 0x1c, 0xeb, 0x4a, 0x10, 0xba, 0x60, 0x89, 0x6e,
	0xb3, 0x21, 0x21, 0x94, 0xcd, 0x84, 0x0f, 0xf4, 0xc9, 0x0b, 0xd0, 0x12, 0x3a, 0xe1, 0x5d, 0xb0,
	0x97, 0x1f, 0x93, 0x79, 0xfc, 0x66, 0x11, 0xbf, 0x75, 0x46, 0x38, 0x01, 0x7d, 0xf6, 0xea, 0x43,
	0xec, 0x30, 0x04, 0x30, 0x6f, 0x62, 0x85, 0xc7, 0x68, 0x83, 0xf1, 0x69, 0xf1, 0x3e, 0xe6, 0x8e,
	0x76, 0xfd, 0x83, 0x81, 0x2e, 0x1d, 0x70, 0x0a, 0xd6, 0x39, 0x3f, 0xf4, 0xa3, 0x3f, 0xf3, 0x8f,
	0x2e, 0xa3, 0xf5, 0xbc, 0xbf, 0x27, 0x86, 0x54, 0x82, 0x11, 0xce, 0xc1, 0x4a, 0xfe, 0x2d, 0x75,
	0xf9, 0x38, 0xef, 0xf1, 0x7f, 0x26, 0xfa, 0xa3, 0x83, 0xd1, 0x6b, 0xf3, 0xb3, 0x2e, 0x9b, 0x6b,
	0x53, 0x3d, 0xec, 0xf9, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9a, 0x32, 0x25, 0x29, 0xad, 0x02,
	0x00, 0x00,
}