// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: proto/teams/list_user_teams.proto

package teams_pb

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

type ListUserTeamsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The id of the user.
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Limit  int64  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int64  `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *ListUserTeamsRequest) Reset() {
	*x = ListUserTeamsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_teams_list_user_teams_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUserTeamsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserTeamsRequest) ProtoMessage() {}

func (x *ListUserTeamsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_teams_list_user_teams_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserTeamsRequest.ProtoReflect.Descriptor instead.
func (*ListUserTeamsRequest) Descriptor() ([]byte, []int) {
	return file_proto_teams_list_user_teams_proto_rawDescGZIP(), []int{0}
}

func (x *ListUserTeamsRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ListUserTeamsRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListUserTeamsRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type ListUserTeamsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The teams that the user is a member of.
	Teams []*Team `protobuf:"bytes,1,rep,name=teams,proto3" json:"teams,omitempty"`
}

func (x *ListUserTeamsResponse) Reset() {
	*x = ListUserTeamsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_teams_list_user_teams_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUserTeamsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserTeamsResponse) ProtoMessage() {}

func (x *ListUserTeamsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_teams_list_user_teams_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserTeamsResponse.ProtoReflect.Descriptor instead.
func (*ListUserTeamsResponse) Descriptor() ([]byte, []int) {
	return file_proto_teams_list_user_teams_proto_rawDescGZIP(), []int{1}
}

func (x *ListUserTeamsResponse) GetTeams() []*Team {
	if x != nil {
		return x.Teams
	}
	return nil
}

var File_proto_teams_list_user_teams_proto protoreflect.FileDescriptor

var file_proto_teams_list_user_teams_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x2f, 0x6c, 0x69,
	0x73, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x1a, 0x18, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5d, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x54, 0x65, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x22, 0x3a, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54,
	0x65, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x05,
	0x74, 0x65, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x74, 0x65,
	0x61, 0x6d, 0x73, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x05, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x32,
	0x5d, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x65, 0x61, 0x6d, 0x73,
	0x12, 0x4c, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x65, 0x61, 0x6d,
	0x73, 0x12, 0x1b, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x54, 0x65, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54,
	0x65, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x19,
	0x5a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x61, 0x6d, 0x73,
	0x3b, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_teams_list_user_teams_proto_rawDescOnce sync.Once
	file_proto_teams_list_user_teams_proto_rawDescData = file_proto_teams_list_user_teams_proto_rawDesc
)

func file_proto_teams_list_user_teams_proto_rawDescGZIP() []byte {
	file_proto_teams_list_user_teams_proto_rawDescOnce.Do(func() {
		file_proto_teams_list_user_teams_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_teams_list_user_teams_proto_rawDescData)
	})
	return file_proto_teams_list_user_teams_proto_rawDescData
}

var file_proto_teams_list_user_teams_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_teams_list_user_teams_proto_goTypes = []any{
	(*ListUserTeamsRequest)(nil),  // 0: teams.ListUserTeamsRequest
	(*ListUserTeamsResponse)(nil), // 1: teams.ListUserTeamsResponse
	(*Team)(nil),                  // 2: teams.Team
}
var file_proto_teams_list_user_teams_proto_depIdxs = []int32{
	2, // 0: teams.ListUserTeamsResponse.teams:type_name -> teams.Team
	0, // 1: teams.ListUserTeams.ListUserTeams:input_type -> teams.ListUserTeamsRequest
	1, // 2: teams.ListUserTeams.ListUserTeams:output_type -> teams.ListUserTeamsResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_teams_list_user_teams_proto_init() }
func file_proto_teams_list_user_teams_proto_init() {
	if File_proto_teams_list_user_teams_proto != nil {
		return
	}
	file_proto_teams_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_teams_list_user_teams_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ListUserTeamsRequest); i {
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
		file_proto_teams_list_user_teams_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ListUserTeamsResponse); i {
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
			RawDescriptor: file_proto_teams_list_user_teams_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_teams_list_user_teams_proto_goTypes,
		DependencyIndexes: file_proto_teams_list_user_teams_proto_depIdxs,
		MessageInfos:      file_proto_teams_list_user_teams_proto_msgTypes,
	}.Build()
	File_proto_teams_list_user_teams_proto = out.File
	file_proto_teams_list_user_teams_proto_rawDesc = nil
	file_proto_teams_list_user_teams_proto_goTypes = nil
	file_proto_teams_list_user_teams_proto_depIdxs = nil
}
