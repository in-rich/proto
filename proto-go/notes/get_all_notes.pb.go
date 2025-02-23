// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.21.12
// source: proto/notes/get_all_notes.proto

package notes_pb

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

type GetAllNotesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Limit         int64                  `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset        int64                  `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAllNotesRequest) Reset() {
	*x = GetAllNotesRequest{}
	mi := &file_proto_notes_get_all_notes_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAllNotesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllNotesRequest) ProtoMessage() {}

func (x *GetAllNotesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_notes_get_all_notes_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllNotesRequest.ProtoReflect.Descriptor instead.
func (*GetAllNotesRequest) Descriptor() ([]byte, []int) {
	return file_proto_notes_get_all_notes_proto_rawDescGZIP(), []int{0}
}

func (x *GetAllNotesRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetAllNotesRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type GetAllNotesResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The GetAll of retrieved notes.
	Notes         []*Note `protobuf:"bytes,1,rep,name=notes,proto3" json:"notes,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAllNotesResponse) Reset() {
	*x = GetAllNotesResponse{}
	mi := &file_proto_notes_get_all_notes_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAllNotesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllNotesResponse) ProtoMessage() {}

func (x *GetAllNotesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_notes_get_all_notes_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllNotesResponse.ProtoReflect.Descriptor instead.
func (*GetAllNotesResponse) Descriptor() ([]byte, []int) {
	return file_proto_notes_get_all_notes_proto_rawDescGZIP(), []int{1}
}

func (x *GetAllNotesResponse) GetNotes() []*Note {
	if x != nil {
		return x.Notes
	}
	return nil
}

var File_proto_notes_get_all_notes_proto protoreflect.FileDescriptor

var file_proto_notes_get_all_notes_proto_rawDesc = string([]byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2f, 0x67, 0x65,
	0x74, 0x5f, 0x61, 0x6c, 0x6c, 0x5f, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x1a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x42, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4e, 0x6f, 0x74, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x38, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x4e, 0x6f, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a,
	0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x6e,
	0x6f, 0x74, 0x65, 0x73, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73,
	0x32, 0x55, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x12,
	0x46, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x19,
	0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4e, 0x6f, 0x74,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6e, 0x6f, 0x74, 0x65,
	0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x19, 0x5a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2d, 0x67, 0x6f, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x3b, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x5f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_proto_notes_get_all_notes_proto_rawDescOnce sync.Once
	file_proto_notes_get_all_notes_proto_rawDescData []byte
)

func file_proto_notes_get_all_notes_proto_rawDescGZIP() []byte {
	file_proto_notes_get_all_notes_proto_rawDescOnce.Do(func() {
		file_proto_notes_get_all_notes_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_notes_get_all_notes_proto_rawDesc), len(file_proto_notes_get_all_notes_proto_rawDesc)))
	})
	return file_proto_notes_get_all_notes_proto_rawDescData
}

var file_proto_notes_get_all_notes_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_notes_get_all_notes_proto_goTypes = []any{
	(*GetAllNotesRequest)(nil),  // 0: notes.GetAllNotesRequest
	(*GetAllNotesResponse)(nil), // 1: notes.GetAllNotesResponse
	(*Note)(nil),                // 2: notes.Note
}
var file_proto_notes_get_all_notes_proto_depIdxs = []int32{
	2, // 0: notes.GetAllNotesResponse.notes:type_name -> notes.Note
	0, // 1: notes.GetAllNotes.GetAllNotes:input_type -> notes.GetAllNotesRequest
	1, // 2: notes.GetAllNotes.GetAllNotes:output_type -> notes.GetAllNotesResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_notes_get_all_notes_proto_init() }
func file_proto_notes_get_all_notes_proto_init() {
	if File_proto_notes_get_all_notes_proto != nil {
		return
	}
	file_proto_notes_common_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_notes_get_all_notes_proto_rawDesc), len(file_proto_notes_get_all_notes_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_notes_get_all_notes_proto_goTypes,
		DependencyIndexes: file_proto_notes_get_all_notes_proto_depIdxs,
		MessageInfos:      file_proto_notes_get_all_notes_proto_msgTypes,
	}.Build()
	File_proto_notes_get_all_notes_proto = out.File
	file_proto_notes_get_all_notes_proto_goTypes = nil
	file_proto_notes_get_all_notes_proto_depIdxs = nil
}
