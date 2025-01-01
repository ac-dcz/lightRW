package logic

import (
	"context"
	"github.com/ac-dcz/lightRW/apps/reply/rpc/reply"
	"strconv"

	"github.com/ac-dcz/lightRW/apps/reply/api/internal/svc"
	"github.com/ac-dcz/lightRW/apps/reply/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProposeReplyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProposeReplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProposeReplyLogic {
	return &ProposeReplyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProposeReplyLogic) ProposeReply(req *types.ProposeReplyReq) (resp *types.ProposeReplyResp, err error) {
	// todo: handel error
	t, _ := l.ctx.Value("uid").(string)
	mid, _ := strconv.ParseUint(t, 10, 64)
	// todo: level is ok ?

	if r, err := l.svcCtx.ReplyRpc.ProposeReply(l.ctx, &reply.ProposeReplyReq{
		Mid:      mid,
		ReviewId: req.ReviewId,
		Sku:      req.Sku,
		StoreId:  req.StoreId,
		Content:  req.Content,
		HasImage: uint32(req.HasImage),
		ImageCDN: req.ImageCDN,
	}); err != nil {
		l.Errorf("propose reply rpc error: %v", err)
		return nil, err
	} else {
		resp = &types.ProposeReplyResp{
			ReplyId: r.ReplyId,
			Status:  uint8(r.Status),
		}
	}

	return
}
