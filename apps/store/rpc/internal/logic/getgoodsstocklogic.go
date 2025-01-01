package logic

import (
	"context"
	stderr "errors"
	gmodel "github.com/ac-dcz/lightRW/apps/goods/model"
	"github.com/ac-dcz/lightRW/common/codes"
	"github.com/ac-dcz/lightRW/common/errors"

	"github.com/ac-dcz/lightRW/apps/store/rpc/internal/svc"
	"github.com/ac-dcz/lightRW/apps/store/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGoodsStockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGoodsStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGoodsStockLogic {
	return &GetGoodsStockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGoodsStockLogic) GetGoodsStock(in *pb.GoodsStockReq) (*pb.GoodsStockResp, error) {

	if data, err := l.svcCtx.GoodsStoreModel.FindOneByStoreIdSku(l.ctx, in.StoreId, in.Sku); stderr.Is(err, gmodel.ErrNotFound) {
		l.Logger.Errorf("find goods stock err: %v", err)
		return nil, errors.New(codes.GoodsNotFound, "goods not in store")
	} else if err != nil {
		l.Logger.Errorf("find goods stock err: %v", err)
		return nil, errors.New(codes.InternalError, err.Error())
	} else {
		return &pb.GoodsStockResp{
			Stock: data.Stock,
		}, nil
	}
}
