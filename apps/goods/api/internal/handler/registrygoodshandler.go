package handler

import (
	"github.com/ac-dcz/lightRW/common/codes"
	"github.com/ac-dcz/lightRW/common/errors"
	"github.com/ac-dcz/lightRW/common/jwt"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"

	"github.com/ac-dcz/lightRW/apps/goods/api/internal/logic"
	"github.com/ac-dcz/lightRW/apps/goods/api/internal/svc"
	"github.com/ac-dcz/lightRW/apps/goods/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 注册商品
func RegistryGoodsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegistryGoodsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		//解析token
		if token, err := jwt.ParseTokenFromRequest(&jwt.Option{
			AccessSecret: svcCtx.Config.Auth.AccessSecret,
			AccessExpire: svcCtx.Config.Auth.AccessExpire,
		}, r); err != nil {
			logx.Error(err)
			httpx.ErrorCtx(r.Context(), w, errors.New(codes.InvalidToken, err.Error()))
			return
		} else {
			req.Token = token
		}

		if err := svcCtx.Validator.Struct(&req); err != nil {
			httpx.ErrorCtx(r.Context(), w, errors.New(codes.InvalidParams, "invalid parameters"))
			return
		}

		l := logic.NewRegistryGoodsLogic(r.Context(), svcCtx)
		resp, err := l.RegistryGoods(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
