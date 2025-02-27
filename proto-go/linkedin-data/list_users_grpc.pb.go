// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: proto/linkedin-data/list_users.proto

package linkedin_data_pb

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
	ListUsers_ListUsers_FullMethodName = "/linkedinData.ListUsers/ListUsers"
)

// ListUsersClient is the client API for ListUsers service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ListUsersClient interface {
	// Retrieves a list of users LinkedIn profile data.
	ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error)
}

type listUsersClient struct {
	cc grpc.ClientConnInterface
}

func NewListUsersClient(cc grpc.ClientConnInterface) ListUsersClient {
	return &listUsersClient{cc}
}

func (c *listUsersClient) ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListUsersResponse)
	err := c.cc.Invoke(ctx, ListUsers_ListUsers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ListUsersServer is the server API for ListUsers service.
// All implementations must embed UnimplementedListUsersServer
// for forward compatibility.
type ListUsersServer interface {
	// Retrieves a list of users LinkedIn profile data.
	ListUsers(context.Context, *ListUsersRequest) (*ListUsersResponse, error)
	mustEmbedUnimplementedListUsersServer()
}

// UnimplementedListUsersServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedListUsersServer struct{}

func (UnimplementedListUsersServer) ListUsers(context.Context, *ListUsersRequest) (*ListUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsers not implemented")
}
func (UnimplementedListUsersServer) mustEmbedUnimplementedListUsersServer() {}
func (UnimplementedListUsersServer) testEmbeddedByValue()                   {}

// UnsafeListUsersServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ListUsersServer will
// result in compilation errors.
type UnsafeListUsersServer interface {
	mustEmbedUnimplementedListUsersServer()
}

func RegisterListUsersServer(s grpc.ServiceRegistrar, srv ListUsersServer) {
	// If the following call pancis, it indicates UnimplementedListUsersServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ListUsers_ServiceDesc, srv)
}

func _ListUsers_ListUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListUsersServer).ListUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ListUsers_ListUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListUsersServer).ListUsers(ctx, req.(*ListUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ListUsers_ServiceDesc is the grpc.ServiceDesc for ListUsers service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ListUsers_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "linkedinData.ListUsers",
	HandlerType: (*ListUsersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListUsers",
			Handler:    _ListUsers_ListUsers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/linkedin-data/list_users.proto",
}
