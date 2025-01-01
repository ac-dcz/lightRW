package model

import (
	"context"
	stderr "errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ReviewModel = (*customReviewModel)(nil)

type (
	// ReviewModel is an interface to be customized, add more methods here,
	// and implement the added methods in customReviewModel.
	ReviewModel interface {
		reviewModel
		FindManyByUid(ctx context.Context, uid uint64) ([]*Review, error)
		FindManyBySSku(ctx context.Context, storeId uint64, sku string) ([]*Review, error)
		UpdateStatus(ctx context.Context, reviewId uint64, status uint32, opReason string) error
	}

	customReviewModel struct {
		*defaultReviewModel
	}
)

// NewReviewModel returns a model for the database table.
func NewReviewModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ReviewModel {
	return &customReviewModel{
		defaultReviewModel: newReviewModel(conn, c, opts...),
	}
}

func (m *customReviewModel) FindManyByUid(ctx context.Context, uid uint64) ([]*Review, error) {
	reviews := make([]*Review, 0)
	query := fmt.Sprintf("select %s from %s where uid = ?", reviewRows, m.table)
	if err := m.QueryRowsNoCacheCtx(ctx, &reviews, query, uid); err != nil {
		if stderr.Is(err, sqlc.ErrNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return reviews, nil
}

func (m *customReviewModel) FindManyBySSku(ctx context.Context, storeId uint64, sku string) ([]*Review, error) {
	reviews := make([]*Review, 0)
	query := fmt.Sprintf("select %s from %s where store_id = ? and sku = ?", reviewRows, m.table)
	if err := m.QueryRowsNoCacheCtx(ctx, &reviews, query, storeId, sku); err != nil {
		if stderr.Is(err, sqlc.ErrNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return reviews, nil
}

func (m *customReviewModel) UpdateStatus(ctx context.Context, reviewId uint64, status uint32, opReason string) error {
	query := fmt.Sprintf("update %s set status = ? and op_reason = ? where review_id = ?", m.table)
	if _, err := m.CachedConn.ExecNoCacheCtx(ctx, query, status, opReason, reviewId); err != nil {
		return err
	} else {
		return nil
	}
}
