package mq

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"github.com/zeromicro/go-zero/core/logx"
	"sync/atomic"
)

type ReaderConf struct {
	Brokers []string
	Topics  []string
	GroupId string
}

type Reader struct {
	r         *kafka.Reader
	closeFlag atomic.Bool
	done      chan struct{}
}

func NewReader(conf *ReaderConf) *Reader {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        conf.Brokers,
		GroupID:        conf.GroupId,
		GroupTopics:    conf.Topics,
		MaxAttempts:    10,
		IsolationLevel: kafka.ReadCommitted,
		CommitInterval: 0, //同步提交
	})

	return &Reader{r: r, done: make(chan struct{})}
}

func (r *Reader) Close() error {
	if r.closeFlag.CompareAndSwap(false, true) {
		return fmt.Errorf("mq reader already close")
	}
	close(r.done)
	return r.r.Close()
}

func (r *Reader) Run(ctx context.Context, handle func(messages kafka.Message) error) error {
	for {
		select {
		case <-r.done:
			return nil
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		msg, err := r.r.FetchMessage(ctx)
		if err != nil {
			logx.Errorw("mq read error", logx.Field("err", err))
			return err
		}
		if err := handle(msg); err != nil {
			logx.Errorw("mq reader handle msg error", logx.Field("err", err))
			return err
		}
		if err := r.r.CommitMessages(ctx, msg); err != nil {
			logx.Errorw("mq reader commit error", logx.Field("err", err))
			return err
		}
	}
}
