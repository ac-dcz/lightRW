package logic

import (
	"context"
	omodel "github.com/ac-dcz/lightRW/apps/order/model"
	"github.com/ac-dcz/lightRW/apps/order/rpc/order"
	"github.com/ac-dcz/lightRW/apps/review/rpc/pb"
	"github.com/ac-dcz/lightRW/apps/review/rpc/review"
	"github.com/ac-dcz/lightRW/common/codes"
	"github.com/ac-dcz/lightRW/common/errors"
	"strconv"

	"github.com/ac-dcz/lightRW/apps/review/api/internal/svc"
	"github.com/ac-dcz/lightRW/apps/review/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProposeReviewLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProposeReviewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProposeReviewLogic {
	return &ProposeReviewLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProposeReviewLogic) ProposeReview(req *types.ProposeReviewReq) (resp *types.ProposeReviewResp, err error) {
	//TODO: Error Handle
	t, _ := l.ctx.Value("uid").(string)
	uid, _ := strconv.ParseUint(t, 10, 64)

	//Step1: 查询Order是否已经支付
	info, err := l.svcCtx.OrderRpc.OrderInfo(l.ctx, &order.OrderInfoReq{
		OrderId: req.OrderId,
		Uid:     uid,
	})
	if err != nil {
		l.Error(err)
		return nil, err
	}
	if info.Status != int32(omodel.Pay) {
		l.Errorf("order status invalid")
		return nil, errors.New(codes.InvalidOrderStatus, "order status invalid")
	}

	flag := false
	for _, entry := range info.Entries {
		if entry.StoreId == req.StoreId && entry.Sku == req.Sku {
			flag = true
			break
		}
	}
	if !flag {
		l.Errorf("sku not in order")
		return nil, errors.New(codes.SkuNotInOrder, "sku not in order")
	}
	//Step2:rpc call review

	if r, err := l.svcCtx.ReviewRpc.ProposeReview(l.ctx, &review.ProposeReviewReq{
		Uid:        uid,
		OrderId:    req.OrderId,
		StoreId:    req.StoreId,
		Sku:        req.Sku,
		Level:      pb.ScoreLevel(req.Level),
		GoodsDesc:  req.GoodsDesc,
		HasImage:   req.HasImage,
		ImageCDN:   req.ImageCDN,
		StoreScore: uint32(req.StoreScore),
	}); err != nil {
		l.Errorf("proposeReview err: %v", err)
		return nil, err
	} else {
		resp = &types.ProposeReviewResp{
			ReviewId: r.ReviewId,
			Status:   uint8(r.Status),
		}
	}

	return
}
