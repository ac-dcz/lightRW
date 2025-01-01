package logic

import (
	"context"
	stderr "errors"
	"github.com/ac-dcz/lightRW/apps/genid/rpc/genid"
	gmodel "github.com/ac-dcz/lightRW/apps/goods/model"
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

	//Step2:
	//1. store_id sku 是否存在
	//2. 库存是否足够?
	//3. 减库存
	//4. 插入订单
	//5. TODO: 绑定超时事件，超时未支付归还库存

	//lock
	for {
		select {
		case <-l.ctx.Done():
			return nil, l.ctx.Err()
		default:
		}
		if ok, err := l.svcCtx.BizLocker.AcquireCtx(l.ctx); err != nil {
			l.Logger.Errorf("AcquireCtx err: %v", err)
			return nil, errors.New(codes.InternalError, err.Error())
		} else if ok {
			break
		}
	}
	defer func() {
		_, _ = l.svcCtx.BizLocker.ReleaseCtx(l.ctx)
	}()

	orders := make([]*model.Orders, 0)
	for _, entry := range in.Entries {
		if gs, err := l.svcCtx.GStoreModel.FindOneByStoreIdSku(l.ctx, entry.StoreId, entry.Sku); stderr.Is(err, gmodel.ErrNotFound) {
			return nil, errors.New(codes.InvalidStoreIdAndSku, "invalid store id and sku")
		} else if gs.Stock < uint64(entry.Nums) {
			return nil, errors.New(codes.StockNotEnough, "stock not enough")
		}
		orders = append(orders, &model.Orders{
			OrderId: OrderId.Id,
			Uid:     in.Uid,
			StoreId: entry.StoreId,
			Sku:     entry.Sku,
			Num:     uint64(entry.Nums),
			Status:  model.UnPay,
		})
	}
	l.ctx = context.WithValue(l.ctx, "goods_store", l.svcCtx.GStoreModel.TableName())
	if err := l.svcCtx.OrderModel.InsertWithTx(l.ctx, orders...); err != nil {
		l.Errorf("OrderModel.InsertWithTx err: %v", err)
		return nil, errors.New(codes.InternalError, err.Error())
	}

	return &pb.CreateOrderResp{
		OrderId: OrderId.Id,
		Status:  int32(model.UnPay),
	}, nil

}
