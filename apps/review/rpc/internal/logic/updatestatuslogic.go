package logic

import (
	"context"
	"github.com/ac-dcz/lightRW/common/codes"
	"github.com/ac-dcz/lightRW/common/errors"

	"github.com/ac-dcz/lightRW/apps/review/rpc/internal/svc"
	"github.com/ac-dcz/lightRW/apps/review/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStatusLogic {
	return &UpdateStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateStatus 更新review status
func (l *UpdateStatusLogic) UpdateStatus(in *pb.UpdateStatusReq) (*pb.UpdateStatusResp, error) {

	if err := l.svcCtx.ReviewModel.UpdateStatus(l.ctx, in.ReviewId, in.Status, in.OpReason); err != nil {
		l.Logger.Errorf("update status failed: %v", err)
		return nil, errors.New(codes.InternalError, err.Error())
	}

	return &pb.UpdateStatusResp{
		ReviewId: in.ReviewId,
		Status:   in.Status,
	}, nil
}
