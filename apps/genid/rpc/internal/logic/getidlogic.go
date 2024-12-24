package logic

import (
	"context"
	"github.com/ac-dcz/lightRW/common/codes"
	"github.com/ac-dcz/lightRW/common/errors"

	"github.com/ac-dcz/lightRW/apps/genid/rpc/internal/svc"
	"github.com/ac-dcz/lightRW/apps/genid/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetIdLogic {
	return &GetIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetIdLogic) GetId(in *pb.GetIdReq) (*pb.GetIdResp, error) {
	id, err := l.svcCtx.Flake.NextID()
	if err != nil {
		l.Logger.Errorf("flake.nextID failed: %v", err)
		return nil, errors.New(codes.InternalError, err.Error())
	}

	return &pb.GetIdResp{
		Id: id,
	}, nil
}
