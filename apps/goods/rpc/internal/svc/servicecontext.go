package svc

import (
	"github.com/ac-dcz/lightRW/apps/goods/model"
	"github.com/ac-dcz/lightRW/apps/goods/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config     config.Config
	GoodsModel model.GoodsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		GoodsModel: model.NewGoodsModel(sqlx.NewMysql(c.GoodsModelDSN)),
	}
}
