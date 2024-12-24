package server

import (
	"context"
	"github.com/ac-dcz/lightRW/apps/genid/rpc/internal/config"
	"github.com/ac-dcz/lightRW/apps/genid/rpc/internal/svc"
	"github.com/ac-dcz/lightRW/apps/genid/rpc/pb"
	"github.com/zeromicro/go-zero/core/conf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"net"
	"testing"
)

var listener *bufconn.Listener

func init() {
	c := &config.Config{}
	if err := conf.Load("../../etc/genid.yaml", c); err != nil {
		panic(err)
	}
	svcCtx := svc.NewServiceContext(*c)
	listener = bufconn.Listen(1024 * 1024)
	server := grpc.NewServer()
	pb.RegisterGenIdServer(server, NewGenIdServer(svcCtx))
	go func() {
		if err := server.Serve(listener); err != nil {
			panic(err)
		}
	}()
}

func TestGenIdServer_GetId(t *testing.T) {
	conn, err := grpc.NewClient(
		"127.0.0.1:8080",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithContextDialer(func(ctx context.Context, addr string) (net.Conn, error) {
			return listener.DialContext(ctx)
		}),
	)
	if err != nil {
		t.Fatal(err)
	}
	client := pb.NewGenIdClient(conn)
	if resp, err := client.GetId(context.Background(), &pb.GetIdReq{}); err != nil {
		t.Fatal(err)
	} else {
		t.Logf("resp: %v", resp.Id)
	}
}
