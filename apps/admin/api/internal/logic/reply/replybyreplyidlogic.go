package reply

import (
	"context"
	"github.com/ac-dcz/lightRW/common/codes"
	"github.com/ac-dcz/lightRW/common/errors"

	"github.com/ac-dcz/lightRW/apps/admin/api/internal/svc"
	"github.com/ac-dcz/lightRW/apps/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReplyByReplyIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReplyByReplyIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReplyByReplyIdLogic {
	return &ReplyByReplyIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReplyByReplyIdLogic) ReplyByReplyId(req *types.ReplyByReplyIdReq) (resp *types.ReplyResp, err error) {

	if items, err := l.svcCtx.EsReplyModel.FindByReplyId(l.ctx, req.ReplyId, req.Page, req.Size); err != nil {
		l.Logger.Errorf("find by replyId error: %v", err)
		return nil, errors.New(codes.InternalError, err.Error())
	} else {
		datas := make([]map[string]any, len(items))
		for i, item := range items {
			datas[i] = item
		}
		resp = &types.ReplyResp{
			Items: datas,
		}
	}

	return
}
