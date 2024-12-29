package logic

import (
	"context"
	"github.com/ac-dcz/lightRW/apps/reply/rpc/reply"
	"strconv"

	"github.com/ac-dcz/lightRW/apps/reply/api/internal/svc"
	"github.com/ac-dcz/lightRW/apps/reply/api/internal/types"

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

func (l *ReplyByMidLogic) ReplyByMid() (resp *types.ReplyByMidResp, err error) {
	// todo: handel error
	t, _ := l.ctx.Value("uid").(string)
	mid, _ := strconv.ParseUint(t, 10, 64)
	// todo: level is ok ?

	if r, err := l.svcCtx.ReplyRpc.ReplyByMid(l.ctx, &reply.ReplyByMidReq{
		Mid: mid,
	}); err != nil {
		l.Errorf("call ReplyByMid error: %v", err)
		return nil, err
	} else {
		infos := make([]types.ReplyInfo, len(r.Infos))
		for i, info := range r.Infos {
			infos[i] = types.ReplyInfo{
				ReplyId:  info.ReplyId,
				Mid:      info.Mid,
				ReviewId: info.ReviewId,
				StoreId:  info.StoreId,
				Sku:      info.Sku,
				Status:   uint8(info.Status),
				Content:  info.Content,
				HasImage: info.HasImage,
				ImageCDN: info.ImageCDN,
				CreateAt: info.CreateAt,
				IsDel:    uint8(info.IsDel),
			}
		}
		resp = &types.ReplyByMidResp{
			Infos: infos,
		}
	}

	return
}
