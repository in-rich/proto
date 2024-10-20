// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
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
	SearchMessages_SearchMessages_FullMethodName = "/search.SearchMessages/SearchMessages"
)

// SearchMessagesClient is the client API for SearchMessages service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SearchMessagesClient interface {
	SearchMessages(ctx context.Context, in *SearchMessagesRequest, opts ...grpc.CallOption) (*SearchMessagesResponse, error)
}

type searchMessagesClient struct {
	cc grpc.ClientConnInterface
}

func NewSearchMessagesClient(cc grpc.ClientConnInterface) SearchMessagesClient {
	return &searchMessagesClient{cc}
}

func (c *searchMessagesClient) SearchMessages(ctx context.Context, in *SearchMessagesRequest, opts ...grpc.CallOption) (*SearchMessagesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchMessagesResponse)
	err := c.cc.Invoke(ctx, SearchMessages_SearchMessages_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchMessagesServer is the server API for SearchMessages service.
// All implementations must embed UnimplementedSearchMessagesServer
// for forward compatibility.
type SearchMessagesServer interface {
	SearchMessages(context.Context, *SearchMessagesRequest) (*SearchMessagesResponse, error)
	mustEmbedUnimplementedSearchMessagesServer()
}

// UnimplementedSearchMessagesServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSearchMessagesServer struct{}

func (UnimplementedSearchMessagesServer) SearchMessages(context.Context, *SearchMessagesRequest) (*SearchMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchMessages not implemented")
}
func (UnimplementedSearchMessagesServer) mustEmbedUnimplementedSearchMessagesServer() {}
func (UnimplementedSearchMessagesServer) testEmbeddedByValue()                        {}

// UnsafeSearchMessagesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SearchMessagesServer will
// result in compilation errors.
type UnsafeSearchMessagesServer interface {
	mustEmbedUnimplementedSearchMessagesServer()
}

func RegisterSearchMessagesServer(s grpc.ServiceRegistrar, srv SearchMessagesServer) {
	// If the following call pancis, it indicates UnimplementedSearchMessagesServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SearchMessages_ServiceDesc, srv)
}

func _SearchMessages_SearchMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchMessagesServer).SearchMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SearchMessages_SearchMessages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchMessagesServer).SearchMessages(ctx, req.(*SearchMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SearchMessages_ServiceDesc is the grpc.ServiceDesc for SearchMessages service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SearchMessages_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "search.SearchMessages",
	HandlerType: (*SearchMessagesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchMessages",
			Handler:    _SearchMessages_SearchMessages_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/search/search_messages.proto",
}
