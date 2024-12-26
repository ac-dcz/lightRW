package logic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return &pb.GoodsStockResp{}, nil
}
