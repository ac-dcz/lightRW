package svc

import (
	"github.com/ac-dcz/lightRW/apps/admin/api/internal/config"
	"github.com/ac-dcz/lightRW/apps/admin/model"
	"github.com/ac-dcz/lightRW/apps/reply/rpc/reply"
	"github.com/ac-dcz/lightRW/apps/review/rpc/review"
	"github.com/ac-dcz/lightRW/common/interceptor"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	ReplyRpc      reply.Reply
	ReviewRpc     review.Review
	EsReplyModel  model.ReplyEsModel
	EsReviewModel model.ReviewEsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	replyEsCli, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: c.ReplyEsConf.Addresses,
	})
	if err != nil {
		panic(err)
	}
	reviewEsCli, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: c.ReviewEsConf.Addresses,
	})
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config:        c,
		EsReplyModel:  model.NewReplyEsModel(replyEsCli),
		EsReviewModel: model.NewReviewEsModel(reviewEsCli),
		ReplyRpc:      reply.NewReply(zrpc.MustNewClient(c.ReplyRpcConf, zrpc.WithUnaryClientInterceptor(interceptor.ErrorForClient()))),
		ReviewRpc:     review.NewReview(zrpc.MustNewClient(c.ReviewRpcConf, zrpc.WithUnaryClientInterceptor(interceptor.ErrorForClient()))),
	}
}
