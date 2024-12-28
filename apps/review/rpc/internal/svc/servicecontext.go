package svc

import (
	"github.com/ac-dcz/lightRW/apps/genid/rpc/genid"
	gmodel "github.com/ac-dcz/lightRW/apps/goods/model"
	"github.com/ac-dcz/lightRW/apps/order/rpc/order"
	"github.com/ac-dcz/lightRW/apps/review/model"
	"github.com/ac-dcz/lightRW/apps/review/rpc/internal/config"
	"github.com/ac-dcz/lightRW/common/interceptor"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	ReviewModel model.ReviewModel
	GenIdRpc    genid.GenId
	GStoreModel gmodel.GoodsStoreModel
	OrderRpc    order.Order
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		ReviewModel: model.NewReviewModel(sqlx.NewMysql(c.ReviewModelDSN), c.CacheConf),
		GenIdRpc:    genid.NewGenId(zrpc.MustNewClient(c.GenIdRpcConf, zrpc.WithUnaryClientInterceptor(interceptor.ErrorForClient()))),
		GStoreModel: gmodel.NewGoodsStoreModel(sqlx.NewMysql(c.GStoreModelDSN)),
		OrderRpc:    order.NewOrder(zrpc.MustNewClient(c.OrderRpcConf)),
	}
}
