package logic

import (
	"context"
	"github.com/ac-dcz/lightRW/apps/goods/rpc/goodsrpc"
	"github.com/ac-dcz/lightRW/apps/store/rpc/store"
	"google.golang.org/grpc/metadata"

	"github.com/ac-dcz/lightRW/apps/store/api/internal/svc"
	"github.com/ac-dcz/lightRW/apps/store/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddGoodsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddGoodsLogic {
	return &AddGoodsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddGoodsLogic) AddGoods(req *types.AddGoodsReq) error {

	//Step1. sku是否存在
	if _, err := l.svcCtx.GoodsRpc.GoodsInfo(l.ctx, &goodsrpc.GoodsInfoReq{
		Sku: req.Sku,
	}); err != nil {
		return err
	}

	//Step2.
	l.ctx = metadata.AppendToOutgoingContext(l.ctx, "token", req.Token)
	if _, err := l.svcCtx.StoreRpc.AddGoods(l.ctx, &store.AddGoodsReq{
		Sku:     req.Sku,
		StoreId: req.StoreId,
		Stock:   req.Stock,
	}); err != nil {
		l.Logger.Error(err)
		return err
	}

	return nil
}
