// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.21.12
// source: proto/discussions/common.proto

package discussions_pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type Message struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The target of a note indicates what type of LinkedIn profile this message belongs to.
	Target string `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	// The vanity name of the LinkedIn profile, as it appears in its LinkedIn profile URL.
	PublicIdentifier string `protobuf:"bytes,2,opt,name=public_identifier,json=publicIdentifier,proto3" json:"public_identifier,omitempty"`
	// The message content.
	Content string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	// The timestamp of the message creation.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// The id of the note's author.
	AuthorId string `protobuf:"bytes,5,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	// The ID of the team.
	TeamId string `protobuf:"bytes,6,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	// The ID of the message.
	MessageId     string `protobuf:"bytes,7,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Message) Reset() {
	*x = Message{}
	mi := &file_proto_discussions_common_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_proto_discussions_common_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_proto_discussions_common_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *Message) GetPublicIdentifier() string {
	if x != nil {
		return x.PublicIdentifier
	}
	return ""
}

func (x *Message) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Message) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Message) GetAuthorId() string {
	if x != nil {
		return x.AuthorId
	}
	return ""
}

func (x *Message) GetTeamId() string {
	if x != nil {
		return x.TeamId
	}
	return ""
}

func (x *Message) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

type Discussion struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The target of a note indicates what type of LinkedIn profile this discussion belongs to.
	Target string `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	// The vanity name of the LinkedIn profile, as it appears in its LinkedIn profile URL.
	PublicIdentifier string `protobuf:"bytes,2,opt,name=public_identifier,json=publicIdentifier,proto3" json:"public_identifier,omitempty"`
	// The ID of the team.
	TeamId string `protobuf:"bytes,3,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	// The timestamp of the latest message creation.
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Discussion) Reset() {
	*x = Discussion{}
	mi := &file_proto_discussions_common_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Discussion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Discussion) ProtoMessage() {}

func (x *Discussion) ProtoReflect() protoreflect.Message {
	mi := &file_proto_discussions_common_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Discussion.ProtoReflect.Descriptor instead.
func (*Discussion) Descriptor() ([]byte, []int) {
	return file_proto_discussions_common_proto_rawDescGZIP(), []int{1}
}

func (x *Discussion) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *Discussion) GetPublicIdentifier() string {
	if x != nil {
		return x.PublicIdentifier
	}
	return ""
}

func (x *Discussion) GetTeamId() string {
	if x != nil {
		return x.TeamId
	}
	return ""
}

func (x *Discussion) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type DiscussionReadStatus struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The target of a note indicates what type of LinkedIn profile this discussion belongs to.
	Target string `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	// The vanity name of the LinkedIn profile, as it appears in its LinkedIn profile URL.
	PublicIdentifier string `protobuf:"bytes,2,opt,name=public_identifier,json=publicIdentifier,proto3" json:"public_identifier,omitempty"`
	// The ID of the team.
	TeamId string `protobuf:"bytes,3,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	// The ID of the user.
	UserId string `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// The ID of the latest read message.
	LatestReadMessageId string `protobuf:"bytes,5,opt,name=latest_read_message_id,json=latestReadMessageId,proto3" json:"latest_read_message_id,omitempty"`
	// Whether the discussion has unread messages.
	HasUnreadMessages bool `protobuf:"varint,6,opt,name=has_unread_messages,json=hasUnreadMessages,proto3" json:"has_unread_messages,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *DiscussionReadStatus) Reset() {
	*x = DiscussionReadStatus{}
	mi := &file_proto_discussions_common_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DiscussionReadStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscussionReadStatus) ProtoMessage() {}

func (x *DiscussionReadStatus) ProtoReflect() protoreflect.Message {
	mi := &file_proto_discussions_common_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscussionReadStatus.ProtoReflect.Descriptor instead.
func (*DiscussionReadStatus) Descriptor() ([]byte, []int) {
	return file_proto_discussions_common_proto_rawDescGZIP(), []int{2}
}

func (x *DiscussionReadStatus) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *DiscussionReadStatus) GetPublicIdentifier() string {
	if x != nil {
		return x.PublicIdentifier
	}
	return ""
}

func (x *DiscussionReadStatus) GetTeamId() string {
	if x != nil {
		return x.TeamId
	}
	return ""
}

func (x *DiscussionReadStatus) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *DiscussionReadStatus) GetLatestReadMessageId() string {
	if x != nil {
		return x.LatestReadMessageId
	}
	return ""
}

func (x *DiscussionReadStatus) GetHasUnreadMessages() bool {
	if x != nil {
		return x.HasUnreadMessages
	}
	return false
}

var File_proto_discussions_common_proto protoreflect.FileDescriptor

var file_proto_discussions_common_proto_rawDesc = string([]byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0b, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf8,
	0x01, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x12, 0x2b, 0x0a, 0x11, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x22, 0xa5, 0x01, 0x0a, 0x0a, 0x44, 0x69,
	0x73, 0x63, 0x75, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x12, 0x2b, 0x0a, 0x11, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x17, 0x0a,
	0x07, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x22, 0xf2, 0x01, 0x0a, 0x14, 0x44, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x12, 0x2b, 0x0a, 0x11, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12,
	0x17, 0x0a, 0x07, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x33, 0x0a, 0x16, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x61, 0x64,
	0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x13, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x52, 0x65, 0x61, 0x64, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x13, 0x68, 0x61, 0x73, 0x5f, 0x75, 0x6e,
	0x72, 0x65, 0x61, 0x64, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x11, 0x68, 0x61, 0x73, 0x55, 0x6e, 0x72, 0x65, 0x61, 0x64, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x42, 0x25, 0x5a, 0x23, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d,
	0x67, 0x6f, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x3b, 0x64,
	0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_proto_discussions_common_proto_rawDescOnce sync.Once
	file_proto_discussions_common_proto_rawDescData []byte
)

func file_proto_discussions_common_proto_rawDescGZIP() []byte {
	file_proto_discussions_common_proto_rawDescOnce.Do(func() {
		file_proto_discussions_common_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_discussions_common_proto_rawDesc), len(file_proto_discussions_common_proto_rawDesc)))
	})
	return file_proto_discussions_common_proto_rawDescData
}

var file_proto_discussions_common_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_discussions_common_proto_goTypes = []any{
	(*Message)(nil),               // 0: discussions.Message
	(*Discussion)(nil),            // 1: discussions.Discussion
	(*DiscussionReadStatus)(nil),  // 2: discussions.DiscussionReadStatus
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_proto_discussions_common_proto_depIdxs = []int32{
	3, // 0: discussions.Message.created_at:type_name -> google.protobuf.Timestamp
	3, // 1: discussions.Discussion.updated_at:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_discussions_common_proto_init() }
func file_proto_discussions_common_proto_init() {
	if File_proto_discussions_common_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_discussions_common_proto_rawDesc), len(file_proto_discussions_common_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_discussions_common_proto_goTypes,
		DependencyIndexes: file_proto_discussions_common_proto_depIdxs,
		MessageInfos:      file_proto_discussions_common_proto_msgTypes,
	}.Build()
	File_proto_discussions_common_proto = out.File
	file_proto_discussions_common_proto_goTypes = nil
	file_proto_discussions_common_proto_depIdxs = nil
}
