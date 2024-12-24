package logic

import (
	"context"

	"github.com/ac-dcz/lightRW/apps/goods/api/internal/svc"
	"github.com/ac-dcz/lightRW/apps/goods/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GoodsInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取商品元信息
func NewGoodsInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GoodsInfoLogic {
	return &GoodsInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GoodsInfoLogic) GoodsInfo(req *types.GoodsInfoReq) (resp *types.GoodsInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
