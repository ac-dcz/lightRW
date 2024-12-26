package logic

import (
	"context"
	"github.com/ac-dcz/lightRW/apps/user/rpc/userrpc"

	"github.com/ac-dcz/lightRW/apps/enroll/api/internal/svc"
	"github.com/ac-dcz/lightRW/apps/enroll/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SignUpLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSignUpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignUpLogic {
	return &SignUpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SignUpLogic) SignUp(req *types.SignUpReq) (resp *types.SignUpResp, err error) {
	if resp, err := l.svcCtx.UserRpc.SignUp(l.ctx, &userrpc.SignUpReq{
		Tel:      req.Tel,
		Code:     req.Code,
		NickName: req.NickName,
		Pass:     req.Pass,
		Level:    req.Level,
	}); err != nil {
		l.Logger.Errorf("SignUp rpc call error: %v", err)
		return nil, err
	} else {
		return &types.SignUpResp{
			User: types.User{
				NickName: resp.User.NickName,
				Level:    resp.User.Level,
				Tel:      resp.User.Tel,
				Uid:      resp.User.Uid,
			},
		}, nil
	}
}
