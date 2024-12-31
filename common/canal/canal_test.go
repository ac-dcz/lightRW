package canal

import (
	"context"
	"fmt"
	"github.com/ac-dcz/lightRW/common/mq"
	"github.com/segmentio/kafka-go"
	"testing"
)

var writer *mq.Writer

func init() {
	writer = mq.NewWriter(&mq.WriterConf{
		Topic:   "example-test-0",
		Brokers: []string{"127.0.0.1:9092"},
	})
}

func TestClient_Run(t *testing.T) {
	client, err := NewClient(&ClientConf{
		Host:        "127.0.0.1",
		Port:        11111,
		Destination: "example",
		Subscribe:   "canal_test\\..*",
	})
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	t.Fatal(client.Run(context.Background(), kafkaTestLogic))
}

func consumerTestLogic(records ...*Record) error {
	for _, record := range records {
		fmt.Println("=============================================")
		fmt.Printf("DataBase: %s Table: %s EntryType: %v \n", record.DataBase, record.Table, record.Type)

		fmt.Println(">>> BeforeColumns:")
		for _, column := range record.BeforeColumns {
			fmt.Printf("\t name: %s \n", column.Name)
			fmt.Printf("\t value: %s \n", column.Value)
			fmt.Printf("\t isnull: %v \n", column.IsNull)
			fmt.Printf("\t isupdate: %v \n", column.IsUpdate)
		}
		fmt.Println()
		fmt.Println(">>> AfterColumns:")
		for _, column := range record.AfterColumns {
			fmt.Printf("\t name: %s \n", column.Name)
			fmt.Printf("\t value: %s \n", column.Value)
			fmt.Printf("\t isnull: %v \n", column.IsNull)
			fmt.Printf("\t isupdate: %v \n", column.IsUpdate)
		}

	}
	return nil
}

func kafkaTestLogic(records ...*Record) error {
	for _, record := range records {
		if value, err := record.Encode(); err != nil {
			return err
		} else {
			if err := writer.Write(context.Background(), &kafka.Message{Value: value}, func(msg *kafka.Message, err error) {

			}); err != nil {
				return err
			}
		}
	}
	return nil
}
