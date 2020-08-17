// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: external/infra_proxy/request/environments.proto

package request

import (
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
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

type Environments struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Chef organization ID.
	OrgId string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	// Chef Infra Server ID.
	ServerId string `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
}

func (x *Environments) Reset() {
	*x = Environments{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_infra_proxy_request_environments_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Environments) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Environments) ProtoMessage() {}

func (x *Environments) ProtoReflect() protoreflect.Message {
	mi := &file_external_infra_proxy_request_environments_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Environments.ProtoReflect.Descriptor instead.
func (*Environments) Descriptor() ([]byte, []int) {
	return file_external_infra_proxy_request_environments_proto_rawDescGZIP(), []int{0}
}

func (x *Environments) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *Environments) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

type Environment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Chef organization ID.
	OrgId string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	// Chef Infra Server ID.
	ServerId string `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	// Environment name.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Environment) Reset() {
	*x = Environment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_infra_proxy_request_environments_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Environment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Environment) ProtoMessage() {}

func (x *Environment) ProtoReflect() protoreflect.Message {
	mi := &file_external_infra_proxy_request_environments_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Environment.ProtoReflect.Descriptor instead.
func (*Environment) Descriptor() ([]byte, []int) {
	return file_external_infra_proxy_request_environments_proto_rawDescGZIP(), []int{1}
}

func (x *Environment) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *Environment) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (x *Environment) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateEnvironment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Chef organization ID.
	OrgId string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	// Chef Infra Server ID.
	ServerId string `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	// Environment name.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Environment description.
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	// Class name.
	JsonClass string `protobuf:"bytes,5,opt,name=json_class,json=jsonClass,proto3" json:"json_class,omitempty"`
	// Environment versioned cookbooks constraints.
	CookbookVersions map[string]string `protobuf:"bytes,6,rep,name=cookbook_versions,json=cookbookVersions,proto3" json:"cookbook_versions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Environment default attributes JSON.
	DefaultAttributes *_struct.Struct `protobuf:"bytes,7,opt,name=default_attributes,json=defaultAttributes,proto3" json:"default_attributes,omitempty"`
	// Environment override attributes JSON.
	OverrideAttributes *_struct.Struct `protobuf:"bytes,8,opt,name=override_attributes,json=overrideAttributes,proto3" json:"override_attributes,omitempty"`
}

func (x *CreateEnvironment) Reset() {
	*x = CreateEnvironment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_infra_proxy_request_environments_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEnvironment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEnvironment) ProtoMessage() {}

func (x *CreateEnvironment) ProtoReflect() protoreflect.Message {
	mi := &file_external_infra_proxy_request_environments_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEnvironment.ProtoReflect.Descriptor instead.
func (*CreateEnvironment) Descriptor() ([]byte, []int) {
	return file_external_infra_proxy_request_environments_proto_rawDescGZIP(), []int{2}
}

func (x *CreateEnvironment) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *CreateEnvironment) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (x *CreateEnvironment) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateEnvironment) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateEnvironment) GetJsonClass() string {
	if x != nil {
		return x.JsonClass
	}
	return ""
}

func (x *CreateEnvironment) GetCookbookVersions() map[string]string {
	if x != nil {
		return x.CookbookVersions
	}
	return nil
}

func (x *CreateEnvironment) GetDefaultAttributes() *_struct.Struct {
	if x != nil {
		return x.DefaultAttributes
	}
	return nil
}

func (x *CreateEnvironment) GetOverrideAttributes() *_struct.Struct {
	if x != nil {
		return x.OverrideAttributes
	}
	return nil
}

