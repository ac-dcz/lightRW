// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: order.proto

package order

import (
	"context"

	"github.com/ac-dcz/lightRW/apps/order/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreateOrderReq  = pb.CreateOrderReq
	CreateOrderResp = pb.CreateOrderResp
	OrderEntry      = pb.OrderEntry
	OrderInfoReq    = pb.OrderInfoReq
	OrderInfoResp   = pb.OrderInfoResp
	OrderStatusReq  = pb.OrderStatusReq
	OrderStatusResp = pb.OrderStatusResp
	PayOrderReq     = pb.PayOrderReq
	PayOrderResp    = pb.PayOrderResp

	Order interface {
		CreateOrder(ctx context.Context, in *CreateOrderReq, opts ...grpc.CallOption) (*CreateOrderResp, error)
		PayOrder(ctx context.Context, in *PayOrderReq, opts ...grpc.CallOption) (*PayOrderResp, error)
		OrderInfo(ctx context.Context, in *OrderInfoReq, opts ...grpc.CallOption) (*OrderInfoResp, error)
		OrderStatus(ctx context.Context, in *OrderStatusReq, opts ...grpc.CallOption) (*OrderStatusResp, error)
	}

	defaultOrder struct {
		cli zrpc.Client
	}
)

func NewOrder(cli zrpc.Client) Order {
	return &defaultOrder{
		cli: cli,
	}
}

func (m *defaultOrder) CreateOrder(ctx context.Context, in *CreateOrderReq, opts ...grpc.CallOption) (*CreateOrderResp, error) {
	client := pb.NewOrderClient(m.cli.Conn())
	return client.CreateOrder(ctx, in, opts...)
}

func (m *defaultOrder) PayOrder(ctx context.Context, in *PayOrderReq, opts ...grpc.CallOption) (*PayOrderResp, error) {
	client := pb.NewOrderClient(m.cli.Conn())
	return client.PayOrder(ctx, in, opts...)
}

func (m *defaultOrder) OrderInfo(ctx context.Context, in *OrderInfoReq, opts ...grpc.CallOption) (*OrderInfoResp, error) {
	client := pb.NewOrderClient(m.cli.Conn())
	return client.OrderInfo(ctx, in, opts...)
}

func (m *defaultOrder) OrderStatus(ctx context.Context, in *OrderStatusReq, opts ...grpc.CallOption) (*OrderStatusResp, error) {
	client := pb.NewOrderClient(m.cli.Conn())
	return client.OrderStatus(ctx, in, opts...)
}
