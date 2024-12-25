package interceptor

import (
	"context"
	"github.com/ac-dcz/lightRW/common/codes"
	"github.com/ac-dcz/lightRW/common/errors"
	"github.com/ac-dcz/lightRW/common/jwt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func AuthForServer(opt *jwt.Option) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

		if md, ok := metadata.FromIncomingContext(ctx); !ok {
			return nil, errors.ToStatus(errors.New(codes.NotFoundMetaData, "not found metadata")).Err()
		} else {
			if token, ok := md["token"]; ok && len(token) > 0 { //仅对提供token的做检查
				if claims, err := jwt.VerifyToken(opt, token[0]); err != nil {
					return nil, errors.ToStatus(errors.New(codes.InvalidToken, err.Error())).Err()
				} else {
					if uid, ok := claims["uid"]; ok {
						ctx = context.WithValue(ctx, "uid", uid)
					} else {
						return nil, errors.ToStatus(errors.New(codes.InvalidTokenPayLoad, "not found uid")).Err()
					}
					if level, ok := claims["level"]; ok {
						ctx = context.WithValue(ctx, "level", level)
					} else {
						return nil, errors.ToStatus(errors.New(codes.InvalidTokenPayLoad, "not found level")).Err()
					}
				}
			}
		}

		resp, err := handler(ctx, req)
		if err != nil {
			return resp, errors.ToStatus(err).Err()
		}
		return resp, nil
	}
}
