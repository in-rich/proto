// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.25.4
// source: proto/teams/list_user_teams.proto

package teams_pb

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
	ListUserTeams_ListUserTeams_FullMethodName = "/teams.ListUserTeams/ListUserTeams"
)

// ListUserTeamsClient is the client API for ListUserTeams service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ListUserTeamsClient interface {
	// Retrieves the teams that a user is a member of.
	ListUserTeams(ctx context.Context, in *ListUserTeamsRequest, opts ...grpc.CallOption) (*ListUserTeamsResponse, error)
}

type listUserTeamsClient struct {
	cc grpc.ClientConnInterface
}

func NewListUserTeamsClient(cc grpc.ClientConnInterface) ListUserTeamsClient {
	return &listUserTeamsClient{cc}
}

func (c *listUserTeamsClient) ListUserTeams(ctx context.Context, in *ListUserTeamsRequest, opts ...grpc.CallOption) (*ListUserTeamsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListUserTeamsResponse)
	err := c.cc.Invoke(ctx, ListUserTeams_ListUserTeams_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ListUserTeamsServer is the server API for ListUserTeams service.
// All implementations must embed UnimplementedListUserTeamsServer
// for forward compatibility.
type ListUserTeamsServer interface {
	// Retrieves the teams that a user is a member of.
	ListUserTeams(context.Context, *ListUserTeamsRequest) (*ListUserTeamsResponse, error)
	mustEmbedUnimplementedListUserTeamsServer()
}

// UnimplementedListUserTeamsServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedListUserTeamsServer struct{}

func (UnimplementedListUserTeamsServer) ListUserTeams(context.Context, *ListUserTeamsRequest) (*ListUserTeamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUserTeams not implemented")
}
func (UnimplementedListUserTeamsServer) mustEmbedUnimplementedListUserTeamsServer() {}
func (UnimplementedListUserTeamsServer) testEmbeddedByValue()                       {}

// UnsafeListUserTeamsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ListUserTeamsServer will
// result in compilation errors.
type UnsafeListUserTeamsServer interface {
	mustEmbedUnimplementedListUserTeamsServer()
}

func RegisterListUserTeamsServer(s grpc.ServiceRegistrar, srv ListUserTeamsServer) {
	// If the following call pancis, it indicates UnimplementedListUserTeamsServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ListUserTeams_ServiceDesc, srv)
}

func _ListUserTeams_ListUserTeams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserTeamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListUserTeamsServer).ListUserTeams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ListUserTeams_ListUserTeams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListUserTeamsServer).ListUserTeams(ctx, req.(*ListUserTeamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ListUserTeams_ServiceDesc is the grpc.ServiceDesc for ListUserTeams service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ListUserTeams_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "teams.ListUserTeams",
	HandlerType: (*ListUserTeamsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListUserTeams",
			Handler:    _ListUserTeams_ListUserTeams_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/teams/list_user_teams.proto",
}
