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

type CreateOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderLogic {
	return &CreateOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateOrderLogic) CreateOrder(req *types.CreateOrderReq) (resp *types.CreateOrderResp, err error) {
	//TODO: handle error
	t, _ := l.ctx.Value("uid").(string)
	uid, _ := strconv.ParseUint(t, 10, 64)

	entries := make([]*order.OrderEntry, 0)
	for _, entry := range req.Entries {
		entries = append(entries, &order.OrderEntry{
			StoreId: entry.StoreId,
			Sku:     entry.Sku,
			Nums:    entry.Nums,
		})
	}

	l.ctx = metadata.AppendToOutgoingContext(l.ctx, "token", req.Token)
	if r, err := l.svcCtx.OrderRpc.CreateOrder(l.ctx, &order.CreateOrderReq{
		Uid:     uid,
		Entries: entries,
	}); err != nil {
		l.Error(err)
		return nil, err
	} else {
		resp = &types.CreateOrderResp{
			OrderId: r.OrderId,
			Status:  r.Status,
		}
	}
	return
}
