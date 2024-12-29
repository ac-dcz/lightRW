package logic

import (
	"context"
	"github.com/ac-dcz/lightRW/apps/review/rpc/review"

	"github.com/ac-dcz/lightRW/apps/review/api/internal/svc"
	"github.com/ac-dcz/lightRW/apps/review/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReviewByReviewIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReviewByReviewIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReviewByReviewIdLogic {
	return &ReviewByReviewIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReviewByReviewIdLogic) ReviewByReviewId(req *types.ReviewByReviewIdReq) (resp *types.ReviewByReviewIdResp, err error) {
	if r, err := l.svcCtx.ReviewRpc.ReviewByReviewId(l.ctx, &review.ReviewByReviewIdReq{
		ReviewId: req.ReviewId,
	}); err != nil {
		l.Errorf("ReviewByReviewId - error: %v", err)
		return nil, err
	} else {
		resp = &types.ReviewByReviewIdResp{
			Info: types.ReviewInfo{
				Uid:        r.Info.Uid,
				OrderId:    r.Info.OrderId,
				ReviewId:   r.Info.ReviewId,
				StoreId:    r.Info.StoreId,
				Sku:        r.Info.Sku,
				GoodDesc:   r.Info.GoodsDesc,
				HasImage:   r.Info.HasImage,
				ImageCDN:   r.Info.ImageCDN,
				Status:     uint8(r.Info.Status),
				StoreScore: uint8(r.Info.StoreScore),
				CreateAt:   r.Info.CreateAt,
				Level:      uint8(r.Info.Level),
				IsDel:      uint8(r.Info.IsDel),
			},
		}
	}
	return
}
