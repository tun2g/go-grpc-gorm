package src

import (
	"app/src/config"
	"app/src/shared/interceptors"
	"app/src/lib/logger"
	"context"

	"net"

	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type Server struct {
	grpc   *grpc.Server
	db     *gorm.DB
	ctx    *context.Context
	config *config.Config
}

var log = logger.NewLogger("Server")

func NewServer(dbConnection *gorm.DB, config *config.Config) (*Server, error) {
	ctx := context.Background()	

	grpcSv := grpc.NewServer(
		grpc.UnaryInterceptor(interceptors.GlobalExceptionInterceptor()),
	)

	if config.ReflectionEnabled {
		reflection.Register(grpcSv)
		log.Printf("Server is enabled reflection! Using `grpcui -plaintext localhost:%v` to open GPRC UI", config.AppPort)
	}

	server := &Server{
		db:     dbConnection,
		ctx:    &ctx,
		config: config,
		grpc:   grpcSv,
	}

	return server, nil
}

func (s *Server) Serve(lis net.Listener) error {
	return s.grpc.Serve(lis)
}
