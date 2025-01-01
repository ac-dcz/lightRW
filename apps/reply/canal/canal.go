package main

import (
	"context"
	"flag"
	"github.com/ac-dcz/lightRW/apps/review/canal/internal/config"
	"github.com/ac-dcz/lightRW/apps/review/canal/internal/server"
	"github.com/ac-dcz/lightRW/apps/review/canal/internal/svc"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
)

var path = flag.String("config", "./etc/canal.yaml", "config")

func main() {
	flag.Parse()
	c := &config.Config{}
	if err := conf.Load(*path, c); err != nil {
		panic(err)
	}
	if err := logx.SetUp(c.LogConf); err != nil {
		panic(err)
	}

	svcCtx, err := svc.NewServiceContext(*c)
	if err != nil {
		panic(err)
	}
	s := server.NewServer(svcCtx)
	defer s.Stop()
	logx.Info("Canal start...")
	s.Run(context.Background())
}
