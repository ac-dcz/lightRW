package svc

import (
	"github.com/ac-dcz/lightRW/apps/order/api/internal/config"
	"github.com/ac-dcz/lightRW/apps/order/api/internal/middleware"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config    config.Config
	RateLimit rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		RateLimit: middleware.NewRateLimitMiddleware().Handle,
	}
}
