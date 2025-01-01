package logic

import (
	"context"
	"github.com/ac-dcz/lightRW/common/codes"
	"github.com/ac-dcz/lightRW/common/errors"

	"github.com/ac-dcz/lightRW/apps/reply/rpc/internal/svc"
	"github.com/ac-dcz/lightRW/apps/reply/rpc/pb"

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

func (l *UpdateStatusLogic) UpdateStatus(in *pb.UpdateStatusReq) (*pb.UpdateStatusResp, error) {

	if err := l.svcCtx.ReplyModel.UpdateStatus(l.ctx, in.ReplyId, in.Status, in.OpReason); err != nil {
		l.Logger.Errorf("update status err:%v", err)
		return nil, errors.New(codes.InternalError, err.Error())
	}

	return &pb.UpdateStatusResp{
		ReplyId: in.ReplyId,
		Status:  in.Status,
	}, nil
}
