package logic

import (
	"context"
	stderr "errors"
	"github.com/ac-dcz/lightRW/apps/order/model"
	"github.com/ac-dcz/lightRW/apps/order/rpc/internal/svc"
	"github.com/ac-dcz/lightRW/apps/order/rpc/pb"
	"github.com/ac-dcz/lightRW/common/codes"
	"github.com/ac-dcz/lightRW/common/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderInfoLogic {
	return &OrderInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderInfoLogic) OrderInfo(in *pb.OrderInfoReq) (*pb.OrderInfoResp, error) {

	datas, err := l.svcCtx.OrderModel.FindOrdersByOrderIdUId(l.ctx, in.OrderId, in.Uid)
	if stderr.Is(err, model.ErrNotFound) {
		return nil, errors.New(codes.OrderNotFound, "order not found")
	} else if err != nil {
		l.Errorf("FindOrdersByOrderIdUId Error: %v", err.Error())
		return nil, errors.New(codes.InternalError, err.Error())
	}

	resp := &pb.OrderInfoResp{}
	entries := make([]*pb.OrderEntry, len(datas))
	for _, data := range datas {
		entries = append(entries, &pb.OrderEntry{
			StoreId: data.StoreId,
			Sku:     data.Sku,
			Nums:    int32(data.Num),
		})
		resp.Id = data.Id
		resp.OrderId = data.OrderId
		resp.Status = int32(data.Status)
		resp.CreateAt = data.CreateAt.Format("2006-01-02 15:04:05")
	}

	return resp, nil
}
