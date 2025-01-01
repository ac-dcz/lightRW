package logic

import (
	"context"
	stderr "errors"
	"github.com/ac-dcz/lightRW/apps/reply/model"
	"github.com/ac-dcz/lightRW/common/codes"
	"github.com/ac-dcz/lightRW/common/errors"

	"github.com/ac-dcz/lightRW/apps/reply/rpc/internal/svc"
	"github.com/ac-dcz/lightRW/apps/reply/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReplyByReplyIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewReplyByReplyIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReplyByReplyIdLogic {
	return &ReplyByReplyIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ReplyByReplyIdLogic) ReplyByReplyId(in *pb.ReplyByReplyIdReq) (*pb.ReplyByReplyIdResp, error) {

	reply, err := l.svcCtx.ReplyModel.FindOneByReplyId(l.ctx, in.ReplyId)
	if stderr.Is(err, model.ErrNotFound) || reply == nil {
		return nil, errors.New(codes.NotFoundReply, "not found reply")
	} else if err != nil {
		l.Errorf("FindOneByReplyId err: %v", err)
		return nil, errors.New(codes.InternalError, err.Error())
	}
	info := &pb.ReplyInfo{
		ReplyId:  reply.ReplyId,
		Mid:      reply.Mid,
		StoreId:  reply.StoreId,
		Sku:      reply.Sku,
		ReviewId: reply.ReviewId,
		Content:  reply.ReplyContent,
		HasImage: uint32(reply.HasImage),
		ImageCDN: reply.ImageJson,
		Status:   uint32(reply.Status),
		CreateAt: reply.CreateAt.Format("2006-01-02 15:04:05"),
		IsDel:    uint32(reply.IsDel),
	}

	return &pb.ReplyByReplyIdResp{
		Info: info,
	}, nil
}
