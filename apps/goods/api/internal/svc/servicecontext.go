package svc

import (
	"github.com/ac-dcz/lightRW/apps/goods/api/internal/config"
	"github.com/ac-dcz/lightRW/apps/goods/rpc/goodsrpc"
	"github.com/go-playground/validator/v10"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	GoodRpc   goodsrpc.GoodsRpc
	Validator *validator.Validate
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		GoodRpc:   goodsrpc.NewGoodsRpc(zrpc.MustNewClient(c.GoodsRpcConf)),
		Validator: validator.New(validator.WithRequiredStructEnabled()),
	}
}
