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

type ReviewByReviewIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewReviewByReviewIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReviewByReviewIdLogic {
	return &ReviewByReviewIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ReviewByReviewId 查找某一条评论的信息
func (l *ReviewByReviewIdLogic) ReviewByReviewId(in *pb.ReviewByReviewIdReq) (*pb.ReviewByReviewIdResp, error) {

	info, err := l.svcCtx.ReviewModel.FindOneByReviewId(l.ctx, in.ReviewId)
	if stderr.Is(err, model.ErrNotFound) {
		return nil, errors.New(codes.NotFoundReview, "not found review")
	}
	if err != nil {
		l.Errorf("FindOneByReviewId err: %v", err)
		return nil, errors.New(codes.InternalError, err.Error())
	}

	return &pb.ReviewByReviewIdResp{
		Info: &pb.ReviewInfo{
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
		},
	}, nil
}
