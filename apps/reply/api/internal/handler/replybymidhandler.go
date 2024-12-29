package handler

import (
	"net/http"

	"github.com/ac-dcz/lightRW/apps/reply/api/internal/logic"
	"github.com/ac-dcz/lightRW/apps/reply/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ReplyByMidHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewReplyByMidLogic(r.Context(), svcCtx)
		resp, err := l.ReplyByMid()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
