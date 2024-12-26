package svc

import (
	gmodel "github.com/ac-dcz/lightRW/apps/goods/model"
	"github.com/ac-dcz/lightRW/apps/store/model"
	"github.com/ac-dcz/lightRW/apps/store/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config          config.Config
	StoreModel      model.StoreModel
	GoodsStoreModel gmodel.GoodsStoreModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		StoreModel:      model.NewStoreModel(sqlx.NewMysql(c.StoreModelDSN), c.StoreCacheConf),
		GoodsStoreModel: gmodel.NewGoodsStoreModel(sqlx.NewMysql(c.GoodsStoreModelDSN)),
	}
}
