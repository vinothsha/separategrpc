// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: createproto/create.proto

package createproto

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

// BlogCreateServiceClient is the client API for BlogCreateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BlogCreateServiceClient interface {
	CreateBlog(ctx context.Context, in *Blog, opts ...grpc.CallOption) (*BlogId, error)
}

type blogCreateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBlogCreateServiceClient(cc grpc.ClientConnInterface) BlogCreateServiceClient {
	return &blogCreateServiceClient{cc}
}

func (c *blogCreateServiceClient) CreateBlog(ctx context.Context, in *Blog, opts ...grpc.CallOption) (*BlogId, error) {
	out := new(BlogId)
	err := c.cc.Invoke(ctx, "/separategrpc.BlogCreateService/CreateBlog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlogCreateServiceServer is the server API for BlogCreateService service.
// All implementations must embed UnimplementedBlogCreateServiceServer
// for forward compatibility
type BlogCreateServiceServer interface {
	CreateBlog(context.Context, *Blog) (*BlogId, error)
	mustEmbedUnimplementedBlogCreateServiceServer()
}

// UnimplementedBlogCreateServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBlogCreateServiceServer struct {
}

func (UnimplementedBlogCreateServiceServer) CreateBlog(context.Context, *Blog) (*BlogId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBlog not implemented")
}
func (UnimplementedBlogCreateServiceServer) mustEmbedUnimplementedBlogCreateServiceServer() {}

// UnsafeBlogCreateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BlogCreateServiceServer will
// result in compilation errors.
type UnsafeBlogCreateServiceServer interface {
	mustEmbedUnimplementedBlogCreateServiceServer()
}

func RegisterBlogCreateServiceServer(s grpc.ServiceRegistrar, srv BlogCreateServiceServer) {
	s.RegisterService(&BlogCreateService_ServiceDesc, srv)
}

func _BlogCreateService_CreateBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Blog)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogCreateServiceServer).CreateBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/separategrpc.BlogCreateService/CreateBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogCreateServiceServer).CreateBlog(ctx, req.(*Blog))
	}
	return interceptor(ctx, in, info, handler)
}

// BlogCreateService_ServiceDesc is the grpc.ServiceDesc for BlogCreateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BlogCreateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "separategrpc.BlogCreateService",
	HandlerType: (*BlogCreateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBlog",
			Handler:    _BlogCreateService_CreateBlog_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "createproto/create.proto",
}
