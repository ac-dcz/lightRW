package canal

import (
	"context"
	"github.com/withlin/canal-go/client"
	"github.com/withlin/canal-go/protocol"
	pbe "github.com/withlin/canal-go/protocol/entry"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/protobuf/proto"
	"strconv"
	"strings"
	"time"
)

type ClientConf struct {
	Host        string `json:",optional,default=127.0.0.1"`
	Port        int    `json:",optional,default=11111"`
	User        string `json:",optional"`
	Pass        string `json:",optional"`
	Destination string `json:",optional,default=example"`
	Subscribe   string `json:",optional,default=.*\\..*"`
	ClientId    int    `json:",optional,default=1001"`
}

type Client struct {
	c    *ClientConf
	conn *client.SimpleCanalConnector
}

func NewClient(c *ClientConf) (*Client, error) {
	canalConn := client.NewSimpleCanalConnector(
		c.Host,
		c.Port,
		c.User,
		c.Pass,
		c.Destination,
		60000,
		60*60*1000,
	)

	canalConn.ClientIdentity.ClientId = c.ClientId
	if err := canalConn.Connect(); err != nil {
		return nil, err
	}
	if err := canalConn.Subscribe(c.Subscribe); err != nil {
		return nil, err
	}
	return &Client{
		c:    c,
		conn: canalConn,
	}, nil
}

func (c *Client) Close() error {
	return c.conn.DisConnection()
}

func (c *Client) Run(ctx context.Context, handle func(records ...*Record) error) error {

	errChan := make(chan error, 1)
	msgChan := make(chan *protocol.Message, 100)
	done := make(chan struct{})

	go func() {
		for {
			select {
			case <-done:
				return
			default:
			}
			msg, err := c.conn.GetWithOutAck(100, nil, nil)
			if err != nil {
				errChan <- err
				return
			}
			//没有数据
			if msg.Id == -1 || len(msg.Entries) == 0 {
				time.Sleep(300 * time.Millisecond)
				continue
			}
			msgChan <- msg
		}
	}()

	for {
		select {
		case <-ctx.Done():
			close(done)
			return ctx.Err()
		case err := <-errChan:
			close(done)
			return err
		case msg, ok := <-msgChan:
			if !ok {
				return nil
			}
			logx.Debugf("canal entry msg id: %v", msg.Id)
			if records, err := c.parseMsg(msg); err != nil {
				return err
			} else {
				if err := handle(records...); err != nil {
					return err
				}
				if err := c.conn.Ack(msg.Id); err != nil {
					return err
				}
			}
		}
	}
}

func (c *Client) parseMsg(msg *protocol.Message) ([]*Record, error) {

	records := make([]*Record, 0)

	for i := range msg.Entries {
		entry := &msg.Entries[i]
		if entry.GetEntryType() == pbe.EntryType_TRANSACTIONBEGIN || entry.GetEntryType() == pbe.EntryType_TRANSACTIONEND {
			continue
		}
		rowChange := new(pbe.RowChange)
		if err := proto.Unmarshal(entry.GetStoreValue(), rowChange); err != nil {
			logx.Errorf("protocol unmarshal error: %v", err)
			return nil, err
		}
		eventType := rowChange.GetEventType()
		header := entry.GetHeader()
		for _, row := range rowChange.RowDatas {
			record := &Record{
				DataBase: header.GetSchemaName(),
				Table:    header.GetTableName(),
			}
			switch eventType {
			case pbe.EventType_DELETE:
				record.Type = DeleteType
				record.BeforeColumns = c.parseRowData(row.GetBeforeColumns())
			case pbe.EventType_INSERT:
				record.Type = InsertType
				record.AfterColumns = c.parseRowData(row.GetAfterColumns())
			case pbe.EventType_UPDATE:
				record.Type = UpdateType
				record.BeforeColumns = c.parseRowData(row.GetBeforeColumns())
				record.AfterColumns = c.parseRowData(row.GetAfterColumns())
			default:
				//not attention
			}
			records = append(records, record)
		}
	}

	return records, nil
}

func (c *Client) parseRowData(datas []*pbe.Column) []*Column {
	columns := make([]*Column, len(datas))
	for i, data := range datas {
		columns[i] = &Column{
			Name:      data.Name,
			Value:     convertMysqlTypeToGo(data.MysqlType, data.Value),
			MySqlType: data.MysqlType,
			IsNull:    data.GetIsNull(),
			IsUpdate:  data.GetUpdated(),
			IsKey:     data.GetIsKey(),
		}
	}
	return columns
}

func convertMysqlTypeToGo(mysqlType, value string) any {
	if strings.HasPrefix(mysqlType, "tinyint") ||
		strings.HasPrefix(mysqlType, "int") ||
		strings.HasPrefix(mysqlType, "bigint") ||
		strings.HasPrefix(mysqlType, "smallint") {
		if v, err := strconv.ParseInt(value, 10, 64); err == nil {
			return v
		}
	}
	if strings.HasPrefix(mysqlType, "decimal") ||
		strings.HasPrefix(mysqlType, "float") ||
		strings.HasPrefix(mysqlType, "double") {
		if v, err := strconv.ParseFloat(value, 64); err == nil {
			return v
		}
	}
	return value
}
