// Code generated by goctl. DO NOT EDIT.
// versions:
//  goctl version: 1.7.3

package reply

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
	replyFieldNames          = builder.RawFieldNames(&Reply{})
	replyRows                = strings.Join(replyFieldNames, ",")
	replyRowsExpectAutoSet   = strings.Join(stringx.Remove(replyFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	replyRowsWithPlaceHolder = strings.Join(stringx.Remove(replyFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheReplyIdPrefix = "cache:reply:id:"
)

type (
	replyModel interface {
		Insert(ctx context.Context, data *Reply) (sql.Result, error)
		FindOne(ctx context.Context, id uint64) (*Reply, error)
		Update(ctx context.Context, data *Reply) error
		Delete(ctx context.Context, id uint64) error
	}

	defaultReplyModel struct {
		sqlc.CachedConn
		table string
	}

	Reply struct {
		Id           uint64    `db:"id"`            // id
		ReplyId      uint64    `db:"reply_id"`      // 回复id
		Mid          uint64    `db:"mid"`           // 商家id
		StoreId      uint64    `db:"store_id"`      // 店铺id
		Sku          uint64    `db:"sku"`           // sku
		ReviewId     uint64    `db:"review_id"`     // 评价id
		ReplyContent string    `db:"reply_content"` // 回复内容
		HasImage     uint64    `db:"has_image"`     // 0无/1有
		ImageJson    string    `db:"image_json"`    // image json
		Status       uint64    `db:"status"`        // 状态:10待审核；20审核通过；30审核不通过；40隐藏
		OpReason     string    `db:"op_reason"`     // 运营审核拒绝原因
		IsDel        uint64    `db:"is_del"`        // 0否/1是
		CreateAt     time.Time `db:"create_at"`     // 创建时间
		UpdateAt     time.Time `db:"update_at"`     // 更新时间
	}
)

func newReplyModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultReplyModel {
	return &defaultReplyModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`reply`",
	}
}

func (m *defaultReplyModel) Delete(ctx context.Context, id uint64) error {
	replyIdKey := fmt.Sprintf("%s%v", cacheReplyIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, replyIdKey)
	return err
}

func (m *defaultReplyModel) FindOne(ctx context.Context, id uint64) (*Reply, error) {
	replyIdKey := fmt.Sprintf("%s%v", cacheReplyIdPrefix, id)
	var resp Reply
	err := m.QueryRowCtx(ctx, &resp, replyIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", replyRows, m.table)
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

func (m *defaultReplyModel) Insert(ctx context.Context, data *Reply) (sql.Result, error) {
	replyIdKey := fmt.Sprintf("%s%v", cacheReplyIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, replyRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.ReplyId, data.Mid, data.StoreId, data.Sku, data.ReviewId, data.ReplyContent, data.HasImage, data.ImageJson, data.Status, data.OpReason, data.IsDel)
	}, replyIdKey)
	return ret, err
}

func (m *defaultReplyModel) Update(ctx context.Context, data *Reply) error {
	replyIdKey := fmt.Sprintf("%s%v", cacheReplyIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, replyRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.ReplyId, data.Mid, data.StoreId, data.Sku, data.ReviewId, data.ReplyContent, data.HasImage, data.ImageJson, data.Status, data.OpReason, data.IsDel, data.Id)
	}, replyIdKey)
	return err
}

func (m *defaultReplyModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheReplyIdPrefix, primary)
}

func (m *defaultReplyModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", replyRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultReplyModel) tableName() string {
	return m.table
}