// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: proto/linkedin-data/get_company.proto

package linkedin_data_pb

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
	GetCompany_GetCompany_FullMethodName = "/linkedinData.GetCompany/GetCompany"
)

// GetCompanyClient is the client API for GetCompany service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GetCompanyClient interface {
	// Retrieves a company's LinkedIn profile data.
	GetCompany(ctx context.Context, in *GetCompanyRequest, opts ...grpc.CallOption) (*Company, error)
}

type getCompanyClient struct {
	cc grpc.ClientConnInterface
}

func NewGetCompanyClient(cc grpc.ClientConnInterface) GetCompanyClient {
	return &getCompanyClient{cc}
}

func (c *getCompanyClient) GetCompany(ctx context.Context, in *GetCompanyRequest, opts ...grpc.CallOption) (*Company, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Company)
	err := c.cc.Invoke(ctx, GetCompany_GetCompany_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetCompanyServer is the server API for GetCompany service.
// All implementations must embed UnimplementedGetCompanyServer
// for forward compatibility.
type GetCompanyServer interface {
	// Retrieves a company's LinkedIn profile data.
	GetCompany(context.Context, *GetCompanyRequest) (*Company, error)
	mustEmbedUnimplementedGetCompanyServer()
}

// UnimplementedGetCompanyServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGetCompanyServer struct{}

func (UnimplementedGetCompanyServer) GetCompany(context.Context, *GetCompanyRequest) (*Company, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompany not implemented")
}
func (UnimplementedGetCompanyServer) mustEmbedUnimplementedGetCompanyServer() {}
func (UnimplementedGetCompanyServer) testEmbeddedByValue()                    {}

// UnsafeGetCompanyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GetCompanyServer will
// result in compilation errors.
type UnsafeGetCompanyServer interface {
	mustEmbedUnimplementedGetCompanyServer()
}

func RegisterGetCompanyServer(s grpc.ServiceRegistrar, srv GetCompanyServer) {
	// If the following call pancis, it indicates UnimplementedGetCompanyServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&GetCompany_ServiceDesc, srv)
}

func _GetCompany_GetCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetCompanyServer).GetCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GetCompany_GetCompany_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetCompanyServer).GetCompany(ctx, req.(*GetCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GetCompany_ServiceDesc is the grpc.ServiceDesc for GetCompany service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GetCompany_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "linkedinData.GetCompany",
	HandlerType: (*GetCompanyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCompany",
			Handler:    _GetCompany_GetCompany_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/linkedin-data/get_company.proto",
}
