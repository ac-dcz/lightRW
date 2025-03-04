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

type ReplyByMidLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewReplyByMidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReplyByMidLogic {
	return &ReplyByMidLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ReplyByMidLogic) ReplyByMid(in *pb.ReplyByMidReq) (*pb.ReplyByMidResp, error) {
	replyList, err := l.svcCtx.ReplyModel.FindManyByMid(l.ctx, in.Mid)
	if stderr.Is(err, model.ErrNotFound) || replyList == nil || len(replyList) == 0 {
		l.Logger.Errorf("not found reply")
		return nil, errors.New(codes.NotFoundReply, "not found reply")
	} else if err != nil {
		l.Errorf(err.Error())
		return nil, errors.New(codes.InternalError, err.Error())
	}

	infos := make([]*pb.ReplyInfo, len(replyList))
	for i, reply := range replyList {
		infos[i] = &pb.ReplyInfo{
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
	}

	return &pb.ReplyByMidResp{
		Infos: infos,
	}, nil
}
