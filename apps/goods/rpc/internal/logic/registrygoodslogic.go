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

type RegistryGoodsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegistryGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegistryGoodsLogic {
	return &RegistryGoodsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegistryGoodsLogic) RegistryGoods(in *pb.RegistryGoodsReq) (*pb.RegistryGoodsResp, error) {
	//Step0: 参数有效性由api层检验
	//Step1: 检验sku是由唯一
	if _, err := l.svcCtx.GoodsModel.FindOneBySku(l.ctx, in.Sku); !stderr.Is(err, model.ErrNotFound) {
		if err != nil {
			l.Logger.Errorf("RegistryGoods FindOneBySku err: %v", err)
			return nil, errors.New(codes.InternalError, err.Error())
		}
		return nil, errors.New(codes.SkuAlreadyExists, "sku is already exists")
	}
	//Step2: 更新数据库
	goods := model.Goods{
		Sku:  in.Sku,
		Name: in.Name,
	}
	if r, err := l.svcCtx.GoodsModel.Insert(l.ctx, &goods); err != nil {
		l.Logger.Errorf("RegistryGoods Insert err: %v", err)
		return nil, errors.New(codes.InternalError, err.Error())
	} else {
		id, _ := r.LastInsertId()
		return &pb.RegistryGoodsResp{
			Goods: &pb.Goods{
				GoodsId: uint64(id),
				Name:    in.Name,
				Sku:     in.Sku,
			},
		}, nil
	}
}
