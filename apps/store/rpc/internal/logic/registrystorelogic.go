package logic

import (
	"context"
	"github.com/ac-dcz/lightRW/apps/store/model"
	"github.com/ac-dcz/lightRW/common/codes"
	"github.com/ac-dcz/lightRW/common/errors"
	"strconv"

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

	//Step1: 插入新数据

	//Step1.1 从ctx中获取uid
	t, ok := l.ctx.Value("uid").(string)
	if !ok {
		return nil, errors.New(codes.InternalError, "uid not found in context")
	}
	uid, _ := strconv.ParseUint(t, 10, 64)

	data := &model.Store{
		StoreId: in.StoreId,
		Name:    in.Name,
		Uid:     uid,
	}
	if r, err := l.svcCtx.StoreModel.Insert(l.ctx, data); err != nil {
		l.Logger.Errorf("Store Insert Error: %s", err.Error())
		return nil, errors.New(codes.InternalError, err.Error())
	} else {
		id, _ := r.LastInsertId()
		return &pb.RegistryStoreResp{
			Info: &pb.StoreInfo{
				Id:      uint64(id),
				StoreId: data.StoreId,
				Name:    data.Name,
				Uid:     data.Uid,
			},
		}, nil
	}
}
