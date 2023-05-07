// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: protos/clock/clock.proto

package clock

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type TimeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TimeRequest) Reset() {
	*x = TimeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_clock_clock_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeRequest) ProtoMessage() {}

func (x *TimeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_clock_clock_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeRequest.ProtoReflect.Descriptor instead.
func (*TimeRequest) Descriptor() ([]byte, []int) {
	return file_protos_clock_clock_proto_rawDescGZIP(), []int{0}
}

type TimeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Time *timestamp.Timestamp `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *TimeResponse) Reset() {
	*x = TimeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_clock_clock_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeResponse) ProtoMessage() {}

func (x *TimeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_clock_clock_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeResponse.ProtoReflect.Descriptor instead.
func (*TimeResponse) Descriptor() ([]byte, []int) {
	return file_protos_clock_clock_proto_rawDescGZIP(), []int{1}
}

func (x *TimeResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TimeResponse) GetTime() *timestamp.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

var File_protos_clock_clock_proto protoreflect.FileDescriptor

var file_protos_clock_clock_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x2f, 0x63,
	0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x63, 0x6c, 0x6f, 0x63,
	0x6b, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x0d, 0x0a, 0x0b, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x4e, 0x0a, 0x0c, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d,
	0x65, 0x32, 0x37, 0x0a, 0x05, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x2e, 0x0a, 0x03, 0x4e, 0x6f,
	0x77, 0x12, 0x12, 0x2e, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x47, 0x6f, 0x6f, 0x64, 0x2f, 0x76,
	0x38, 0x2d, 0x72, 0x70, 0x63, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_clock_clock_proto_rawDescOnce sync.Once
	file_protos_clock_clock_proto_rawDescData = file_protos_clock_clock_proto_rawDesc
)

func file_protos_clock_clock_proto_rawDescGZIP() []byte {
	file_protos_clock_clock_proto_rawDescOnce.Do(func() {
		file_protos_clock_clock_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_clock_clock_proto_rawDescData)
	})
	return file_protos_clock_clock_proto_rawDescData
}

var file_protos_clock_clock_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protos_clock_clock_proto_goTypes = []interface{}{
	(*TimeRequest)(nil),         // 0: clock.TimeRequest
	(*TimeResponse)(nil),        // 1: clock.TimeResponse
	(*timestamp.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_protos_clock_clock_proto_depIdxs = []int32{
	2, // 0: clock.TimeResponse.time:type_name -> google.protobuf.Timestamp
	0, // 1: clock.Clock.Now:input_type -> clock.TimeRequest
	1, // 2: clock.Clock.Now:output_type -> clock.TimeResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_clock_clock_proto_init() }
func file_protos_clock_clock_proto_init() {
	if File_protos_clock_clock_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_clock_clock_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeRequest); i {
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
		file_protos_clock_clock_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeResponse); i {
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
			RawDescriptor: file_protos_clock_clock_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_clock_clock_proto_goTypes,
		DependencyIndexes: file_protos_clock_clock_proto_depIdxs,
		MessageInfos:      file_protos_clock_clock_proto_msgTypes,
	}.Build()
	File_protos_clock_clock_proto = out.File
	file_protos_clock_clock_proto_rawDesc = nil
	file_protos_clock_clock_proto_goTypes = nil
	file_protos_clock_clock_proto_depIdxs = nil
}