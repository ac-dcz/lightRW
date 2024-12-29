package handler

import (
	"net/http"

	"github.com/ac-dcz/lightRW/apps/reply/api/internal/logic"
	"github.com/ac-dcz/lightRW/apps/reply/api/internal/svc"
	"github.com/ac-dcz/lightRW/apps/reply/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ReplyByReplyIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ReplyByReplyIdReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewReplyByReplyIdLogic(r.Context(), svcCtx)
		resp, err := l.ReplyByReplyId(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
