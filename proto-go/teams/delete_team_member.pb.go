// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0
// source: proto/teams/delete_team_member.proto

package teams_pb

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

type DeleteTeamMemberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The id of the team.
	TeamId string `protobuf:"bytes,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	// The id of the user.
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *DeleteTeamMemberRequest) Reset() {
	*x = DeleteTeamMemberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_teams_delete_team_member_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTeamMemberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTeamMemberRequest) ProtoMessage() {}

func (x *DeleteTeamMemberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_teams_delete_team_member_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTeamMemberRequest.ProtoReflect.Descriptor instead.
func (*DeleteTeamMemberRequest) Descriptor() ([]byte, []int) {
	return file_proto_teams_delete_team_member_proto_rawDescGZIP(), []int{0}
}

func (x *DeleteTeamMemberRequest) GetTeamId() string {
	if x != nil {
		return x.TeamId
	}
	return ""
}

func (x *DeleteTeamMemberRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

var File_proto_teams_delete_team_member_proto protoreflect.FileDescriptor

var file_proto_teams_delete_team_member_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x2f, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x5f, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4b, 0x0a, 0x17, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x32, 0x60, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x54, 0x65, 0x61, 0x6d, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x4c, 0x0a, 0x10, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x1e, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65,
	0x61, 0x6d, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x19, 0x5a, 0x17, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x3b, 0x74, 0x65, 0x61, 0x6d,
	0x73, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_teams_delete_team_member_proto_rawDescOnce sync.Once
	file_proto_teams_delete_team_member_proto_rawDescData = file_proto_teams_delete_team_member_proto_rawDesc
)

func file_proto_teams_delete_team_member_proto_rawDescGZIP() []byte {
	file_proto_teams_delete_team_member_proto_rawDescOnce.Do(func() {
		file_proto_teams_delete_team_member_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_teams_delete_team_member_proto_rawDescData)
	})
	return file_proto_teams_delete_team_member_proto_rawDescData
}

var file_proto_teams_delete_team_member_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_teams_delete_team_member_proto_goTypes = []any{
	(*DeleteTeamMemberRequest)(nil), // 0: teams.DeleteTeamMemberRequest
	(*emptypb.Empty)(nil),           // 1: google.protobuf.Empty
}
var file_proto_teams_delete_team_member_proto_depIdxs = []int32{
	0, // 0: teams.DeleteTeamMember.DeleteTeamMember:input_type -> teams.DeleteTeamMemberRequest
	1, // 1: teams.DeleteTeamMember.DeleteTeamMember:output_type -> google.protobuf.Empty
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_teams_delete_team_member_proto_init() }
func file_proto_teams_delete_team_member_proto_init() {
	if File_proto_teams_delete_team_member_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_teams_delete_team_member_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteTeamMemberRequest); i {
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
			RawDescriptor: file_proto_teams_delete_team_member_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_teams_delete_team_member_proto_goTypes,
		DependencyIndexes: file_proto_teams_delete_team_member_proto_depIdxs,
		MessageInfos:      file_proto_teams_delete_team_member_proto_msgTypes,
	}.Build()
	File_proto_teams_delete_team_member_proto = out.File
	file_proto_teams_delete_team_member_proto_rawDesc = nil
	file_proto_teams_delete_team_member_proto_goTypes = nil
	file_proto_teams_delete_team_member_proto_depIdxs = nil
}
