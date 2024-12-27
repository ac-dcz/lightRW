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
	OrderModelDSN  string
	OrderCacheConf cache.CacheConf
	GenIdRpcConf   zrpc.RpcClientConf
}
