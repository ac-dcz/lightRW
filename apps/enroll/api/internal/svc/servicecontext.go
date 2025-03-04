package svc

import (
	"github.com/ac-dcz/lightRW/apps/code/rpc/verifycode"
	"github.com/ac-dcz/lightRW/apps/enroll/api/internal/config"
	"github.com/ac-dcz/lightRW/apps/user/rpc/userrpc"
	"github.com/ac-dcz/lightRW/common/interceptor"
	"github.com/go-playground/validator/v10"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	Validator *validator.Validate
	UserRpc   userrpc.UserRpc
	CodeRpc   verifycode.VerifyCode
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		Validator: validator.New(validator.WithRequiredStructEnabled()),
		UserRpc:   userrpc.NewUserRpc(zrpc.MustNewClient(c.UserRpcConf, zrpc.WithUnaryClientInterceptor(interceptor.ErrorForClient()))),
		CodeRpc:   verifycode.NewVerifyCode(zrpc.MustNewClient(c.VerifyCodeRpcConf, zrpc.WithUnaryClientInterceptor(interceptor.ErrorForClient()))),
	}
}
