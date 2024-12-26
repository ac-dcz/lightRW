package main

import (
	"flag"
	"fmt"
	"github.com/ac-dcz/lightRW/common/interceptor"
	"github.com/ac-dcz/lightRW/common/jwt"

	"github.com/ac-dcz/lightRW/apps/goods/rpc/internal/config"
	"github.com/ac-dcz/lightRW/apps/goods/rpc/internal/server"
	"github.com/ac-dcz/lightRW/apps/goods/rpc/internal/svc"
	"github.com/ac-dcz/lightRW/apps/goods/rpc/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/goods.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterGoodsRpcServer(grpcServer, server.NewGoodsRpcServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	s.AddUnaryInterceptors(interceptor.ErrorForServer(), interceptor.AuthForServer(&jwt.Option{
		AccessSecret: c.TokenAuth.AccessSecret,
		AccessExpire: c.TokenAuth.AccessExpire,
	}))

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
