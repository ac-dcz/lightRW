package logic

import (
	"context"

	"github.com/ac-dcz/lightRW/apps/store/rpc/internal/svc"
	"github.com/ac-dcz/lightRW/apps/store/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStoreInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStoreInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStoreInfoLogic {
	return &GetStoreInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetStoreInfoLogic) GetStoreInfo(in *pb.StoreInfoReq) (*pb.StoreInfoResp, error) {
	// todo: add your logic here and delete this line

	return &pb.StoreInfoResp{}, nil
}
