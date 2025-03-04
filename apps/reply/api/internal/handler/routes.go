// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package handler

import (
	"net/http"

	"github.com/ac-dcz/lightRW/apps/reply/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.RateLimit},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/get/mid",
					Handler: ReplyByMidHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/get/replyid",
					Handler: ReplyByReplyIdHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/get/reviewid",
					Handler: ReplyByReviewIdHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/get/ssku",
					Handler: ReplyBySSkuHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/propose",
					Handler: ProposeReplyHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/v1/reply"),
	)
}
