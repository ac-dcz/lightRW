package review

import (
	"context"
	"github.com/ac-dcz/lightRW/common/codes"
	"github.com/ac-dcz/lightRW/common/errors"

	"github.com/ac-dcz/lightRW/apps/admin/api/internal/svc"
	"github.com/ac-dcz/lightRW/apps/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReviewByUidLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReviewByUidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReviewByUidLogic {
	return &ReviewByUidLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReviewByUidLogic) ReviewByUid(req *types.ReviewByUidReq) (resp *types.ReviewResp, err error) {
	if items, err := l.svcCtx.EsReviewModel.FindByUid(l.ctx, req.Uid, req.Page, req.Size); err != nil {
		l.Errorf("find by uid error: %v", err)
		return nil, errors.New(codes.InternalError, err.Error())
	} else {
		datas := make([]map[string]any, len(items))
		for i, item := range items {
			datas[i] = item
		}
		resp = &types.ReviewResp{
			Items: datas,
		}
	}

	return
}
