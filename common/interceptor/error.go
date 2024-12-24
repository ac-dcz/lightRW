package interceptor

import (
	"context"
	"github.com/ac-dcz/lightRW/common/errors"
	"google.golang.org/grpc"
)

func ErrorForClient() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		err := invoker(ctx, method, req, reply, cc, opts...)
		if err != nil {
			return errors.FromError(err)
		}
		return nil
	}
}

func ErrorForServer() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		resp, err := handler(ctx, req)
		if err != nil {
			return resp, errors.ToStatus(err).Err()
		}
		return resp, nil
	}
}
