package svc

import (
	"github.com/ac-dcz/lightRW/apps/genid/rpc/genid"
	"github.com/ac-dcz/lightRW/apps/goods/rpc/goodsrpc"
	"github.com/ac-dcz/lightRW/apps/store/api/internal/config"
	"github.com/ac-dcz/lightRW/apps/store/rpc/store"
	"github.com/ac-dcz/lightRW/common/interceptor"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	GenIdRpc genid.GenId
	GoodsRpc goodsrpc.GoodsRpc
	StoreRpc store.Store
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		GenIdRpc: genid.NewGenId(zrpc.MustNewClient(c.GenIdConf, zrpc.WithUnaryClientInterceptor(interceptor.ErrorForClient()))),
		GoodsRpc: goodsrpc.NewGoodsRpc(zrpc.MustNewClient(c.GoodsConf, zrpc.WithUnaryClientInterceptor(interceptor.ErrorForClient()))),
		StoreRpc: store.NewStore(zrpc.MustNewClient(c.StoreConf, zrpc.WithUnaryClientInterceptor(interceptor.ErrorForClient()))),
	}
}
