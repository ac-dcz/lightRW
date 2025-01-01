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

type PayOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPayOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PayOrderLogic {
	return &PayOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PayOrderLogic) PayOrder(in *pb.PayOrderReq) (*pb.PayOrderResp, error) {

	//Step1: 判断订单是否存在
	order, err := l.svcCtx.OrderModel.FindOneByOrderId(l.ctx, in.OrderId)
	if stderr.Is(err, model.ErrNotFound) || order == nil {
		return nil, errors.New(codes.OrderNotFound, "order not found")
	} else if err != nil {
		l.Errorf("PayOrder err: %v", err)
		return nil, errors.New(codes.InternalError, err.Error())
	} else if order.Status == model.Pay {
		l.Errorf("PayOrder already pay order: %v", order.OrderId)
		return nil, errors.New(codes.OrderPayed, "order has been pay already")
	} else if order.Status == model.Expired {
		return nil, errors.New(codes.OrderExpire, "order expired")
	}

	//Step2: 跟新订单状态
	_, err = l.svcCtx.OrderModel.UpdateOrdersForStatus(l.ctx, in.OrderId, in.Uid, model.Pay)
	if err != nil {
		l.Errorf("PayOrder err: %v", err)
		return nil, errors.New(codes.InternalError, err.Error())
	}

	return &pb.PayOrderResp{
		OrderId: in.OrderId,
		Status:  int32(model.Pay),
	}, nil
}