type UpdateEnvironment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Chef organization ID.
	OrgId string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	// Chef Infra Server ID.
	ServerId string `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	// Environment name.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Environment description.
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	// Class name.
	JsonClass string `protobuf:"bytes,5,opt,name=json_class,json=jsonClass,proto3" json:"json_class,omitempty"`
	// Environment versioned cookbooks constraints.
	CookbookVersions map[string]string `protobuf:"bytes,6,rep,name=cookbook_versions,json=cookbookVersions,proto3" json:"cookbook_versions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Environment default attributes JSON.
	DefaultAttributes *_struct.Struct `protobuf:"bytes,7,opt,name=default_attributes,json=defaultAttributes,proto3" json:"default_attributes,omitempty"`
	// Environment override attributes JSON.
	OverrideAttributes *_struct.Struct `protobuf:"bytes,8,opt,name=override_attributes,json=overrideAttributes,proto3" json:"override_attributes,omitempty"`
}

func (x *UpdateEnvironment) Reset() {
	*x = UpdateEnvironment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_infra_proxy_request_environments_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEnvironment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEnvironment) ProtoMessage() {}

func (x *UpdateEnvironment) ProtoReflect() protoreflect.Message {
	mi := &file_external_infra_proxy_request_environments_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEnvironment.ProtoReflect.Descriptor instead.
func (*UpdateEnvironment) Descriptor() ([]byte, []int) {
	return file_external_infra_proxy_request_environments_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateEnvironment) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *UpdateEnvironment) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (x *UpdateEnvironment) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateEnvironment) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateEnvironment) GetJsonClass() string {
	if x != nil {
		return x.JsonClass
	}
	return ""
}

func (x *UpdateEnvironment) GetCookbookVersions() map[string]string {
	if x != nil {
		return x.CookbookVersions
	}
	return nil
}

func (x *UpdateEnvironment) GetDefaultAttributes() *_struct.Struct {
	if x != nil {
		return x.DefaultAttributes
	}
	return nil
}

func (x *UpdateEnvironment) GetOverrideAttributes() *_struct.Struct {
	if x != nil {
		return x.OverrideAttributes
	}
	return nil
}

var File_external_infra_proxy_request_environments_proto protoreflect.FileDescriptor

var file_external_infra_proxy_request_environments_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2f, 0x65,
	0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x25, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x42, 0x0a, 0x0c, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x15, 0x0a, 0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x22, 0x55, 0x0a, 0x0b, 0x45, 0x6e,
	0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x6f, 0x72, 0x67,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0xf0, 0x03, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x76, 0x69,
	0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6a, 0x73, 0x6f, 0x6e, 0x43, 0x6c, 0x61, 0x73, 0x73,
	0x12, 0x7b, 0x0a, 0x11, 0x63, 0x6f, 0x6f, 0x6b, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x4e, 0x2e, 0x63, 0x68,
	0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6f, 0x6b, 0x62, 0x6f, 0x6f, 0x6b, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x10, 0x63, 0x6f, 0x6f,
	0x6b, 0x62, 0x6f, 0x6f, 0x6b, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x46, 0x0a,
	0x12, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x52, 0x11, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x48, 0x0a, 0x13, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64,
	0x65, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x12, 0x6f, 0x76, 0x65,
	0x72, 0x72, 0x69, 0x64, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x1a,
	0x43, 0x0a, 0x15, 0x43, 0x6f, 0x6f, 0x6b, 0x62, 0x6f, 0x6f, 0x6b, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0xf0, 0x03, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45,
	0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x6f, 0x72,
	0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x63, 0x6c, 0x61,
	0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6a, 0x73, 0x6f, 0x6e, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x12, 0x7b, 0x0a, 0x11, 0x63, 0x6f, 0x6f, 0x6b, 0x62, 0x6f, 0x6f, 0x6b, 0x5f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x4e,
	0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x76,
	0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6f, 0x6b, 0x62, 0x6f, 0x6f,
	0x6b, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x10,
	0x63, 0x6f, 0x6f, 0x6b, 0x62, 0x6f, 0x6f, 0x6b, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x46, 0x0a, 0x12, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x61, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x11, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x41, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x48, 0x0a, 0x13, 0x6f, 0x76, 0x65, 0x72,
	0x72, 0x69, 0x64, 0x65, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x12,
	0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x73, 0x1a, 0x43, 0x0a, 0x15, 0x43, 0x6f, 0x6f, 0x6b, 0x62, 0x6f, 0x6f, 0x6b, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d,
	0x61, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_external_infra_proxy_request_environments_proto_rawDescOnce sync.Once
	file_external_infra_proxy_request_environments_proto_rawDescData = file_external_infra_proxy_request_environments_proto_rawDesc
)

