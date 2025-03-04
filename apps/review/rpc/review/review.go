// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: review.proto

package review

import (
	"context"

	"github.com/ac-dcz/lightRW/apps/review/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ProposeReviewReq     = pb.ProposeReviewReq
	ProposeReviewResp    = pb.ProposeReviewResp
	ReviewByReviewIdReq  = pb.ReviewByReviewIdReq
	ReviewByReviewIdResp = pb.ReviewByReviewIdResp
	ReviewBySSkuReq      = pb.ReviewBySSkuReq
	ReviewBySSkuResp     = pb.ReviewBySSkuResp
	ReviewByUidReq       = pb.ReviewByUidReq
	ReviewByUidResp      = pb.ReviewByUidResp
	ReviewInfo           = pb.ReviewInfo
	UpdateStatusReq      = pb.UpdateStatusReq
	UpdateStatusResp     = pb.UpdateStatusResp

	Review interface {
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

	defaultReview struct {
		cli zrpc.Client
	}
)

func NewReview(cli zrpc.Client) Review {
	return &defaultReview{
		cli: cli,
	}
}

// ProposeReview 发表评论
func (m *defaultReview) ProposeReview(ctx context.Context, in *ProposeReviewReq, opts ...grpc.CallOption) (*ProposeReviewResp, error) {
	client := pb.NewReviewClient(m.cli.Conn())
	return client.ProposeReview(ctx, in, opts...)
}

// ReviewByUid 查找某一用户的评论
func (m *defaultReview) ReviewByUid(ctx context.Context, in *ReviewByUidReq, opts ...grpc.CallOption) (*ReviewByUidResp, error) {
	client := pb.NewReviewClient(m.cli.Conn())
	return client.ReviewByUid(ctx, in, opts...)
}

// ReviewBySSku( 查找某一款商品的评论
func (m *defaultReview) ReviewBySSku(ctx context.Context, in *ReviewBySSkuReq, opts ...grpc.CallOption) (*ReviewBySSkuResp, error) {
	client := pb.NewReviewClient(m.cli.Conn())
	return client.ReviewBySSku(ctx, in, opts...)
}

// ReviewByReviewId 查找某一条评论的信息
func (m *defaultReview) ReviewByReviewId(ctx context.Context, in *ReviewByReviewIdReq, opts ...grpc.CallOption) (*ReviewByReviewIdResp, error) {
	client := pb.NewReviewClient(m.cli.Conn())
	return client.ReviewByReviewId(ctx, in, opts...)
}

// UpdateStatus 更新review status
func (m *defaultReview) UpdateStatus(ctx context.Context, in *UpdateStatusReq, opts ...grpc.CallOption) (*UpdateStatusResp, error) {
	client := pb.NewReviewClient(m.cli.Conn())
	return client.UpdateStatus(ctx, in, opts...)
}
