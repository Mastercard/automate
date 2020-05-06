// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/interservice/cfgmgmt/request/stats.proto

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

type NodesCounts struct {
	Filter               []string `protobuf:"bytes,1,rep,name=filter,proto3" json:"filter,omitempty" toml:"filter,omitempty" mapstructure:"filter,omitempty"`
	Start                string   `protobuf:"bytes,2,opt,name=start,proto3" json:"start,omitempty" toml:"start,omitempty" mapstructure:"start,omitempty"`
	End                  string   `protobuf:"bytes,3,opt,name=end,proto3" json:"end,omitempty" toml:"end,omitempty" mapstructure:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *NodesCounts) Reset()         { *m = NodesCounts{} }
func (m *NodesCounts) String() string { return proto.CompactTextString(m) }
func (*NodesCounts) ProtoMessage()    {}
func (*NodesCounts) Descriptor() ([]byte, []int) {
	return fileDescriptor_49f61da66790a24b, []int{0}
}

func (m *NodesCounts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodesCounts.Unmarshal(m, b)
}
func (m *NodesCounts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodesCounts.Marshal(b, m, deterministic)
}
func (m *NodesCounts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodesCounts.Merge(m, src)
}
func (m *NodesCounts) XXX_Size() int {
	return xxx_messageInfo_NodesCounts.Size(m)
}
func (m *NodesCounts) XXX_DiscardUnknown() {
	xxx_messageInfo_NodesCounts.DiscardUnknown(m)
}

var xxx_messageInfo_NodesCounts proto.InternalMessageInfo

func (m *NodesCounts) GetFilter() []string {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *NodesCounts) GetStart() string {
	if m != nil {
		return m.Start
	}
	return ""
}

func (m *NodesCounts) GetEnd() string {
	if m != nil {
		return m.End
	}
	return ""
}

type RunsCounts struct {
	Filter               []string `protobuf:"bytes,1,rep,name=filter,proto3" json:"filter,omitempty" toml:"filter,omitempty" mapstructure:"filter,omitempty"`
	Start                string   `protobuf:"bytes,2,opt,name=start,proto3" json:"start,omitempty" toml:"start,omitempty" mapstructure:"start,omitempty"`
	End                  string   `protobuf:"bytes,3,opt,name=end,proto3" json:"end,omitempty" toml:"end,omitempty" mapstructure:"end,omitempty"`
	NodeId               string   `protobuf:"bytes,4,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty" toml:"node_id,omitempty" mapstructure:"node_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *RunsCounts) Reset()         { *m = RunsCounts{} }
func (m *RunsCounts) String() string { return proto.CompactTextString(m) }
func (*RunsCounts) ProtoMessage()    {}
func (*RunsCounts) Descriptor() ([]byte, []int) {
	return fileDescriptor_49f61da66790a24b, []int{1}
}

func (m *RunsCounts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunsCounts.Unmarshal(m, b)
}
func (m *RunsCounts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunsCounts.Marshal(b, m, deterministic)
}
func (m *RunsCounts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunsCounts.Merge(m, src)
}
func (m *RunsCounts) XXX_Size() int {
	return xxx_messageInfo_RunsCounts.Size(m)
}
func (m *RunsCounts) XXX_DiscardUnknown() {
	xxx_messageInfo_RunsCounts.DiscardUnknown(m)
}

var xxx_messageInfo_RunsCounts proto.InternalMessageInfo

func (m *RunsCounts) GetFilter() []string {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *RunsCounts) GetStart() string {
	if m != nil {
		return m.Start
	}
	return ""
}

func (m *RunsCounts) GetEnd() string {
	if m != nil {
		return m.End
	}
	return ""
}

func (m *RunsCounts) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

type CheckInCountsTimeSeries struct {
	// List of filters to be applied to the time series.
	Filter []string `protobuf:"bytes,1,rep,name=filter,proto3" json:"filter,omitempty" toml:"filter,omitempty" mapstructure:"filter,omitempty"`
	// Number of past days to create the time series
	DaysAgo              int32    `protobuf:"varint,2,opt,name=days_ago,json=daysAgo,proto3" json:"days_ago,omitempty" toml:"days_ago,omitempty" mapstructure:"days_ago,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *CheckInCountsTimeSeries) Reset()         { *m = CheckInCountsTimeSeries{} }
func (m *CheckInCountsTimeSeries) String() string { return proto.CompactTextString(m) }
func (*CheckInCountsTimeSeries) ProtoMessage()    {}
func (*CheckInCountsTimeSeries) Descriptor() ([]byte, []int) {
	return fileDescriptor_49f61da66790a24b, []int{2}
}

func (m *CheckInCountsTimeSeries) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckInCountsTimeSeries.Unmarshal(m, b)
}
func (m *CheckInCountsTimeSeries) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckInCountsTimeSeries.Marshal(b, m, deterministic)
}
func (m *CheckInCountsTimeSeries) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckInCountsTimeSeries.Merge(m, src)
}
func (m *CheckInCountsTimeSeries) XXX_Size() int {
	return xxx_messageInfo_CheckInCountsTimeSeries.Size(m)
}
func (m *CheckInCountsTimeSeries) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckInCountsTimeSeries.DiscardUnknown(m)
}

var xxx_messageInfo_CheckInCountsTimeSeries proto.InternalMessageInfo

func (m *CheckInCountsTimeSeries) GetFilter() []string {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *CheckInCountsTimeSeries) GetDaysAgo() int32 {
	if m != nil {
		return m.DaysAgo
	}
	return 0
}

func init() {
	proto.RegisterType((*NodesCounts)(nil), "chef.automate.domain.cfgmgmt.request.NodesCounts")
	proto.RegisterType((*RunsCounts)(nil), "chef.automate.domain.cfgmgmt.request.RunsCounts")
	proto.RegisterType((*CheckInCountsTimeSeries)(nil), "chef.automate.domain.cfgmgmt.request.CheckInCountsTimeSeries")
}

func init() {
	proto.RegisterFile("api/interservice/cfgmgmt/request/stats.proto", fileDescriptor_49f61da66790a24b)
}

var fileDescriptor_49f61da66790a24b = []byte{
	// 254 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x90, 0xb1, 0x4f, 0xf3, 0x30,
	0x10, 0xc5, 0xd5, 0x2f, 0x5f, 0x53, 0x7a, 0x2c, 0x28, 0x42, 0x34, 0x6c, 0x55, 0xc4, 0xd0, 0x01,
	0xd9, 0x03, 0x13, 0x62, 0x82, 0x4e, 0x95, 0x80, 0x21, 0x30, 0xb1, 0x54, 0x6e, 0x7c, 0x71, 0x2c,
	0xb0, 0xaf, 0xd8, 0x67, 0x24, 0xfe, 0x7b, 0x94, 0x34, 0x8c, 0x88, 0x85, 0xed, 0x7e, 0x4f, 0x7a,
	0xf7, 0xa4, 0x1f, 0x5c, 0xaa, 0xbd, 0x95, 0xd6, 0x33, 0x86, 0x88, 0xe1, 0xc3, 0x36, 0x28, 0x9b,
	0xd6, 0x38, 0xe3, 0x58, 0x06, 0x7c, 0x4f, 0x18, 0x59, 0x46, 0x56, 0x1c, 0xc5, 0x3e, 0x10, 0x53,
	0x71, 0xd1, 0x74, 0xd8, 0x0a, 0x95, 0x98, 0x9c, 0x62, 0x14, 0x9a, 0x9c, 0xb2, 0x5e, 0x8c, 0x0d,
	0x31, 0x36, 0xaa, 0x07, 0x38, 0x7e, 0x24, 0x8d, 0x71, 0x4d, 0xc9, 0x73, 0x2c, 0xce, 0x20, 0x6f,
	0xed, 0x1b, 0x63, 0x28, 0x27, 0xcb, 0x6c, 0x35, 0xaf, 0x47, 0x2a, 0x4e, 0x61, 0x1a, 0x59, 0x05,
	0x2e, 0xff, 0x2d, 0x27, 0xab, 0x79, 0x7d, 0x80, 0xe2, 0x04, 0x32, 0xf4, 0xba, 0xcc, 0x86, 0xac,
	0x3f, 0x2b, 0x04, 0xa8, 0x93, 0xff, 0xa3, 0x6f, 0xc5, 0x02, 0x66, 0x9e, 0x34, 0x6e, 0xad, 0x2e,
	0xff, 0x0f, 0x69, 0xde, 0xe3, 0x46, 0x57, 0xf7, 0xb0, 0x58, 0x77, 0xd8, 0xbc, 0x6e, 0xfc, 0x61,
	0xe9, 0xd9, 0x3a, 0x7c, 0xc2, 0x60, 0xf1, 0xe7, 0xcd, 0x73, 0x38, 0xd2, 0xea, 0x33, 0x6e, 0x95,
	0xa1, 0x61, 0x76, 0x5a, 0xcf, 0x7a, 0xbe, 0x35, 0x74, 0x77, 0xf3, 0x72, 0x6d, 0x2c, 0x77, 0x69,
	0x27, 0x1a, 0x72, 0xb2, 0xd7, 0x26, 0xbf, 0xb5, 0xc9, 0xdf, 0x94, 0xef, 0xf2, 0xc1, 0xf6, 0xd5,
	0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb6, 0x07, 0x8f, 0xbf, 0x9d, 0x01, 0x00, 0x00,
}
