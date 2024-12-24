package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ StoreModel = (*customStoreModel)(nil)

type (
	// StoreModel is an interface to be customized, add more methods here,
	// and implement the added methods in customStoreModel.
	StoreModel interface {
		storeModel
	}

	customStoreModel struct {
		*defaultStoreModel
	}
)

// NewStoreModel returns a model for the database table.
func NewStoreModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) StoreModel {
	return &customStoreModel{
		defaultStoreModel: newStoreModel(conn, c, opts...),
	}
}
