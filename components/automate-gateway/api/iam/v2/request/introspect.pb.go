// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/automate-gateway/api/iam/v2/request/introspect.proto

package request

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

type IntrospectAllReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IntrospectAllReq) Reset()         { *m = IntrospectAllReq{} }
func (m *IntrospectAllReq) String() string { return proto.CompactTextString(m) }
func (*IntrospectAllReq) ProtoMessage()    {}
func (*IntrospectAllReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_51df080d17bfc9ba, []int{0}
}

func (m *IntrospectAllReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntrospectAllReq.Unmarshal(m, b)
}
func (m *IntrospectAllReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntrospectAllReq.Marshal(b, m, deterministic)
}
func (m *IntrospectAllReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntrospectAllReq.Merge(m, src)
}
func (m *IntrospectAllReq) XXX_Size() int {
	return xxx_messageInfo_IntrospectAllReq.Size(m)
}
func (m *IntrospectAllReq) XXX_DiscardUnknown() {
	xxx_messageInfo_IntrospectAllReq.DiscardUnknown(m)
}

var xxx_messageInfo_IntrospectAllReq proto.InternalMessageInfo

type IntrospectSomeReq struct {
	Paths                []string `protobuf:"bytes,1,rep,name=paths,proto3" json:"paths,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IntrospectSomeReq) Reset()         { *m = IntrospectSomeReq{} }
func (m *IntrospectSomeReq) String() string { return proto.CompactTextString(m) }
func (*IntrospectSomeReq) ProtoMessage()    {}
func (*IntrospectSomeReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_51df080d17bfc9ba, []int{1}
}

func (m *IntrospectSomeReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntrospectSomeReq.Unmarshal(m, b)
}
func (m *IntrospectSomeReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntrospectSomeReq.Marshal(b, m, deterministic)
}
func (m *IntrospectSomeReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntrospectSomeReq.Merge(m, src)
}
func (m *IntrospectSomeReq) XXX_Size() int {
	return xxx_messageInfo_IntrospectSomeReq.Size(m)
}
func (m *IntrospectSomeReq) XXX_DiscardUnknown() {
	xxx_messageInfo_IntrospectSomeReq.DiscardUnknown(m)
}

var xxx_messageInfo_IntrospectSomeReq proto.InternalMessageInfo

func (m *IntrospectSomeReq) GetPaths() []string {
	if m != nil {
		return m.Paths
	}
	return nil
}

type IntrospectReq struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Parameters           []string `protobuf:"bytes,2,rep,name=parameters,proto3" json:"parameters,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IntrospectReq) Reset()         { *m = IntrospectReq{} }
func (m *IntrospectReq) String() string { return proto.CompactTextString(m) }
func (*IntrospectReq) ProtoMessage()    {}
func (*IntrospectReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_51df080d17bfc9ba, []int{2}
}

func (m *IntrospectReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntrospectReq.Unmarshal(m, b)
}
func (m *IntrospectReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntrospectReq.Marshal(b, m, deterministic)
}
func (m *IntrospectReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntrospectReq.Merge(m, src)
}
func (m *IntrospectReq) XXX_Size() int {
	return xxx_messageInfo_IntrospectReq.Size(m)
}
func (m *IntrospectReq) XXX_DiscardUnknown() {
	xxx_messageInfo_IntrospectReq.DiscardUnknown(m)
}

var xxx_messageInfo_IntrospectReq proto.InternalMessageInfo

func (m *IntrospectReq) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *IntrospectReq) GetParameters() []string {
	if m != nil {
		return m.Parameters
	}
	return nil
}

func init() {
	proto.RegisterType((*IntrospectAllReq)(nil), "chef.automate.api.iam.v2.IntrospectAllReq")
	proto.RegisterType((*IntrospectSomeReq)(nil), "chef.automate.api.iam.v2.IntrospectSomeReq")
	proto.RegisterType((*IntrospectReq)(nil), "chef.automate.api.iam.v2.IntrospectReq")
}

func init() {
	proto.RegisterFile("components/automate-gateway/api/iam/v2/request/introspect.proto", fileDescriptor_51df080d17bfc9ba)
}

var fileDescriptor_51df080d17bfc9ba = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x8f, 0x3d, 0x4b, 0x03, 0x41,
	0x10, 0x86, 0x39, 0xbf, 0x20, 0x03, 0x82, 0x2e, 0x16, 0x57, 0x49, 0xb8, 0x2a, 0x16, 0xee, 0x40,
	0xfc, 0x01, 0xa2, 0x16, 0x92, 0x76, 0xed, 0xec, 0x26, 0xcb, 0x98, 0x5b, 0xc8, 0xde, 0xec, 0xed,
	0xce, 0x45, 0xfc, 0xf7, 0x72, 0x27, 0xe6, 0x6c, 0xed, 0x66, 0xde, 0x79, 0x5e, 0x98, 0x07, 0x1e,
	0xbd, 0xc4, 0x24, 0x1d, 0x77, 0x5a, 0x90, 0x06, 0x95, 0x48, 0xca, 0xf7, 0x3b, 0x52, 0xfe, 0xa4,
	0x2f, 0xa4, 0x14, 0x30, 0x50, 0xc4, 0xc3, 0x1a, 0x33, 0xf7, 0x03, 0x17, 0xc5, 0xd0, 0x69, 0x96,
	0x92, 0xd8, 0xab, 0x4d, 0x59, 0x54, 0x4c, 0xed, 0x5b, 0xfe, 0xb0, 0xbf, 0x55, 0x4b, 0x29, 0xd8,
	0x40, 0xd1, 0x1e, 0xd6, 0x8d, 0x81, 0xab, 0xcd, 0x91, 0x7e, 0xda, 0xef, 0x1d, 0xf7, 0xcd, 0x1d,
	0x5c, 0xcf, 0xd9, 0x9b, 0x44, 0x76, 0xdc, 0x9b, 0x1b, 0x38, 0x4f, 0xa4, 0x6d, 0xa9, 0xab, 0xe5,
	0xe9, 0x6a, 0xe1, 0x7e, 0x96, 0xe6, 0x05, 0x2e, 0x67, 0x74, 0xc4, 0x0c, 0x9c, 0x8d, 0x97, 0xba,
	0x5a, 0x56, 0xab, 0x85, 0x9b, 0x66, 0x73, 0x0b, 0x90, 0x28, 0x53, 0x64, 0xe5, 0x5c, 0xea, 0x93,
	0xa9, 0xff, 0x27, 0x79, 0xde, 0xbc, 0xbf, 0xee, 0x82, 0xb6, 0xc3, 0xd6, 0x7a, 0x89, 0x38, 0xbe,
	0x7a, 0xb4, 0xc4, 0xff, 0x99, 0x6f, 0x2f, 0x26, 0xdf, 0x87, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x08, 0x66, 0xea, 0xad, 0x32, 0x01, 0x00, 0x00,
}