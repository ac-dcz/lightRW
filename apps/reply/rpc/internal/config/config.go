package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	ReplyModelDSN string
	ReviewRpcConf zrpc.RpcClientConf
	GenIdRpcConf  zrpc.RpcClientConf
	TokenAuth     struct {
		AccessSecret string
		AccessExpire int64
	}
	CacheConf cache.CacheConf
}
