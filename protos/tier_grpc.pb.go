// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: tier.proto

package jstreams_server

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
	TierService_GetTierIcon_FullMethodName = "/tier.TierService/GetTierIcon"
)

// TierServiceClient is the client API for TierService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TierServiceClient interface {
	GetTierIcon(ctx context.Context, in *GetTierIconRequest, opts ...grpc.CallOption) (*GetTierIconResponse, error)
}

type tierServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTierServiceClient(cc grpc.ClientConnInterface) TierServiceClient {
	return &tierServiceClient{cc}
}

func (c *tierServiceClient) GetTierIcon(ctx context.Context, in *GetTierIconRequest, opts ...grpc.CallOption) (*GetTierIconResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTierIconResponse)
	err := c.cc.Invoke(ctx, TierService_GetTierIcon_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TierServiceServer is the server API for TierService service.
// All implementations must embed UnimplementedTierServiceServer
// for forward compatibility.
type TierServiceServer interface {
	GetTierIcon(context.Context, *GetTierIconRequest) (*GetTierIconResponse, error)
	mustEmbedUnimplementedTierServiceServer()
}

// UnimplementedTierServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTierServiceServer struct{}

func (UnimplementedTierServiceServer) GetTierIcon(context.Context, *GetTierIconRequest) (*GetTierIconResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTierIcon not implemented")
}
func (UnimplementedTierServiceServer) mustEmbedUnimplementedTierServiceServer() {}
func (UnimplementedTierServiceServer) testEmbeddedByValue()                     {}

// UnsafeTierServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TierServiceServer will
// result in compilation errors.
type UnsafeTierServiceServer interface {
	mustEmbedUnimplementedTierServiceServer()
}

func RegisterTierServiceServer(s grpc.ServiceRegistrar, srv TierServiceServer) {
	// If the following call pancis, it indicates UnimplementedTierServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TierService_ServiceDesc, srv)
}

func _TierService_GetTierIcon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTierIconRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TierServiceServer).GetTierIcon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TierService_GetTierIcon_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TierServiceServer).GetTierIcon(ctx, req.(*GetTierIconRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TierService_ServiceDesc is the grpc.ServiceDesc for TierService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TierService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tier.TierService",
	HandlerType: (*TierServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTierIcon",
			Handler:    _TierService_GetTierIcon_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tier.proto",
}