package logic

import (
	"context"
	"github.com/ac-dcz/lightRW/apps/reply/rpc/reply"

	"github.com/ac-dcz/lightRW/apps/reply/api/internal/svc"
	"github.com/ac-dcz/lightRW/apps/reply/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReplyByReviewIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReplyByReviewIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReplyByReviewIdLogic {
	return &ReplyByReviewIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReplyByReviewIdLogic) ReplyByReviewId(req *types.ReplyByReviewIdReq) (resp *types.ReplyByReviewIdResp, err error) {
	if r, err := l.svcCtx.ReplyRpc.ReplyByReviewId(l.ctx, &reply.ReplyByReviewIdReq{
		ReviewId: req.ReviewId,
	}); err != nil {
		l.Errorf("call ReplyByMid error: %v", err)
		return nil, err
	} else {
		infos := make([]types.ReplyInfo, len(r.Infos))
		for i, info := range r.Infos {
			infos[i] = types.ReplyInfo{
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
			}
		}
		resp = &types.ReplyByReviewIdResp{
			Infos: infos,
		}
	}

	return
}
