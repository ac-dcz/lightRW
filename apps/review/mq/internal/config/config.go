package config

import (
	"github.com/ac-dcz/lightRW/common/mq"
	"github.com/zeromicro/go-zero/core/logx"
)

type Config struct {
	logx.LogConf
	ReaderReviewConf mq.ReaderConf
	EsClientConf     struct {
		Addresses []string
	}
}
