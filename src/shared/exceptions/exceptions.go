package exceptions

import (
	pb "app/proto/exceptions"
	"app/src/lib/logger"
	constants "app/src/shared/constants"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type GrpcError struct {
	Code        codes.Code
	ErrorDetail *pb.ErrorDetail
}

func (e GrpcError) Error() string {
	return "GRPC Server error"
}

func ThrowGrpcError(code codes.Code, errorDetail *pb.ErrorDetail) GrpcError {
	return GrpcError{
		Code:        code,
		ErrorDetail: errorDetail,
	}
}

var log = logger.NewLogger("Exceptions")

func HandleGrpcError(err error, ctx *context.Context) error {
	md, _ := metadata.FromIncomingContext(*ctx)
	requestIds := md.Get(constants.RequestIDKey)
	requestId := "unknown"
	if len(requestIds) > 0 {
		requestId = requestIds[0]
	}

	var errorResponse pb.GRPCErrorResponse
	var statusErr *status.Status

	if customErr, ok := err.(GrpcError); ok {
		errorResponse.Code = int32(customErr.Code)
		statusErr = status.New(codes.Code(errorResponse.Code), "GRPC Server error")
		errorResponse.ErrorDetail = customErr.ErrorDetail

	} else {
		statusErr = status.New(codes.Internal, "GRPC Server error")
	}

	errorResponse.ErrorDetail.RequestId = requestId

	log.Error(errorResponse.ErrorDetail)
	statusErr, _ = statusErr.WithDetails(errorResponse.ErrorDetail)

	return statusErr.Err()
}
