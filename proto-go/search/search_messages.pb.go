// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0
// source: proto/search/search_messages.proto

package search_pb

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

type SearchMessagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Search string entered by the user.
	Search string `protobuf:"bytes,1,opt,name=search,proto3" json:"search,omitempty"`
	// The id of the team.
	TeamId string `protobuf:"bytes,2,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	// The maximum number of messages to retrieve.
	Limit int64 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	// The offset to start retrieving messages.
	Offset int64 `protobuf:"varint,4,opt,name=offset,proto3" json:"offset,omitempty"`
	// The id of the user.
	UserId string `protobuf:"bytes,5,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// Whether we want one message or not.
	OneMessageByTeam bool `protobuf:"varint,6,opt,name=one_message_by_team,json=oneMessageByTeam,proto3" json:"one_message_by_team,omitempty"`
}

func (x *SearchMessagesRequest) Reset() {
	*x = SearchMessagesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_search_search_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchMessagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchMessagesRequest) ProtoMessage() {}

func (x *SearchMessagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_search_search_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchMessagesRequest.ProtoReflect.Descriptor instead.
func (*SearchMessagesRequest) Descriptor() ([]byte, []int) {
	return file_proto_search_search_messages_proto_rawDescGZIP(), []int{0}
}

func (x *SearchMessagesRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

func (x *SearchMessagesRequest) GetTeamId() string {
	if x != nil {
		return x.TeamId
	}
	return ""
}

func (x *SearchMessagesRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *SearchMessagesRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *SearchMessagesRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *SearchMessagesRequest) GetOneMessageByTeam() bool {
	if x != nil {
		return x.OneMessageByTeam
	}
	return false
}

type SearchMessagesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of retrieved messages.
	Messages []*Message `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *SearchMessagesResponse) Reset() {
	*x = SearchMessagesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_search_search_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchMessagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchMessagesResponse) ProtoMessage() {}

func (x *SearchMessagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_search_search_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchMessagesResponse.ProtoReflect.Descriptor instead.
func (*SearchMessagesResponse) Descriptor() ([]byte, []int) {
	return file_proto_search_search_messages_proto_rawDescGZIP(), []int{1}
}

func (x *SearchMessagesResponse) GetMessages() []*Message {
	if x != nil {
		return x.Messages
	}
	return nil
}

var File_proto_search_search_messages_proto protoreflect.FileDescriptor

var file_proto_search_search_messages_proto_rawDesc = []byte{
	0x0a, 0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2f, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x1a, 0x19, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbe, 0x01, 0x0a, 0x15, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x61,
	0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d,
	0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x13, 0x6f, 0x6e, 0x65,
	0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x62, 0x79, 0x5f, 0x74, 0x65, 0x61, 0x6d,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x6f, 0x6e, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x42, 0x79, 0x54, 0x65, 0x61, 0x6d, 0x22, 0x45, 0x0a, 0x16, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2b, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x32,
	0x63, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x12, 0x51, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x12, 0x1d, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x1b, 0x5a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x6f,
	0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x3b, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_search_search_messages_proto_rawDescOnce sync.Once
	file_proto_search_search_messages_proto_rawDescData = file_proto_search_search_messages_proto_rawDesc
)

func file_proto_search_search_messages_proto_rawDescGZIP() []byte {
	file_proto_search_search_messages_proto_rawDescOnce.Do(func() {
		file_proto_search_search_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_search_search_messages_proto_rawDescData)
	})
	return file_proto_search_search_messages_proto_rawDescData
}

var file_proto_search_search_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_search_search_messages_proto_goTypes = []any{
	(*SearchMessagesRequest)(nil),  // 0: search.SearchMessagesRequest
	(*SearchMessagesResponse)(nil), // 1: search.SearchMessagesResponse
	(*Message)(nil),                // 2: search.Message
}
var file_proto_search_search_messages_proto_depIdxs = []int32{
	2, // 0: search.SearchMessagesResponse.messages:type_name -> search.Message
	0, // 1: search.SearchMessages.SearchMessages:input_type -> search.SearchMessagesRequest
	1, // 2: search.SearchMessages.SearchMessages:output_type -> search.SearchMessagesResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_search_search_messages_proto_init() }
func file_proto_search_search_messages_proto_init() {
	if File_proto_search_search_messages_proto != nil {
		return
	}
	file_proto_search_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_search_search_messages_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*SearchMessagesRequest); i {
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
		file_proto_search_search_messages_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*SearchMessagesResponse); i {
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
			RawDescriptor: file_proto_search_search_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_search_search_messages_proto_goTypes,
		DependencyIndexes: file_proto_search_search_messages_proto_depIdxs,
		MessageInfos:      file_proto_search_search_messages_proto_msgTypes,
	}.Build()
	File_proto_search_search_messages_proto = out.File
	file_proto_search_search_messages_proto_rawDesc = nil
	file_proto_search_search_messages_proto_goTypes = nil
	file_proto_search_search_messages_proto_depIdxs = nil
}
