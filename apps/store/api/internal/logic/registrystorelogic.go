package logic

import (
	"context"
	"github.com/ac-dcz/lightRW/apps/genid/rpc/genid"
	"github.com/ac-dcz/lightRW/apps/store/rpc/store"
	"google.golang.org/grpc/metadata"

	"github.com/ac-dcz/lightRW/apps/store/api/internal/svc"
	"github.com/ac-dcz/lightRW/apps/store/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegistryStoreLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegistryStoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegistryStoreLogic {
	return &RegistryStoreLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegistryStoreLogic) RegistryStore(req *types.RegistryStoreReq) (resp *types.RegistryStoreResp, err error) {

	//Step1: genid
	r, err := l.svcCtx.GenIdRpc.GetId(l.ctx, &genid.GetIdReq{})
	if err != nil {
		l.Logger.Error(err)
		return nil, err
	}

	//Step2: registry
	l.ctx = metadata.AppendToOutgoingContext(l.ctx, "token", req.Token)
	if resp, err := l.svcCtx.StoreRpc.RegistryStore(l.ctx, &store.RegistryStoreReq{
		StoreId: r.Id,
		Name:    req.Name,
	}); err != nil {
		l.Logger.Error(err)
		return nil, err
	} else {
		return &types.RegistryStoreResp{
			StoreInfo: types.StoreInfo{
				StoreId:    resp.Info.StoreId,
				Name:       resp.Info.Name,
				Id:         resp.Info.Id,
				CreateDate: resp.Info.CreateDate,
				Uid:        resp.Info.Uid,
			},
		}, nil
	}
}
