package model

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/ac-dcz/lightRW/common/canal"
	"github.com/elastic/go-elasticsearch/v7"
	"strconv"
)

type EsReplyModel interface {
	Insert(ctx context.Context, id string, data map[string]any) error
	Update(ctx context.Context, id string, data map[string]any) error
}
type (
	defaultEsReplyModel struct {
		esCli *elasticsearch.Client
		index string
	}
)

func NewEsReplyModel(esCli *elasticsearch.Client) EsReplyModel {
	return &defaultEsReplyModel{
		esCli: esCli,
		index: "reply",
	}
}

func (m *defaultEsReplyModel) Insert(ctx context.Context, id string, data map[string]any) error {
	buf := bytes.NewBuffer(nil)
	if err := json.NewEncoder(buf).Encode(data); err != nil {
		return err
	}

	resp, err := m.esCli.Index(
		m.index,
		buf,
		m.esCli.Index.WithDocumentID(id),
		m.esCli.Index.WithContext(ctx),
		m.esCli.Index.WithPretty(),
	)
	if err != nil {
		return err
	} else if resp.IsError() {
		return errors.New(resp.String())
	}
	return nil
}

func (m *defaultEsReplyModel) Update(ctx context.Context, id string, data map[string]any) error {
	buf := bytes.NewBuffer(nil)
	if err := json.NewEncoder(buf).Encode(data); err != nil {
		return err
	}

	resp, err := m.esCli.Index(
		m.index,
		buf,
		m.esCli.Index.WithDocumentID(id),
		m.esCli.Index.WithContext(ctx),
	)
	if err != nil {
		return err
	} else if resp.IsError() {
		return errors.New(resp.String())
	}
	return nil
}

func ParseToReply(record *canal.Record) (string, map[string]any, error) {
	id := ""
	data := map[string]any{}
	for _, column := range record.AfterColumns {
		data[column.Name] = column.Value
		if column.Name == "id" {
			if t, ok := column.Value.(int64); !ok {
				return "", nil, errors.New("id type error")
			} else {
				id = strconv.FormatInt(t, 10)
			}
		}
	}
	return id, data, nil
}
