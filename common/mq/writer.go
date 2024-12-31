package mq

import (
	"context"
	"errors"
	"github.com/segmentio/kafka-go"
	"github.com/zeromicro/go-zero/core/logx"
	"sync/atomic"
	"time"
)

type WriterConf struct {
	Brokers []string
	Topic   string
}

type callBackMsg struct {
	ctx      context.Context
	msg      *kafka.Message
	callback func(msg *kafka.Message, err error)
}

type Writer struct {
	w         *kafka.Writer
	msgChan   chan *callBackMsg
	closeFlag atomic.Bool
	done      chan struct{}
}

func NewWriter(conf *WriterConf) *Writer {
	w := &kafka.Writer{
		Addr:                   kafka.TCP(conf.Brokers...),
		Topic:                  conf.Topic,
		RequiredAcks:           kafka.RequireAll,
		MaxAttempts:            10,
		AllowAutoTopicCreation: true,
		Balancer:               &kafka.LeastBytes{},
	}
	writer := &Writer{w: w, msgChan: make(chan *callBackMsg, 100), done: make(chan struct{})}

	go writer.run()

	return writer
}

func (w *Writer) run() {
	for {
		select {
		case <-w.done:
			return
		case msg := <-w.msgChan:
			for {
				err := w.w.WriteMessages(msg.ctx, *msg.msg)
				if errors.Is(err, kafka.LeaderNotAvailable) {
					time.Sleep(time.Millisecond * 100)
					continue
				} else if err != nil {
					logx.Errorw("mq write message error", logx.Field("err", err))
					if w.closeFlag.CompareAndSwap(false, true) {
						close(w.msgChan)
					}
					msg.callback(msg.msg, err)
				}
				break
			}
		}
	}
}

func (w *Writer) Write(ctx context.Context, msg *kafka.Message, callback func(msg *kafka.Message, err error)) error {
	if w.closeFlag.Load() {
		return errors.New("mq writer already close")
	}
	w.msgChan <- &callBackMsg{ctx: ctx, msg: msg, callback: callback}
	return nil
}

func (w *Writer) Close() error {
	if !w.closeFlag.CompareAndSwap(false, true) {
		return errors.New("mq writer already close")
	}
	close(w.msgChan)
	close(w.done)
	return w.w.Close()
}
