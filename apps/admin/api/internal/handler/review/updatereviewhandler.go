package review

import (
	"net/http"

	"github.com/ac-dcz/lightRW/apps/admin/api/internal/logic/review"
	"github.com/ac-dcz/lightRW/apps/admin/api/internal/svc"
	"github.com/ac-dcz/lightRW/apps/admin/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateReviewHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateReviewReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := review.NewUpdateReviewLogic(r.Context(), svcCtx)
		resp, err := l.UpdateReview(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
