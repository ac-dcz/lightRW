// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package handler

import (
	"net/http"

	"github.com/ac-dcz/lightRW/apps/goods/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				// 获取商品元信息
				Method:  http.MethodGet,
				Path:    "/info/:sku",
				Handler: GoodsInfoHandler(serverCtx),
			},
			{
				// 注册商品
				Method:  http.MethodPost,
				Path:    "/registry",
				Handler: RegistryGoodsHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.auth.AccessSecret),
		rest.WithPrefix("/api/v1/goods"),
	)
}
