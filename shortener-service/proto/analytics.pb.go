// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.0--rc2
// source: proto/analytics.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type VisitRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Shortcode     string                 `protobuf:"bytes,1,opt,name=shortcode,proto3" json:"shortcode,omitempty"`
	Url           string                 `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Ip            string                 `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`
	UserAgent     string                 `protobuf:"bytes,4,opt,name=user_agent,json=userAgent,proto3" json:"user_agent,omitempty"`
	Timestamp     int64                  `protobuf:"varint,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VisitRequest) Reset() {
	*x = VisitRequest{}
	mi := &file_proto_analytics_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VisitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VisitRequest) ProtoMessage() {}

func (x *VisitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_analytics_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VisitRequest.ProtoReflect.Descriptor instead.
func (*VisitRequest) Descriptor() ([]byte, []int) {
	return file_proto_analytics_proto_rawDescGZIP(), []int{0}
}

func (x *VisitRequest) GetShortcode() string {
	if x != nil {
		return x.Shortcode
	}
	return ""
}

func (x *VisitRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *VisitRequest) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *VisitRequest) GetUserAgent() string {
	if x != nil {
		return x.UserAgent
	}
	return ""
}

func (x *VisitRequest) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type VisitResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        string                 `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VisitResponse) Reset() {
	*x = VisitResponse{}
	mi := &file_proto_analytics_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VisitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VisitResponse) ProtoMessage() {}

func (x *VisitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_analytics_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VisitResponse.ProtoReflect.Descriptor instead.
func (*VisitResponse) Descriptor() ([]byte, []int) {
	return file_proto_analytics_proto_rawDescGZIP(), []int{1}
}

func (x *VisitResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_proto_analytics_proto protoreflect.FileDescriptor

const file_proto_analytics_proto_rawDesc = "" +
	"\n" +
	"\x15proto/analytics.proto\x12\tanalytics\"\x8b\x01\n" +
	"\fVisitRequest\x12\x1c\n" +
	"\tshortcode\x18\x01 \x01(\tR\tshortcode\x12\x10\n" +
	"\x03url\x18\x02 \x01(\tR\x03url\x12\x0e\n" +
	"\x02ip\x18\x03 \x01(\tR\x02ip\x12\x1d\n" +
	"\n" +
	"user_agent\x18\x04 \x01(\tR\tuserAgent\x12\x1c\n" +
	"\ttimestamp\x18\x05 \x01(\x03R\ttimestamp\"'\n" +
	"\rVisitResponse\x12\x16\n" +
	"\x06status\x18\x01 \x01(\tR\x06status2J\n" +
	"\tAnalytics\x12=\n" +
	"\bLogVisit\x12\x17.analytics.VisitRequest\x1a\x18.analytics.VisitResponseB\x1fZ\x1danalytics-service/proto;protob\x06proto3"

var (
	file_proto_analytics_proto_rawDescOnce sync.Once
	file_proto_analytics_proto_rawDescData []byte
)

func file_proto_analytics_proto_rawDescGZIP() []byte {
	file_proto_analytics_proto_rawDescOnce.Do(func() {
		file_proto_analytics_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_analytics_proto_rawDesc), len(file_proto_analytics_proto_rawDesc)))
	})
	return file_proto_analytics_proto_rawDescData
}

var file_proto_analytics_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_analytics_proto_goTypes = []any{
	(*VisitRequest)(nil),  // 0: analytics.VisitRequest
	(*VisitResponse)(nil), // 1: analytics.VisitResponse
}
var file_proto_analytics_proto_depIdxs = []int32{
	0, // 0: analytics.Analytics.LogVisit:input_type -> analytics.VisitRequest
	1, // 1: analytics.Analytics.LogVisit:output_type -> analytics.VisitResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_analytics_proto_init() }
func file_proto_analytics_proto_init() {
	if File_proto_analytics_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_analytics_proto_rawDesc), len(file_proto_analytics_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_analytics_proto_goTypes,
		DependencyIndexes: file_proto_analytics_proto_depIdxs,
		MessageInfos:      file_proto_analytics_proto_msgTypes,
	}.Build()
	File_proto_analytics_proto = out.File
	file_proto_analytics_proto_goTypes = nil
	file_proto_analytics_proto_depIdxs = nil
}
