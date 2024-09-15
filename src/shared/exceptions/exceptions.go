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
	Message string
	Code    codes.Code
	IssueId string
}

func (e GrpcError) Error() string {
	return e.Message
}

func ThrowGrpcError(code codes.Code, message string, issueId string) GrpcError {
	return GrpcError{
		Message: message,
		Code:    code,
		IssueId: issueId,
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
	log.Error(err)

	var errorResponse pb.GRPCErrorResponse
	var statusErr *status.Status
	errorResponse.RequestId = requestId

	if customErr, ok := err.(GrpcError); ok {
		errorResponse.Message = customErr.Message
		errorResponse.IssueId = customErr.IssueId
		errorResponse.Code = int32(customErr.Code)
		statusErr = status.New(codes.Code(errorResponse.Code), errorResponse.Message)
	} else {
		errorResponse.Message = err.Error()
		statusErr = status.New(codes.Internal, "Internal Server Error")
	}

	statusErr, _ = statusErr.WithDetails(&errorResponse)
	return statusErr.Err()
}
