// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: proto/search/upsert_note.proto

package search_pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UpsertNoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The id of the note.
	NoteId string `protobuf:"bytes,1,opt,name=note_id,json=noteId,proto3" json:"note_id,omitempty"`
	// The id of the note's author.
	AuthorId string `protobuf:"bytes,2,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	// The note content.
	Content string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	// The vanity name of the LinkedIn profile, as it appears in its LinkedIn profile URL.
	PublicIdentifier string `protobuf:"bytes,4,opt,name=public_identifier,json=publicIdentifier,proto3" json:"public_identifier,omitempty"`
	// The full name of the target, as it appears in LinkedIn.
	TargetName string                 `protobuf:"bytes,5,opt,name=target_name,json=targetName,proto3" json:"target_name,omitempty"`
	UpdatedAt  *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3,oneof" json:"updated_at,omitempty"`
}

func (x *UpsertNoteRequest) Reset() {
	*x = UpsertNoteRequest{}
	mi := &file_proto_search_upsert_note_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpsertNoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertNoteRequest) ProtoMessage() {}

func (x *UpsertNoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_search_upsert_note_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertNoteRequest.ProtoReflect.Descriptor instead.
func (*UpsertNoteRequest) Descriptor() ([]byte, []int) {
	return file_proto_search_upsert_note_proto_rawDescGZIP(), []int{0}
}

func (x *UpsertNoteRequest) GetNoteId() string {
	if x != nil {
		return x.NoteId
	}
	return ""
}

func (x *UpsertNoteRequest) GetAuthorId() string {
	if x != nil {
		return x.AuthorId
	}
	return ""
}

func (x *UpsertNoteRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *UpsertNoteRequest) GetPublicIdentifier() string {
	if x != nil {
		return x.PublicIdentifier
	}
	return ""
}

func (x *UpsertNoteRequest) GetTargetName() string {
	if x != nil {
		return x.TargetName
	}
	return ""
}

func (x *UpsertNoteRequest) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type UpsertNoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The saved value of the note. It returns nothing if the note was deleted (no content).
	Note *Note `protobuf:"bytes,1,opt,name=note,proto3,oneof" json:"note,omitempty"`
}

func (x *UpsertNoteResponse) Reset() {
	*x = UpsertNoteResponse{}
	mi := &file_proto_search_upsert_note_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpsertNoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertNoteResponse) ProtoMessage() {}

func (x *UpsertNoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_search_upsert_note_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertNoteResponse.ProtoReflect.Descriptor instead.
func (*UpsertNoteResponse) Descriptor() ([]byte, []int) {
	return file_proto_search_upsert_note_proto_rawDescGZIP(), []int{1}
}

func (x *UpsertNoteResponse) GetNote() *Note {
	if x != nil {
		return x.Note
	}
	return nil
}

var File_proto_search_upsert_note_proto protoreflect.FileDescriptor

var file_proto_search_upsert_note_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2f, 0x75,
	0x70, 0x73, 0x65, 0x72, 0x74, 0x5f, 0x6e, 0x6f, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x1a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x80, 0x02, 0x0a, 0x11, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x4e,
	0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f,
	0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x74,
	0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x2b, 0x0a, 0x11, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3e, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x00, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x22, 0x44, 0x0a, 0x12, 0x55, 0x70, 0x73, 0x65, 0x72,
	0x74, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a,
	0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x6f, 0x74,
	0x65, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x6f, 0x74, 0x65, 0x32, 0x53, 0x0a,
	0x0a, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x12, 0x45, 0x0a, 0x0a, 0x55,
	0x70, 0x73, 0x65, 0x72, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x55, 0x70,
	0x73, 0x65, 0x72, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x1b, 0x5a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x6f, 0x2f, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x3b, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_search_upsert_note_proto_rawDescOnce sync.Once
	file_proto_search_upsert_note_proto_rawDescData = file_proto_search_upsert_note_proto_rawDesc
)

func file_proto_search_upsert_note_proto_rawDescGZIP() []byte {
	file_proto_search_upsert_note_proto_rawDescOnce.Do(func() {
		file_proto_search_upsert_note_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_search_upsert_note_proto_rawDescData)
	})
	return file_proto_search_upsert_note_proto_rawDescData
}

var file_proto_search_upsert_note_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_search_upsert_note_proto_goTypes = []any{
	(*UpsertNoteRequest)(nil),     // 0: search.UpsertNoteRequest
	(*UpsertNoteResponse)(nil),    // 1: search.UpsertNoteResponse
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
	(*Note)(nil),                  // 3: search.Note
}
var file_proto_search_upsert_note_proto_depIdxs = []int32{
	2, // 0: search.UpsertNoteRequest.updated_at:type_name -> google.protobuf.Timestamp
	3, // 1: search.UpsertNoteResponse.note:type_name -> search.Note
	0, // 2: search.UpsertNote.UpsertNote:input_type -> search.UpsertNoteRequest
	1, // 3: search.UpsertNote.UpsertNote:output_type -> search.UpsertNoteResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_search_upsert_note_proto_init() }
func file_proto_search_upsert_note_proto_init() {
	if File_proto_search_upsert_note_proto != nil {
		return
	}
	file_proto_search_common_proto_init()
	file_proto_search_upsert_note_proto_msgTypes[0].OneofWrappers = []any{}
	file_proto_search_upsert_note_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_search_upsert_note_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_search_upsert_note_proto_goTypes,
		DependencyIndexes: file_proto_search_upsert_note_proto_depIdxs,
		MessageInfos:      file_proto_search_upsert_note_proto_msgTypes,
	}.Build()
	File_proto_search_upsert_note_proto = out.File
	file_proto_search_upsert_note_proto_rawDesc = nil
	file_proto_search_upsert_note_proto_goTypes = nil
	file_proto_search_upsert_note_proto_depIdxs = nil
}
