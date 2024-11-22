// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.4
// source: proto/teams/list_team_members.proto

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

type ListTeamMembersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The id of the team.
	TeamId string `protobuf:"bytes,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	Limit  int64  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int64  `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *ListTeamMembersRequest) Reset() {
	*x = ListTeamMembersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_teams_list_team_members_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTeamMembersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTeamMembersRequest) ProtoMessage() {}

func (x *ListTeamMembersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_teams_list_team_members_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTeamMembersRequest.ProtoReflect.Descriptor instead.
func (*ListTeamMembersRequest) Descriptor() ([]byte, []int) {
	return file_proto_teams_list_team_members_proto_rawDescGZIP(), []int{0}
}

func (x *ListTeamMembersRequest) GetTeamId() string {
	if x != nil {
		return x.TeamId
	}
	return ""
}

func (x *ListTeamMembersRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListTeamMembersRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type ListTeamMembersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The members of the team.
	Members []*TeamMember `protobuf:"bytes,1,rep,name=members,proto3" json:"members,omitempty"`
}

func (x *ListTeamMembersResponse) Reset() {
	*x = ListTeamMembersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_teams_list_team_members_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTeamMembersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTeamMembersResponse) ProtoMessage() {}

func (x *ListTeamMembersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_teams_list_team_members_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTeamMembersResponse.ProtoReflect.Descriptor instead.
func (*ListTeamMembersResponse) Descriptor() ([]byte, []int) {
	return file_proto_teams_list_team_members_proto_rawDescGZIP(), []int{1}
}

func (x *ListTeamMembersResponse) GetMembers() []*TeamMember {
	if x != nil {
		return x.Members
	}
	return nil
}

var File_proto_teams_list_team_members_proto protoreflect.FileDescriptor

var file_proto_teams_list_team_members_proto_rawDesc = []byte{
	0x0a, 0x23, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x2f, 0x6c, 0x69,
	0x73, 0x74, 0x5f, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x1a, 0x18, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5f, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65,
	0x61, 0x6d, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x46, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x54,
	0x65, 0x61, 0x6d, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2b, 0x0a, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x54, 0x65, 0x61, 0x6d,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x32,
	0x65, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x73, 0x12, 0x52, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x1d, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x54, 0x65, 0x61, 0x6d, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x19, 0x5a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d,
	0x67, 0x6f, 0x2f, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x3b, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x5f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_teams_list_team_members_proto_rawDescOnce sync.Once
	file_proto_teams_list_team_members_proto_rawDescData = file_proto_teams_list_team_members_proto_rawDesc
)

func file_proto_teams_list_team_members_proto_rawDescGZIP() []byte {
	file_proto_teams_list_team_members_proto_rawDescOnce.Do(func() {
		file_proto_teams_list_team_members_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_teams_list_team_members_proto_rawDescData)
	})
	return file_proto_teams_list_team_members_proto_rawDescData
}

var file_proto_teams_list_team_members_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_teams_list_team_members_proto_goTypes = []any{
	(*ListTeamMembersRequest)(nil),  // 0: teams.ListTeamMembersRequest
	(*ListTeamMembersResponse)(nil), // 1: teams.ListTeamMembersResponse
	(*TeamMember)(nil),              // 2: teams.TeamMember
}
var file_proto_teams_list_team_members_proto_depIdxs = []int32{
	2, // 0: teams.ListTeamMembersResponse.members:type_name -> teams.TeamMember
	0, // 1: teams.ListTeamMembers.ListTeamMembers:input_type -> teams.ListTeamMembersRequest
	1, // 2: teams.ListTeamMembers.ListTeamMembers:output_type -> teams.ListTeamMembersResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_teams_list_team_members_proto_init() }
func file_proto_teams_list_team_members_proto_init() {
	if File_proto_teams_list_team_members_proto != nil {
		return
	}
	file_proto_teams_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_teams_list_team_members_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ListTeamMembersRequest); i {
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
		file_proto_teams_list_team_members_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ListTeamMembersResponse); i {
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
			RawDescriptor: file_proto_teams_list_team_members_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_teams_list_team_members_proto_goTypes,
		DependencyIndexes: file_proto_teams_list_team_members_proto_depIdxs,
		MessageInfos:      file_proto_teams_list_team_members_proto_msgTypes,
	}.Build()
	File_proto_teams_list_team_members_proto = out.File
	file_proto_teams_list_team_members_proto_rawDesc = nil
	file_proto_teams_list_team_members_proto_goTypes = nil
	file_proto_teams_list_team_members_proto_depIdxs = nil
}
