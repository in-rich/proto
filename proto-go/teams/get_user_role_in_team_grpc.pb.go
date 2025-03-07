// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: proto/teams/get_user_role_in_team.proto

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
	GetUserRoleInTeam_GetUserRoleInTeam_FullMethodName = "/teams.GetUserRoleInTeam/GetUserRoleInTeam"
)

// GetUserRoleInTeamClient is the client API for GetUserRoleInTeam service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GetUserRoleInTeamClient interface {
	// Retrieves the role of a user in a team.
	GetUserRoleInTeam(ctx context.Context, in *GetUserRoleInTeamRequest, opts ...grpc.CallOption) (*GetUserRoleInTeamResponse, error)
}

type getUserRoleInTeamClient struct {
	cc grpc.ClientConnInterface
}

func NewGetUserRoleInTeamClient(cc grpc.ClientConnInterface) GetUserRoleInTeamClient {
	return &getUserRoleInTeamClient{cc}
}

func (c *getUserRoleInTeamClient) GetUserRoleInTeam(ctx context.Context, in *GetUserRoleInTeamRequest, opts ...grpc.CallOption) (*GetUserRoleInTeamResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserRoleInTeamResponse)
	err := c.cc.Invoke(ctx, GetUserRoleInTeam_GetUserRoleInTeam_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetUserRoleInTeamServer is the server API for GetUserRoleInTeam service.
// All implementations must embed UnimplementedGetUserRoleInTeamServer
// for forward compatibility.
type GetUserRoleInTeamServer interface {
	// Retrieves the role of a user in a team.
	GetUserRoleInTeam(context.Context, *GetUserRoleInTeamRequest) (*GetUserRoleInTeamResponse, error)
	mustEmbedUnimplementedGetUserRoleInTeamServer()
}

// UnimplementedGetUserRoleInTeamServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGetUserRoleInTeamServer struct{}

func (UnimplementedGetUserRoleInTeamServer) GetUserRoleInTeam(context.Context, *GetUserRoleInTeamRequest) (*GetUserRoleInTeamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserRoleInTeam not implemented")
}
func (UnimplementedGetUserRoleInTeamServer) mustEmbedUnimplementedGetUserRoleInTeamServer() {}
func (UnimplementedGetUserRoleInTeamServer) testEmbeddedByValue()                           {}

// UnsafeGetUserRoleInTeamServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GetUserRoleInTeamServer will
// result in compilation errors.
type UnsafeGetUserRoleInTeamServer interface {
	mustEmbedUnimplementedGetUserRoleInTeamServer()
}

func RegisterGetUserRoleInTeamServer(s grpc.ServiceRegistrar, srv GetUserRoleInTeamServer) {
	// If the following call pancis, it indicates UnimplementedGetUserRoleInTeamServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&GetUserRoleInTeam_ServiceDesc, srv)
}

func _GetUserRoleInTeam_GetUserRoleInTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRoleInTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetUserRoleInTeamServer).GetUserRoleInTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GetUserRoleInTeam_GetUserRoleInTeam_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetUserRoleInTeamServer).GetUserRoleInTeam(ctx, req.(*GetUserRoleInTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GetUserRoleInTeam_ServiceDesc is the grpc.ServiceDesc for GetUserRoleInTeam service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GetUserRoleInTeam_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "teams.GetUserRoleInTeam",
	HandlerType: (*GetUserRoleInTeamServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserRoleInTeam",
			Handler:    _GetUserRoleInTeam_GetUserRoleInTeam_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/teams/get_user_role_in_team.proto",
}
