package review

import (
	"net/http"

	"github.com/ac-dcz/lightRW/apps/admin/api/internal/logic/review"
	"github.com/ac-dcz/lightRW/apps/admin/api/internal/svc"
	"github.com/ac-dcz/lightRW/apps/admin/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ReviewBySSkuHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ReviewBySSkuReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := review.NewReviewBySSkuLogic(r.Context(), svcCtx)
		resp, err := l.ReviewBySSku(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
