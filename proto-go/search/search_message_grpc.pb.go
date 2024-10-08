// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.25.4
// source: proto/search/search_messages.proto

package search_pb

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
	SearchMessage_SearchMessage_FullMethodName = "/search.SearchMessage/SearchMessage"
)

// SearchMessageClient is the client API for SearchMessage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SearchMessageClient interface {
	SearchMessage(ctx context.Context, in *SearchMessageRequest, opts ...grpc.CallOption) (*Message, error)
}

type searchMessageClient struct {
	cc grpc.ClientConnInterface
}

func NewSearchMessageClient(cc grpc.ClientConnInterface) SearchMessageClient {
	return &searchMessageClient{cc}
}

func (c *searchMessageClient) SearchMessage(ctx context.Context, in *SearchMessageRequest, opts ...grpc.CallOption) (*Message, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Message)
	err := c.cc.Invoke(ctx, SearchMessage_SearchMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchMessageServer is the server API for SearchMessage service.
// All implementations must embed UnimplementedSearchMessageServer
// for forward compatibility.
type SearchMessageServer interface {
	SearchMessage(context.Context, *SearchMessageRequest) (*Message, error)
	mustEmbedUnimplementedSearchMessageServer()
}

// UnimplementedSearchMessageServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSearchMessageServer struct{}

func (UnimplementedSearchMessageServer) SearchMessage(context.Context, *SearchMessageRequest) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchMessage not implemented")
}
func (UnimplementedSearchMessageServer) mustEmbedUnimplementedSearchMessageServer() {}
func (UnimplementedSearchMessageServer) testEmbeddedByValue()                       {}

// UnsafeSearchMessageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SearchMessageServer will
// result in compilation errors.
type UnsafeSearchMessageServer interface {
	mustEmbedUnimplementedSearchMessageServer()
}

func RegisterSearchMessageServer(s grpc.ServiceRegistrar, srv SearchMessageServer) {
	// If the following call pancis, it indicates UnimplementedSearchMessageServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SearchMessage_ServiceDesc, srv)
}

func _SearchMessage_SearchMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchMessageServer).SearchMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SearchMessage_SearchMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchMessageServer).SearchMessage(ctx, req.(*SearchMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SearchMessage_ServiceDesc is the grpc.ServiceDesc for SearchMessage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SearchMessage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "search.SearchMessage",
	HandlerType: (*SearchMessageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchMessage",
			Handler:    _SearchMessage_SearchMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/search/search_messages.proto",
}
