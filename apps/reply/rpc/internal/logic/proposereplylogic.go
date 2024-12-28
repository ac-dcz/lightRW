package logic

import (
	"context"
	"github.com/ac-dcz/lightRW/apps/genid/rpc/genid"
	"github.com/ac-dcz/lightRW/apps/reply/model"
	"github.com/ac-dcz/lightRW/apps/review/rpc/review"
	"github.com/ac-dcz/lightRW/common/codes"
	"github.com/ac-dcz/lightRW/common/errors"

	"github.com/ac-dcz/lightRW/apps/reply/rpc/internal/svc"
	"github.com/ac-dcz/lightRW/apps/reply/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProposeReplyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProposeReplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProposeReplyLogic {
	return &ProposeReplyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProposeReplyLogic) ProposeReply(in *pb.ProposeReplyReq) (*pb.ProposeReplyResp, error) {

	//Step1: 检查对应的ReviewId是否存在
	r, err := l.svcCtx.ReviewRpc.ReviewByReviewId(l.ctx, &review.ReviewByReviewIdReq{
		ReviewId: in.ReviewId,
	})
	if err != nil {
		l.Error(err)
		return nil, err
	}
	//Step2: 判断review是否已经审批通过
	if r.Info.GetStatus() != model.AuditSuc {
		return nil, errors.New(codes.ReviewNotPassAudit, "review not pass audit")
	}
	if r.Info.StoreId != in.StoreId || r.Info.Sku != in.Sku {
		return nil, errors.New(codes.InvalidSSku, "invalid storeId and sku")
	}
	//Step3: 生成评论ID
	replyId, err := l.svcCtx.GenIdRpc.GetId(l.ctx, &genid.GetIdReq{})
	if err != nil {
		l.Error(err)
		return nil, err
	}
	//Step4: 存入数据库
	reply := &model.Reply{
		ReplyId:      replyId.Id,
		ReviewId:     in.ReviewId,
		StoreId:      in.StoreId,
		Sku:          in.Sku,
		Mid:          in.Mid,
		ReplyContent: in.Content,
		HasImage:     uint64(in.HasImage),
		ImageJson:    in.ImageCDN,
		Status:       model.AuditSuc, //商家回复 先显示后审核
	}

	if _, err := l.svcCtx.ReplyModel.Insert(l.ctx, reply); err != nil {
		l.Errorf("insert reply error: %v", err)
		return nil, errors.New(codes.InternalError, err.Error())
	}

	return &pb.ProposeReplyResp{
		ReplyId: reply.Id,
		Status:  model.AuditSuc,
	}, nil
}
