package interceptors

import (
	"app/src/lib/logger"
	"app/src/shared/constants"
	"context"
	"encoding/json"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var log = logger.NewLogger("LogInterceptor")

func LogInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		start := time.Now()

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			md = metadata.New(nil)
		}
		requestIDs := md.Get(constants.RequestIDKey)
		requestID := ""
		if len(requestIDs) > 0 {
			requestID = requestIDs[0]
		}

		reqJSON, _ := json.Marshal(req)
		log.WithFields(logrus.Fields{
			"method":     info.FullMethod,
			"request_id": requestID,
			"params":     string(reqJSON),
			"start_time": start,
		}).Info("--------Incoming request")

		resp, err := handler(ctx, req)

		duration := time.Since(start)

		if err != nil {
			st, _ := status.FromError(err)
			log.WithFields(logrus.Fields{
				"method":     info.FullMethod,
				"request_id": requestID,
				"error":      err,
				"status":     st.Code().String(),
				"duration":   duration,
			}).Error("--------Request failed")
		} else {
			log.WithFields(logrus.Fields{
				"method":     info.FullMethod,
				"request_id": requestID,
				"duration":   duration,
			}).Info("--------Request completed")
		}

		return resp, err
	}
}
