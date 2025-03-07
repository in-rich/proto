// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: proto/discussions/list_discussion_messages.proto

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
	ListDiscussionMessages_ListDiscussionMessages_FullMethodName = "/discussions.ListDiscussionMessages/ListDiscussionMessages"
)

// ListDiscussionMessagesClient is the client API for ListDiscussionMessages service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ListDiscussionMessagesClient interface {
	// Retrieves a discussion.
	ListDiscussionMessages(ctx context.Context, in *ListDiscussionMessagesRequest, opts ...grpc.CallOption) (*ListDiscussionMessagesResponse, error)
}

type listDiscussionMessagesClient struct {
	cc grpc.ClientConnInterface
}

func NewListDiscussionMessagesClient(cc grpc.ClientConnInterface) ListDiscussionMessagesClient {
	return &listDiscussionMessagesClient{cc}
}

func (c *listDiscussionMessagesClient) ListDiscussionMessages(ctx context.Context, in *ListDiscussionMessagesRequest, opts ...grpc.CallOption) (*ListDiscussionMessagesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListDiscussionMessagesResponse)
	err := c.cc.Invoke(ctx, ListDiscussionMessages_ListDiscussionMessages_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ListDiscussionMessagesServer is the server API for ListDiscussionMessages service.
// All implementations must embed UnimplementedListDiscussionMessagesServer
// for forward compatibility.
type ListDiscussionMessagesServer interface {
	// Retrieves a discussion.
	ListDiscussionMessages(context.Context, *ListDiscussionMessagesRequest) (*ListDiscussionMessagesResponse, error)
	mustEmbedUnimplementedListDiscussionMessagesServer()
}

// UnimplementedListDiscussionMessagesServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedListDiscussionMessagesServer struct{}

func (UnimplementedListDiscussionMessagesServer) ListDiscussionMessages(context.Context, *ListDiscussionMessagesRequest) (*ListDiscussionMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDiscussionMessages not implemented")
}
func (UnimplementedListDiscussionMessagesServer) mustEmbedUnimplementedListDiscussionMessagesServer() {
}
func (UnimplementedListDiscussionMessagesServer) testEmbeddedByValue() {}

// UnsafeListDiscussionMessagesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ListDiscussionMessagesServer will
// result in compilation errors.
type UnsafeListDiscussionMessagesServer interface {
	mustEmbedUnimplementedListDiscussionMessagesServer()
}

func RegisterListDiscussionMessagesServer(s grpc.ServiceRegistrar, srv ListDiscussionMessagesServer) {
	// If the following call pancis, it indicates UnimplementedListDiscussionMessagesServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ListDiscussionMessages_ServiceDesc, srv)
}

func _ListDiscussionMessages_ListDiscussionMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDiscussionMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListDiscussionMessagesServer).ListDiscussionMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ListDiscussionMessages_ListDiscussionMessages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListDiscussionMessagesServer).ListDiscussionMessages(ctx, req.(*ListDiscussionMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ListDiscussionMessages_ServiceDesc is the grpc.ServiceDesc for ListDiscussionMessages service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ListDiscussionMessages_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "discussions.ListDiscussionMessages",
	HandlerType: (*ListDiscussionMessagesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListDiscussionMessages",
			Handler:    _ListDiscussionMessages_ListDiscussionMessages_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/discussions/list_discussion_messages.proto",
}
