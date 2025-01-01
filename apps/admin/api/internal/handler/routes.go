// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package handler

import (
	"net/http"

	reply "github.com/ac-dcz/lightRW/apps/admin/api/internal/handler/reply"
	review "github.com/ac-dcz/lightRW/apps/admin/api/internal/handler/review"
	"github.com/ac-dcz/lightRW/apps/admin/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/key",
				Handler: reply.ReplyByKeyHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/mid",
				Handler: reply.ReplyByMidHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/reviewid",
				Handler: reply.ReplyByReplyIdHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/reviewid",
				Handler: reply.ReplyByReviewIdHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/ssku",
				Handler: reply.ReplyBySSkuHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/update/status",
				Handler: reply.UpdateReplyHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/v1/admin/reply"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/key",
				Handler: review.ReviewByKeyHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/reviewid",
				Handler: review.ReviewByReviewIdHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/ssku",
				Handler: review.ReviewBySSkuHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/uid",
				Handler: review.ReviewByUidHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/update/status",
				Handler: review.UpdateReviewHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/v1/admin/review"),
	)
}
