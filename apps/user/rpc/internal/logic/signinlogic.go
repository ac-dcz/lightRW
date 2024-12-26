package logic

import (
	"context"
	"github.com/ac-dcz/lightRW/apps/user/model"
	"github.com/ac-dcz/lightRW/common/codes"
	"github.com/ac-dcz/lightRW/common/errors"
	"github.com/ac-dcz/lightRW/common/utils"

	"github.com/ac-dcz/lightRW/apps/user/rpc/internal/svc"
	"github.com/ac-dcz/lightRW/apps/user/rpc/pb"

	stderr "errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type SignInLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSignInLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignInLogic {
	return &SignInLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SignInLogic) SignIn(in *pb.SignInReq) (*pb.SignInResp, error) {
	//Step0: 参数有效性由api层检验
	//Step1: 查找Tel
	user, err := l.svcCtx.UserModel.FindOneByTel(l.ctx, in.Tel)
	if stderr.Is(err, model.ErrNotFound) {
		return nil, errors.New(codes.TelNotRegistry, "tel is not registry")
	} else if err != nil {
		l.Logger.Errorf("UserModel.FindOneByTel err: %v", err)
		return nil, errors.New(codes.InternalError, err.Error())
	}

	//Step2: 判断密码是否相等
	if pwd := utils.MD5(passWordPrefix + in.Pass); pwd != user.Password {
		l.Logger.Errorf("Password Error")
		return nil, errors.New(codes.PassWordError, "password error")
	}

	return &pb.SignInResp{
		User: &pb.User{
			Uid:      user.Id,
			NickName: user.NickName,
			Tel:      user.Tel,
			Level:    uint32(user.Level[0]),
		},
	}, nil
}
