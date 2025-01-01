package logic

import (
	"context"
	"github.com/ac-dcz/lightRW/apps/store/rpc/store"

	"github.com/ac-dcz/lightRW/apps/store/api/internal/svc"
	"github.com/ac-dcz/lightRW/apps/store/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStoreInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetStoreInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStoreInfoLogic {
	return &GetStoreInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStoreInfoLogic) GetStoreInfo(req *types.StoreInfoReq) (resp *types.StoreInfoResp, err error) {

	if info, err := l.svcCtx.StoreRpc.GetStoreInfo(l.ctx, &store.StoreInfoReq{
		StoreId: req.StoreId,
	}); err != nil {
		l.Error(err)
		return nil, err
	} else {
		goodsInfos := make([]types.GoodsInfo, 0)
		for _, goods := range info.Info.GoodsInfos {
			goodsInfos = append(goodsInfos, types.GoodsInfo{
				Sku:   goods.Sku,
				Stock: goods.Stock,
			})
		}
		resp = &types.StoreInfoResp{
			StoreInfo: types.StoreInfo{
				StoreId:    info.Info.StoreId,
				Name:       info.Info.Name,
				Id:         info.Info.Id,
				CreateDate: info.Info.CreateDate,
				Uid:        info.Info.Uid,
				GoodsInfos: goodsInfos,
			},
		}
	}

	return
}
