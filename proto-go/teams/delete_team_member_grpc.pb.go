// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: proto/teams/delete_team_member.proto

package teams_pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	DeleteTeamMember_DeleteTeamMember_FullMethodName = "/teams.DeleteTeamMember/DeleteTeamMember"
)

// DeleteTeamMemberClient is the client API for DeleteTeamMember service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeleteTeamMemberClient interface {
	// Deletes a member of a team.
	DeleteTeamMember(ctx context.Context, in *DeleteTeamMemberRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type deleteTeamMemberClient struct {
	cc grpc.ClientConnInterface
}

func NewDeleteTeamMemberClient(cc grpc.ClientConnInterface) DeleteTeamMemberClient {
	return &deleteTeamMemberClient{cc}
}

func (c *deleteTeamMemberClient) DeleteTeamMember(ctx context.Context, in *DeleteTeamMemberRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, DeleteTeamMember_DeleteTeamMember_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeleteTeamMemberServer is the server API for DeleteTeamMember service.
// All implementations must embed UnimplementedDeleteTeamMemberServer
// for forward compatibility.
type DeleteTeamMemberServer interface {
	// Deletes a member of a team.
	DeleteTeamMember(context.Context, *DeleteTeamMemberRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedDeleteTeamMemberServer()
}

// UnimplementedDeleteTeamMemberServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDeleteTeamMemberServer struct{}

func (UnimplementedDeleteTeamMemberServer) DeleteTeamMember(context.Context, *DeleteTeamMemberRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTeamMember not implemented")
}
func (UnimplementedDeleteTeamMemberServer) mustEmbedUnimplementedDeleteTeamMemberServer() {}
func (UnimplementedDeleteTeamMemberServer) testEmbeddedByValue()                          {}

// UnsafeDeleteTeamMemberServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeleteTeamMemberServer will
// result in compilation errors.
type UnsafeDeleteTeamMemberServer interface {
	mustEmbedUnimplementedDeleteTeamMemberServer()
}

func RegisterDeleteTeamMemberServer(s grpc.ServiceRegistrar, srv DeleteTeamMemberServer) {
	// If the following call pancis, it indicates UnimplementedDeleteTeamMemberServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DeleteTeamMember_ServiceDesc, srv)
}

func _DeleteTeamMember_DeleteTeamMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTeamMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeleteTeamMemberServer).DeleteTeamMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeleteTeamMember_DeleteTeamMember_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeleteTeamMemberServer).DeleteTeamMember(ctx, req.(*DeleteTeamMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DeleteTeamMember_ServiceDesc is the grpc.ServiceDesc for DeleteTeamMember service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DeleteTeamMember_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "teams.DeleteTeamMember",
	HandlerType: (*DeleteTeamMemberServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeleteTeamMember",
			Handler:    _DeleteTeamMember_DeleteTeamMember_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/teams/delete_team_member.proto",
}
