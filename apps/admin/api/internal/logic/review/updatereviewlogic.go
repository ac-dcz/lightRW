package review

import (
	"context"
	"github.com/ac-dcz/lightRW/apps/review/rpc/review"
	"github.com/ac-dcz/lightRW/common/codes"
	"github.com/ac-dcz/lightRW/common/errors"

	"github.com/ac-dcz/lightRW/apps/admin/api/internal/svc"
	"github.com/ac-dcz/lightRW/apps/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateReviewLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateReviewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateReviewLogic {
	return &UpdateReviewLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateReviewLogic) UpdateReview(req *types.UpdateReviewReq) (resp *types.UpdateReviewResp, err error) {

	if _, err := l.svcCtx.ReviewRpc.UpdateStatus(l.ctx, &review.UpdateStatusReq{
		ReviewId: req.ReviewId,
		Status:   uint32(req.Status),
		OpReason: req.OpReason,
	}); err != nil {
		l.Logger.Errorf("UpdateReview err:%v", err)
		return nil, errors.New(codes.InternalError, err.Error())
	} else {
		resp = &types.UpdateReviewResp{}
	}

	return
}
