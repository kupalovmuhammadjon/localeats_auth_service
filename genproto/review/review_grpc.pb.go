// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: review.proto

package review

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

// ReviewClient is the client API for Review service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReviewClient interface {
	CreateReview(ctx context.Context, in *ReqCreateReview, opts ...grpc.CallOption) (*ReviewInfo, error)
	GetReviewsByKitchenId(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*Reviews, error)
	DeleteComment(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Void, error)
	ValidateReviewId(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Void, error)
}

type reviewClient struct {
	cc grpc.ClientConnInterface
}

func NewReviewClient(cc grpc.ClientConnInterface) ReviewClient {
	return &reviewClient{cc}
}

func (c *reviewClient) CreateReview(ctx context.Context, in *ReqCreateReview, opts ...grpc.CallOption) (*ReviewInfo, error) {
	out := new(ReviewInfo)
	err := c.cc.Invoke(ctx, "/review.Review/CreateReview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewClient) GetReviewsByKitchenId(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*Reviews, error) {
	out := new(Reviews)
	err := c.cc.Invoke(ctx, "/review.Review/GetReviewsByKitchenId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewClient) DeleteComment(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/review.Review/DeleteComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewClient) ValidateReviewId(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/review.Review/ValidateReviewId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReviewServer is the server API for Review service.
// All implementations must embed UnimplementedReviewServer
// for forward compatibility
type ReviewServer interface {
	CreateReview(context.Context, *ReqCreateReview) (*ReviewInfo, error)
	GetReviewsByKitchenId(context.Context, *Filter) (*Reviews, error)
	DeleteComment(context.Context, *Id) (*Void, error)
	ValidateReviewId(context.Context, *Id) (*Void, error)
	mustEmbedUnimplementedReviewServer()
}

// UnimplementedReviewServer must be embedded to have forward compatible implementations.
type UnimplementedReviewServer struct {
}

func (UnimplementedReviewServer) CreateReview(context.Context, *ReqCreateReview) (*ReviewInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateReview not implemented")
}
func (UnimplementedReviewServer) GetReviewsByKitchenId(context.Context, *Filter) (*Reviews, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReviewsByKitchenId not implemented")
}
func (UnimplementedReviewServer) DeleteComment(context.Context, *Id) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteComment not implemented")
}
func (UnimplementedReviewServer) ValidateReviewId(context.Context, *Id) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateReviewId not implemented")
}
func (UnimplementedReviewServer) mustEmbedUnimplementedReviewServer() {}

// UnsafeReviewServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReviewServer will
// result in compilation errors.
type UnsafeReviewServer interface {
	mustEmbedUnimplementedReviewServer()
}

func RegisterReviewServer(s grpc.ServiceRegistrar, srv ReviewServer) {
	s.RegisterService(&Review_ServiceDesc, srv)
}

func _Review_CreateReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqCreateReview)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServer).CreateReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/review.Review/CreateReview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServer).CreateReview(ctx, req.(*ReqCreateReview))
	}
	return interceptor(ctx, in, info, handler)
}

func _Review_GetReviewsByKitchenId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Filter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServer).GetReviewsByKitchenId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/review.Review/GetReviewsByKitchenId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServer).GetReviewsByKitchenId(ctx, req.(*Filter))
	}
	return interceptor(ctx, in, info, handler)
}

func _Review_DeleteComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServer).DeleteComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/review.Review/DeleteComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServer).DeleteComment(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _Review_ValidateReviewId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServer).ValidateReviewId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/review.Review/ValidateReviewId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServer).ValidateReviewId(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

// Review_ServiceDesc is the grpc.ServiceDesc for Review service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Review_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "review.Review",
	HandlerType: (*ReviewServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateReview",
			Handler:    _Review_CreateReview_Handler,
		},
		{
			MethodName: "GetReviewsByKitchenId",
			Handler:    _Review_GetReviewsByKitchenId_Handler,
		},
		{
			MethodName: "DeleteComment",
			Handler:    _Review_DeleteComment_Handler,
		},
		{
			MethodName: "ValidateReviewId",
			Handler:    _Review_ValidateReviewId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "review.proto",
}
