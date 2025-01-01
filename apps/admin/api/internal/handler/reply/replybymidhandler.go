package reply

import (
	"net/http"

	"github.com/ac-dcz/lightRW/apps/admin/api/internal/logic/reply"
	"github.com/ac-dcz/lightRW/apps/admin/api/internal/svc"
	"github.com/ac-dcz/lightRW/apps/admin/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ReplyByMidHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ReplyByMidReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := reply.NewReplyByMidLogic(r.Context(), svcCtx)
		resp, err := l.ReplyByMid(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
