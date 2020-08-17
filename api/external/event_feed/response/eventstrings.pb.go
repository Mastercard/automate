// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: external/event_feed/response/eventstrings.proto

package response

import (
	proto "github.com/golang/protobuf/proto"
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

type EventStrings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Strings      []*EventString `protobuf:"bytes,1,rep,name=strings,proto3" json:"strings,omitempty"`
	Start        string         `protobuf:"bytes,2,opt,name=start,proto3" json:"start,omitempty"`
	End          string         `protobuf:"bytes,3,opt,name=end,proto3" json:"end,omitempty"`
	HoursBetween int32          `protobuf:"varint,4,opt,name=hours_between,json=hoursBetween,proto3" json:"hours_between,omitempty"`
}

func (x *EventStrings) Reset() {
	*x = EventStrings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_event_feed_response_eventstrings_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventStrings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventStrings) ProtoMessage() {}

func (x *EventStrings) ProtoReflect() protoreflect.Message {
	mi := &file_external_event_feed_response_eventstrings_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventStrings.ProtoReflect.Descriptor instead.
func (*EventStrings) Descriptor() ([]byte, []int) {
	return file_external_event_feed_response_eventstrings_proto_rawDescGZIP(), []int{0}
}

func (x *EventStrings) GetStrings() []*EventString {
	if x != nil {
		return x.Strings
	}
	return nil
}

func (x *EventStrings) GetStart() string {
	if x != nil {
		return x.Start
	}
	return ""
}

func (x *EventStrings) GetEnd() string {
	if x != nil {
		return x.End
	}
	return ""
}

func (x *EventStrings) GetHoursBetween() int32 {
	if x != nil {
		return x.HoursBetween
	}
	return 0
}

type EventString struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Collection  []*EventCollection `protobuf:"bytes,1,rep,name=collection,proto3" json:"collection,omitempty"`
	EventAction string             `protobuf:"bytes,2,opt,name=event_action,json=eventAction,proto3" json:"event_action,omitempty"`
}

func (x *EventString) Reset() {
	*x = EventString{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_event_feed_response_eventstrings_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventString) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventString) ProtoMessage() {}

func (x *EventString) ProtoReflect() protoreflect.Message {
	mi := &file_external_event_feed_response_eventstrings_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventString.ProtoReflect.Descriptor instead.
func (*EventString) Descriptor() ([]byte, []int) {
	return file_external_event_feed_response_eventstrings_proto_rawDescGZIP(), []int{1}
}

func (x *EventString) GetCollection() []*EventCollection {
	if x != nil {
		return x.Collection
	}
	return nil
}

func (x *EventString) GetEventAction() string {
	if x != nil {
		return x.EventAction
	}
	return ""
}

type EventCollection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventsCount []*EventCount `protobuf:"bytes,1,rep,name=events_count,json=eventsCount,proto3" json:"events_count,omitempty"`
}

func (x *EventCollection) Reset() {
	*x = EventCollection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_event_feed_response_eventstrings_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventCollection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventCollection) ProtoMessage() {}

func (x *EventCollection) ProtoReflect() protoreflect.Message {
	mi := &file_external_event_feed_response_eventstrings_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventCollection.ProtoReflect.Descriptor instead.
func (*EventCollection) Descriptor() ([]byte, []int) {
	return file_external_event_feed_response_eventstrings_proto_rawDescGZIP(), []int{2}
}

func (x *EventCollection) GetEventsCount() []*EventCount {
	if x != nil {
		return x.EventsCount
	}
	return nil
}

var File_external_event_feed_response_eventstrings_proto protoreflect.FileDescriptor

var file_external_event_feed_response_eventstrings_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x5f, 0x66, 0x65, 0x65, 0x64, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x25, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x2e,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x28, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x2f, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xa9, 0x01, 0x0a, 0x0c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x73, 0x12, 0x4c, 0x0a, 0x07, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f,
	0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x66,
	0x65, 0x65, 0x64, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x68, 0x6f, 0x75,
	0x72, 0x73, 0x5f, 0x62, 0x65, 0x74, 0x77, 0x65, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0c, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x42, 0x65, 0x74, 0x77, 0x65, 0x65, 0x6e, 0x22, 0x88,
	0x01, 0x0a, 0x0b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x56,
	0x0a, 0x0a, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x36, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61,
	0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x66, 0x65, 0x65,
	0x64, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x63, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x67, 0x0a, 0x0f, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x54, 0x0a, 0x0c,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x31, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61,
	0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x66, 0x65, 0x65,
	0x64, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x0b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x68, 0x65, 0x66, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_external_event_feed_response_eventstrings_proto_rawDescOnce sync.Once
	file_external_event_feed_response_eventstrings_proto_rawDescData = file_external_event_feed_response_eventstrings_proto_rawDesc
)

func file_external_event_feed_response_eventstrings_proto_rawDescGZIP() []byte {
	file_external_event_feed_response_eventstrings_proto_rawDescOnce.Do(func() {
		file_external_event_feed_response_eventstrings_proto_rawDescData = protoimpl.X.CompressGZIP(file_external_event_feed_response_eventstrings_proto_rawDescData)
	})
	return file_external_event_feed_response_eventstrings_proto_rawDescData
}

var file_external_event_feed_response_eventstrings_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_external_event_feed_response_eventstrings_proto_goTypes = []interface{}{
	(*EventStrings)(nil),    // 0: chef.automate.api.event_feed.response.EventStrings
	(*EventString)(nil),     // 1: chef.automate.api.event_feed.response.EventString
	(*EventCollection)(nil), // 2: chef.automate.api.event_feed.response.EventCollection
	(*EventCount)(nil),      // 3: chef.automate.api.event_feed.response.EventCount
}
var file_external_event_feed_response_eventstrings_proto_depIdxs = []int32{
	1, // 0: chef.automate.api.event_feed.response.EventStrings.strings:type_name -> chef.automate.api.event_feed.response.EventString
	2, // 1: chef.automate.api.event_feed.response.EventString.collection:type_name -> chef.automate.api.event_feed.response.EventCollection
	3, // 2: chef.automate.api.event_feed.response.EventCollection.events_count:type_name -> chef.automate.api.event_feed.response.EventCount
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_external_event_feed_response_eventstrings_proto_init() }
func file_external_event_feed_response_eventstrings_proto_init() {
	if File_external_event_feed_response_eventstrings_proto != nil {
		return
	}
	file_external_event_feed_response_event_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_external_event_feed_response_eventstrings_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventStrings); i {
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
		file_external_event_feed_response_eventstrings_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventString); i {
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
		file_external_event_feed_response_eventstrings_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventCollection); i {
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
			RawDescriptor: file_external_event_feed_response_eventstrings_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_external_event_feed_response_eventstrings_proto_goTypes,
		DependencyIndexes: file_external_event_feed_response_eventstrings_proto_depIdxs,
		MessageInfos:      file_external_event_feed_response_eventstrings_proto_msgTypes,
	}.Build()
	File_external_event_feed_response_eventstrings_proto = out.File
	file_external_event_feed_response_eventstrings_proto_rawDesc = nil
	file_external_event_feed_response_eventstrings_proto_goTypes = nil
	file_external_event_feed_response_eventstrings_proto_depIdxs = nil
}
