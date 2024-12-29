package handler

import (
	"net/http"

	"github.com/ac-dcz/lightRW/apps/reply/api/internal/logic"
	"github.com/ac-dcz/lightRW/apps/reply/api/internal/svc"
	"github.com/ac-dcz/lightRW/apps/reply/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ReplyBySSkuHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ReplyBySSkuReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewReplyBySSkuLogic(r.Context(), svcCtx)
		resp, err := l.ReplyBySSku(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
