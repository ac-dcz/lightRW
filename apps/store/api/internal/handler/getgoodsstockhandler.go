package handler

import (
	"github.com/ac-dcz/lightRW/common/codes"
	"github.com/ac-dcz/lightRW/common/errors"
	"net/http"

	"github.com/ac-dcz/lightRW/apps/store/api/internal/logic"
	"github.com/ac-dcz/lightRW/apps/store/api/internal/svc"
	"github.com/ac-dcz/lightRW/apps/store/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetGoodsStockHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GoodsStockReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, errors.New(codes.InvalidParams, err.Error()))
			return
		}

		l := logic.NewGetGoodsStockLogic(r.Context(), svcCtx)
		resp, err := l.GetGoodsStock(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
