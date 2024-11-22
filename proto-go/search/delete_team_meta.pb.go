// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.4
// source: proto/search/delete_team_meta.proto

package search_pb

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

type DeleteTeamMetaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the team.
	TeamId string `protobuf:"bytes,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	// The ID of the user.
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *DeleteTeamMetaRequest) Reset() {
	*x = DeleteTeamMetaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_search_delete_team_meta_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTeamMetaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTeamMetaRequest) ProtoMessage() {}

func (x *DeleteTeamMetaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_search_delete_team_meta_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTeamMetaRequest.ProtoReflect.Descriptor instead.
func (*DeleteTeamMetaRequest) Descriptor() ([]byte, []int) {
	return file_proto_search_delete_team_meta_proto_rawDescGZIP(), []int{0}
}

func (x *DeleteTeamMetaRequest) GetTeamId() string {
	if x != nil {
		return x.TeamId
	}
	return ""
}

func (x *DeleteTeamMetaRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

var File_proto_search_delete_team_meta_proto protoreflect.FileDescriptor

var file_proto_search_delete_team_meta_proto_rawDesc = []byte{
	0x0a, 0x23, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2f, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x49, 0x0a, 0x15, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x32, 0x5b, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54,
	0x65, 0x61, 0x6d, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x49, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x54, 0x65, 0x61, 0x6d, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x1d, 0x2e, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x4d, 0x65, 0x74,
	0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x42, 0x1b, 0x5a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x6f, 0x2f, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x3b, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_search_delete_team_meta_proto_rawDescOnce sync.Once
	file_proto_search_delete_team_meta_proto_rawDescData = file_proto_search_delete_team_meta_proto_rawDesc
)

func file_proto_search_delete_team_meta_proto_rawDescGZIP() []byte {
	file_proto_search_delete_team_meta_proto_rawDescOnce.Do(func() {
		file_proto_search_delete_team_meta_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_search_delete_team_meta_proto_rawDescData)
	})
	return file_proto_search_delete_team_meta_proto_rawDescData
}

var file_proto_search_delete_team_meta_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_search_delete_team_meta_proto_goTypes = []any{
	(*DeleteTeamMetaRequest)(nil), // 0: search.DeleteTeamMetaRequest
	(*emptypb.Empty)(nil),         // 1: google.protobuf.Empty
}
var file_proto_search_delete_team_meta_proto_depIdxs = []int32{
	0, // 0: search.DeleteTeamMeta.DeleteTeamMeta:input_type -> search.DeleteTeamMetaRequest
	1, // 1: search.DeleteTeamMeta.DeleteTeamMeta:output_type -> google.protobuf.Empty
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_search_delete_team_meta_proto_init() }
func file_proto_search_delete_team_meta_proto_init() {
	if File_proto_search_delete_team_meta_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_search_delete_team_meta_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteTeamMetaRequest); i {
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
			RawDescriptor: file_proto_search_delete_team_meta_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_search_delete_team_meta_proto_goTypes,
		DependencyIndexes: file_proto_search_delete_team_meta_proto_depIdxs,
		MessageInfos:      file_proto_search_delete_team_meta_proto_msgTypes,
	}.Build()
	File_proto_search_delete_team_meta_proto = out.File
	file_proto_search_delete_team_meta_proto_rawDesc = nil
	file_proto_search_delete_team_meta_proto_goTypes = nil
	file_proto_search_delete_team_meta_proto_depIdxs = nil
}
