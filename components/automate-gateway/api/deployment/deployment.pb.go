// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: automate-gateway/api/deployment/deployment.proto

package deployment

import (
	context "context"
	_ "github.com/chef/automate/components/automate-grpc/protoc-gen-policy/iam"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Version message
//
// The manifest version constructed with:
// * build_timestamp
type Version struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BuildTimestamp string `protobuf:"bytes,1,opt,name=build_timestamp,json=buildTimestamp,proto3" json:"build_timestamp,omitempty"`
}

func (x *Version) Reset() {
	*x = Version{}
	if protoimpl.UnsafeEnabled {
		mi := &file_automate_gateway_api_deployment_deployment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Version) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Version) ProtoMessage() {}

func (x *Version) ProtoReflect() protoreflect.Message {
	mi := &file_automate_gateway_api_deployment_deployment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Version.ProtoReflect.Descriptor instead.
func (*Version) Descriptor() ([]byte, []int) {
	return file_automate_gateway_api_deployment_deployment_proto_rawDescGZIP(), []int{0}
}

func (x *Version) GetBuildTimestamp() string {
	if x != nil {
		return x.BuildTimestamp
	}
	return ""
}

type ServiceVersionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ServiceVersionsRequest) Reset() {
	*x = ServiceVersionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_automate_gateway_api_deployment_deployment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceVersionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceVersionsRequest) ProtoMessage() {}

