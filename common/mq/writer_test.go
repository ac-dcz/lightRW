package mq

import (
	"context"
	"github.com/segmentio/kafka-go"
	"testing"
	"time"
)

func TestWriter_Write(t *testing.T) {
	w := NewWriter(&WriterConf{
		Brokers: []string{"127.0.0.1:9092"},
		Topic:   "example-test-0",
	})

	ticker := time.NewTicker(time.Second * 3)

	for range ticker.C {
		msg := &kafka.Message{
			Value: []byte("test kafka"),
		}
		if err := w.Write(context.Background(), msg, func(msg *kafka.Message, err error) {
			t.Errorf("write message error: %s", err)
		}); err != nil {
			t.Fatal(err)
		}
	}

}
