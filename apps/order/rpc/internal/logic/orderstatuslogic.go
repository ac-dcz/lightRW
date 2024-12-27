package logic

import (
	"context"
	stderr "errors"
	"github.com/ac-dcz/lightRW/apps/order/model"
	"github.com/ac-dcz/lightRW/common/codes"
	"github.com/ac-dcz/lightRW/common/errors"

	"github.com/ac-dcz/lightRW/apps/order/rpc/internal/svc"
	"github.com/ac-dcz/lightRW/apps/order/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderStatusLogic {
	return &OrderStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderStatusLogic) OrderStatus(in *pb.OrderStatusReq) (*pb.OrderStatusResp, error) {

	data, err := l.svcCtx.OrderModel.FindOneByOrderId(l.ctx, in.OrderId)
	if stderr.Is(err, model.ErrNotFound) {
		return nil, errors.New(codes.OrderNotFound, "order not found")
	} else if err != nil {
		l.Errorf("FindOneByOrderId error: %v", err)
		return nil, errors.New(codes.InternalError, err.Error())
	}

	return &pb.OrderStatusResp{
		Status: int32(data.Status),
	}, nil
}
