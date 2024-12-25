package handler

import (
	"github.com/ac-dcz/lightRW/common/codes"
	"github.com/ac-dcz/lightRW/common/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"

	"github.com/ac-dcz/lightRW/apps/goods/api/internal/logic"
	"github.com/ac-dcz/lightRW/apps/goods/api/internal/svc"
	"github.com/ac-dcz/lightRW/apps/goods/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 获取商品元信息
func GoodsInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GoodsInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		if err := svcCtx.Validate.Struct(&req); err != nil {
			logx.Errorf("invalid parameters")
			httpx.ErrorCtx(r.Context(), w, errors.New(codes.InvalidParams, "invalid parameters"))
			return
		}
		l := logic.NewGoodsInfoLogic(r.Context(), svcCtx)
		resp, err := l.GoodsInfo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
