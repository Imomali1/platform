// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: platform/student/analytics.proto

package student

import (
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

type Request1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request string `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *Request1) Reset() {
	*x = Request1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platform_student_analytics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request1) ProtoMessage() {}

func (x *Request1) ProtoReflect() protoreflect.Message {
	mi := &file_platform_student_analytics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request1.ProtoReflect.Descriptor instead.
func (*Request1) Descriptor() ([]byte, []int) {
	return file_platform_student_analytics_proto_rawDescGZIP(), []int{0}
}

func (x *Request1) GetRequest() string {
	if x != nil {
		return x.Request
	}
	return ""
}

type Response1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reply string `protobuf:"bytes,1,opt,name=reply,proto3" json:"reply,omitempty"`
}

func (x *Response1) Reset() {
	*x = Response1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platform_student_analytics_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response1) ProtoMessage() {}

func (x *Response1) ProtoReflect() protoreflect.Message {
	mi := &file_platform_student_analytics_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response1.ProtoReflect.Descriptor instead.
func (*Response1) Descriptor() ([]byte, []int) {
	return file_platform_student_analytics_proto_rawDescGZIP(), []int{1}
}

func (x *Response1) GetReply() string {
	if x != nil {
		return x.Reply
	}
	return ""
}

var File_platform_student_analytics_proto protoreflect.FileDescriptor

var file_platform_student_analytics_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x73, 0x74, 0x75, 0x64, 0x65,
	0x6e, 0x74, 0x2f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x24, 0x0a, 0x08, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x31, 0x12, 0x18,
	0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x21, 0x0a, 0x09, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x31, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x32, 0x37, 0x0a, 0x10, 0x41,
	0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x23, 0x0a, 0x08, 0x53, 0x6f, 0x6d, 0x65, 0x46, 0x75, 0x6e, 0x63, 0x12, 0x09, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x31, 0x1a, 0x0a, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x31, 0x22, 0x00, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x49, 0x6d, 0x6f, 0x6d, 0x61, 0x6c, 0x69, 0x31, 0x2f, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x73, 0x74,
	0x75, 0x64, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_platform_student_analytics_proto_rawDescOnce sync.Once
	file_platform_student_analytics_proto_rawDescData = file_platform_student_analytics_proto_rawDesc
)

func file_platform_student_analytics_proto_rawDescGZIP() []byte {
	file_platform_student_analytics_proto_rawDescOnce.Do(func() {
		file_platform_student_analytics_proto_rawDescData = protoimpl.X.CompressGZIP(file_platform_student_analytics_proto_rawDescData)
	})
	return file_platform_student_analytics_proto_rawDescData
}

var file_platform_student_analytics_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_platform_student_analytics_proto_goTypes = []interface{}{
	(*Request1)(nil),  // 0: Request1
	(*Response1)(nil), // 1: Response1
}
var file_platform_student_analytics_proto_depIdxs = []int32{
	0, // 0: AnalyticsService.SomeFunc:input_type -> Request1
	1, // 1: AnalyticsService.SomeFunc:output_type -> Response1
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_platform_student_analytics_proto_init() }
func file_platform_student_analytics_proto_init() {
	if File_platform_student_analytics_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_platform_student_analytics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request1); i {
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
		file_platform_student_analytics_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response1); i {
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
			RawDescriptor: file_platform_student_analytics_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_platform_student_analytics_proto_goTypes,
		DependencyIndexes: file_platform_student_analytics_proto_depIdxs,
		MessageInfos:      file_platform_student_analytics_proto_msgTypes,
	}.Build()
	File_platform_student_analytics_proto = out.File
	file_platform_student_analytics_proto_rawDesc = nil
	file_platform_student_analytics_proto_goTypes = nil
	file_platform_student_analytics_proto_depIdxs = nil
}
