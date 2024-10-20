// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: proto/search/search_notes.proto

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
	SearchNotes_SearchNotes_FullMethodName = "/search.SearchNotes/SearchNotes"
)

// SearchNotesClient is the client API for SearchNotes service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SearchNotesClient interface {
	SearchNotes(ctx context.Context, in *SearchNotesRequest, opts ...grpc.CallOption) (*SearchNotesResponse, error)
}

type searchNotesClient struct {
	cc grpc.ClientConnInterface
}

func NewSearchNotesClient(cc grpc.ClientConnInterface) SearchNotesClient {
	return &searchNotesClient{cc}
}

func (c *searchNotesClient) SearchNotes(ctx context.Context, in *SearchNotesRequest, opts ...grpc.CallOption) (*SearchNotesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchNotesResponse)
	err := c.cc.Invoke(ctx, SearchNotes_SearchNotes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchNotesServer is the server API for SearchNotes service.
// All implementations must embed UnimplementedSearchNotesServer
// for forward compatibility.
type SearchNotesServer interface {
	SearchNotes(context.Context, *SearchNotesRequest) (*SearchNotesResponse, error)
	mustEmbedUnimplementedSearchNotesServer()
}

// UnimplementedSearchNotesServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSearchNotesServer struct{}

func (UnimplementedSearchNotesServer) SearchNotes(context.Context, *SearchNotesRequest) (*SearchNotesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchNotes not implemented")
}
func (UnimplementedSearchNotesServer) mustEmbedUnimplementedSearchNotesServer() {}
func (UnimplementedSearchNotesServer) testEmbeddedByValue()                     {}

// UnsafeSearchNotesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SearchNotesServer will
// result in compilation errors.
type UnsafeSearchNotesServer interface {
	mustEmbedUnimplementedSearchNotesServer()
}

func RegisterSearchNotesServer(s grpc.ServiceRegistrar, srv SearchNotesServer) {
	// If the following call pancis, it indicates UnimplementedSearchNotesServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SearchNotes_ServiceDesc, srv)
}

func _SearchNotes_SearchNotes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchNotesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchNotesServer).SearchNotes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SearchNotes_SearchNotes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchNotesServer).SearchNotes(ctx, req.(*SearchNotesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SearchNotes_ServiceDesc is the grpc.ServiceDesc for SearchNotes service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SearchNotes_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "search.SearchNotes",
	HandlerType: (*SearchNotesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchNotes",
			Handler:    _SearchNotes_SearchNotes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/search/search_notes.proto",
}
