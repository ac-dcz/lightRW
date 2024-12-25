package logic

import (
	"context"
	"github.com/ac-dcz/lightRW/apps/goods/rpc/goodsrpc"
	"google.golang.org/grpc/metadata"

	"github.com/ac-dcz/lightRW/apps/goods/api/internal/svc"
	"github.com/ac-dcz/lightRW/apps/goods/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegistryGoodsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewRegistryGoodsLogic 注册商品
func NewRegistryGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegistryGoodsLogic {
	return &RegistryGoodsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegistryGoodsLogic) RegistryGoods(req *types.RegistryGoodsReq) (resp *types.RegistryGoodsResp, err error) {
	//添加metadata
	l.ctx = metadata.AppendToOutgoingContext(l.ctx, "token", req.Token)
	if resp, err := l.svcCtx.GoodRpc.RegistryGoods(l.ctx, &goodsrpc.RegistryGoodsReq{
		Sku:  req.Sku,
		Name: req.Name,
	}); err != nil {
		return nil, err
	} else {
		return &types.RegistryGoodsResp{
			Goods: types.Goods{
				GoodsId: resp.Goods.GoodsId,
				Name:    resp.Goods.Name,
				Sku:     resp.Goods.Sku,
			},
		}, nil
	}
}
