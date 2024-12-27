package logic

import (
	"context"
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

	//Step1: 获取未支付订单
	//datas, err := l.svcCtx.OrderModel.FindOrdersWithStatus(l.ctx, in.OrderId, in.Uid, model.UnPay)
	//if stderr.Is(err, model.ErrNotFound) {
	//	return &pb.PayOrderResp{
	//		OrderId: in.OrderId,
	//	}, nil
	//} else if err != nil {
	//	l.Errorf("Logic.PayOrder err: %v", err)
	//	return nil, errors.New(codes.InternalError, err.Error())
	//}

	r, err := l.svcCtx.OrderModel.UpdateOrdersForStatus(l.ctx, in.OrderId, in.Uid, model.Pay)
	if err != nil {
		l.Errorf("Logic.PayOrder err: %v", err)
		return nil, errors.New(codes.InternalError, err.Error())
	} else if r == 0 {
		//TODO:
	}

	return &pb.PayOrderResp{
		OrderId: in.OrderId,
	}, nil
}
