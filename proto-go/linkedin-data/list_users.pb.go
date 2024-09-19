// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: proto/linkedin-data/list_users.proto

package linkedin_data_pb

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

type ListUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The vanity name of the users, as they appear in their LinkedIn profile URL.
	PublicIdentifiers []string `protobuf:"bytes,1,rep,name=public_identifiers,json=publicIdentifiers,proto3" json:"public_identifiers,omitempty"`
}

func (x *ListUsersRequest) Reset() {
	*x = ListUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_linkedin_data_list_users_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUsersRequest) ProtoMessage() {}

func (x *ListUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_linkedin_data_list_users_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUsersRequest.ProtoReflect.Descriptor instead.
func (*ListUsersRequest) Descriptor() ([]byte, []int) {
	return file_proto_linkedin_data_list_users_proto_rawDescGZIP(), []int{0}
}

func (x *ListUsersRequest) GetPublicIdentifiers() []string {
	if x != nil {
		return x.PublicIdentifiers
	}
	return nil
}

type ListUsersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of users.
	Users []*User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *ListUsersResponse) Reset() {
	*x = ListUsersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_linkedin_data_list_users_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUsersResponse) ProtoMessage() {}

func (x *ListUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_linkedin_data_list_users_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUsersResponse.ProtoReflect.Descriptor instead.
func (*ListUsersResponse) Descriptor() ([]byte, []int) {
	return file_proto_linkedin_data_list_users_proto_rawDescGZIP(), []int{1}
}

func (x *ListUsersResponse) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

var File_proto_linkedin_data_list_users_proto protoreflect.FileDescriptor

var file_proto_linkedin_data_list_users_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x69, 0x6e,
	0x2d, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x69, 0x6e,
	0x44, 0x61, 0x74, 0x61, 0x1a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x69, 0x6e, 0x6b,
	0x65, 0x64, 0x69, 0x6e, 0x2d, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x41, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x12, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x11, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x22, 0x3d, 0x0a, 0x11, 0x4c, 0x69, 0x73,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28,
	0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x69, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x32, 0x5b, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x4e, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x73, 0x12, 0x1e, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x69, 0x6e, 0x44, 0x61, 0x74,
	0x61, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x69, 0x6e, 0x44, 0x61, 0x74,
	0x61, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x29, 0x5a, 0x27, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67,
	0x6f, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x69, 0x6e, 0x2d, 0x64, 0x61, 0x74, 0x61, 0x3b,
	0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x69, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_linkedin_data_list_users_proto_rawDescOnce sync.Once
	file_proto_linkedin_data_list_users_proto_rawDescData = file_proto_linkedin_data_list_users_proto_rawDesc
)

func file_proto_linkedin_data_list_users_proto_rawDescGZIP() []byte {
	file_proto_linkedin_data_list_users_proto_rawDescOnce.Do(func() {
		file_proto_linkedin_data_list_users_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_linkedin_data_list_users_proto_rawDescData)
	})
	return file_proto_linkedin_data_list_users_proto_rawDescData
}

var file_proto_linkedin_data_list_users_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_linkedin_data_list_users_proto_goTypes = []any{
	(*ListUsersRequest)(nil),  // 0: linkedinData.ListUsersRequest
	(*ListUsersResponse)(nil), // 1: linkedinData.ListUsersResponse
	(*User)(nil),              // 2: linkedinData.User
}
var file_proto_linkedin_data_list_users_proto_depIdxs = []int32{
	2, // 0: linkedinData.ListUsersResponse.users:type_name -> linkedinData.User
	0, // 1: linkedinData.ListUsers.ListUsers:input_type -> linkedinData.ListUsersRequest
	1, // 2: linkedinData.ListUsers.ListUsers:output_type -> linkedinData.ListUsersResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_linkedin_data_list_users_proto_init() }
func file_proto_linkedin_data_list_users_proto_init() {
	if File_proto_linkedin_data_list_users_proto != nil {
		return
	}
	file_proto_linkedin_data_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_linkedin_data_list_users_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ListUsersRequest); i {
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
		file_proto_linkedin_data_list_users_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ListUsersResponse); i {
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
			RawDescriptor: file_proto_linkedin_data_list_users_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_linkedin_data_list_users_proto_goTypes,
		DependencyIndexes: file_proto_linkedin_data_list_users_proto_depIdxs,
		MessageInfos:      file_proto_linkedin_data_list_users_proto_msgTypes,
	}.Build()
	File_proto_linkedin_data_list_users_proto = out.File
	file_proto_linkedin_data_list_users_proto_rawDesc = nil
	file_proto_linkedin_data_list_users_proto_goTypes = nil
	file_proto_linkedin_data_list_users_proto_depIdxs = nil
}
