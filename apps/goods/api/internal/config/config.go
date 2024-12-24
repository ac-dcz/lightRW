package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	auth struct {
		AccessSecret string
		AccessExpire int64
	}
}
