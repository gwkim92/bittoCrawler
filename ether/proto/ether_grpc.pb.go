// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: ether.proto

package ether

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	BlockService_ImportLatestBlock_FullMethodName = "/ether.BlockService/ImportLatestBlock"
)

// BlockServiceClient is the client API for BlockService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BlockServiceClient interface {
	ImportLatestBlock(ctx context.Context, in *ImportBlockRequest, opts ...grpc.CallOption) (*ImportBlockResponse, error)
}

type blockServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBlockServiceClient(cc grpc.ClientConnInterface) BlockServiceClient {
	return &blockServiceClient{cc}
}

func (c *blockServiceClient) ImportLatestBlock(ctx context.Context, in *ImportBlockRequest, opts ...grpc.CallOption) (*ImportBlockResponse, error) {
	out := new(ImportBlockResponse)
	err := c.cc.Invoke(ctx, BlockService_ImportLatestBlock_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlockServiceServer is the server API for BlockService service.
// All implementations must embed UnimplementedBlockServiceServer
// for forward compatibility
type BlockServiceServer interface {
	ImportLatestBlock(context.Context, *ImportBlockRequest) (*ImportBlockResponse, error)
	mustEmbedUnimplementedBlockServiceServer()
}

// UnimplementedBlockServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBlockServiceServer struct {
}

func (UnimplementedBlockServiceServer) ImportLatestBlock(context.Context, *ImportBlockRequest) (*ImportBlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImportLatestBlock not implemented")
}
func (UnimplementedBlockServiceServer) mustEmbedUnimplementedBlockServiceServer() {}

// UnsafeBlockServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BlockServiceServer will
// result in compilation errors.
type UnsafeBlockServiceServer interface {
	mustEmbedUnimplementedBlockServiceServer()
}

func RegisterBlockServiceServer(s grpc.ServiceRegistrar, srv BlockServiceServer) {
	s.RegisterService(&BlockService_ServiceDesc, srv)
}

func _BlockService_ImportLatestBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImportBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockServiceServer).ImportLatestBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlockService_ImportLatestBlock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockServiceServer).ImportLatestBlock(ctx, req.(*ImportBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BlockService_ServiceDesc is the grpc.ServiceDesc for BlockService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BlockService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ether.BlockService",
	HandlerType: (*BlockServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ImportLatestBlock",
			Handler:    _BlockService_ImportLatestBlock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ether.proto",
}