func (x *ServiceVersionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_automate_gateway_api_deployment_deployment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceVersionsRequest.ProtoReflect.Descriptor instead.
func (*ServiceVersionsRequest) Descriptor() ([]byte, []int) {
	return file_automate_gateway_api_deployment_deployment_proto_rawDescGZIP(), []int{1}
}

type ServiceVersionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Services []*ServiceVersion `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
}

func (x *ServiceVersionsResponse) Reset() {
	*x = ServiceVersionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_automate_gateway_api_deployment_deployment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceVersionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceVersionsResponse) ProtoMessage() {}

func (x *ServiceVersionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_automate_gateway_api_deployment_deployment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceVersionsResponse.ProtoReflect.Descriptor instead.
func (*ServiceVersionsResponse) Descriptor() ([]byte, []int) {
	return file_automate_gateway_api_deployment_deployment_proto_rawDescGZIP(), []int{2}
}

func (x *ServiceVersionsResponse) GetServices() []*ServiceVersion {
	if x != nil {
		return x.Services
	}
	return nil
}

type ServiceVersion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Origin  string `protobuf:"bytes,2,opt,name=origin,proto3" json:"origin,omitempty"`
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	Release string `protobuf:"bytes,4,opt,name=release,proto3" json:"release,omitempty"`
}

func (x *ServiceVersion) Reset() {
	*x = ServiceVersion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_automate_gateway_api_deployment_deployment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceVersion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceVersion) ProtoMessage() {}

func (x *ServiceVersion) ProtoReflect() protoreflect.Message {
	mi := &file_automate_gateway_api_deployment_deployment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceVersion.ProtoReflect.Descriptor instead.
func (*ServiceVersion) Descriptor() ([]byte, []int) {
	return file_automate_gateway_api_deployment_deployment_proto_rawDescGZIP(), []int{3}
}

func (x *ServiceVersion) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ServiceVersion) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

func (x *ServiceVersion) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ServiceVersion) GetRelease() string {
	if x != nil {
		return x.Release
	}
	return ""
}

var File_automate_gateway_api_deployment_deployment_proto protoreflect.FileDescriptor

var file_automate_gateway_api_deployment_deployment_proto_rawDesc = []byte{
	0x0a, 0x30, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1c, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x35, 0x61, 0x75, 0x74,
	0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x69, 0x61, 0x6d,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x32, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a,
	0x0f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x18, 0x0a, 0x16, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x63, 0x0a, 0x17, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x08, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e,
	0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0x70, 0x0a, 0x0e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a,
	0x07, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x32, 0x98, 0x03, 0x0a, 0x0a, 0x44, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x9f, 0x01, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x25, 0x2e,
	0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x22, 0x52, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x30, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x8a, 0xb5, 0x18,
	0x18, 0x0a, 0x16, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x3a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x3a, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x8a, 0xb5, 0x18, 0x1b, 0x12, 0x19, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x3a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x3a, 0x67, 0x65, 0x74, 0x12, 0xe7, 0x01, 0x0a, 0x0f, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x34, 0x2e, 0x63,
	0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x35, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61,
	0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x67, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x25, 0x12, 0x23, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x30, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x8a, 0xb5, 0x18, 0x18, 0x0a, 0x16, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x3a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3a, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x8a, 0xb5, 0x18, 0x1c, 0x12, 0x1a, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x3a, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x3a, 0x6c, 0x69,
	0x73, 0x74, 0x42, 0x45, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x68, 0x65, 0x66, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2f, 0x63,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61,
	0x74, 0x65, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_automate_gateway_api_deployment_deployment_proto_rawDescOnce sync.Once
	file_automate_gateway_api_deployment_deployment_proto_rawDescData = file_automate_gateway_api_deployment_deployment_proto_rawDesc
)

func file_automate_gateway_api_deployment_deployment_proto_rawDescGZIP() []byte {
	file_automate_gateway_api_deployment_deployment_proto_rawDescOnce.Do(func() {
		file_automate_gateway_api_deployment_deployment_proto_rawDescData = protoimpl.X.CompressGZIP(file_automate_gateway_api_deployment_deployment_proto_rawDescData)
	})
	return file_automate_gateway_api_deployment_deployment_proto_rawDescData
}

var file_automate_gateway_api_deployment_deployment_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_automate_gateway_api_deployment_deployment_proto_goTypes = []interface{}{
	(*Version)(nil),                 // 0: chef.automate.api.deployment.Version
	(*ServiceVersionsRequest)(nil),  // 1: chef.automate.api.deployment.ServiceVersionsRequest
	(*ServiceVersionsResponse)(nil), // 2: chef.automate.api.deployment.ServiceVersionsResponse
	(*ServiceVersion)(nil),          // 3: chef.automate.api.deployment.ServiceVersion
	(*empty.Empty)(nil),             // 4: google.protobuf.Empty
}
var file_automate_gateway_api_deployment_deployment_proto_depIdxs = []int32{
	3, // 0: chef.automate.api.deployment.ServiceVersionsResponse.services:type_name -> chef.automate.api.deployment.ServiceVersion
	4, // 1: chef.automate.api.deployment.Deployment.GetVersion:input_type -> google.protobuf.Empty
	1, // 2: chef.automate.api.deployment.Deployment.ServiceVersions:input_type -> chef.automate.api.deployment.ServiceVersionsRequest
	0, // 3: chef.automate.api.deployment.Deployment.GetVersion:output_type -> chef.automate.api.deployment.Version
	2, // 4: chef.automate.api.deployment.Deployment.ServiceVersions:output_type -> chef.automate.api.deployment.ServiceVersionsResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_automate_gateway_api_deployment_deployment_proto_init() }
func file_automate_gateway_api_deployment_deployment_proto_init() {
	if File_automate_gateway_api_deployment_deployment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_automate_gateway_api_deployment_deployment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Version); i {
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
		file_automate_gateway_api_deployment_deployment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceVersionsRequest); i {
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
		file_automate_gateway_api_deployment_deployment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceVersionsResponse); i {
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
		file_automate_gateway_api_deployment_deployment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceVersion); i {
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
			RawDescriptor: file_automate_gateway_api_deployment_deployment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_automate_gateway_api_deployment_deployment_proto_goTypes,
		DependencyIndexes: file_automate_gateway_api_deployment_deployment_proto_depIdxs,
		MessageInfos:      file_automate_gateway_api_deployment_deployment_proto_msgTypes,
	}.Build()
	File_automate_gateway_api_deployment_deployment_proto = out.File
	file_automate_gateway_api_deployment_deployment_proto_rawDesc = nil
	file_automate_gateway_api_deployment_deployment_proto_goTypes = nil
	file_automate_gateway_api_deployment_deployment_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DeploymentClient is the client API for Deployment service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DeploymentClient interface {
	GetVersion(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Version, error)
	ServiceVersions(ctx context.Context, in *ServiceVersionsRequest, opts ...grpc.CallOption) (*ServiceVersionsResponse, error)
}

type deploymentClient struct {
	cc grpc.ClientConnInterface
}

func NewDeploymentClient(cc grpc.ClientConnInterface) DeploymentClient {
	return &deploymentClient{cc}
}

func (c *deploymentClient) GetVersion(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Version, error) {
	out := new(Version)
	err := c.cc.Invoke(ctx, "/chef.automate.api.deployment.Deployment/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deploymentClient) ServiceVersions(ctx context.Context, in *ServiceVersionsRequest, opts ...grpc.CallOption) (*ServiceVersionsResponse, error) {
	out := new(ServiceVersionsResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.api.deployment.Deployment/ServiceVersions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeploymentServer is the server API for Deployment service.
type DeploymentServer interface {
	GetVersion(context.Context, *empty.Empty) (*Version, error)
	ServiceVersions(context.Context, *ServiceVersionsRequest) (*ServiceVersionsResponse, error)
}

// UnimplementedDeploymentServer can be embedded to have forward compatible implementations.
type UnimplementedDeploymentServer struct {
}

func (*UnimplementedDeploymentServer) GetVersion(context.Context, *empty.Empty) (*Version, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVersion not implemented")
}
func (*UnimplementedDeploymentServer) ServiceVersions(context.Context, *ServiceVersionsRequest) (*ServiceVersionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServiceVersions not implemented")
}

func RegisterDeploymentServer(s *grpc.Server, srv DeploymentServer) {
	s.RegisterService(&_Deployment_serviceDesc, srv)
}

func _Deployment_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeploymentServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.deployment.Deployment/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeploymentServer).GetVersion(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Deployment_ServiceVersions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceVersionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeploymentServer).ServiceVersions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.deployment.Deployment/ServiceVersions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeploymentServer).ServiceVersions(ctx, req.(*ServiceVersionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Deployment_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.api.deployment.Deployment",
	HandlerType: (*DeploymentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVersion",
			Handler:    _Deployment_GetVersion_Handler,
		},
		{
			MethodName: "ServiceVersions",
			Handler:    _Deployment_ServiceVersions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "automate-gateway/api/deployment/deployment.proto",
}
