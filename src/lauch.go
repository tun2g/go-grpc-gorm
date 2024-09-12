package src

import (
	userRepository "app/src/api/user/repositories/impl"
	authController "app/src/api/auth"

	authPb "app/proto/auth"
)

func (s *Server) RegisterService(){

	// Repositories
	userRepo := userRepository.NewUserRepository(s.db)

	// Service
	
	
	// Controller
	var authCtl = authController.NewAuthController(
		&userRepo,
	)


	// Register ProtoBuf
	authPb.RegisterAuthControllerServer(s.grpc, authCtl)
}