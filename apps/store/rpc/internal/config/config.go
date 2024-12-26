package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	TokenAuth struct {
		AccessSecret string
		AccessExpire int64
	}
	StoreModelDSN      string
	StoreCacheConf     cache.ClusterConf
	GoodsStoreModelDSN string
}
