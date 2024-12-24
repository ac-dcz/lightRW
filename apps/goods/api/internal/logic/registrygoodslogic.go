package logic

import (
	"context"

	"github.com/ac-dcz/lightRW/apps/goods/api/internal/svc"
	"github.com/ac-dcz/lightRW/apps/goods/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegistryGoodsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 注册商品
func NewRegistryGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegistryGoodsLogic {
	return &RegistryGoodsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegistryGoodsLogic) RegistryGoods(req *types.RegistryGoodsReq) (resp *types.RegistryGoodsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
