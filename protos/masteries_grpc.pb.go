// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: masteries.proto

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
	MasteriesService_GetChampionMasteriesByPuuid_FullMethodName           = "/masteries.MasteriesService/GetChampionMasteriesByPuuid"
	MasteriesService_GetChampionMasteriesByPuuidByChampion_FullMethodName = "/masteries.MasteriesService/GetChampionMasteriesByPuuidByChampion"
)

// MasteriesServiceClient is the client API for MasteriesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MasteriesServiceClient interface {
	GetChampionMasteriesByPuuid(ctx context.Context, in *GetChampionMasteriesRequeset, opts ...grpc.CallOption) (*GetChampionMasteriesResponse, error)
	GetChampionMasteriesByPuuidByChampion(ctx context.Context, in *GetChampionMasteriesByChampionRequeset, opts ...grpc.CallOption) (*GetChampionMasteriesByChampionResponse, error)
}

type masteriesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMasteriesServiceClient(cc grpc.ClientConnInterface) MasteriesServiceClient {
	return &masteriesServiceClient{cc}
}

func (c *masteriesServiceClient) GetChampionMasteriesByPuuid(ctx context.Context, in *GetChampionMasteriesRequeset, opts ...grpc.CallOption) (*GetChampionMasteriesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetChampionMasteriesResponse)
	err := c.cc.Invoke(ctx, MasteriesService_GetChampionMasteriesByPuuid_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masteriesServiceClient) GetChampionMasteriesByPuuidByChampion(ctx context.Context, in *GetChampionMasteriesByChampionRequeset, opts ...grpc.CallOption) (*GetChampionMasteriesByChampionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetChampionMasteriesByChampionResponse)
	err := c.cc.Invoke(ctx, MasteriesService_GetChampionMasteriesByPuuidByChampion_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MasteriesServiceServer is the server API for MasteriesService service.
// All implementations must embed UnimplementedMasteriesServiceServer
// for forward compatibility.
type MasteriesServiceServer interface {
	GetChampionMasteriesByPuuid(context.Context, *GetChampionMasteriesRequeset) (*GetChampionMasteriesResponse, error)
	GetChampionMasteriesByPuuidByChampion(context.Context, *GetChampionMasteriesByChampionRequeset) (*GetChampionMasteriesByChampionResponse, error)
	mustEmbedUnimplementedMasteriesServiceServer()
}

// UnimplementedMasteriesServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMasteriesServiceServer struct{}

func (UnimplementedMasteriesServiceServer) GetChampionMasteriesByPuuid(context.Context, *GetChampionMasteriesRequeset) (*GetChampionMasteriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChampionMasteriesByPuuid not implemented")
}
func (UnimplementedMasteriesServiceServer) GetChampionMasteriesByPuuidByChampion(context.Context, *GetChampionMasteriesByChampionRequeset) (*GetChampionMasteriesByChampionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChampionMasteriesByPuuidByChampion not implemented")
}
func (UnimplementedMasteriesServiceServer) mustEmbedUnimplementedMasteriesServiceServer() {}
func (UnimplementedMasteriesServiceServer) testEmbeddedByValue()                          {}

// UnsafeMasteriesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MasteriesServiceServer will
// result in compilation errors.
type UnsafeMasteriesServiceServer interface {
	mustEmbedUnimplementedMasteriesServiceServer()
}

func RegisterMasteriesServiceServer(s grpc.ServiceRegistrar, srv MasteriesServiceServer) {
	// If the following call pancis, it indicates UnimplementedMasteriesServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MasteriesService_ServiceDesc, srv)
}

func _MasteriesService_GetChampionMasteriesByPuuid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChampionMasteriesRequeset)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasteriesServiceServer).GetChampionMasteriesByPuuid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MasteriesService_GetChampionMasteriesByPuuid_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasteriesServiceServer).GetChampionMasteriesByPuuid(ctx, req.(*GetChampionMasteriesRequeset))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasteriesService_GetChampionMasteriesByPuuidByChampion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChampionMasteriesByChampionRequeset)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasteriesServiceServer).GetChampionMasteriesByPuuidByChampion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MasteriesService_GetChampionMasteriesByPuuidByChampion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasteriesServiceServer).GetChampionMasteriesByPuuidByChampion(ctx, req.(*GetChampionMasteriesByChampionRequeset))
	}
	return interceptor(ctx, in, info, handler)
}

// MasteriesService_ServiceDesc is the grpc.ServiceDesc for MasteriesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MasteriesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "masteries.MasteriesService",
	HandlerType: (*MasteriesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetChampionMasteriesByPuuid",
			Handler:    _MasteriesService_GetChampionMasteriesByPuuid_Handler,
		},
		{
			MethodName: "GetChampionMasteriesByPuuidByChampion",
			Handler:    _MasteriesService_GetChampionMasteriesByPuuidByChampion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "masteries.proto",
}