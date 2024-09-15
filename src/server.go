package src

import (
	"app/src/config"
	"app/src/database"
	"app/src/lib/logger"
	"app/src/shared/interceptors"
	"context"
	"fmt"

	"net"

	"github.com/urfave/cli"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
)

type Server struct {
	grpc   *grpc.Server
	db     *gorm.DB
	ctx    *context.Context
	config *config.Config
}

var log = logger.NewLogger("Server")

func (s *Server) Serve(lis net.Listener) error {
	return s.grpc.Serve(lis)
}

func NewServer(dbConnection *gorm.DB, config *config.Config) (*Server, error) {
	ctx := context.Background()

	grpcSv := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptors.MetadataInterceptor(),
			interceptors.LogInterceptor(),
			interceptors.GlobalExceptionInterceptor(),
		),
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

func StartServer() cli.Command {
	cli := cli.Command{
		Name:  "server",
		Usage: "send example tasks ",
		Action: func(c *cli.Context) error {
			log := logger.NewLogger("Main")

			connection := database.InitDB()

			s, err := NewServer(connection, &config.AppConfiguration)

			if err != nil {
				log.Errorf("Error creating server: %v", err)
			}

			s.RegisterService()

			lis, err := net.Listen("tcp", fmt.Sprintf(":%v", config.AppConfiguration.AppPort))
			if err != nil {
				log.Fatalf("failed to listen: %v", err)
			}

			log.Println("Server is running on port", config.AppConfiguration.AppPort)
			if err := s.Serve(lis); err != nil {
				log.Fatalf("failed to serve: %v", err)
			}
			return nil
		},
	}
	return cli
}
