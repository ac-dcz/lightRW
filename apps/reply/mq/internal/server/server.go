package server

import (
	"context"
	"github.com/ac-dcz/lightRW/apps/review/mq/internal/logic"
	"github.com/ac-dcz/lightRW/apps/review/mq/internal/svc"
)

type Server struct {
	svcCtx *svc.ServiceContext
}

func NewServer(svcCtx *svc.ServiceContext) *Server {
	return &Server{
		svcCtx: svcCtx,
	}
}

func (s *Server) Run(ctx context.Context) error {
	l := logic.NewMqLogic(ctx, s.svcCtx)
	return s.svcCtx.Reader.Run(ctx, l.Handle)
}

func (s *Server) Stop() error {
	return s.svcCtx.Reader.Close()
}