func file_external_infra_proxy_request_environments_proto_rawDescGZIP() []byte {
	file_external_infra_proxy_request_environments_proto_rawDescOnce.Do(func() {
		file_external_infra_proxy_request_environments_proto_rawDescData = protoimpl.X.CompressGZIP(file_external_infra_proxy_request_environments_proto_rawDescData)
	})
	return file_external_infra_proxy_request_environments_proto_rawDescData
}

var file_external_infra_proxy_request_environments_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_external_infra_proxy_request_environments_proto_goTypes = []interface{}{
	(*Environments)(nil),      // 0: chef.automate.api.infra_proxy.request.Environments
	(*Environment)(nil),       // 1: chef.automate.api.infra_proxy.request.Environment
	(*CreateEnvironment)(nil), // 2: chef.automate.api.infra_proxy.request.CreateEnvironment
	(*UpdateEnvironment)(nil), // 3: chef.automate.api.infra_proxy.request.UpdateEnvironment
	nil,                       // 4: chef.automate.api.infra_proxy.request.CreateEnvironment.CookbookVersionsEntry
	nil,                       // 5: chef.automate.api.infra_proxy.request.UpdateEnvironment.CookbookVersionsEntry
	(*_struct.Struct)(nil),    // 6: google.protobuf.Struct
}
var file_external_infra_proxy_request_environments_proto_depIdxs = []int32{
	4, // 0: chef.automate.api.infra_proxy.request.CreateEnvironment.cookbook_versions:type_name -> chef.automate.api.infra_proxy.request.CreateEnvironment.CookbookVersionsEntry
	6, // 1: chef.automate.api.infra_proxy.request.CreateEnvironment.default_attributes:type_name -> google.protobuf.Struct
	6, // 2: chef.automate.api.infra_proxy.request.CreateEnvironment.override_attributes:type_name -> google.protobuf.Struct
	5, // 3: chef.automate.api.infra_proxy.request.UpdateEnvironment.cookbook_versions:type_name -> chef.automate.api.infra_proxy.request.UpdateEnvironment.CookbookVersionsEntry
	6, // 4: chef.automate.api.infra_proxy.request.UpdateEnvironment.default_attributes:type_name -> google.protobuf.Struct
	6, // 5: chef.automate.api.infra_proxy.request.UpdateEnvironment.override_attributes:type_name -> google.protobuf.Struct
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_external_infra_proxy_request_environments_proto_init() }
func file_external_infra_proxy_request_environments_proto_init() {
	if File_external_infra_proxy_request_environments_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_external_infra_proxy_request_environments_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Environments); i {
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
		file_external_infra_proxy_request_environments_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Environment); i {
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
		file_external_infra_proxy_request_environments_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEnvironment); i {
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
		file_external_infra_proxy_request_environments_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEnvironment); i {
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
			RawDescriptor: file_external_infra_proxy_request_environments_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_external_infra_proxy_request_environments_proto_goTypes,
		DependencyIndexes: file_external_infra_proxy_request_environments_proto_depIdxs,
		MessageInfos:      file_external_infra_proxy_request_environments_proto_msgTypes,
	}.Build()
	File_external_infra_proxy_request_environments_proto = out.File
	file_external_infra_proxy_request_environments_proto_rawDesc = nil
	file_external_infra_proxy_request_environments_proto_goTypes = nil
	file_external_infra_proxy_request_environments_proto_depIdxs = nil
}
