package logic

import (
	"context"
	stderr "errors"
	"github.com/ac-dcz/lightRW/apps/review/model"
	"github.com/ac-dcz/lightRW/common/codes"
	"github.com/ac-dcz/lightRW/common/errors"

	"github.com/ac-dcz/lightRW/apps/review/rpc/internal/svc"
	"github.com/ac-dcz/lightRW/apps/review/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReviewByUidLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewReviewByUidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReviewByUidLogic {
	return &ReviewByUidLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ReviewByUid 查找某一用户的评论
func (l *ReviewByUidLogic) ReviewByUid(in *pb.ReviewByUidReq) (*pb.ReviewByUidResp, error) {

	datas, err := l.svcCtx.ReviewModel.FindManyByUid(l.ctx, in.Uid)

	if stderr.Is(err, model.ErrNotFound) || datas == nil || len(datas) == 0 {
		return nil, errors.New(codes.NotFoundReview, "not found review")
	} else if err != nil {
		l.Errorf("FindManyBySSku err: %v", err)
		return nil, errors.New(codes.InternalError, err.Error())
	}
	infos := make([]*pb.ReviewInfo, len(datas))
	for i, info := range datas {
		infos[i] = &pb.ReviewInfo{
			ReviewId:   info.ReviewId,
			Uid:        info.Uid,
			OrderId:    info.OrderId,
			Sku:        info.Sku,
			StoreId:    info.StoreId,
			GoodsDesc:  info.GoodsDesc,
			Level:      pb.ScoreLevel(info.Score),
			HasImage:   info.HasImage == 1,
			ImageCDN:   info.ImageJson,
			StoreScore: uint32(info.StoreScore),
			Status:     info.Status,
			IsDel:      uint32(info.IsDel),
			CreateAt:   info.CreateAt.Format("2006-01-02 15:04:05"),
		}
	}

	return &pb.ReviewByUidResp{
		Infos: infos,
	}, nil
}
