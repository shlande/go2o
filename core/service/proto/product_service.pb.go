// Code generated by protoc-gen-go. DO NOT EDIT.
// source: product_service.proto

package proto // import "."

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProductServiceClient is the client API for ProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProductServiceClient interface {
	// 删除产品
	DeleteProduct(ctx context.Context, in *DeleteProductRequest, opts ...grpc.CallOption) (*Result, error)
}

type productServiceClient struct {
	cc *grpc.ClientConn
}

func NewProductServiceClient(cc *grpc.ClientConn) ProductServiceClient {
	return &productServiceClient{cc}
}

func (c *productServiceClient) DeleteProduct(ctx context.Context, in *DeleteProductRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/ProductService/DeleteProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServiceServer is the server API for ProductService service.
type ProductServiceServer interface {
	// 删除产品
	DeleteProduct(context.Context, *DeleteProductRequest) (*Result, error)
}

func RegisterProductServiceServer(s *grpc.Server, srv ProductServiceServer) {
	s.RegisterService(&_ProductService_serviceDesc, srv)
}

func _ProductService_DeleteProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).DeleteProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProductService/DeleteProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).DeleteProduct(ctx, req.(*DeleteProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProductService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ProductService",
	HandlerType: (*ProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeleteProduct",
			Handler:    _ProductService_DeleteProduct_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "product_service.proto",
}

func init() {
	proto.RegisterFile("product_service.proto", fileDescriptor_product_service_36558956a8a226fa)
}

var fileDescriptor_product_service_36558956a8a226fa = []byte{
	// 127 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x28, 0xca, 0x4f,
	0x29, 0x4d, 0x2e, 0x89, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x97, 0xe2, 0x49, 0xcf, 0xc9, 0x4f, 0x4a, 0xcc, 0x81, 0xf2, 0x44, 0x73, 0x53, 0x8b, 0x8b,
	0x13, 0xd3, 0x53, 0xf5, 0xa1, 0x8a, 0x21, 0xc2, 0x46, 0xce, 0x5c, 0x7c, 0x01, 0x10, 0x81, 0x60,
	0x88, 0x66, 0x21, 0x43, 0x2e, 0x5e, 0x97, 0xd4, 0x9c, 0xd4, 0x92, 0x54, 0xa8, 0xb8, 0x90, 0xa8,
	0x1e, 0x0a, 0x3f, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x44, 0x8a, 0x5d, 0x2f, 0x28, 0xb5, 0xb8,
	0x34, 0xa7, 0x44, 0x89, 0xc1, 0x89, 0x33, 0x8a, 0x5d, 0xcf, 0x1a, 0x6c, 0x5e, 0x12, 0x1b, 0x98,
	0x32, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xc5, 0xb0, 0x81, 0x56, 0x94, 0x00, 0x00, 0x00,
}
