package mq

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"testing"
)

func TestReader_Run(t *testing.T) {

	r := NewReader(&ReaderConf{
		Brokers: []string{"127.0.0.1:9092"},
		Topics:  []string{"example-test-0"},
		GroupId: "example-test-0-c",
	})
	defer r.Close()
	if err := r.Run(context.Background(), func(msg kafka.Message) error {
		fmt.Println("Key: ", string(msg.Key))
		fmt.Println("Value: ", string(msg.Value))
		return nil
	}); err != nil {
		t.Fatal(err)
	}

}
