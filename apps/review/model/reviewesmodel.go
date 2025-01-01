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

type EsReviewModel interface {
	Insert(ctx context.Context, id string, data map[string]any) error
	Update(ctx context.Context, id string, data map[string]any) error
}
type (
	defaultEsReviewModel struct {
		esCli *elasticsearch.Client
		index string
	}

	customEsReviewModel struct {
		*defaultEsReviewModel
	}
)

func NewEsReviewModel(esCli *elasticsearch.Client) EsReviewModel {
	return &customEsReviewModel{
		defaultEsReviewModel: newEsReviewModel(esCli),
	}
}

func newEsReviewModel(esCli *elasticsearch.Client) *defaultEsReviewModel {
	return &defaultEsReviewModel{
		esCli: esCli,
		index: "review",
	}
}

func (m *defaultEsReviewModel) Insert(ctx context.Context, id string, data map[string]any) error {
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

func (m *defaultEsReviewModel) Update(ctx context.Context, id string, data map[string]any) error {
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

func ParseToReview(record *canal.Record) (string, map[string]any, error) {
	id := ""
	data := map[string]any{}
	for _, column := range record.AfterColumns {
		data[column.Name] = column.Value
		if column.Name == "id" {
			if t, ok := column.Value.(float64); !ok {
				if t, ok := column.Value.(string); !ok {
					return "", nil, errors.New("id type error")
				} else {
					id = t
				}
			} else {
				id = strconv.FormatInt(int64(t), 10)
			}
		}
	}
	return id, data, nil
}

//ES:Mapping
/*
PUT /review
{
  "mappings": {
    "properties":{
      "id":{
        "type": "long"
      },
      "review_id":{
        "type": "long"
      },
      "uid":{
        "type": "long"
      },
      "order_id":{
        "type": "long"
      },
      "store_id":{
        "type": "long"
      },
      "sku":{
        "type": "keyword"
      },
      "score":{
        "type": "short"
      },
      "good_desc":{
        "type": "text",
        "analyzer": "ik_smart"
      },
      "has_image":{
        "type": "short",
        "index": false
      },
      "image_json":{
        "type": "keyword",
        "index": false
      },
      "store_score":{
        "type": "short",
        "index": false
      },
      "is_reply":{
        "type": "short",
        "index": false
      },
      "status":{
        "type": "short"
      },
      "op_reason":{
        "type": "text",
        "index": false
      },
      "goods_snapshot":{
        "type": "text",
        "index": false
      },
      "is_del":{
        "type": "short",
        "index": false
      },
      "create_at":{
        "type": "date",
        "format": ["yyyy-MM-dd HH:mm:ss"]
      },
      "update_at":{
        "type": "date",
        "format": ["yyyy-MM-dd HH:mm:ss"]
      }
    }
  }
}
*/
