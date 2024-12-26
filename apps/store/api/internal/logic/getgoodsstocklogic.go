package logic

import (
	"context"
	"github.com/ac-dcz/lightRW/apps/store/rpc/store"

	"github.com/ac-dcz/lightRW/apps/store/api/internal/svc"
	"github.com/ac-dcz/lightRW/apps/store/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGoodsStockLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetGoodsStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGoodsStockLogic {
	return &GetGoodsStockLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGoodsStockLogic) GetGoodsStock(req *types.GoodsStockReq) (resp *types.GoodsStockResp, err error) {

	if r, err := l.svcCtx.StoreRpc.GetGoodsStock(l.ctx, &store.GoodsStockReq{
		StoreId: req.StoreId,
		Sku:     req.Sku,
	}); err != nil {
		l.Error(err)
		return nil, err
	} else {
		resp = &types.GoodsStockResp{
			Stock: r.Stock,
		}
	}
	return
}
