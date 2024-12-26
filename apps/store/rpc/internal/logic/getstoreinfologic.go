package logic

import (
	"context"
	stderr "errors"
	"github.com/ac-dcz/lightRW/apps/store/model"
	"github.com/ac-dcz/lightRW/common/codes"
	"github.com/ac-dcz/lightRW/common/errors"
	"time"

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

	if data, err := l.svcCtx.StoreModel.FindOneByStoreId(l.ctx, in.StoreId); stderr.Is(err, model.ErrNotFound) {
		l.Logger.Errorf("GetStoreInfoLogic.findOneByStoreId err: %v", err)
		return nil, errors.New(codes.StoreNotRegistry, "store nor registry")
	} else if err != nil {
		l.Logger.Errorf("GetStoreInfoLogic.findOneByStoreId err: %v", err)
		return nil, errors.New(codes.InternalError, err.Error())
	} else {
		return &pb.StoreInfoResp{
			Info: &pb.StoreInfo{
				Id:         data.Id,
				StoreId:    data.StoreId,
				Name:       data.Name,
				Uid:        data.Uid,
				CreateDate: data.CreateAt.Format(time.DateTime),
			},
		}, nil
	}
}
