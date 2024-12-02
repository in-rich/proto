// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: proto/subscription/can_update_note.proto

package subscription_pb

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
	CanUpdateNote_CanUpdateNote_FullMethodName = "/subscription.CanUpdateNote/CanUpdateNote"
)

// CanUpdateNoteClient is the client API for CanUpdateNote service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CanUpdateNoteClient interface {
	// Check if a note can be updated. If at least one edit is available, count a new edit and return the
	// remaining number of edits.
	CanUpdateNote(ctx context.Context, in *CanUpdateNoteRequest, opts ...grpc.CallOption) (*CanUpdateNoteResponse, error)
}

type canUpdateNoteClient struct {
	cc grpc.ClientConnInterface
}

func NewCanUpdateNoteClient(cc grpc.ClientConnInterface) CanUpdateNoteClient {
	return &canUpdateNoteClient{cc}
}

func (c *canUpdateNoteClient) CanUpdateNote(ctx context.Context, in *CanUpdateNoteRequest, opts ...grpc.CallOption) (*CanUpdateNoteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CanUpdateNoteResponse)
	err := c.cc.Invoke(ctx, CanUpdateNote_CanUpdateNote_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CanUpdateNoteServer is the server API for CanUpdateNote service.
// All implementations must embed UnimplementedCanUpdateNoteServer
// for forward compatibility.
type CanUpdateNoteServer interface {
	// Check if a note can be updated. If at least one edit is available, count a new edit and return the
	// remaining number of edits.
	CanUpdateNote(context.Context, *CanUpdateNoteRequest) (*CanUpdateNoteResponse, error)
	mustEmbedUnimplementedCanUpdateNoteServer()
}

// UnimplementedCanUpdateNoteServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCanUpdateNoteServer struct{}

func (UnimplementedCanUpdateNoteServer) CanUpdateNote(context.Context, *CanUpdateNoteRequest) (*CanUpdateNoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CanUpdateNote not implemented")
}
func (UnimplementedCanUpdateNoteServer) mustEmbedUnimplementedCanUpdateNoteServer() {}
func (UnimplementedCanUpdateNoteServer) testEmbeddedByValue()                       {}

// UnsafeCanUpdateNoteServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CanUpdateNoteServer will
// result in compilation errors.
type UnsafeCanUpdateNoteServer interface {
	mustEmbedUnimplementedCanUpdateNoteServer()
}

func RegisterCanUpdateNoteServer(s grpc.ServiceRegistrar, srv CanUpdateNoteServer) {
	// If the following call pancis, it indicates UnimplementedCanUpdateNoteServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CanUpdateNote_ServiceDesc, srv)
}

func _CanUpdateNote_CanUpdateNote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CanUpdateNoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CanUpdateNoteServer).CanUpdateNote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CanUpdateNote_CanUpdateNote_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CanUpdateNoteServer).CanUpdateNote(ctx, req.(*CanUpdateNoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CanUpdateNote_ServiceDesc is the grpc.ServiceDesc for CanUpdateNote service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CanUpdateNote_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "subscription.CanUpdateNote",
	HandlerType: (*CanUpdateNoteServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CanUpdateNote",
			Handler:    _CanUpdateNote_CanUpdateNote_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/subscription/can_update_note.proto",
}
