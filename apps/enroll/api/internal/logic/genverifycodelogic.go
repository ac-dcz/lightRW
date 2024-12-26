package logic

import (
	"context"
	"github.com/ac-dcz/lightRW/apps/code/rpc/verifycode"

	"github.com/ac-dcz/lightRW/apps/enroll/api/internal/svc"
	"github.com/ac-dcz/lightRW/apps/enroll/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenVerifyCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGenVerifyCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenVerifyCodeLogic {
	return &GenVerifyCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GenVerifyCodeLogic) GenVerifyCode(req *types.GenVerifyCodeReq) error {
	if resp, err := l.svcCtx.CodeRpc.GenCode(l.ctx, &verifycode.GenCodeReq{
		Tel: req.Tel,
	}); err != nil {
		l.Logger.Error(err)
		return err
	} else {
		l.Infow("Verify Code", logx.Field("code", resp.Code))
	}
	return nil
}
