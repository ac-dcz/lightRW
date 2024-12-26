package handler

import (
	"github.com/ac-dcz/lightRW/common/codes"
	"github.com/ac-dcz/lightRW/common/errors"
	"net/http"

	"github.com/ac-dcz/lightRW/apps/enroll/api/internal/logic"
	"github.com/ac-dcz/lightRW/apps/enroll/api/internal/svc"
	"github.com/ac-dcz/lightRW/apps/enroll/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GenVerifyCodeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GenVerifyCodeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		if err := svcCtx.Validator.Struct(&req); err != nil {
			httpx.ErrorCtx(r.Context(), w, errors.New(codes.InvalidParams, "invalid params"))
			return
		}
		l := logic.NewGenVerifyCodeLogic(r.Context(), svcCtx)
		err := l.GenVerifyCode(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, "")
		}
	}
}