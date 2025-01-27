// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.25.4
// source: proto/search/upsert_reminder.proto

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
	UpsertReminder_UpsertReminder_FullMethodName = "/search.UpsertReminder/UpsertReminder"
)

// UpsertReminderClient is the client API for UpsertReminder service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UpsertReminderClient interface {
	// Upsert a reminder.
	UpsertReminder(ctx context.Context, in *UpsertReminderRequest, opts ...grpc.CallOption) (*UpsertReminderResponse, error)
}

type upsertReminderClient struct {
	cc grpc.ClientConnInterface
}

func NewUpsertReminderClient(cc grpc.ClientConnInterface) UpsertReminderClient {
	return &upsertReminderClient{cc}
}

func (c *upsertReminderClient) UpsertReminder(ctx context.Context, in *UpsertReminderRequest, opts ...grpc.CallOption) (*UpsertReminderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpsertReminderResponse)
	err := c.cc.Invoke(ctx, UpsertReminder_UpsertReminder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UpsertReminderServer is the server API for UpsertReminder service.
// All implementations must embed UnimplementedUpsertReminderServer
// for forward compatibility.
type UpsertReminderServer interface {
	// Upsert a reminder.
	UpsertReminder(context.Context, *UpsertReminderRequest) (*UpsertReminderResponse, error)
	mustEmbedUnimplementedUpsertReminderServer()
}

// UnimplementedUpsertReminderServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUpsertReminderServer struct{}

func (UnimplementedUpsertReminderServer) UpsertReminder(context.Context, *UpsertReminderRequest) (*UpsertReminderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertReminder not implemented")
}
func (UnimplementedUpsertReminderServer) mustEmbedUnimplementedUpsertReminderServer() {}
func (UnimplementedUpsertReminderServer) testEmbeddedByValue()                        {}

// UnsafeUpsertReminderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UpsertReminderServer will
// result in compilation errors.
type UnsafeUpsertReminderServer interface {
	mustEmbedUnimplementedUpsertReminderServer()
}

func RegisterUpsertReminderServer(s grpc.ServiceRegistrar, srv UpsertReminderServer) {
	// If the following call pancis, it indicates UnimplementedUpsertReminderServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UpsertReminder_ServiceDesc, srv)
}

func _UpsertReminder_UpsertReminder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertReminderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpsertReminderServer).UpsertReminder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UpsertReminder_UpsertReminder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpsertReminderServer).UpsertReminder(ctx, req.(*UpsertReminderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UpsertReminder_ServiceDesc is the grpc.ServiceDesc for UpsertReminder service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UpsertReminder_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "search.UpsertReminder",
	HandlerType: (*UpsertReminderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpsertReminder",
			Handler:    _UpsertReminder_UpsertReminder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/search/upsert_reminder.proto",
}
