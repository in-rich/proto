// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.25.4
// source: proto/reminders/get_reminder_by_id.proto

package reminders_pb

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
	GetReminderByID_GetReminderByID_FullMethodName = "/reminders.GetReminderByID/GetReminderByID"
)

// GetReminderByIDClient is the client API for GetReminderByID service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GetReminderByIDClient interface {
	// Retrieves a reminder.
	GetReminderByID(ctx context.Context, in *GetReminderByIDRequest, opts ...grpc.CallOption) (*Reminder, error)
}

type getReminderByIDClient struct {
	cc grpc.ClientConnInterface
}

func NewGetReminderByIDClient(cc grpc.ClientConnInterface) GetReminderByIDClient {
	return &getReminderByIDClient{cc}
}

func (c *getReminderByIDClient) GetReminderByID(ctx context.Context, in *GetReminderByIDRequest, opts ...grpc.CallOption) (*Reminder, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Reminder)
	err := c.cc.Invoke(ctx, GetReminderByID_GetReminderByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetReminderByIDServer is the server API for GetReminderByID service.
// All implementations must embed UnimplementedGetReminderByIDServer
// for forward compatibility.
type GetReminderByIDServer interface {
	// Retrieves a reminder.
	GetReminderByID(context.Context, *GetReminderByIDRequest) (*Reminder, error)
	mustEmbedUnimplementedGetReminderByIDServer()
}

// UnimplementedGetReminderByIDServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGetReminderByIDServer struct{}

func (UnimplementedGetReminderByIDServer) GetReminderByID(context.Context, *GetReminderByIDRequest) (*Reminder, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReminderByID not implemented")
}
func (UnimplementedGetReminderByIDServer) mustEmbedUnimplementedGetReminderByIDServer() {}
func (UnimplementedGetReminderByIDServer) testEmbeddedByValue()                         {}

// UnsafeGetReminderByIDServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GetReminderByIDServer will
// result in compilation errors.
type UnsafeGetReminderByIDServer interface {
	mustEmbedUnimplementedGetReminderByIDServer()
}

func RegisterGetReminderByIDServer(s grpc.ServiceRegistrar, srv GetReminderByIDServer) {
	// If the following call pancis, it indicates UnimplementedGetReminderByIDServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&GetReminderByID_ServiceDesc, srv)
}

func _GetReminderByID_GetReminderByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReminderByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetReminderByIDServer).GetReminderByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GetReminderByID_GetReminderByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetReminderByIDServer).GetReminderByID(ctx, req.(*GetReminderByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GetReminderByID_ServiceDesc is the grpc.ServiceDesc for GetReminderByID service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GetReminderByID_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "reminders.GetReminderByID",
	HandlerType: (*GetReminderByIDServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetReminderByID",
			Handler:    _GetReminderByID_GetReminderByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/reminders/get_reminder_by_id.proto",
}