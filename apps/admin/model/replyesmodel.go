package model

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/elastic/go-elasticsearch/v7"
)

//type EsReply struct {
//	Id           uint64 `json:"id"`            // id
//	ReplyId      uint64 `json:"reply_id"`      // 回复id
//	Mid          uint64 `json:"mid"`           // 商家id
//	StoreId      uint64 `json:"store_id"`      // 店铺id
//	Sku          string `json:"sku"`           // sku
//	ReviewId     uint64 `json:"review_id"`     // 评价id
//	ReplyContent string `json:"reply_content"` // 回复内容
//	HasImage     uint64 `json:"has_image"`     // 0无/1有
//	ImageJson    string `json:"image_json"`    // image json
//	Status       uint64 `json:"status"`        // 状态:10待审核；20审核通过；30审核不通过；40隐藏
//	OpReason     string `json:"op_reason"`     // 运营审核拒绝原因
//	IsDel        uint64 `json:"is_del"`        // 0否/1是
//	CreateAt     string `json:"create_at"`     // 创建时间
//	UpdateAt     string `json:"update_at"`     // 更新时间
//}

type ReplyEsModel interface {
	FindByMid(ctx context.Context, mid uint64, page, pageSize int) ([]EsModel, error)
	FindByKey(ctx context.Context, key string, page, pageSize int) ([]EsModel, error)
	FindBySSku(ctx context.Context, storeId uint64, sku string, page, pageSize int) ([]EsModel, error)
	FindByReplyId(ctx context.Context, replyId uint64, page, pageSize int) ([]EsModel, error)
	FindByReviewId(ctx context.Context, reviewId uint64, page, pageSize int) ([]EsModel, error)
}

type defaultReplyEsModel struct {
	esCli *elasticsearch.Client
	index string
}

func NewReplyEsModel(esCli *elasticsearch.Client) ReplyEsModel {
	return &defaultReplyEsModel{
		esCli: esCli,
		index: "reply",
	}
}

func (m *defaultReplyEsModel) FindByMid(ctx context.Context, mid uint64, page, pageSize int) ([]EsModel, error) {
	body := M{
		"query": M{
			"term": M{
				"mid": mid,
			},
		},
		"from": (page - 1) * pageSize,
		"size": pageSize,
		"sort": []M{
			{
				"create_at": "desc",
			},
			{
				"mid": "asc",
			},
		},
	}
	return m.doSearch(ctx, body)
}

func (m *defaultReplyEsModel) FindByKey(ctx context.Context, key string, page, pageSize int) ([]EsModel, error) {
	body := M{
		"query": M{
			"match": M{
				"reply_content": key,
			},
		},
		"from": (page - 1) * pageSize,
		"size": pageSize,
		"sort": []M{
			{
				"create_at": "desc",
			},
			{
				"mid": "asc",
			},
		},
	}
	return m.doSearch(ctx, body)
}

func (m *defaultReplyEsModel) FindBySSku(ctx context.Context, storeId uint64, sku string, page, pageSize int) ([]EsModel, error) {
	body := M{
		"query": M{
			"bool": M{
				"must": []M{
					{
						"term": M{
							"sku": sku,
						},
					},
					{
						"term": M{
							"store_id": storeId,
						},
					},
				},
			},
		},
		"from": (page - 1) * pageSize,
		"size": pageSize,
		"sort": []M{
			{
				"create_at": "desc",
			},
			{
				"mid": "asc",
			},
		},
	}
	return m.doSearch(ctx, body)
}

func (m *defaultReplyEsModel) FindByReplyId(ctx context.Context, replyId uint64, page, pageSize int) ([]EsModel, error) {
	body := M{
		"query": M{
			"term": M{
				"reply_id": replyId,
			},
		},
		"from": (page - 1) * pageSize,
		"size": pageSize,
		"sort": []M{
			{
				"create_at": "desc",
			},
			{
				"mid": "asc",
			},
		},
	}
	return m.doSearch(ctx, body)
}

func (m *defaultReplyEsModel) FindByReviewId(ctx context.Context, reviewId uint64, page, pageSize int) ([]EsModel, error) {
	body := M{
		"query": M{
			"term": M{
				"review_id": reviewId,
			},
		},
		"from": (page - 1) * pageSize,
		"size": pageSize,
		"sort": []M{
			{
				"create_at": "desc",
			},
			{
				"mid": "asc",
			},
		},
	}
	return m.doSearch(ctx, body)
}

func (m *defaultReplyEsModel) doSearch(ctx context.Context, body M) ([]EsModel, error) {
	buf := bytes.NewBuffer([]byte{})
	if err := json.NewEncoder(buf).Encode(body); err != nil {
		return nil, err
	}
	resp, err := m.esCli.Search(
		m.esCli.Search.WithContext(ctx),
		m.esCli.Search.WithIndex(m.index),
		m.esCli.Search.WithPretty(),
		m.esCli.Search.WithBody(buf),
	)
	if err != nil {
		return nil, err
	} else if resp.IsError() {
		return nil, errors.New(resp.String())
	} else {
		return parseRespToEsModel(resp)
	}
}
