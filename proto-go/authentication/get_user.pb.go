// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.21.12
// source: proto/authentication/get_user.proto

package authentication_pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type GetUserRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The firebase user ID of the user.
	FirebaseUid   string `protobuf:"bytes,1,opt,name=firebase_uid,json=firebaseUid,proto3" json:"firebase_uid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserRequest) Reset() {
	*x = GetUserRequest{}
	mi := &file_proto_authentication_get_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRequest) ProtoMessage() {}

func (x *GetUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_authentication_get_user_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRequest.ProtoReflect.Descriptor instead.
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return file_proto_authentication_get_user_proto_rawDescGZIP(), []int{0}
}

func (x *GetUserRequest) GetFirebaseUid() string {
	if x != nil {
		return x.FirebaseUid
	}
	return ""
}

var File_proto_authentication_get_user_proto protoreflect.FileDescriptor

var file_proto_authentication_get_user_proto_rawDesc = string([]byte{
	0x0a, 0x23, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x33, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x69,
	0x72, 0x65, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x66, 0x69, 0x72, 0x65, 0x62, 0x61, 0x73, 0x65, 0x55, 0x69, 0x64, 0x32, 0x4c, 0x0a,
	0x07, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x41, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x1e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x00, 0x42, 0x2b, 0x5a, 0x29, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3b, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_proto_authentication_get_user_proto_rawDescOnce sync.Once
	file_proto_authentication_get_user_proto_rawDescData []byte
)

func file_proto_authentication_get_user_proto_rawDescGZIP() []byte {
	file_proto_authentication_get_user_proto_rawDescOnce.Do(func() {
		file_proto_authentication_get_user_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_authentication_get_user_proto_rawDesc), len(file_proto_authentication_get_user_proto_rawDesc)))
	})
	return file_proto_authentication_get_user_proto_rawDescData
}

var file_proto_authentication_get_user_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_authentication_get_user_proto_goTypes = []any{
	(*GetUserRequest)(nil), // 0: authentication.GetUserRequest
	(*User)(nil),           // 1: authentication.User
}
var file_proto_authentication_get_user_proto_depIdxs = []int32{
	0, // 0: authentication.GetUser.GetUser:input_type -> authentication.GetUserRequest
	1, // 1: authentication.GetUser.GetUser:output_type -> authentication.User
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_authentication_get_user_proto_init() }
func file_proto_authentication_get_user_proto_init() {
	if File_proto_authentication_get_user_proto != nil {
		return
	}
	file_proto_authentication_common_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_authentication_get_user_proto_rawDesc), len(file_proto_authentication_get_user_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_authentication_get_user_proto_goTypes,
		DependencyIndexes: file_proto_authentication_get_user_proto_depIdxs,
		MessageInfos:      file_proto_authentication_get_user_proto_msgTypes,
	}.Build()
	File_proto_authentication_get_user_proto = out.File
	file_proto_authentication_get_user_proto_goTypes = nil
	file_proto_authentication_get_user_proto_depIdxs = nil
}
