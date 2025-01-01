package logic

import (
	"context"
	"github.com/ac-dcz/lightRW/apps/review/canal/internal/svc"
	"github.com/ac-dcz/lightRW/common/canal"
	"github.com/segmentio/kafka-go"
	"github.com/zeromicro/go-zero/core/logx"
)

type MqLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMqLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MqLogic {
	return &MqLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MqLogic) Handle(records ...*canal.Record) error {
	for _, record := range records {
		logx.Infow("canal record", logx.Field("header", map[string]any{
			"database": record.DataBase,
			"table":    record.Table,
			"type":     record.Type,
		}))
		data, err := record.Encode()
		if err != nil {
			logx.Errorw("record parse error", logx.Field("err", err))
			return err
		}
		logx.Infow("canal record", logx.Field("data", string(data)))

		//write to kafka
		if err := l.svcCtx.KqWriter.Write(l.ctx, &kafka.Message{
			Value: data,
		}, func(msg *kafka.Message, err error) {
			logx.Errorw("record write kafka error",
				logx.Field("err", err),
				logx.Field("data", string(data)),
			)
		}); err != nil {
			logx.Errorw("record write kafka error", logx.Field("err", err))
			return err
		}

	}
	return nil
}
