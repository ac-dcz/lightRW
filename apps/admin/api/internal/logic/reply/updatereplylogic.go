package reply

import (
	"context"
	"github.com/ac-dcz/lightRW/apps/reply/rpc/reply"

	"github.com/ac-dcz/lightRW/apps/admin/api/internal/svc"
	"github.com/ac-dcz/lightRW/apps/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateReplyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateReplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateReplyLogic {
	return &UpdateReplyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateReplyLogic) UpdateReply(req *types.UpdateReplyReq) (resp *types.UpdateReplyResp, err error) {
	if _, err := l.svcCtx.ReplyRpc.UpdateStatus(l.ctx, &reply.UpdateStatusReq{
		ReplyId:  req.ReplyId,
		Status:   uint32(req.Status),
		OpReason: req.OpReason,
	}); err != nil {
		l.Logger.Errorf("UpdateReply - error: %v", err)
		return nil, err
	} else {
		resp = &types.UpdateReplyResp{}
	}
	return
}
