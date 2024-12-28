package logic

import (
	"context"
	"github.com/ac-dcz/lightRW/apps/genid/rpc/genid"
	omodel "github.com/ac-dcz/lightRW/apps/order/model"
	"github.com/ac-dcz/lightRW/apps/order/rpc/order"
	"github.com/ac-dcz/lightRW/apps/review/model"
	"github.com/ac-dcz/lightRW/apps/review/rpc/internal/svc"
	"github.com/ac-dcz/lightRW/apps/review/rpc/pb"
	"github.com/ac-dcz/lightRW/common/codes"
	"github.com/ac-dcz/lightRW/common/errors"
	"github.com/ac-dcz/lightRW/common/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProposeReviewLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProposeReviewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProposeReviewLogic {
	return &ProposeReviewLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ProposeReview 发表评论
func (l *ProposeReviewLogic) ProposeReview(in *pb.ProposeReviewReq) (*pb.ProposeReviewResp, error) {
	//step1: 检查store_id,sku是否存在在订单中
	//step2: 检查该用户是否已经完成该订单
	info, err := l.svcCtx.OrderRpc.OrderInfo(l.ctx, &order.OrderInfoReq{
		OrderId: in.OrderId,
		Uid:     in.Uid,
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
		if entry.StoreId == in.StoreId && entry.Sku == in.Sku {
			flag = true
			break
		}
	}
	if !flag {
		l.Errorf("sku not in order")
		return nil, errors.New(codes.SkuNotInOrder, "sku not in order")
	}
	//step3: 生成评论id
	r, err := l.svcCtx.GenIdRpc.GetId(l.ctx, &genid.GetIdReq{})
	if err != nil {
		l.Error(err)
		return nil, err
	}
	//step4: 存入评论
	review := &model.Review{
		ReviewId:   r.Id,
		Uid:        in.Uid,
		OrderId:    in.OrderId,
		StoreId:    in.StoreId,
		Sku:        in.Sku,
		Score:      uint64(in.Level),
		GoodsDesc:  in.GoodsDesc,
		StoreScore: uint64(in.StoreScore),
		HasImage:   uint64(utils.BoolToInt(in.HasImage)),
		ImageJson:  in.ImageCDN,
	}
	if _, err := l.svcCtx.ReviewModel.Insert(l.ctx, review); err != nil {
		l.Errorf("review model insert error: %v", err)
		return nil, errors.New(codes.InternalError, err.Error())
	}

	return &pb.ProposeReviewResp{
		ReviewId: review.ReviewId,
		Status:   model.AuditIng,
	}, nil
}
