package svc

import (
	"github.com/ac-dcz/lightRW/apps/genid/rpc/genid"
	gmodel "github.com/ac-dcz/lightRW/apps/goods/model"
	"github.com/ac-dcz/lightRW/apps/order/model"
	"github.com/ac-dcz/lightRW/apps/order/rpc/internal/config"
	"github.com/ac-dcz/lightRW/common/interceptor"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	OrderModel  model.OrdersModel
	GStoreModel gmodel.GoodsStoreModel
	GenIdRpc    genid.GenId
	BizRedis    *redis.Redis
	BizLocker   *redis.RedisLock
}

func NewServiceContext(c config.Config) *ServiceContext {
	bizRedis := redis.MustNewRedis(c.BizRedisConf)
	bizLocker := redis.NewRedisLock(bizRedis, "sku_locker")
	return &ServiceContext{
		Config:      c,
		OrderModel:  model.NewOrdersModel(sqlx.NewMysql(c.OrderModelDSN), c.OrderCacheConf),
		GStoreModel: gmodel.NewGoodsStoreModel(sqlx.NewMysql(c.GStoreModelDSN)),
		GenIdRpc:    genid.NewGenId(zrpc.MustNewClient(c.GenIdRpcConf, zrpc.WithUnaryClientInterceptor(interceptor.ErrorForClient()))),
		BizRedis:    bizRedis,
		BizLocker:   bizLocker,
	}
}
