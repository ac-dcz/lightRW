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
	userFieldNames          = builder.RawFieldNames(&User{})
	userRows                = strings.Join(userFieldNames, ",")
	userRowsExpectAutoSet   = strings.Join(stringx.Remove(userFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userRowsWithPlaceHolder = strings.Join(stringx.Remove(userFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheUserIdPrefix  = "cache:user:id:"
	cacheUserTelPrefix = "cache:user:tel:"
)

type (
	userModel interface {
		Insert(ctx context.Context, data *User) (sql.Result, error)
		FindOne(ctx context.Context, id uint64) (*User, error)
		FindOneByTel(ctx context.Context, tel string) (*User, error)
		Update(ctx context.Context, data *User) error
		Delete(ctx context.Context, id uint64) error
	}

	defaultUserModel struct {
		sqlc.CachedConn
		table string
	}

	User struct {
		Id       uint64    `db:"id"`        // id
		NickName string    `db:"nick_name"` // 昵称
		Tel      string    `db:"tel"`       // 手机号
		Password string    `db:"password"`  // 密码
		Level    []byte      `db:"level"`     // 1普通用户/2商家/4管理员
		Status   uint64    `db:"status"`    // 0有效/1注销
		CreateAt time.Time `db:"create_at"` // 创建时间
		UpdateAt time.Time `db:"update_at"` // 更新时间
	}
)

func newUserModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultUserModel {
	return &defaultUserModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`user`",
	}
}

func (m *defaultUserModel) Delete(ctx context.Context, id uint64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	userIdKey := fmt.Sprintf("%s%v", cacheUserIdPrefix, id)
	userTelKey := fmt.Sprintf("%s%v", cacheUserTelPrefix, data.Tel)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, userIdKey, userTelKey)
	return err
}

func (m *defaultUserModel) FindOne(ctx context.Context, id uint64) (*User, error) {
	userIdKey := fmt.Sprintf("%s%v", cacheUserIdPrefix, id)
	var resp User
	err := m.QueryRowCtx(ctx, &resp, userIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userRows, m.table)
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

func (m *defaultUserModel) FindOneByTel(ctx context.Context, tel string) (*User, error) {
	userTelKey := fmt.Sprintf("%s%v", cacheUserTelPrefix, tel)
	var resp User
	err := m.QueryRowIndexCtx(ctx, &resp, userTelKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `tel` = ? limit 1", userRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, tel); err != nil {
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

func (m *defaultUserModel) Insert(ctx context.Context, data *User) (sql.Result, error) {
	userIdKey := fmt.Sprintf("%s%v", cacheUserIdPrefix, data.Id)
	userTelKey := fmt.Sprintf("%s%v", cacheUserTelPrefix, data.Tel)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, userRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.NickName, data.Tel, data.Password, data.Level, data.Status)
	}, userIdKey, userTelKey)
	return ret, err
}

func (m *defaultUserModel) Update(ctx context.Context, newData *User) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	userIdKey := fmt.Sprintf("%s%v", cacheUserIdPrefix, data.Id)
	userTelKey := fmt.Sprintf("%s%v", cacheUserTelPrefix, data.Tel)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.NickName, newData.Tel, newData.Password, newData.Level, newData.Status, newData.Id)
	}, userIdKey, userTelKey)
	return err
}

func (m *defaultUserModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheUserIdPrefix, primary)
}

func (m *defaultUserModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUserModel) tableName() string {
	return m.table
}
