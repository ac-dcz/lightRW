package logic

import (
	"context"

	"github.com/ac-dcz/lightRW/apps/order/api/internal/svc"
	"github.com/ac-dcz/lightRW/apps/order/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PayOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPayOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PayOrderLogic {
	return &PayOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PayOrderLogic) PayOrder(req *types.PayOrderReq) (resp *types.PayOrderResp, err error) {
	// todo: add your logic here and delete this line

	return
}