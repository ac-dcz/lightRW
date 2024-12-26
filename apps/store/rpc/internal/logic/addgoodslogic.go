package logic

import (
	"context"
	gmodel "github.com/ac-dcz/lightRW/apps/goods/model"
	"github.com/ac-dcz/lightRW/common/codes"
	"github.com/ac-dcz/lightRW/common/errors"

	"github.com/ac-dcz/lightRW/apps/store/rpc/internal/svc"
	"github.com/ac-dcz/lightRW/apps/store/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddGoodsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddGoodsLogic {
	return &AddGoodsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddGoodsLogic) AddGoods(in *pb.AddGoodsReq) (*pb.AddGoodsResp, error) {
	data := &gmodel.GoodsStore{
		StoreId: in.StoreId,
		Stock:   in.Stock,
		Sku:     in.Sku,
	}
	if _, err := l.svcCtx.GoodsStoreModel.Insert(l.ctx, data); err != nil {
		l.Logger.Errorf("AddGoods err:%v", err)
		return nil, errors.New(codes.InternalError, err.Error())
	} else {
		return &pb.AddGoodsResp{}, nil
	}
}
