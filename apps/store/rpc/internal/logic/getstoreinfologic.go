package logic

import (
	"context"
	stderr "errors"
	gmodel "github.com/ac-dcz/lightRW/apps/goods/model"
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

	//Step1: find store
	if data, err := l.svcCtx.StoreModel.FindOneByStoreId(l.ctx, in.StoreId); stderr.Is(err, model.ErrNotFound) {
		l.Logger.Errorf("GetStoreInfoLogic.findOneByStoreId err: %v", err)
		return nil, errors.New(codes.StoreNotRegistry, "store nor registry")
	} else if err != nil {
		l.Logger.Errorf("GetStoreInfoLogic.findOneByStoreId err: %v", err)
		return nil, errors.New(codes.InternalError, err.Error())
	} else {
		//Step2: find goods in store
		goods, err := l.svcCtx.GoodsStoreModel.FindManyByStoreId(l.ctx, in.StoreId)
		if err != nil && !stderr.Is(err, gmodel.ErrNotFound) {
			l.Errorf("GetStoreInfoLogic.findOneByStoreId err: %v", err)
			return nil, errors.New(codes.InternalError, err.Error())
		}
		infos := make([]*pb.GoodsInfo, 0)
		for _, item := range goods {
			infos = append(infos, &pb.GoodsInfo{
				Sku:   item.Sku,
				Stock: item.Stock,
			})
		}

		return &pb.StoreInfoResp{
			Info: &pb.StoreInfo{
				Id:         data.Id,
				StoreId:    data.StoreId,
				Name:       data.Name,
				Uid:        data.Uid,
				CreateDate: data.CreateAt.Format(time.DateTime),
				GoodsInfos: infos,
			},
		}, nil
	}
}
