package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	VerifyCodeConf zrpc.RpcClientConf
	UserModelDSN   string
	UserModelCache cache.CacheConf
}
