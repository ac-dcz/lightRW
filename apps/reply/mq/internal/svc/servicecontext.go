package svc

import (
	"github.com/ac-dcz/lightRW/apps/review/model"
	"github.com/ac-dcz/lightRW/apps/review/mq/internal/config"
	"github.com/ac-dcz/lightRW/common/mq"
	"github.com/elastic/go-elasticsearch/v7"
)

type ServiceContext struct {
	C             config.Config
	Reader        *mq.Reader
	EsReviewModel model.EsReviewModel
}

func NewServiceContext(config config.Config) (*ServiceContext, error) {
	esCli, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: config.EsClientConf.Addresses,
	})
	if err != nil {
		return nil, err
	}
	return &ServiceContext{
		C:             config,
		Reader:        mq.NewReader(&config.ReaderReviewConf),
		EsReviewModel: model.NewEsReviewModel(esCli),
	}, nil
}
