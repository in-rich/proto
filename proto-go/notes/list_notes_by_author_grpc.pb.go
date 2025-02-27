// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: proto/notes/list_notes_by_author.proto

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
	ListNotesByAuthor_ListNotesByAuthor_FullMethodName = "/notes.ListNotesByAuthor/ListNotesByAuthor"
)

// ListNotesByAuthorClient is the client API for ListNotesByAuthor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ListNotesByAuthorClient interface {
	// Retrieves a note.
	ListNotesByAuthor(ctx context.Context, in *ListNotesByAuthorRequest, opts ...grpc.CallOption) (*ListNotesByAuthorResponse, error)
}

type listNotesByAuthorClient struct {
	cc grpc.ClientConnInterface
}

func NewListNotesByAuthorClient(cc grpc.ClientConnInterface) ListNotesByAuthorClient {
	return &listNotesByAuthorClient{cc}
}

func (c *listNotesByAuthorClient) ListNotesByAuthor(ctx context.Context, in *ListNotesByAuthorRequest, opts ...grpc.CallOption) (*ListNotesByAuthorResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListNotesByAuthorResponse)
	err := c.cc.Invoke(ctx, ListNotesByAuthor_ListNotesByAuthor_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ListNotesByAuthorServer is the server API for ListNotesByAuthor service.
// All implementations must embed UnimplementedListNotesByAuthorServer
// for forward compatibility.
type ListNotesByAuthorServer interface {
	// Retrieves a note.
	ListNotesByAuthor(context.Context, *ListNotesByAuthorRequest) (*ListNotesByAuthorResponse, error)
	mustEmbedUnimplementedListNotesByAuthorServer()
}

// UnimplementedListNotesByAuthorServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedListNotesByAuthorServer struct{}

func (UnimplementedListNotesByAuthorServer) ListNotesByAuthor(context.Context, *ListNotesByAuthorRequest) (*ListNotesByAuthorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNotesByAuthor not implemented")
}
func (UnimplementedListNotesByAuthorServer) mustEmbedUnimplementedListNotesByAuthorServer() {}
func (UnimplementedListNotesByAuthorServer) testEmbeddedByValue()                           {}

// UnsafeListNotesByAuthorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ListNotesByAuthorServer will
// result in compilation errors.
type UnsafeListNotesByAuthorServer interface {
	mustEmbedUnimplementedListNotesByAuthorServer()
}

func RegisterListNotesByAuthorServer(s grpc.ServiceRegistrar, srv ListNotesByAuthorServer) {
	// If the following call pancis, it indicates UnimplementedListNotesByAuthorServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ListNotesByAuthor_ServiceDesc, srv)
}

func _ListNotesByAuthor_ListNotesByAuthor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNotesByAuthorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListNotesByAuthorServer).ListNotesByAuthor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ListNotesByAuthor_ListNotesByAuthor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListNotesByAuthorServer).ListNotesByAuthor(ctx, req.(*ListNotesByAuthorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ListNotesByAuthor_ServiceDesc is the grpc.ServiceDesc for ListNotesByAuthor service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ListNotesByAuthor_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "notes.ListNotesByAuthor",
	HandlerType: (*ListNotesByAuthorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListNotesByAuthor",
			Handler:    _ListNotesByAuthor_ListNotesByAuthor_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/notes/list_notes_by_author.proto",
}
