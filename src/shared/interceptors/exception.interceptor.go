package interceptors

import (
	"context"

	exceptions "app/src/shared/exceptions"

	"google.golang.org/grpc"
)

func GlobalExceptionInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		resp, err := handler(ctx, req)
		if err != nil {
			grpcErr := (err)
			return nil, exceptions.HandleGrpcError(grpcErr, &ctx)
		}
		return resp, nil
	}
}
