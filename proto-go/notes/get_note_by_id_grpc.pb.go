// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: proto/notes/get_note_by_id.proto

package notes_pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	GetNoteByID_GetNoteByID_FullMethodName = "/notes.GetNoteByID/GetNoteByID"
)

// GetNoteByIDClient is the client API for GetNoteByID service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GetNoteByIDClient interface {
	// Retrieves a note.
	GetNoteByID(ctx context.Context, in *GetNoteByIDRequest, opts ...grpc.CallOption) (*Note, error)
}

type getNoteByIDClient struct {
	cc grpc.ClientConnInterface
}

func NewGetNoteByIDClient(cc grpc.ClientConnInterface) GetNoteByIDClient {
	return &getNoteByIDClient{cc}
}

func (c *getNoteByIDClient) GetNoteByID(ctx context.Context, in *GetNoteByIDRequest, opts ...grpc.CallOption) (*Note, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Note)
	err := c.cc.Invoke(ctx, GetNoteByID_GetNoteByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetNoteByIDServer is the server API for GetNoteByID service.
// All implementations must embed UnimplementedGetNoteByIDServer
// for forward compatibility.
type GetNoteByIDServer interface {
	// Retrieves a note.
	GetNoteByID(context.Context, *GetNoteByIDRequest) (*Note, error)
	mustEmbedUnimplementedGetNoteByIDServer()
}

// UnimplementedGetNoteByIDServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGetNoteByIDServer struct{}

func (UnimplementedGetNoteByIDServer) GetNoteByID(context.Context, *GetNoteByIDRequest) (*Note, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNoteByID not implemented")
}
func (UnimplementedGetNoteByIDServer) mustEmbedUnimplementedGetNoteByIDServer() {}
func (UnimplementedGetNoteByIDServer) testEmbeddedByValue()                     {}

// UnsafeGetNoteByIDServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GetNoteByIDServer will
// result in compilation errors.
type UnsafeGetNoteByIDServer interface {
	mustEmbedUnimplementedGetNoteByIDServer()
}

func RegisterGetNoteByIDServer(s grpc.ServiceRegistrar, srv GetNoteByIDServer) {
	// If the following call pancis, it indicates UnimplementedGetNoteByIDServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&GetNoteByID_ServiceDesc, srv)
}

func _GetNoteByID_GetNoteByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNoteByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetNoteByIDServer).GetNoteByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GetNoteByID_GetNoteByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetNoteByIDServer).GetNoteByID(ctx, req.(*GetNoteByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GetNoteByID_ServiceDesc is the grpc.ServiceDesc for GetNoteByID service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GetNoteByID_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "notes.GetNoteByID",
	HandlerType: (*GetNoteByIDServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNoteByID",
			Handler:    _GetNoteByID_GetNoteByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/notes/get_note_by_id.proto",
}
