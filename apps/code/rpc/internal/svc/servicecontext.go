package svc

import (
	"github.com/ac-dcz/lightRW/apps/code/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config config.Config
	BizRds *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		BizRds: redis.MustNewRedis(c.BizRdsConf),
	}
}
