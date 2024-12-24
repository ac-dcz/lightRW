package server

import (
	"context"
	"github.com/ac-dcz/lightRW/apps/code/rpc/internal/config"
	"github.com/ac-dcz/lightRW/apps/code/rpc/internal/svc"
	"github.com/ac-dcz/lightRW/apps/code/rpc/pb"
	"github.com/zeromicro/go-zero/core/conf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"net"
	"testing"
)

var listener *bufconn.Listener

func init() {
	listener = bufconn.Listen(1024 * 1024)
	c := config.Config{}
	if err := conf.Load("../../etc/code.yaml", &c); err != nil {
		panic(err)
	}
	svcCtx := svc.NewServiceContext(c)
	s := grpc.NewServer()
	pb.RegisterVerifyCodeServer(s, NewVerifyCodeServer(svcCtx))
	go func() {
		if err := s.Serve(listener); err != nil {
			panic(err)
		}
	}()
}

func TestVerifyCodeServer_GenCode(t *testing.T) {
	conn, err := grpc.NewClient(
		"127.0.0.1:8080",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithContextDialer(func(ctx context.Context, s string) (net.Conn, error) {
			return listener.DialContext(ctx)
		}),
	)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := pb.NewVerifyCodeClient(conn)
	if resp, err := c.GenCode(context.Background(), &pb.GenCodeReq{
		Tel: "12771695264",
	}); err != nil {
		t.Fatal(err)
	} else {
		t.Logf("resp: %v", resp.Code)
	}
}

func TestVerifyCodeServer_VerifyCode(t *testing.T) {
	conn, err := grpc.NewClient(
		"127.0.0.1:8080",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithContextDialer(func(ctx context.Context, s string) (net.Conn, error) {
			return listener.DialContext(ctx)
		}),
	)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := pb.NewVerifyCodeClient(conn)
	if resp, err := c.VerifyCode(context.Background(), &pb.VerifyCodeReq{
		Tel:  "12771695264",
		Code: "638173",
	}); err != nil {
		t.Fatal(err)
	} else {
		t.Logf("resp: %v", resp.Success)
	}
}
