// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: proto/teams/get_team.proto

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
	GetTeam_GetTeam_FullMethodName = "/teams.GetTeam/GetTeam"
)

// GetTeamClient is the client API for GetTeam service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GetTeamClient interface {
	// Get the info of a team.
	GetTeam(ctx context.Context, in *GetTeamRequest, opts ...grpc.CallOption) (*Team, error)
}

type getTeamClient struct {
	cc grpc.ClientConnInterface
}

func NewGetTeamClient(cc grpc.ClientConnInterface) GetTeamClient {
	return &getTeamClient{cc}
}

func (c *getTeamClient) GetTeam(ctx context.Context, in *GetTeamRequest, opts ...grpc.CallOption) (*Team, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Team)
	err := c.cc.Invoke(ctx, GetTeam_GetTeam_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetTeamServer is the server API for GetTeam service.
// All implementations must embed UnimplementedGetTeamServer
// for forward compatibility.
type GetTeamServer interface {
	// Get the info of a team.
	GetTeam(context.Context, *GetTeamRequest) (*Team, error)
	mustEmbedUnimplementedGetTeamServer()
}

// UnimplementedGetTeamServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGetTeamServer struct{}

func (UnimplementedGetTeamServer) GetTeam(context.Context, *GetTeamRequest) (*Team, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTeam not implemented")
}
func (UnimplementedGetTeamServer) mustEmbedUnimplementedGetTeamServer() {}
func (UnimplementedGetTeamServer) testEmbeddedByValue()                 {}

// UnsafeGetTeamServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GetTeamServer will
// result in compilation errors.
type UnsafeGetTeamServer interface {
	mustEmbedUnimplementedGetTeamServer()
}

func RegisterGetTeamServer(s grpc.ServiceRegistrar, srv GetTeamServer) {
	// If the following call pancis, it indicates UnimplementedGetTeamServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&GetTeam_ServiceDesc, srv)
}

func _GetTeam_GetTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetTeamServer).GetTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GetTeam_GetTeam_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetTeamServer).GetTeam(ctx, req.(*GetTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GetTeam_ServiceDesc is the grpc.ServiceDesc for GetTeam service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GetTeam_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "teams.GetTeam",
	HandlerType: (*GetTeamServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTeam",
			Handler:    _GetTeam_GetTeam_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/teams/get_team.proto",
}
