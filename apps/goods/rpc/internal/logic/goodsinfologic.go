package logic

import (
	"context"
	stderr "errors"
	"github.com/ac-dcz/lightRW/apps/goods/model"
	"github.com/ac-dcz/lightRW/common/codes"
	"github.com/ac-dcz/lightRW/common/errors"

	"github.com/ac-dcz/lightRW/apps/goods/rpc/internal/svc"
	"github.com/ac-dcz/lightRW/apps/goods/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GoodsInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGoodsInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GoodsInfoLogic {
	return &GoodsInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GoodsInfoLogic) GoodsInfo(in *pb.GoodsInfoReq) (*pb.GoodsInfoResp, error) {
	goods, err := l.svcCtx.GoodsModel.FindOneBySku(l.ctx, in.Sku)
	if stderr.Is(err, model.ErrNotFound) {
		return nil, errors.New(codes.SkuNotRegistry, "sku is not registry")
	} else if err != nil {
		l.Logger.Errorf("GoodsInfoLogic FindOneBySku err: %v", err)
		return nil, errors.New(codes.InternalError, err.Error())
	}

	return &pb.GoodsInfoResp{
		Goods: &pb.Goods{
			GoodsId: goods.Id,
			Sku:     goods.Sku,
			Name:    goods.Name,
		},
	}, nil
}
