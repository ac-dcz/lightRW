package config

import (
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	GoodsModelDSN string
	TokenAuth     struct {
		AccessSecret string
		AccessExpire int64
	}
}
