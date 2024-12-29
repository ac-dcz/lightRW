package handler

import (
	"net/http"

	"github.com/ac-dcz/lightRW/apps/review/api/internal/logic"
	"github.com/ac-dcz/lightRW/apps/review/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ReviewByUidHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewReviewByUidLogic(r.Context(), svcCtx)
		resp, err := l.ReviewByUid()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
