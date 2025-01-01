package model

import (
	"database/sql"
	stderr "errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"golang.org/x/net/context"
	"strconv"
)

var _ OrdersModel = (*customOrdersModel)(nil)

type (
	// OrdersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOrdersModel.
	OrdersModel interface {
		ordersModel
		InsertWithTx(ctx context.Context, orders ...*Orders) error
		FindOrdersByOrderIdUId(ctx context.Context, orderId, uId uint64) ([]*Orders, error)
		FindOneByOrderId(ctx context.Context, orderId uint64) (*Orders, error)
		UpdateOrdersForStatus(ctx context.Context, orderId, uId, status uint64) (int64, error)
	}

	customOrdersModel struct {
		*defaultOrdersModel
	}
)

// NewOrdersModel returns a model for the database table.
func NewOrdersModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) OrdersModel {
	return &customOrdersModel{
		defaultOrdersModel: newOrdersModel(conn, c, opts...),
	}
}

func (m *customOrdersModel) InsertWithTx(ctx context.Context, orders ...*Orders) error {
	gstoreTable, ok := ctx.Value("goods_store").(string)
	if !ok {
		return fmt.Errorf("not found goods_store table name")
	}

	query1 := fmt.Sprintf("update %s set stock = stock - ? where store_id = ? and sku = ?", gstoreTable)
	query2 := fmt.Sprintf("insert into %s(%s) values(?, ?, ?, ?, ?, ?, ?)", m.table, ordersRowsExpectAutoSet)

	err := m.CachedConn.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		for _, data := range orders {
			//减库存
			if _, err := session.ExecCtx(ctx, query1, data.Num, data.StoreId, data.Sku); err != nil {
				return err
			}
			// 插入新数据 不会破会缓存一致性
			if _, err := session.ExecCtx(ctx, query2, data.OrderId, data.Uid,
				data.StoreId, data.Sku, data.Num, data.Price, data.Status); err != nil {
				return err
			}
		}
		return nil
	})
	return err
}

func (m *customOrdersModel) FindOneByOrderId(ctx context.Context, orderId uint64) (*Orders, error) {
	var data Orders
	key := cacheOrdersOrderIdPrefix + strconv.FormatUint(orderId, 10)
	query := fmt.Sprintf("select %s from %s where order_id = ?", ordersRows, m.table)
	if err := m.QueryRowCtx(ctx, &data, key, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		return conn.QueryRowCtx(ctx, v, query, orderId)
	}); err != nil {
		if stderr.Is(err, sqlc.ErrNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &data, nil
}

func (m *customOrdersModel) UpdateOrdersForStatus(ctx context.Context, orderId, uId, status uint64) (int64, error) {
	key := cacheOrdersOrderIdPrefix + strconv.FormatUint(orderId, 10)
	query := fmt.Sprintf("update %s set status=? where order_id = ? and uid = ? and status = 0", m.table)
	r, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
		return conn.ExecCtx(ctx, query, status, orderId, uId)
	}, key)
	if err != nil {
		return 0, err
	}
	return r.RowsAffected()
}

func (m *customOrdersModel) FindOrdersByOrderIdUId(ctx context.Context, orderId, uId uint64) ([]*Orders, error) {
	var datas []*Orders
	query := fmt.Sprintf("select %s from %s where order_id = ? and uid = ?", ordersRows, m.table)
	if err := m.QueryRowsNoCacheCtx(ctx, &datas, query, orderId, uId); err != nil {
		if stderr.Is(err, sqlc.ErrNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return datas, nil
}
