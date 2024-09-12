package interceptors

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GlobalExceptionInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		resp, err := handler(ctx, req)
		if err != nil {
			st, ok := status.FromError(err)
			if !ok {
				st = status.New(codes.Internal, "An internal error occurred")
			}
			return nil, st.Err()
		}
		return resp, nil
	}
}
