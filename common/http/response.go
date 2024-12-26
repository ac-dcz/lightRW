package http

import (
	"context"
	"github.com/ac-dcz/lightRW/common/codes"
	"github.com/ac-dcz/lightRW/common/errors"
	"github.com/go-playground/validator/v10"
	"github.com/zeromicro/go-zero/rest/httpx"
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

type myValidator struct {
	Validator *validator.Validate
}

func (v *myValidator) Validate(_ *http.Request, data any) error {
	return v.Validator.Struct(data)
}

func ValidateHandler(opts ...validator.Option) httpx.Validator {
	return &myValidator{Validator: validator.New(opts...)}
}
