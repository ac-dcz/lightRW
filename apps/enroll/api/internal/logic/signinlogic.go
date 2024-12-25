package logic

import (
	"context"
	"github.com/ac-dcz/lightRW/apps/user/rpc/userrpc"

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
		return &types.SignInResp{
			User: types.User{
				NickName: resp.User.NickName,
				Tel:      resp.User.Tel,
				Level:    resp.User.Level,
				Uid:      resp.User.Uid,
			},
		}, nil
	}
}
