package model

import (
	"context"
	stderr "errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ GoodsStoreModel = (*customGoodsStoreModel)(nil)

type (
	// GoodsStoreModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGoodsStoreModel.
	GoodsStoreModel interface {
		goodsStoreModel
		withSession(session sqlx.Session) GoodsStoreModel
		TableName() string
		FindManyByStoreId(ctx context.Context, storeId uint64) ([]*GoodsStore, error)
	}

	customGoodsStoreModel struct {
		*defaultGoodsStoreModel
	}
)

// NewGoodsStoreModel returns a model for the database table.
func NewGoodsStoreModel(conn sqlx.SqlConn) GoodsStoreModel {
	return &customGoodsStoreModel{
		defaultGoodsStoreModel: newGoodsStoreModel(conn),
	}
}

func (m *customGoodsStoreModel) withSession(session sqlx.Session) GoodsStoreModel {
	return NewGoodsStoreModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customGoodsStoreModel) TableName() string {
	return m.table
}

func (m *customGoodsStoreModel) FindManyByStoreId(ctx context.Context, storeId uint64) ([]*GoodsStore, error) {
	var datas []*GoodsStore
	query := fmt.Sprintf("select %s from %s where store_id = ?", goodsStoreRows, m.table)
	if err := m.conn.QueryRowsCtx(ctx, &datas, query, storeId); err != nil {
		if stderr.Is(err, sqlx.ErrNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return datas, nil
}
