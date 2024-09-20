// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.25.4
// source: proto/notes/list_notes_by_id.proto

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
	ListNotesByID_ListNotesByID_FullMethodName = "/notes.ListNotesByID/ListNotesByID"
)

// ListNotesByIDClient is the client API for ListNotesByID service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ListNotesByIDClient interface {
	// Retrieves a note.
	ListNotesByID(ctx context.Context, in *ListNotesByIDRequest, opts ...grpc.CallOption) (*ListNotesByIDResponse, error)
}

type listNotesByIDClient struct {
	cc grpc.ClientConnInterface
}

func NewListNotesByIDClient(cc grpc.ClientConnInterface) ListNotesByIDClient {
	return &listNotesByIDClient{cc}
}

func (c *listNotesByIDClient) ListNotesByID(ctx context.Context, in *ListNotesByIDRequest, opts ...grpc.CallOption) (*ListNotesByIDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListNotesByIDResponse)
	err := c.cc.Invoke(ctx, ListNotesByID_ListNotesByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ListNotesByIDServer is the server API for ListNotesByID service.
// All implementations must embed UnimplementedListNotesByIDServer
// for forward compatibility.
type ListNotesByIDServer interface {
	// Retrieves a note.
	ListNotesByID(context.Context, *ListNotesByIDRequest) (*ListNotesByIDResponse, error)
	mustEmbedUnimplementedListNotesByIDServer()
}

// UnimplementedListNotesByIDServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedListNotesByIDServer struct{}

func (UnimplementedListNotesByIDServer) ListNotesByID(context.Context, *ListNotesByIDRequest) (*ListNotesByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNotesByID not implemented")
}
func (UnimplementedListNotesByIDServer) mustEmbedUnimplementedListNotesByIDServer() {}
func (UnimplementedListNotesByIDServer) testEmbeddedByValue()                       {}

// UnsafeListNotesByIDServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ListNotesByIDServer will
// result in compilation errors.
type UnsafeListNotesByIDServer interface {
	mustEmbedUnimplementedListNotesByIDServer()
}

func RegisterListNotesByIDServer(s grpc.ServiceRegistrar, srv ListNotesByIDServer) {
	// If the following call pancis, it indicates UnimplementedListNotesByIDServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ListNotesByID_ServiceDesc, srv)
}

func _ListNotesByID_ListNotesByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNotesByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListNotesByIDServer).ListNotesByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ListNotesByID_ListNotesByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListNotesByIDServer).ListNotesByID(ctx, req.(*ListNotesByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ListNotesByID_ServiceDesc is the grpc.ServiceDesc for ListNotesByID service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ListNotesByID_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "notes.ListNotesByID",
	HandlerType: (*ListNotesByIDServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListNotesByID",
			Handler:    _ListNotesByID_ListNotesByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/notes/list_notes_by_id.proto",
}