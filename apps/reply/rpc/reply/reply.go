// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: reply.proto

package reply

import (
	"context"

	"github.com/ac-dcz/lightRW/apps/reply/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ProposeReplyReq     = pb.ProposeReplyReq
	ProposeReplyResp    = pb.ProposeReplyResp
	ReplyByMidReq       = pb.ReplyByMidReq
	ReplyByMidResp      = pb.ReplyByMidResp
	ReplyByReplyIdReq   = pb.ReplyByReplyIdReq
	ReplyByReplyIdResp  = pb.ReplyByReplyIdResp
	ReplyByReviewIdReq  = pb.ReplyByReviewIdReq
	ReplyByReviewIdResp = pb.ReplyByReviewIdResp
	ReplyBySSkuReq      = pb.ReplyBySSkuReq
	ReplyBySSkuResp     = pb.ReplyBySSkuResp
	ReplyInfo           = pb.ReplyInfo
	UpdateStatusReq     = pb.UpdateStatusReq
	UpdateStatusResp    = pb.UpdateStatusResp

	Reply interface {
		ProposeReply(ctx context.Context, in *ProposeReplyReq, opts ...grpc.CallOption) (*ProposeReplyResp, error)
		ReplyByReplyId(ctx context.Context, in *ReplyByReplyIdReq, opts ...grpc.CallOption) (*ReplyByReplyIdResp, error)
		ReplyBySSku(ctx context.Context, in *ReplyBySSkuReq, opts ...grpc.CallOption) (*ReplyBySSkuResp, error)
		ReplyByReviewId(ctx context.Context, in *ReplyByReviewIdReq, opts ...grpc.CallOption) (*ReplyByReviewIdResp, error)
		ReplyByMid(ctx context.Context, in *ReplyByMidReq, opts ...grpc.CallOption) (*ReplyByMidResp, error)
		UpdateStatus(ctx context.Context, in *UpdateStatusReq, opts ...grpc.CallOption) (*UpdateStatusResp, error)
	}

	defaultReply struct {
		cli zrpc.Client
	}
)

func NewReply(cli zrpc.Client) Reply {
	return &defaultReply{
		cli: cli,
	}
}

func (m *defaultReply) ProposeReply(ctx context.Context, in *ProposeReplyReq, opts ...grpc.CallOption) (*ProposeReplyResp, error) {
	client := pb.NewReplyClient(m.cli.Conn())
	return client.ProposeReply(ctx, in, opts...)
}

func (m *defaultReply) ReplyByReplyId(ctx context.Context, in *ReplyByReplyIdReq, opts ...grpc.CallOption) (*ReplyByReplyIdResp, error) {
	client := pb.NewReplyClient(m.cli.Conn())
	return client.ReplyByReplyId(ctx, in, opts...)
}

func (m *defaultReply) ReplyBySSku(ctx context.Context, in *ReplyBySSkuReq, opts ...grpc.CallOption) (*ReplyBySSkuResp, error) {
	client := pb.NewReplyClient(m.cli.Conn())
	return client.ReplyBySSku(ctx, in, opts...)
}

func (m *defaultReply) ReplyByReviewId(ctx context.Context, in *ReplyByReviewIdReq, opts ...grpc.CallOption) (*ReplyByReviewIdResp, error) {
	client := pb.NewReplyClient(m.cli.Conn())
	return client.ReplyByReviewId(ctx, in, opts...)
}

func (m *defaultReply) ReplyByMid(ctx context.Context, in *ReplyByMidReq, opts ...grpc.CallOption) (*ReplyByMidResp, error) {
	client := pb.NewReplyClient(m.cli.Conn())
	return client.ReplyByMid(ctx, in, opts...)
}

func (m *defaultReply) UpdateStatus(ctx context.Context, in *UpdateStatusReq, opts ...grpc.CallOption) (*UpdateStatusResp, error) {
	client := pb.NewReplyClient(m.cli.Conn())
	return client.UpdateStatus(ctx, in, opts...)
}
