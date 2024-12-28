package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ GoodsStoreModel = (*customGoodsStoreModel)(nil)

type (
	// GoodsStoreModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGoodsStoreModel.
	GoodsStoreModel interface {
		goodsStoreModel
		withSession(session sqlx.Session) GoodsStoreModel
		TableName() string
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
