// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.23.1
// source: review.proto

package pb

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
	Review_ProposeReview_FullMethodName    = "/pb.Review/ProposeReview"
	Review_ReviewByUid_FullMethodName      = "/pb.Review/ReviewByUid"
	Review_ReviewBySSku_FullMethodName     = "/pb.Review/ReviewBySSku"
	Review_ReviewByReviewId_FullMethodName = "/pb.Review/ReviewByReviewId"
	Review_UpdateStatus_FullMethodName     = "/pb.Review/UpdateStatus"
)

// ReviewClient is the client API for Review service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReviewClient interface {
	// ProposeReview 发表评论
	ProposeReview(ctx context.Context, in *ProposeReviewReq, opts ...grpc.CallOption) (*ProposeReviewResp, error)
	// ReviewByUid 查找某一用户的评论
	ReviewByUid(ctx context.Context, in *ReviewByUidReq, opts ...grpc.CallOption) (*ReviewByUidResp, error)
	// ReviewBySSku( 查找某一款商品的评论
	ReviewBySSku(ctx context.Context, in *ReviewBySSkuReq, opts ...grpc.CallOption) (*ReviewBySSkuResp, error)
	// ReviewByReviewId 查找某一条评论的信息
	ReviewByReviewId(ctx context.Context, in *ReviewByReviewIdReq, opts ...grpc.CallOption) (*ReviewByReviewIdResp, error)
	// UpdateStatus 更新review status
	UpdateStatus(ctx context.Context, in *UpdateStatusReq, opts ...grpc.CallOption) (*UpdateStatusResp, error)
}

type reviewClient struct {
	cc grpc.ClientConnInterface
}

func NewReviewClient(cc grpc.ClientConnInterface) ReviewClient {
	return &reviewClient{cc}
}

func (c *reviewClient) ProposeReview(ctx context.Context, in *ProposeReviewReq, opts ...grpc.CallOption) (*ProposeReviewResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProposeReviewResp)
	err := c.cc.Invoke(ctx, Review_ProposeReview_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewClient) ReviewByUid(ctx context.Context, in *ReviewByUidReq, opts ...grpc.CallOption) (*ReviewByUidResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReviewByUidResp)
	err := c.cc.Invoke(ctx, Review_ReviewByUid_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewClient) ReviewBySSku(ctx context.Context, in *ReviewBySSkuReq, opts ...grpc.CallOption) (*ReviewBySSkuResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReviewBySSkuResp)
	err := c.cc.Invoke(ctx, Review_ReviewBySSku_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewClient) ReviewByReviewId(ctx context.Context, in *ReviewByReviewIdReq, opts ...grpc.CallOption) (*ReviewByReviewIdResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReviewByReviewIdResp)
	err := c.cc.Invoke(ctx, Review_ReviewByReviewId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewClient) UpdateStatus(ctx context.Context, in *UpdateStatusReq, opts ...grpc.CallOption) (*UpdateStatusResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateStatusResp)
	err := c.cc.Invoke(ctx, Review_UpdateStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReviewServer is the server API for Review service.
// All implementations must embed UnimplementedReviewServer
// for forward compatibility.
type ReviewServer interface {
	// ProposeReview 发表评论
	ProposeReview(context.Context, *ProposeReviewReq) (*ProposeReviewResp, error)
	// ReviewByUid 查找某一用户的评论
	ReviewByUid(context.Context, *ReviewByUidReq) (*ReviewByUidResp, error)
	// ReviewBySSku( 查找某一款商品的评论
	ReviewBySSku(context.Context, *ReviewBySSkuReq) (*ReviewBySSkuResp, error)
	// ReviewByReviewId 查找某一条评论的信息
	ReviewByReviewId(context.Context, *ReviewByReviewIdReq) (*ReviewByReviewIdResp, error)
	// UpdateStatus 更新review status
	UpdateStatus(context.Context, *UpdateStatusReq) (*UpdateStatusResp, error)
	mustEmbedUnimplementedReviewServer()
}

// UnimplementedReviewServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedReviewServer struct{}

func (UnimplementedReviewServer) ProposeReview(context.Context, *ProposeReviewReq) (*ProposeReviewResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProposeReview not implemented")
}
func (UnimplementedReviewServer) ReviewByUid(context.Context, *ReviewByUidReq) (*ReviewByUidResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReviewByUid not implemented")
}
func (UnimplementedReviewServer) ReviewBySSku(context.Context, *ReviewBySSkuReq) (*ReviewBySSkuResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReviewBySSku not implemented")
}
func (UnimplementedReviewServer) ReviewByReviewId(context.Context, *ReviewByReviewIdReq) (*ReviewByReviewIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReviewByReviewId not implemented")
}
func (UnimplementedReviewServer) UpdateStatus(context.Context, *UpdateStatusReq) (*UpdateStatusResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStatus not implemented")
}
func (UnimplementedReviewServer) mustEmbedUnimplementedReviewServer() {}
func (UnimplementedReviewServer) testEmbeddedByValue()                {}

// UnsafeReviewServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReviewServer will
// result in compilation errors.
type UnsafeReviewServer interface {
	mustEmbedUnimplementedReviewServer()
}

func RegisterReviewServer(s grpc.ServiceRegistrar, srv ReviewServer) {
	// If the following call pancis, it indicates UnimplementedReviewServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Review_ServiceDesc, srv)
}

func _Review_ProposeReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProposeReviewReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServer).ProposeReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Review_ProposeReview_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServer).ProposeReview(ctx, req.(*ProposeReviewReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Review_ReviewByUid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReviewByUidReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServer).ReviewByUid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Review_ReviewByUid_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServer).ReviewByUid(ctx, req.(*ReviewByUidReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Review_ReviewBySSku_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReviewBySSkuReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServer).ReviewBySSku(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Review_ReviewBySSku_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServer).ReviewBySSku(ctx, req.(*ReviewBySSkuReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Review_ReviewByReviewId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReviewByReviewIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServer).ReviewByReviewId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Review_ReviewByReviewId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServer).ReviewByReviewId(ctx, req.(*ReviewByReviewIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Review_UpdateStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStatusReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServer).UpdateStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Review_UpdateStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServer).UpdateStatus(ctx, req.(*UpdateStatusReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Review_ServiceDesc is the grpc.ServiceDesc for Review service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Review_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Review",
	HandlerType: (*ReviewServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProposeReview",
			Handler:    _Review_ProposeReview_Handler,
		},
		{
			MethodName: "ReviewByUid",
			Handler:    _Review_ReviewByUid_Handler,
		},
		{
			MethodName: "ReviewBySSku",
			Handler:    _Review_ReviewBySSku_Handler,
		},
		{
			MethodName: "ReviewByReviewId",
			Handler:    _Review_ReviewByReviewId_Handler,
		},
		{
			MethodName: "UpdateStatus",
			Handler:    _Review_UpdateStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "review.proto",
}
