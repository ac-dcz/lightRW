package model

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/elastic/go-elasticsearch/v7"
)

//type EsReview struct {
//	Id            uint64 `json:"id"`
//	ReviewId      uint64 `json:"review_id"`      // 评论id
//	Uid           uint64 `json:"uid"`            // 用户id
//	OrderId       uint64 `json:"order_id"`       // 订单id
//	StoreId       uint64 `json:"store_id"`       // 店铺id
//	Sku           string `json:"sku"`            // sku
//	Score         uint64 `json:"score"`          // 0差评/1中评/2好评
//	GoodsDesc     string `json:"goods_desc"`     // 商品描述
//	HasImage      uint64 `json:"has_image"`      // 1有/0无
//	ImageJson     string `json:"image_json"`     // image json
//	StoreScore    uint64 `json:"store_score"`    // 1-5星
//	IsReply       uint64 `json:"is_reply"`       // 0否/1是
//	Status        uint64 `json:"status"`         // 状态:10待审核；20审核通过；30审核不通过；40隐藏
//	OpReason      string `json:"op_reason"`      // 运营审核拒绝原因
//	GoodsSnapshot string `json:"goods_snapshot"` // 商品快照信息
//	IsDel         uint64 `json:"is_del"`         // 0否/1是
//	CreateAt      string `json:"create_at"`      // 创建时间
//	UpdateAt      string `json:"update_at"`      // 更新时间
//}

type ReviewEsModel interface {
	FindByUid(ctx context.Context, uid uint64, page, pageSize int) ([]EsModel, error)
	FindByKey(ctx context.Context, key string, page, pageSize int) ([]EsModel, error)
	FindBySSku(ctx context.Context, storeId uint64, sku string, page, pageSize int) ([]EsModel, error)
	FindByReviewId(ctx context.Context, reviewId uint64, page, pageSize int) ([]EsModel, error)
}

type defaultEsReviewModel struct {
	esCli *elasticsearch.Client
	index string
}

func NewReviewEsModel(esCli *elasticsearch.Client) ReviewEsModel {
	return &defaultEsReviewModel{
		esCli: esCli,
		index: "review",
	}
}

func (m *defaultEsReviewModel) FindByReviewId(ctx context.Context, reviewId uint64, page, pageSize int) ([]EsModel, error) {
	body := M{
		"query": M{
			"term": M{
				"review_id": M{
					"value": reviewId,
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
				"uid": "asc",
			},
		},
	}
	return m.doSearch(ctx, body)
}

func (m *defaultEsReviewModel) FindByUid(ctx context.Context, uid uint64, page, pageSize int) ([]EsModel, error) {
	body := M{
		"query": M{
			"term": M{
				"uid": M{
					"value": uid,
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
				"uid": "asc",
			},
		},
	}
	return m.doSearch(ctx, body)
}

func (m *defaultEsReviewModel) FindByKey(ctx context.Context, key string, page, pageSize int) ([]EsModel, error) {
	body := M{
		"query": M{
			"match": M{
				"good_desc": key,
			},
		},
		"from": (page - 1) * pageSize,
		"size": pageSize,
		"sort": []M{
			{
				"create_at": "desc",
			},
			{
				"uid": "asc",
			},
		},
	}
	return m.doSearch(ctx, body)
}

func (m *defaultEsReviewModel) FindBySSku(ctx context.Context, storeId uint64, sku string, page, pageSize int) ([]EsModel, error) {
	body := M{
		"query": M{
			"bool": M{
				"filter": []M{
					{
						"term": M{
							"store_id": storeId,
						},
					},
					{
						"term": M{
							"sku": sku,
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
				"uid": "asc",
			},
		},
	}
	return m.doSearch(ctx, body)
}

func (m *defaultEsReviewModel) doSearch(ctx context.Context, body M) ([]EsModel, error) {
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
