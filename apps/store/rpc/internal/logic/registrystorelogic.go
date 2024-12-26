package logic

import (
	"context"

	"github.com/ac-dcz/lightRW/apps/store/rpc/internal/svc"
	"github.com/ac-dcz/lightRW/apps/store/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegistryStoreLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegistryStoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegistryStoreLogic {
	return &RegistryStoreLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegistryStoreLogic) RegistryStore(in *pb.RegistryStoreReq) (*pb.RegistryStoreResp, error) {
	// todo: add your logic here and delete this line

	return &pb.RegistryStoreResp{}, nil
}
