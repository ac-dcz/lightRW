package logic

import (
	"context"
	"github.com/ac-dcz/lightRW/apps/reply/rpc/reply"

	"github.com/ac-dcz/lightRW/apps/reply/api/internal/svc"
	"github.com/ac-dcz/lightRW/apps/reply/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReplyByReplyIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReplyByReplyIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReplyByReplyIdLogic {
	return &ReplyByReplyIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReplyByReplyIdLogic) ReplyByReplyId(req *types.ReplyByReplyIdReq) (resp *types.ReplyByReplyIdResp, err error) {

	if r, err := l.svcCtx.ReplyRpc.ReplyByReplyId(l.ctx, &reply.ReplyByReplyIdReq{
		ReplyId: req.ReplyId,
	}); err != nil {
		l.Errorf("call ReplyByReplyId err:%v", err)
		return nil, err
	} else {
		info := r.Info
		resp = &types.ReplyByReplyIdResp{
			Info: types.ReplyInfo{
				ReplyId:  info.ReplyId,
				Mid:      info.Mid,
				ReviewId: info.ReviewId,
				StoreId:  info.StoreId,
				Sku:      info.Sku,
				Status:   uint8(info.Status),
				Content:  info.Content,
				HasImage: info.HasImage,
				ImageCDN: info.ImageCDN,
				CreateAt: info.CreateAt,
				IsDel:    uint8(info.IsDel),
			},
		}
	}

	return
}
