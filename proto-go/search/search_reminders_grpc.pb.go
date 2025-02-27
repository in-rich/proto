// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: proto/search/search_reminders.proto

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
	SearchReminders_SearchReminder_FullMethodName = "/search.SearchReminders/SearchReminder"
)

// SearchRemindersClient is the client API for SearchReminders service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SearchRemindersClient interface {
	SearchReminder(ctx context.Context, in *SearchRemindersRequest, opts ...grpc.CallOption) (*SearchRemindersResponse, error)
}

type searchRemindersClient struct {
	cc grpc.ClientConnInterface
}

func NewSearchRemindersClient(cc grpc.ClientConnInterface) SearchRemindersClient {
	return &searchRemindersClient{cc}
}

func (c *searchRemindersClient) SearchReminder(ctx context.Context, in *SearchRemindersRequest, opts ...grpc.CallOption) (*SearchRemindersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchRemindersResponse)
	err := c.cc.Invoke(ctx, SearchReminders_SearchReminder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchRemindersServer is the server API for SearchReminders service.
// All implementations must embed UnimplementedSearchRemindersServer
// for forward compatibility.
type SearchRemindersServer interface {
	SearchReminder(context.Context, *SearchRemindersRequest) (*SearchRemindersResponse, error)
	mustEmbedUnimplementedSearchRemindersServer()
}

// UnimplementedSearchRemindersServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSearchRemindersServer struct{}

func (UnimplementedSearchRemindersServer) SearchReminder(context.Context, *SearchRemindersRequest) (*SearchRemindersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchReminder not implemented")
}
func (UnimplementedSearchRemindersServer) mustEmbedUnimplementedSearchRemindersServer() {}
func (UnimplementedSearchRemindersServer) testEmbeddedByValue()                         {}

// UnsafeSearchRemindersServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SearchRemindersServer will
// result in compilation errors.
type UnsafeSearchRemindersServer interface {
	mustEmbedUnimplementedSearchRemindersServer()
}

func RegisterSearchRemindersServer(s grpc.ServiceRegistrar, srv SearchRemindersServer) {
	// If the following call pancis, it indicates UnimplementedSearchRemindersServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SearchReminders_ServiceDesc, srv)
}

func _SearchReminders_SearchReminder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRemindersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchRemindersServer).SearchReminder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SearchReminders_SearchReminder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchRemindersServer).SearchReminder(ctx, req.(*SearchRemindersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SearchReminders_ServiceDesc is the grpc.ServiceDesc for SearchReminders service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SearchReminders_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "search.SearchReminders",
	HandlerType: (*SearchRemindersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchReminder",
			Handler:    _SearchReminders_SearchReminder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/search/search_reminders.proto",
}
