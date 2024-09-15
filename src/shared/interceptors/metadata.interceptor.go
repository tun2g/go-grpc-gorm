package interceptors

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	constants "app/src/shared/constants"
)

func MetadataInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			md = metadata.New(nil)
		}

		requestIDs := md.Get(constants.RequestIDKey)
		var requestID string
		if len(requestIDs) > 0 {
			requestID = requestIDs[0]
		} else {
			requestID = uuid.New().String()
		}

		ctx = metadata.AppendToOutgoingContext(ctx, constants.RequestIDKey, requestID)
		return handler(ctx, req)
	}
}
