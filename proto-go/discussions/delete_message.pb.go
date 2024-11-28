// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: proto/discussions/delete_message.proto

package discussions_pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DeleteMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the message.
	MessageId string `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
}

func (x *DeleteMessageRequest) Reset() {
	*x = DeleteMessageRequest{}
	mi := &file_proto_discussions_delete_message_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMessageRequest) ProtoMessage() {}

func (x *DeleteMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_discussions_delete_message_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMessageRequest.ProtoReflect.Descriptor instead.
func (*DeleteMessageRequest) Descriptor() ([]byte, []int) {
	return file_proto_discussions_delete_message_proto_rawDescGZIP(), []int{0}
}

func (x *DeleteMessageRequest) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

var File_proto_discussions_delete_message_proto protoreflect.FileDescriptor

var file_proto_discussions_delete_message_proto_rawDesc = []byte{
	0x0a, 0x26, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x35, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x32, 0x5d, 0x0a, 0x0d, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x4c, 0x0a, 0x0d, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x21, 0x2e, 0x64, 0x69,
	0x73, 0x63, 0x75, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x25, 0x5a, 0x23, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2d, 0x67, 0x6f, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x3b, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_discussions_delete_message_proto_rawDescOnce sync.Once
	file_proto_discussions_delete_message_proto_rawDescData = file_proto_discussions_delete_message_proto_rawDesc
)

func file_proto_discussions_delete_message_proto_rawDescGZIP() []byte {
	file_proto_discussions_delete_message_proto_rawDescOnce.Do(func() {
		file_proto_discussions_delete_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_discussions_delete_message_proto_rawDescData)
	})
	return file_proto_discussions_delete_message_proto_rawDescData
}

var file_proto_discussions_delete_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_discussions_delete_message_proto_goTypes = []any{
	(*DeleteMessageRequest)(nil), // 0: discussions.DeleteMessageRequest
	(*emptypb.Empty)(nil),        // 1: google.protobuf.Empty
}
var file_proto_discussions_delete_message_proto_depIdxs = []int32{
	0, // 0: discussions.DeleteMessage.DeleteMessage:input_type -> discussions.DeleteMessageRequest
	1, // 1: discussions.DeleteMessage.DeleteMessage:output_type -> google.protobuf.Empty
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_discussions_delete_message_proto_init() }
func file_proto_discussions_delete_message_proto_init() {
	if File_proto_discussions_delete_message_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_discussions_delete_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_discussions_delete_message_proto_goTypes,
		DependencyIndexes: file_proto_discussions_delete_message_proto_depIdxs,
		MessageInfos:      file_proto_discussions_delete_message_proto_msgTypes,
	}.Build()
	File_proto_discussions_delete_message_proto = out.File
	file_proto_discussions_delete_message_proto_rawDesc = nil
	file_proto_discussions_delete_message_proto_goTypes = nil
	file_proto_discussions_delete_message_proto_depIdxs = nil
}
