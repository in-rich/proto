// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.2
// source: proto/linkedin-data/list_companies.proto

package linkedin_data_pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	ListCompanies_ListCompanies_FullMethodName = "/linkedinData.ListCompanies/ListCompanies"
)

// ListCompaniesClient is the client API for ListCompanies service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ListCompaniesClient interface {
	// Retrieves a list of companies LinkedIn profile data.
	ListCompanies(ctx context.Context, in *ListCompaniesRequest, opts ...grpc.CallOption) (*ListCompaniesResponse, error)
}

type listCompaniesClient struct {
	cc grpc.ClientConnInterface
}

func NewListCompaniesClient(cc grpc.ClientConnInterface) ListCompaniesClient {
	return &listCompaniesClient{cc}
}

func (c *listCompaniesClient) ListCompanies(ctx context.Context, in *ListCompaniesRequest, opts ...grpc.CallOption) (*ListCompaniesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListCompaniesResponse)
	err := c.cc.Invoke(ctx, ListCompanies_ListCompanies_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ListCompaniesServer is the server API for ListCompanies service.
// All implementations must embed UnimplementedListCompaniesServer
// for forward compatibility
type ListCompaniesServer interface {
	// Retrieves a list of companies LinkedIn profile data.
	ListCompanies(context.Context, *ListCompaniesRequest) (*ListCompaniesResponse, error)
	mustEmbedUnimplementedListCompaniesServer()
}

// UnimplementedListCompaniesServer must be embedded to have forward compatible implementations.
type UnimplementedListCompaniesServer struct {
}

func (UnimplementedListCompaniesServer) ListCompanies(context.Context, *ListCompaniesRequest) (*ListCompaniesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCompanies not implemented")
}
func (UnimplementedListCompaniesServer) mustEmbedUnimplementedListCompaniesServer() {}

// UnsafeListCompaniesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ListCompaniesServer will
// result in compilation errors.
type UnsafeListCompaniesServer interface {
	mustEmbedUnimplementedListCompaniesServer()
}

func RegisterListCompaniesServer(s grpc.ServiceRegistrar, srv ListCompaniesServer) {
	s.RegisterService(&ListCompanies_ServiceDesc, srv)
}

func _ListCompanies_ListCompanies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCompaniesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListCompaniesServer).ListCompanies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ListCompanies_ListCompanies_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListCompaniesServer).ListCompanies(ctx, req.(*ListCompaniesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ListCompanies_ServiceDesc is the grpc.ServiceDesc for ListCompanies service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ListCompanies_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "linkedinData.ListCompanies",
	HandlerType: (*ListCompaniesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListCompanies",
			Handler:    _ListCompanies_ListCompanies_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/linkedin-data/list_companies.proto",
}
