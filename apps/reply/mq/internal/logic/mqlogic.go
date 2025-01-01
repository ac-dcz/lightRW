package logic

import (
	"context"
	"github.com/ac-dcz/lightRW/apps/review/model"
	"github.com/ac-dcz/lightRW/apps/review/mq/internal/svc"
	"github.com/ac-dcz/lightRW/common/canal"
	"github.com/segmentio/kafka-go"
	"github.com/zeromicro/go-zero/core/logx"
)

type MqLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMqLogic(ctx context.Context, svcCtx *svc.ServiceContext) MqLogic {
	return MqLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MqLogic) Handle(msg kafka.Message) error {
	r := &canal.Record{}
	if err := r.Decode(msg.Value); err != nil {
		logx.Errorw("decode kafka message error", logx.Field("err", err))
		return err
	}
	logx.Infof("handel record %s %s %v ", r.DataBase, r.Table, r.Type)
	id, data, err := model.ParseToReview(r)
	if err != nil {
		logx.Errorw("parse kafka message error", logx.Field("err", err))
		return err
	}
	if r.Type == canal.UpdateType {
		if err := l.svcCtx.EsReviewModel.Update(l.ctx, id, data); err != nil {
			logx.Errorw("update es record error", logx.Field("err", err))
			return err
		}
	} else if r.Type == canal.InsertType {
		if err := l.svcCtx.EsReviewModel.Insert(l.ctx, id, data); err != nil {
			logx.Errorw("insert es record error", logx.Field("err", err))
			return err
		}
	}
	return nil
}
