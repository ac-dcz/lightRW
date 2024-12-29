package logic

import (
	"context"
	"github.com/ac-dcz/lightRW/apps/review/rpc/review"

	"github.com/ac-dcz/lightRW/apps/review/api/internal/svc"
	"github.com/ac-dcz/lightRW/apps/review/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReviewBySSkuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReviewBySSkuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReviewBySSkuLogic {
	return &ReviewBySSkuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReviewBySSkuLogic) ReviewBySSku(req *types.ReviewBySSkuReq) (resp *types.ReviewBySSkuResp, err error) {

	if r, err := l.svcCtx.ReviewRpc.ReviewBySSku(l.ctx, &review.ReviewBySSkuReq{
		StoreId: req.StoreId,
		Sku:     req.Sku,
	}); err != nil {
		l.Error(err)
		return nil, err
	} else {
		infos := make([]types.ReviewInfo, len(r.Infos))
		for i, info := range r.Infos {
			infos[i] = types.ReviewInfo{
				Uid:        info.Uid,
				OrderId:    info.OrderId,
				ReviewId:   info.ReviewId,
				StoreId:    info.StoreId,
				Sku:        info.Sku,
				GoodDesc:   info.GoodsDesc,
				HasImage:   info.HasImage,
				ImageCDN:   info.ImageCDN,
				Status:     uint8(info.Status),
				StoreScore: uint8(info.StoreScore),
				CreateAt:   info.CreateAt,
				Level:      uint8(info.Level),
				IsDel:      uint8(info.IsDel),
			}
		}
		resp = &types.ReviewBySSkuResp{
			Infos: infos,
		}
	}

	return
}
