package logic

import (
	"context"
	"github.com/ac-dcz/lightRW/common/codes"
	"github.com/ac-dcz/lightRW/common/errors"

	"github.com/ac-dcz/lightRW/apps/code/rpc/internal/svc"
	"github.com/ac-dcz/lightRW/apps/code/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVerifyCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyCodeLogic {
	return &VerifyCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *VerifyCodeLogic) VerifyCode(in *pb.VerifyCodeReq) (*pb.VerifyCodeResp, error) {
	codeKey := codePrefix(in.Tel)
	if code, err := l.svcCtx.BizRds.GetCtx(l.ctx, codeKey); err != nil {
		l.Logger.Errorf("verify code err: %s", err.Error())
		return nil, errors.New(codes.InternalError, err.Error())
	} else if code != in.Code {
		return &pb.VerifyCodeResp{
			Success: false,
		}, nil
	}

	return &pb.VerifyCodeResp{
		Success: true,
	}, nil
}
