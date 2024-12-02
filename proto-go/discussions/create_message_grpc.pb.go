// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: proto/discussions/create_message.proto

package discussions_pb

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
	CreateMessage_CreateMessage_FullMethodName = "/discussions.CreateMessage/CreateMessage"
)

// CreateMessageClient is the client API for CreateMessage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CreateMessageClient interface {
	// Creates a message in a discussion.
	CreateMessage(ctx context.Context, in *CreateMessageRequest, opts ...grpc.CallOption) (*Message, error)
}

type createMessageClient struct {
	cc grpc.ClientConnInterface
}

func NewCreateMessageClient(cc grpc.ClientConnInterface) CreateMessageClient {
	return &createMessageClient{cc}
}

func (c *createMessageClient) CreateMessage(ctx context.Context, in *CreateMessageRequest, opts ...grpc.CallOption) (*Message, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Message)
	err := c.cc.Invoke(ctx, CreateMessage_CreateMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CreateMessageServer is the server API for CreateMessage service.
// All implementations must embed UnimplementedCreateMessageServer
// for forward compatibility.
type CreateMessageServer interface {
	// Creates a message in a discussion.
	CreateMessage(context.Context, *CreateMessageRequest) (*Message, error)
	mustEmbedUnimplementedCreateMessageServer()
}

// UnimplementedCreateMessageServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCreateMessageServer struct{}

func (UnimplementedCreateMessageServer) CreateMessage(context.Context, *CreateMessageRequest) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMessage not implemented")
}
func (UnimplementedCreateMessageServer) mustEmbedUnimplementedCreateMessageServer() {}
func (UnimplementedCreateMessageServer) testEmbeddedByValue()                       {}

// UnsafeCreateMessageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CreateMessageServer will
// result in compilation errors.
type UnsafeCreateMessageServer interface {
	mustEmbedUnimplementedCreateMessageServer()
}

func RegisterCreateMessageServer(s grpc.ServiceRegistrar, srv CreateMessageServer) {
	// If the following call pancis, it indicates UnimplementedCreateMessageServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CreateMessage_ServiceDesc, srv)
}

func _CreateMessage_CreateMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreateMessageServer).CreateMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CreateMessage_CreateMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreateMessageServer).CreateMessage(ctx, req.(*CreateMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CreateMessage_ServiceDesc is the grpc.ServiceDesc for CreateMessage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CreateMessage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "discussions.CreateMessage",
	HandlerType: (*CreateMessageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMessage",
			Handler:    _CreateMessage_CreateMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/discussions/create_message.proto",
}
