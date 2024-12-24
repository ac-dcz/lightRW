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
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	storeFieldNames          = builder.RawFieldNames(&Store{})
	storeRows                = strings.Join(storeFieldNames, ",")
	storeRowsExpectAutoSet   = strings.Join(stringx.Remove(storeFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	storeRowsWithPlaceHolder = strings.Join(stringx.Remove(storeFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheStoreIdPrefix      = "cache:store:id:"
	cacheStoreStoreIdPrefix = "cache:store:storeId:"
)

type (
	storeModel interface {
		Insert(ctx context.Context, data *Store) (sql.Result, error)
		FindOne(ctx context.Context, id uint64) (*Store, error)
		FindOneByStoreId(ctx context.Context, storeId uint64) (*Store, error)
		Update(ctx context.Context, data *Store) error
		Delete(ctx context.Context, id uint64) error
	}

	defaultStoreModel struct {
		sqlc.CachedConn
		table string
	}

	Store struct {
		Id       uint64    `db:"id"`
		StoreId  uint64    `db:"store_id"`  // 店铺id
		Name     string    `db:"name"`      // 店铺名称
		Uid      uint64    `db:"uid"`       // 拥有者
		CreatAt  time.Time `db:"creat_at"`  // 创建时间
		UpdateAt time.Time `db:"update_at"` // 更新时间
	}
)

func newStoreModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultStoreModel {
	return &defaultStoreModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`store`",
	}
}

func (m *defaultStoreModel) Delete(ctx context.Context, id uint64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	storeIdKey := fmt.Sprintf("%s%v", cacheStoreIdPrefix, id)
	storeStoreIdKey := fmt.Sprintf("%s%v", cacheStoreStoreIdPrefix, data.StoreId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, storeIdKey, storeStoreIdKey)
	return err
}

func (m *defaultStoreModel) FindOne(ctx context.Context, id uint64) (*Store, error) {
	storeIdKey := fmt.Sprintf("%s%v", cacheStoreIdPrefix, id)
	var resp Store
	err := m.QueryRowCtx(ctx, &resp, storeIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", storeRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultStoreModel) FindOneByStoreId(ctx context.Context, storeId uint64) (*Store, error) {
	storeStoreIdKey := fmt.Sprintf("%s%v", cacheStoreStoreIdPrefix, storeId)
	var resp Store
	err := m.QueryRowIndexCtx(ctx, &resp, storeStoreIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `store_id` = ? limit 1", storeRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, storeId); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultStoreModel) Insert(ctx context.Context, data *Store) (sql.Result, error) {
	storeIdKey := fmt.Sprintf("%s%v", cacheStoreIdPrefix, data.Id)
	storeStoreIdKey := fmt.Sprintf("%s%v", cacheStoreStoreIdPrefix, data.StoreId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, storeRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.StoreId, data.Name, data.Uid, data.CreatAt)
	}, storeIdKey, storeStoreIdKey)
	return ret, err
}

func (m *defaultStoreModel) Update(ctx context.Context, newData *Store) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	storeIdKey := fmt.Sprintf("%s%v", cacheStoreIdPrefix, data.Id)
	storeStoreIdKey := fmt.Sprintf("%s%v", cacheStoreStoreIdPrefix, data.StoreId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, storeRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.StoreId, newData.Name, newData.Uid, newData.CreatAt, newData.Id)
	}, storeIdKey, storeStoreIdKey)
	return err
}

func (m *defaultStoreModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheStoreIdPrefix, primary)
}

func (m *defaultStoreModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", storeRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultStoreModel) tableName() string {
	return m.table
}