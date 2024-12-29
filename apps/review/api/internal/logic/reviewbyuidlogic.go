package logic

import (
	"context"
	"github.com/ac-dcz/lightRW/apps/review/rpc/review"
	"strconv"

	"github.com/ac-dcz/lightRW/apps/review/api/internal/svc"
	"github.com/ac-dcz/lightRW/apps/review/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReviewByUidLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReviewByUidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReviewByUidLogic {
	return &ReviewByUidLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReviewByUidLogic) ReviewByUid() (resp *types.ReviewByUidResp, err error) {

	//TODO: Error Handle
	t, _ := l.ctx.Value("uid").(string)
	uid, _ := strconv.ParseUint(t, 10, 64)

	if r, err := l.svcCtx.ReviewRpc.ReviewByUid(l.ctx, &review.ReviewByUidReq{Uid: uid}); err != nil {
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
		resp = &types.ReviewByUidResp{
			Infos: infos,
		}
	}

	return
}
