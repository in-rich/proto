// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.0
// source: proto/discussions/delete_message.proto

package discussions_pb

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
	DeleteMessage_DeleteMessage_FullMethodName = "/discussions.DeleteMessage/DeleteMessage"
)

// DeleteMessageClient is the client API for DeleteMessage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeleteMessageClient interface {
	// Deletes a message in a discussion.
	DeleteMessage(ctx context.Context, in *DeleteMessageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type deleteMessageClient struct {
	cc grpc.ClientConnInterface
}

func NewDeleteMessageClient(cc grpc.ClientConnInterface) DeleteMessageClient {
	return &deleteMessageClient{cc}
}

func (c *deleteMessageClient) DeleteMessage(ctx context.Context, in *DeleteMessageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, DeleteMessage_DeleteMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeleteMessageServer is the server API for DeleteMessage service.
// All implementations must embed UnimplementedDeleteMessageServer
// for forward compatibility.
type DeleteMessageServer interface {
	// Deletes a message in a discussion.
	DeleteMessage(context.Context, *DeleteMessageRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedDeleteMessageServer()
}

// UnimplementedDeleteMessageServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDeleteMessageServer struct{}

func (UnimplementedDeleteMessageServer) DeleteMessage(context.Context, *DeleteMessageRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMessage not implemented")
}
func (UnimplementedDeleteMessageServer) mustEmbedUnimplementedDeleteMessageServer() {}
func (UnimplementedDeleteMessageServer) testEmbeddedByValue()                       {}

// UnsafeDeleteMessageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeleteMessageServer will
// result in compilation errors.
type UnsafeDeleteMessageServer interface {
	mustEmbedUnimplementedDeleteMessageServer()
}

func RegisterDeleteMessageServer(s grpc.ServiceRegistrar, srv DeleteMessageServer) {
	// If the following call pancis, it indicates UnimplementedDeleteMessageServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DeleteMessage_ServiceDesc, srv)
}

func _DeleteMessage_DeleteMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeleteMessageServer).DeleteMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeleteMessage_DeleteMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeleteMessageServer).DeleteMessage(ctx, req.(*DeleteMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DeleteMessage_ServiceDesc is the grpc.ServiceDesc for DeleteMessage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DeleteMessage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "discussions.DeleteMessage",
	HandlerType: (*DeleteMessageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeleteMessage",
			Handler:    _DeleteMessage_DeleteMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/discussions/delete_message.proto",
}
