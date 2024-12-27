package svc

import (
	"github.com/ac-dcz/lightRW/apps/genid/rpc/genid"
	"github.com/ac-dcz/lightRW/apps/order/model"
	"github.com/ac-dcz/lightRW/apps/order/rpc/internal/config"
	"github.com/ac-dcz/lightRW/common/interceptor"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	OrderModel model.OrdersModel
	GenIdRpc   genid.GenId
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		OrderModel: model.NewOrdersModel(sqlx.NewMysql(c.OrderModelDSN), c.OrderCacheConf),
		GenIdRpc:   genid.NewGenId(zrpc.MustNewClient(c.GenIdRpcConf, zrpc.WithUnaryClientInterceptor(interceptor.ErrorForClient()))),
	}
}
