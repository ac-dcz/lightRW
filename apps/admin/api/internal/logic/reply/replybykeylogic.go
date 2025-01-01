package reply

import (
	"context"
	"github.com/ac-dcz/lightRW/common/codes"
	"github.com/ac-dcz/lightRW/common/errors"

	"github.com/ac-dcz/lightRW/apps/admin/api/internal/svc"
	"github.com/ac-dcz/lightRW/apps/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReplyByKeyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReplyByKeyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReplyByKeyLogic {
	return &ReplyByKeyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReplyByKeyLogic) ReplyByKey(req *types.ReplyByKeyReq) (resp *types.ReplyResp, err error) {

	if items, err := l.svcCtx.EsReplyModel.FindByKey(l.ctx, req.Key, req.Page, req.Size); err != nil {
		l.Logger.Errorf("find by key error: %v", err)
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
