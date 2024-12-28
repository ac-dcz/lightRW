package model

import (
	"database/sql"
	stderr "errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"golang.org/x/net/context"
)

var _ ReplyModel = (*customReplyModel)(nil)

type (
	// ReplyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customReplyModel.
	ReplyModel interface {
		replyModel
		FindManyBySSku(ctx context.Context, storeId uint64, sku string) ([]*Reply, error)
		FindManyByMid(ctx context.Context, mid uint64) ([]*Reply, error)
		FindManyByReviewId(ctx context.Context, reviewId uint64) ([]*Reply, error)
	}

	customReplyModel struct {
		*defaultReplyModel
	}
)

// NewReplyModel returns a model for the database table.
func NewReplyModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ReplyModel {
	return &customReplyModel{
		defaultReplyModel: newReplyModel(conn, c, opts...),
	}
}

func (m *customReplyModel) FindManyBySSku(ctx context.Context, storeId uint64, sku string) ([]*Reply, error) {
	replyList := make([]*Reply, 0)
	query := fmt.Sprintf("select %s from %s where store_id = ? and sku = ?", replyRows, m.table)
	if err := m.QueryRowsNoCacheCtx(ctx, &replyList, query, storeId, sku); err != nil {
		if stderr.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return replyList, nil
}

func (m *customReplyModel) FindManyByMid(ctx context.Context, mid uint64) ([]*Reply, error) {
	replyList := make([]*Reply, 0)
	query := fmt.Sprintf("select %s from %s where mid = ?", replyRows, m.table)
	if err := m.QueryRowsNoCacheCtx(ctx, &replyList, query, mid); err != nil {
		if stderr.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return replyList, nil
}

func (m *customReplyModel) FindManyByReviewId(ctx context.Context, reviewId uint64) ([]*Reply, error) {
	replyList := make([]*Reply, 0)
	query := fmt.Sprintf("select %s from %s where reviewId = ?", replyRows, m.table)
	if err := m.QueryRowsNoCacheCtx(ctx, &replyList, query, reviewId); err != nil {
		if stderr.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return replyList, nil
}
