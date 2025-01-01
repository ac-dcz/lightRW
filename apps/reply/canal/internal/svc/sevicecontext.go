package svc

import (
	"github.com/ac-dcz/lightRW/apps/reply/canal/internal/config"
	"github.com/ac-dcz/lightRW/common/canal"
	"github.com/ac-dcz/lightRW/common/mq"
	"github.com/zeromicro/go-zero/core/logx"
)

type ServiceContext struct {
	C        config.Config
	CanalCli *canal.Client
	KqWriter *mq.Writer
}

func NewServiceContext(config config.Config) (*ServiceContext, error) {
	cli, err := canal.NewClient(&config.CanalConf)
	logx.Debug(config.CanalConf)
	if err != nil {
		return nil, err
	}
	return &ServiceContext{
		C:        config,
		CanalCli: cli,
		KqWriter: mq.NewWriter(&config.ReviewWriterConf),
	}, nil
}
