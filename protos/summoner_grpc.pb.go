// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: summoner.proto

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
	SummonerService_GetSummonerByPuuid_FullMethodName = "/summoner.SummonerService/GetSummonerByPuuid"
	SummonerService_GetSummonerSpells_FullMethodName  = "/summoner.SummonerService/GetSummonerSpells"
)

// SummonerServiceClient is the client API for SummonerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SummonerServiceClient interface {
	GetSummonerByPuuid(ctx context.Context, in *GetSummonerByPuuidRequest, opts ...grpc.CallOption) (*GetSummonerByPuuidResponse, error)
	GetSummonerSpells(ctx context.Context, in *GetSummonerSpellsRequest, opts ...grpc.CallOption) (*GetSummonerSpellsResponse, error)
}

type summonerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSummonerServiceClient(cc grpc.ClientConnInterface) SummonerServiceClient {
	return &summonerServiceClient{cc}
}

func (c *summonerServiceClient) GetSummonerByPuuid(ctx context.Context, in *GetSummonerByPuuidRequest, opts ...grpc.CallOption) (*GetSummonerByPuuidResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSummonerByPuuidResponse)
	err := c.cc.Invoke(ctx, SummonerService_GetSummonerByPuuid_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *summonerServiceClient) GetSummonerSpells(ctx context.Context, in *GetSummonerSpellsRequest, opts ...grpc.CallOption) (*GetSummonerSpellsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSummonerSpellsResponse)
	err := c.cc.Invoke(ctx, SummonerService_GetSummonerSpells_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SummonerServiceServer is the server API for SummonerService service.
// All implementations must embed UnimplementedSummonerServiceServer
// for forward compatibility.
type SummonerServiceServer interface {
	GetSummonerByPuuid(context.Context, *GetSummonerByPuuidRequest) (*GetSummonerByPuuidResponse, error)
	GetSummonerSpells(context.Context, *GetSummonerSpellsRequest) (*GetSummonerSpellsResponse, error)
	mustEmbedUnimplementedSummonerServiceServer()
}

// UnimplementedSummonerServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSummonerServiceServer struct{}

func (UnimplementedSummonerServiceServer) GetSummonerByPuuid(context.Context, *GetSummonerByPuuidRequest) (*GetSummonerByPuuidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSummonerByPuuid not implemented")
}
func (UnimplementedSummonerServiceServer) GetSummonerSpells(context.Context, *GetSummonerSpellsRequest) (*GetSummonerSpellsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSummonerSpells not implemented")
}
func (UnimplementedSummonerServiceServer) mustEmbedUnimplementedSummonerServiceServer() {}
func (UnimplementedSummonerServiceServer) testEmbeddedByValue()                         {}

// UnsafeSummonerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SummonerServiceServer will
// result in compilation errors.
type UnsafeSummonerServiceServer interface {
	mustEmbedUnimplementedSummonerServiceServer()
}

func RegisterSummonerServiceServer(s grpc.ServiceRegistrar, srv SummonerServiceServer) {
	// If the following call pancis, it indicates UnimplementedSummonerServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SummonerService_ServiceDesc, srv)
}

func _SummonerService_GetSummonerByPuuid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSummonerByPuuidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SummonerServiceServer).GetSummonerByPuuid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SummonerService_GetSummonerByPuuid_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SummonerServiceServer).GetSummonerByPuuid(ctx, req.(*GetSummonerByPuuidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SummonerService_GetSummonerSpells_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSummonerSpellsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SummonerServiceServer).GetSummonerSpells(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SummonerService_GetSummonerSpells_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SummonerServiceServer).GetSummonerSpells(ctx, req.(*GetSummonerSpellsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SummonerService_ServiceDesc is the grpc.ServiceDesc for SummonerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SummonerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "summoner.SummonerService",
	HandlerType: (*SummonerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSummonerByPuuid",
			Handler:    _SummonerService_GetSummonerByPuuid_Handler,
		},
		{
			MethodName: "GetSummonerSpells",
			Handler:    _SummonerService_GetSummonerSpells_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "summoner.proto",
}