package reply

import (
	"context"
	"github.com/ac-dcz/lightRW/common/codes"
	"github.com/ac-dcz/lightRW/common/errors"

	"github.com/ac-dcz/lightRW/apps/admin/api/internal/svc"
	"github.com/ac-dcz/lightRW/apps/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReplyByMidLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReplyByMidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReplyByMidLogic {
	return &ReplyByMidLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReplyByMidLogic) ReplyByMid(req *types.ReplyByMidReq) (resp *types.ReplyResp, err error) {
	if items, err := l.svcCtx.EsReplyModel.FindByMid(l.ctx, req.Mid, req.Page, req.Size); err != nil {
		l.Logger.Errorf("find by mid error: %v", err)
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
