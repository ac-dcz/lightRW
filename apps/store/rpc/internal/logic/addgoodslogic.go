package logic

import (
	"context"

	"github.com/ac-dcz/lightRW/apps/store/rpc/internal/svc"
	"github.com/ac-dcz/lightRW/apps/store/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddGoodsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddGoodsLogic {
	return &AddGoodsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddGoodsLogic) AddGoods(in *pb.AddGoodsReq) (*pb.AddGoodsResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddGoodsResp{}, nil
}
