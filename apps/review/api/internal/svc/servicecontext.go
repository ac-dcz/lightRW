package svc

import (
	"github.com/ac-dcz/lightRW/apps/order/rpc/order"
	"github.com/ac-dcz/lightRW/apps/review/api/internal/config"
	"github.com/ac-dcz/lightRW/apps/review/api/internal/middleware"
	"github.com/ac-dcz/lightRW/apps/review/rpc/review"
	"github.com/ac-dcz/lightRW/common/interceptor"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"golang.org/x/time/rate"
)

type ServiceContext struct {
	Config    config.Config
	RateLimit rest.Middleware
	ReviewRpc review.Review
	OrderRpc  order.Order
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		RateLimit: middleware.NewRateLimitMiddleware(rate.NewLimiter(rate.Limit(c.RateLimitConf.Rate), c.RateLimitConf.Size)).Handle,
		ReviewRpc: review.NewReview(zrpc.MustNewClient(c.ReviewRpcConf, zrpc.WithUnaryClientInterceptor(interceptor.ErrorForClient()))),
		OrderRpc:  order.NewOrder(zrpc.MustNewClient(c.OrderRpcConf, zrpc.WithUnaryClientInterceptor(interceptor.ErrorForClient()))),
	}
}
