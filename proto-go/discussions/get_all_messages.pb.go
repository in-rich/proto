// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.4
// source: proto/discussions/get_all_messages.proto

package discussions_pb

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

type GetAllMessagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  int64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int64 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *GetAllMessagesRequest) Reset() {
	*x = GetAllMessagesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_discussions_get_all_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllMessagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllMessagesRequest) ProtoMessage() {}

func (x *GetAllMessagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_discussions_get_all_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllMessagesRequest.ProtoReflect.Descriptor instead.
func (*GetAllMessagesRequest) Descriptor() ([]byte, []int) {
	return file_proto_discussions_get_all_messages_proto_rawDescGZIP(), []int{0}
}

func (x *GetAllMessagesRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetAllMessagesRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type GetAllMessagesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The GetAll of retrieved Messages.
	Messages []*Message `protobuf:"bytes,1,rep,name=Messages,proto3" json:"Messages,omitempty"`
}

func (x *GetAllMessagesResponse) Reset() {
	*x = GetAllMessagesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_discussions_get_all_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllMessagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllMessagesResponse) ProtoMessage() {}

func (x *GetAllMessagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_discussions_get_all_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllMessagesResponse.ProtoReflect.Descriptor instead.
func (*GetAllMessagesResponse) Descriptor() ([]byte, []int) {
	return file_proto_discussions_get_all_messages_proto_rawDescGZIP(), []int{1}
}

func (x *GetAllMessagesResponse) GetMessages() []*Message {
	if x != nil {
		return x.Messages
	}
	return nil
}

var File_proto_discussions_get_all_messages_proto protoreflect.FileDescriptor

var file_proto_discussions_get_all_messages_proto_rawDesc = []byte{
	0x0a, 0x28, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x61, 0x6c, 0x6c, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x64, 0x69, 0x73, 0x63,
	0x75, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64,
	0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x45, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x4a,
	0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x08, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x64, 0x69, 0x73,
	0x63, 0x75, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x08, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x32, 0x6d, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x5b, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x22,
	0x2e, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x23, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x25, 0x5a, 0x23, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2d, 0x67, 0x6f, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x3b, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_discussions_get_all_messages_proto_rawDescOnce sync.Once
	file_proto_discussions_get_all_messages_proto_rawDescData = file_proto_discussions_get_all_messages_proto_rawDesc
)

func file_proto_discussions_get_all_messages_proto_rawDescGZIP() []byte {
	file_proto_discussions_get_all_messages_proto_rawDescOnce.Do(func() {
		file_proto_discussions_get_all_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_discussions_get_all_messages_proto_rawDescData)
	})
	return file_proto_discussions_get_all_messages_proto_rawDescData
}

var file_proto_discussions_get_all_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_discussions_get_all_messages_proto_goTypes = []any{
	(*GetAllMessagesRequest)(nil),  // 0: discussions.GetAllMessagesRequest
	(*GetAllMessagesResponse)(nil), // 1: discussions.GetAllMessagesResponse
	(*Message)(nil),                // 2: discussions.Message
}
var file_proto_discussions_get_all_messages_proto_depIdxs = []int32{
	2, // 0: discussions.GetAllMessagesResponse.Messages:type_name -> discussions.Message
	0, // 1: discussions.GetAllMessages.GetAllMessages:input_type -> discussions.GetAllMessagesRequest
	1, // 2: discussions.GetAllMessages.GetAllMessages:output_type -> discussions.GetAllMessagesResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_discussions_get_all_messages_proto_init() }
func file_proto_discussions_get_all_messages_proto_init() {
	if File_proto_discussions_get_all_messages_proto != nil {
		return
	}
	file_proto_discussions_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_discussions_get_all_messages_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllMessagesRequest); i {
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
		file_proto_discussions_get_all_messages_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllMessagesResponse); i {
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
			RawDescriptor: file_proto_discussions_get_all_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_discussions_get_all_messages_proto_goTypes,
		DependencyIndexes: file_proto_discussions_get_all_messages_proto_depIdxs,
		MessageInfos:      file_proto_discussions_get_all_messages_proto_msgTypes,
	}.Build()
	File_proto_discussions_get_all_messages_proto = out.File
	file_proto_discussions_get_all_messages_proto_rawDesc = nil
	file_proto_discussions_get_all_messages_proto_goTypes = nil
	file_proto_discussions_get_all_messages_proto_depIdxs = nil
}
