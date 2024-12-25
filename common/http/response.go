package http

import (
	"context"
	"github.com/ac-dcz/lightRW/common/codes"
	"github.com/ac-dcz/lightRW/common/errors"
	"net/http"
)

type BaseResponse struct {
	Code codes.Code  `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func OkHandler() func(context.Context, any) any {
	return func(ctx context.Context, a any) any {
		resp := &BaseResponse{
			Code: codes.Ok,
			Msg:  "Ok",
			Data: a,
		}
		return resp
	}
}

func ErrorHandler() func(error) (int, any) {
	return func(err error) (int, any) {
		m := errors.FromError(err)
		resp := &BaseResponse{
			Code: m.Code,
			Msg:  m.Msg,
		}
		return http.StatusOK, resp
	}
}
