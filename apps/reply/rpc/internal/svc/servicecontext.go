package svc

import (
	"github.com/ac-dcz/lightRW/apps/genid/rpc/genid"
	"github.com/ac-dcz/lightRW/apps/reply/model"
	"github.com/ac-dcz/lightRW/apps/reply/rpc/internal/config"
	"github.com/ac-dcz/lightRW/apps/review/rpc/review"
	"github.com/ac-dcz/lightRW/common/interceptor"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	ReplyModel model.ReplyModel
	ReviewRpc  review.Review
	GenIdRpc   genid.GenId
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		ReplyModel: model.NewReplyModel(sqlx.NewMysql(c.ReplyModelDSN), c.CacheConf),
		ReviewRpc:  review.NewReview(zrpc.MustNewClient(c.ReviewRpcConf, zrpc.WithUnaryClientInterceptor(interceptor.ErrorForClient()))),
		GenIdRpc:   genid.NewGenId(zrpc.MustNewClient(c.GenIdRpcConf, zrpc.WithUnaryClientInterceptor(interceptor.ErrorForClient()))),
	}
}
