// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: pb/grpc_try.proto

package pb

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

type WorkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Comments []string `protobuf:"bytes,2,rep,name=comments,proto3" json:"comments,omitempty"`
}

func (x *WorkRequest) Reset() {
	*x = WorkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_grpc_try_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkRequest) ProtoMessage() {}

func (x *WorkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_grpc_try_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkRequest.ProtoReflect.Descriptor instead.
func (*WorkRequest) Descriptor() ([]byte, []int) {
	return file_pb_grpc_try_proto_rawDescGZIP(), []int{0}
}

func (x *WorkRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *WorkRequest) GetComments() []string {
	if x != nil {
		return x.Comments
	}
	return nil
}

type WorkReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Result bool  `protobuf:"varint,2,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *WorkReply) Reset() {
	*x = WorkReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_grpc_try_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkReply) ProtoMessage() {}

func (x *WorkReply) ProtoReflect() protoreflect.Message {
	mi := &file_pb_grpc_try_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkReply.ProtoReflect.Descriptor instead.
func (*WorkReply) Descriptor() ([]byte, []int) {
	return file_pb_grpc_try_proto_rawDescGZIP(), []int{1}
}

func (x *WorkReply) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *WorkReply) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

var File_pb_grpc_try_proto protoreflect.FileDescriptor

var file_pb_grpc_try_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x62, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x39, 0x0a, 0x0b, 0x57, 0x6f, 0x72, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x22, 0x33, 0x0a, 0x09, 0x57, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x35, 0x0a, 0x06, 0x57, 0x6f, 0x72, 0x6b, 0x65,
	0x72, 0x12, 0x2b, 0x0a, 0x07, 0x57, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x0f, 0x2e, 0x70,
	0x62, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e,
	0x70, 0x62, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x24,
	0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4c, 0x6f, 0x72,
	0x64, 0x53, 0x68, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x74, 0x72,
	0x79, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_grpc_try_proto_rawDescOnce sync.Once
	file_pb_grpc_try_proto_rawDescData = file_pb_grpc_try_proto_rawDesc
)

func file_pb_grpc_try_proto_rawDescGZIP() []byte {
	file_pb_grpc_try_proto_rawDescOnce.Do(func() {
		file_pb_grpc_try_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_grpc_try_proto_rawDescData)
	})
	return file_pb_grpc_try_proto_rawDescData
}

var file_pb_grpc_try_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pb_grpc_try_proto_goTypes = []interface{}{
	(*WorkRequest)(nil), // 0: pb.WorkRequest
	(*WorkReply)(nil),   // 1: pb.WorkReply
}
var file_pb_grpc_try_proto_depIdxs = []int32{
	0, // 0: pb.Worker.Working:input_type -> pb.WorkRequest
	1, // 1: pb.Worker.Working:output_type -> pb.WorkReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_grpc_try_proto_init() }
func file_pb_grpc_try_proto_init() {
	if File_pb_grpc_try_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_grpc_try_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkRequest); i {
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
		file_pb_grpc_try_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkReply); i {
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
			RawDescriptor: file_pb_grpc_try_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_grpc_try_proto_goTypes,
		DependencyIndexes: file_pb_grpc_try_proto_depIdxs,
		MessageInfos:      file_pb_grpc_try_proto_msgTypes,
	}.Build()
	File_pb_grpc_try_proto = out.File
	file_pb_grpc_try_proto_rawDesc = nil
	file_pb_grpc_try_proto_goTypes = nil
	file_pb_grpc_try_proto_depIdxs = nil
}
