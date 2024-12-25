package logic

import (
	"context"
	"github.com/ac-dcz/lightRW/apps/user/rpc/userrpc"
	"github.com/ac-dcz/lightRW/common/codes"
	"github.com/ac-dcz/lightRW/common/errors"
	"github.com/ac-dcz/lightRW/common/jwt"
	"strconv"

	"github.com/ac-dcz/lightRW/apps/enroll/api/internal/svc"
	"github.com/ac-dcz/lightRW/apps/enroll/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SignInLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSignInLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignInLogic {
	return &SignInLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SignInLogic) SignIn(req *types.SignInReq) (resp *types.SignInResp, err error) {
	if resp, err := l.svcCtx.UserRpc.SignIn(l.ctx, &userrpc.SignInReq{
		Tel:  req.Tel,
		Pass: req.Pass,
		Code: req.Code,
	}); err != nil {
		l.Logger.Errorf("signIn err:%v", err)
		return nil, err
	} else {

		//build token
		payload := map[string]any{
			"uid":   strconv.FormatUint(resp.User.Uid, 10),
			"level": strconv.FormatUint(uint64(resp.User.Level), 10),
		}
		token, err := jwt.BuildToken(&jwt.Option{
			AccessSecret: l.svcCtx.Config.Auth.AccessSecret,
			AccessExpire: l.svcCtx.Config.Auth.AccessExpire,
		}, payload)
		if err != nil {
			l.Logger.Errorf("build token: %v", err)
			return nil, errors.New(codes.BuildTokenError, "build token error")
		}

		return &types.SignInResp{
			User: types.User{
				NickName: resp.User.NickName,
				Tel:      resp.User.Tel,
				Level:    resp.User.Level,
				Uid:      resp.User.Uid,
			},
			Token: token,
		}, nil
	}
}
