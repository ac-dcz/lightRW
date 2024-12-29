package handler

import (
	"github.com/ac-dcz/lightRW/common/codes"
	"github.com/ac-dcz/lightRW/common/errors"
	"net/http"

	"github.com/ac-dcz/lightRW/apps/review/api/internal/logic"
	"github.com/ac-dcz/lightRW/apps/review/api/internal/svc"
	"github.com/ac-dcz/lightRW/apps/review/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ReviewBySSkuHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ReviewBySSkuReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, errors.New(codes.InternalError, err.Error()))
			return
		}

		l := logic.NewReviewBySSkuLogic(r.Context(), svcCtx)
		resp, err := l.ReviewBySSku(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
