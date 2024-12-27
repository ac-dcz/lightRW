package logic

import (
	"context"
	"github.com/ac-dcz/lightRW/apps/genid/rpc/genid"
	"github.com/ac-dcz/lightRW/apps/order/model"
	"github.com/ac-dcz/lightRW/common/codes"
	"github.com/ac-dcz/lightRW/common/errors"

	"github.com/ac-dcz/lightRW/apps/order/rpc/internal/svc"
	"github.com/ac-dcz/lightRW/apps/order/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderLogic {
	return &CreateOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrderLogic) CreateOrder(in *pb.CreateOrderReq) (*pb.CreateOrderResp, error) {
	//Step1: 获取OrderID
	OrderId, err := l.svcCtx.GenIdRpc.GetId(l.ctx, &genid.GetIdReq{})
	if err != nil {
		l.Errorf("genIdRpc.GetId err: %v", err)
		return nil, err
	}

	//Step2: 存入数据库
	orders := make([]*model.Orders, len(in.Entries))
	for _, entry := range in.Entries {
		orders = append(orders, &model.Orders{
			OrderId: OrderId.Id,
			Uid:     in.Uid,
			StoreId: entry.StoreId,
			Sku:     entry.Sku,
			Num:     uint64(entry.Nums),
			Status:  model.UnPay,
		})
	}
	if err := l.svcCtx.OrderModel.InsertWithTx(l.ctx, orders...); err != nil {
		l.Errorf("OrderModel.InsertWithTx err: %v", err)
		return nil, errors.New(codes.InternalError, err.Error())
	}

	return &pb.CreateOrderResp{
		OrderId: OrderId.Id,
		Status:  int32(model.UnPay),
	}, nil
}
