package svc

import (
	"github.com/ac-dcz/lightRW/apps/reply/api/internal/config"
	"github.com/ac-dcz/lightRW/apps/reply/api/internal/middleware"
	"github.com/ac-dcz/lightRW/apps/reply/rpc/reply"
	"github.com/ac-dcz/lightRW/common/interceptor"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"golang.org/x/time/rate"
)

type ServiceContext struct {
	Config    config.Config
	RateLimit rest.Middleware
	ReplyRpc  reply.Reply
}

func NewServiceContext(c config.Config) *ServiceContext {
	limiter := rate.NewLimiter(rate.Limit(c.RateLimitConf.Rate), c.RateLimitConf.Size)
	return &ServiceContext{
		Config:    c,
		RateLimit: middleware.NewRateLimitMiddleware(limiter).Handle,
		ReplyRpc:  reply.NewReply(zrpc.MustNewClient(c.ReplyRpcConf, zrpc.WithUnaryClientInterceptor(interceptor.ErrorForClient()))),
	}
}
