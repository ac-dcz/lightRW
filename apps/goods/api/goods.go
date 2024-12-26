package main

import (
	"flag"
	"fmt"
	"github.com/ac-dcz/lightRW/common/http"
	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/ac-dcz/lightRW/apps/goods/api/internal/config"
	"github.com/ac-dcz/lightRW/apps/goods/api/internal/handler"
	"github.com/ac-dcz/lightRW/apps/goods/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/goods-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	//将返回格式封装
	httpx.SetErrorHandler(http.ErrorHandler())
	httpx.SetOkHandler(http.OkHandler())
	
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
