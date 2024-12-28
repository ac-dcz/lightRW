package logic

import (
	"context"
	"github.com/ac-dcz/lightRW/apps/order/rpc/order"
	"google.golang.org/grpc/metadata"
	"strconv"

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

	t, _ := l.ctx.Value("uid").(string)
	uid, _ := strconv.ParseUint(t, 10, 64)

	l.ctx = metadata.AppendToOutgoingContext(l.ctx, "token", req.Token)
	if r, err := l.svcCtx.OrderRpc.PayOrder(l.ctx, &order.PayOrderReq{
		OrderId: req.OrderId,
		Uid:     uid,
	}); err != nil {
		l.Logger.Error(err)
		return nil, err
	} else {
		resp = &types.PayOrderResp{
			Status: r.Status,
		}
	}

	return
}
