package svc

import (
	"github.com/ac-dcz/lightRW/apps/code/rpc/verifycode"
	"github.com/ac-dcz/lightRW/apps/user/model"
	"github.com/ac-dcz/lightRW/apps/user/rpc/internal/config"
	"github.com/ac-dcz/lightRW/common/interceptor"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	CodeRpc   verifycode.VerifyCode
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	codeCli := zrpc.MustNewClient(c.VerifyCodeConf, zrpc.WithUnaryClientInterceptor(interceptor.ErrorForClient()))
	userConn := sqlx.NewMysql(c.UserModelDSN)
	return &ServiceContext{
		Config:    c,
		CodeRpc:   verifycode.NewVerifyCode(codeCli),
		UserModel: model.NewUserModel(userConn, c.UserModelCache),
	}
}
