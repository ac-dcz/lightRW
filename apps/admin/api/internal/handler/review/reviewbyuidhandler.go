package review

import (
	"net/http"

	"github.com/ac-dcz/lightRW/apps/admin/api/internal/logic/review"
	"github.com/ac-dcz/lightRW/apps/admin/api/internal/svc"
	"github.com/ac-dcz/lightRW/apps/admin/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ReviewByUidHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ReviewByUidReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := review.NewReviewByUidLogic(r.Context(), svcCtx)
		resp, err := l.ReviewByUid(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
