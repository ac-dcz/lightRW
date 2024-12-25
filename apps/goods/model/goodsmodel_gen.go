// Code generated by goctl. DO NOT EDIT.
// versions:
//  goctl version: 1.7.3

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	goodsFieldNames          = builder.RawFieldNames(&Goods{})
	goodsRows                = strings.Join(goodsFieldNames, ",")
	goodsRowsExpectAutoSet   = strings.Join(stringx.Remove(goodsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	goodsRowsWithPlaceHolder = strings.Join(stringx.Remove(goodsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	goodsModel interface {
		Insert(ctx context.Context, data *Goods) (sql.Result, error)
		FindOne(ctx context.Context, id uint64) (*Goods, error)
		FindOneBySku(ctx context.Context, sku string) (*Goods, error)
		Update(ctx context.Context, data *Goods) error
		Delete(ctx context.Context, id uint64) error
	}

	defaultGoodsModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Goods struct {
		Id       uint64    `db:"id"`        // id
		Sku      string    `db:"sku"`       // sku
		Uid      uint64    `db:"uid"`       // uid
		Name     string    `db:"name"`      // 名称
		CreateAt time.Time `db:"create_at"` // 创建时间
		UpdateAt time.Time `db:"update_at"` // 更新时间
	}
)

func newGoodsModel(conn sqlx.SqlConn) *defaultGoodsModel {
	return &defaultGoodsModel{
		conn:  conn,
		table: "`goods`",
	}
}

func (m *defaultGoodsModel) Delete(ctx context.Context, id uint64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultGoodsModel) FindOne(ctx context.Context, id uint64) (*Goods, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", goodsRows, m.table)
	var resp Goods
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultGoodsModel) FindOneBySku(ctx context.Context, sku string) (*Goods, error) {
	var resp Goods
	query := fmt.Sprintf("select %s from %s where `sku` = ? limit 1", goodsRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, sku)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultGoodsModel) Insert(ctx context.Context, data *Goods) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, goodsRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Sku, data.Uid, data.Name)
	return ret, err
}

func (m *defaultGoodsModel) Update(ctx context.Context, newData *Goods) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, goodsRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.Sku, newData.Uid, newData.Name, newData.Id)
	return err
}

func (m *defaultGoodsModel) tableName() string {
	return m.table
}
