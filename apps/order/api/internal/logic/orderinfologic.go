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

type OrderInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderInfoLogic {
	return &OrderInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderInfoLogic) OrderInfo(req *types.OrderInfoReq) (resp *types.OrderInfoResp, err error) {
	t, _ := l.ctx.Value("uid").(string)
	uid, _ := strconv.ParseUint(t, 10, 64)

	l.ctx = metadata.AppendToOutgoingContext(l.ctx, "token", req.Token)
	if r, err := l.svcCtx.OrderRpc.OrderInfo(l.ctx, &order.OrderInfoReq{
		OrderId: req.OrderId,
		Uid:     uid,
	}); err != nil {
		l.Error(err)
		return nil, err
	} else {
		entries := make([]types.OrderEntry, len(r.Entries))
		for _, entry := range r.Entries {
			entries = append(entries, types.OrderEntry{
				StoreId: entry.StoreId,
				Sku:     entry.Sku,
				Nums:    entry.Nums,
			})
		}
		resp = &types.OrderInfoResp{
			Id:       r.Id,
			OrderId:  r.OrderId,
			Uid:      r.Uid,
			Entries:  entries,
			Status:   r.Status,
			CreateAt: r.CreateAt,
		}
	}

	return
}
