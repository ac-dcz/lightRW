package config

import (
	"github.com/ac-dcz/lightRW/common/canal"
	"github.com/ac-dcz/lightRW/common/mq"
	"github.com/zeromicro/go-zero/core/logx"
)

type Config struct {
	logx.LogConf
	CanalConf        canal.ClientConf
	ReviewWriterConf mq.WriterConf
}
