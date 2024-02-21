// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: eventstream.proto

package eventstreamapi

import (
	_ "github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	pb "github.com/cloudevents/sdk-go/binding/format/protobuf/v2/pb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Event describes an event
type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event *pb.CloudEvent `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eventstream_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_eventstream_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_eventstream_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetEvent() *pb.CloudEvent {
	if x != nil {
		return x.Event
	}
	return nil
}

type PushSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result    string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Received  int32  `protobuf:"varint,2,opt,name=received,proto3" json:"received,omitempty"`
	Processed int32  `protobuf:"varint,3,opt,name=processed,proto3" json:"processed,omitempty"`
}

func (x *PushSummary) Reset() {
	*x = PushSummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eventstream_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushSummary) ProtoMessage() {}

func (x *PushSummary) ProtoReflect() protoreflect.Message {
	mi := &file_eventstream_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushSummary.ProtoReflect.Descriptor instead.
func (*PushSummary) Descriptor() ([]byte, []int) {
	return file_eventstream_proto_rawDescGZIP(), []int{1}
}

func (x *PushSummary) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

func (x *PushSummary) GetReceived() int32 {
	if x != nil {
		return x.Received
	}
	return 0
}

func (x *PushSummary) GetProcessed() int32 {
	if x != nil {
		return x.Processed
	}
	return 0
}

type PingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PingRequest) Reset() {
	*x = PingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eventstream_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRequest) ProtoMessage() {}

func (x *PingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_eventstream_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRequest.ProtoReflect.Descriptor instead.
func (*PingRequest) Descriptor() ([]byte, []int) {
	return file_eventstream_proto_rawDescGZIP(), []int{2}
}

type PongReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PongReply) Reset() {
	*x = PongReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eventstream_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PongReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PongReply) ProtoMessage() {}

func (x *PongReply) ProtoReflect() protoreflect.Message {
	mi := &file_eventstream_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PongReply.ProtoReflect.Descriptor instead.
func (*PongReply) Descriptor() ([]byte, []int) {
	return file_eventstream_proto_rawDescGZIP(), []int{3}
}

var File_eventstream_proto protoreflect.FileDescriptor

var file_eventstream_proto_rawDesc = []byte{
	0x0a, 0x11, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x61, 0x70, 0x69, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x72,
	0x67, 0x6f, 0x70, 0x72, 0x6f, 0x6a, 0x2f, 0x61, 0x72, 0x67, 0x6f, 0x2d, 0x63, 0x64, 0x2f, 0x76,
	0x32, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x62,
	0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x2f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x32, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3c, 0x0a,
	0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x33, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x5f, 0x0a, 0x0b, 0x50,
	0x75, 0x73, 0x68, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x12, 0x1c,
	0x0a, 0x09, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x22, 0x0d, 0x0a, 0x0b,
	0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0b, 0x0a, 0x09, 0x50,
	0x6f, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0x9c, 0x02, 0x0a, 0x0b, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x5c, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x15, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x15, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x12, 0x15, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x28, 0x01, 0x30, 0x01, 0x12, 0x59, 0x0a, 0x04, 0x50, 0x75, 0x73, 0x68, 0x12, 0x15,
	0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x61, 0x70, 0x69, 0x2e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x1b, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x53, 0x75, 0x6d, 0x6d, 0x61,
	0x72, 0x79, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x70, 0x75, 0x73, 0x68, 0x28,
	0x01, 0x12, 0x54, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x1b, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x61, 0x6e, 0x6e, 0x66, 0x69, 0x73, 0x2f, 0x61, 0x72,
	0x67, 0x6f, 0x63, 0x64, 0x2d, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eventstream_proto_rawDescOnce sync.Once
	file_eventstream_proto_rawDescData = file_eventstream_proto_rawDesc
)

func file_eventstream_proto_rawDescGZIP() []byte {
	file_eventstream_proto_rawDescOnce.Do(func() {
		file_eventstream_proto_rawDescData = protoimpl.X.CompressGZIP(file_eventstream_proto_rawDescData)
	})
	return file_eventstream_proto_rawDescData
}

var file_eventstream_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_eventstream_proto_goTypes = []interface{}{
	(*Event)(nil),         // 0: eventstreamapi.Event
	(*PushSummary)(nil),   // 1: eventstreamapi.PushSummary
	(*PingRequest)(nil),   // 2: eventstreamapi.PingRequest
	(*PongReply)(nil),     // 3: eventstreamapi.PongReply
	(*pb.CloudEvent)(nil), // 4: io.cloudevents.v1.CloudEvent
}
var file_eventstream_proto_depIdxs = []int32{
	4, // 0: eventstreamapi.Event.event:type_name -> io.cloudevents.v1.CloudEvent
	0, // 1: eventstreamapi.EventStream.Subscribe:input_type -> eventstreamapi.Event
	0, // 2: eventstreamapi.EventStream.Push:input_type -> eventstreamapi.Event
	2, // 3: eventstreamapi.EventStream.Ping:input_type -> eventstreamapi.PingRequest
	0, // 4: eventstreamapi.EventStream.Subscribe:output_type -> eventstreamapi.Event
	1, // 5: eventstreamapi.EventStream.Push:output_type -> eventstreamapi.PushSummary
	3, // 6: eventstreamapi.EventStream.Ping:output_type -> eventstreamapi.PongReply
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_eventstream_proto_init() }
func file_eventstream_proto_init() {
	if File_eventstream_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eventstream_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_eventstream_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushSummary); i {
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
		file_eventstream_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingRequest); i {
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
		file_eventstream_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PongReply); i {
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
			RawDescriptor: file_eventstream_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eventstream_proto_goTypes,
		DependencyIndexes: file_eventstream_proto_depIdxs,
		MessageInfos:      file_eventstream_proto_msgTypes,
	}.Build()
	File_eventstream_proto = out.File
	file_eventstream_proto_rawDesc = nil
	file_eventstream_proto_goTypes = nil
	file_eventstream_proto_depIdxs = nil
}
