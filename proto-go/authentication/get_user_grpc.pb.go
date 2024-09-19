// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.2
// source: proto/authentication/get_user.proto

package authentication_pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	GetUser_GetUser_FullMethodName = "/authentication.GetUser/GetUser"
)

// GetUserClient is the client API for GetUser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GetUserClient interface {
	// Retrieves a user's data.
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error)
}

type getUserClient struct {
	cc grpc.ClientConnInterface
}

func NewGetUserClient(cc grpc.ClientConnInterface) GetUserClient {
	return &getUserClient{cc}
}

func (c *getUserClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(User)
	err := c.cc.Invoke(ctx, GetUser_GetUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetUserServer is the server API for GetUser service.
// All implementations must embed UnimplementedGetUserServer
// for forward compatibility
type GetUserServer interface {
	// Retrieves a user's data.
	GetUser(context.Context, *GetUserRequest) (*User, error)
	mustEmbedUnimplementedGetUserServer()
}

// UnimplementedGetUserServer must be embedded to have forward compatible implementations.
type UnimplementedGetUserServer struct {
}

func (UnimplementedGetUserServer) GetUser(context.Context, *GetUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedGetUserServer) mustEmbedUnimplementedGetUserServer() {}

// UnsafeGetUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GetUserServer will
// result in compilation errors.
type UnsafeGetUserServer interface {
	mustEmbedUnimplementedGetUserServer()
}

func RegisterGetUserServer(s grpc.ServiceRegistrar, srv GetUserServer) {
	s.RegisterService(&GetUser_ServiceDesc, srv)
}

func _GetUser_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetUserServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GetUser_GetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetUserServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GetUser_ServiceDesc is the grpc.ServiceDesc for GetUser service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GetUser_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "authentication.GetUser",
	HandlerType: (*GetUserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _GetUser_GetUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/authentication/get_user.proto",
}
