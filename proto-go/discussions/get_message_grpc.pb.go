// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.2
// source: proto/discussions/get_message.proto

package discussions_pb

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
	GetMessage_GetMessage_FullMethodName = "/discussions.GetMessage/GetMessage"
)

// GetMessageClient is the client API for GetMessage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GetMessageClient interface {
	// Gets a message in a discussion.
	GetMessage(ctx context.Context, in *GetMessageRequest, opts ...grpc.CallOption) (*Message, error)
}

type getMessageClient struct {
	cc grpc.ClientConnInterface
}

func NewGetMessageClient(cc grpc.ClientConnInterface) GetMessageClient {
	return &getMessageClient{cc}
}

func (c *getMessageClient) GetMessage(ctx context.Context, in *GetMessageRequest, opts ...grpc.CallOption) (*Message, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Message)
	err := c.cc.Invoke(ctx, GetMessage_GetMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetMessageServer is the server API for GetMessage service.
// All implementations must embed UnimplementedGetMessageServer
// for forward compatibility
type GetMessageServer interface {
	// Gets a message in a discussion.
	GetMessage(context.Context, *GetMessageRequest) (*Message, error)
	mustEmbedUnimplementedGetMessageServer()
}

// UnimplementedGetMessageServer must be embedded to have forward compatible implementations.
type UnimplementedGetMessageServer struct {
}

func (UnimplementedGetMessageServer) GetMessage(context.Context, *GetMessageRequest) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessage not implemented")
}
func (UnimplementedGetMessageServer) mustEmbedUnimplementedGetMessageServer() {}

// UnsafeGetMessageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GetMessageServer will
// result in compilation errors.
type UnsafeGetMessageServer interface {
	mustEmbedUnimplementedGetMessageServer()
}

func RegisterGetMessageServer(s grpc.ServiceRegistrar, srv GetMessageServer) {
	s.RegisterService(&GetMessage_ServiceDesc, srv)
}

func _GetMessage_GetMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetMessageServer).GetMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GetMessage_GetMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetMessageServer).GetMessage(ctx, req.(*GetMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GetMessage_ServiceDesc is the grpc.ServiceDesc for GetMessage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GetMessage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "discussions.GetMessage",
	HandlerType: (*GetMessageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMessage",
			Handler:    _GetMessage_GetMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/discussions/get_message.proto",
}
