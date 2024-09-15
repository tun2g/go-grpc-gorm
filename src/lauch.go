package src

import (
	authController "app/src/api/auth"
	userRepository "app/src/api/user/repositories/impl"
	"app/src/shared/jwt"
	"app/src/shared/utils"
	"time"

	authPb "app/proto/auth"
	authService "app/src/api/auth/services/impl"

	"golang.org/x/crypto/bcrypt"
)

func (s *Server) RegisterService() {

	var jwtAccessTokenManager = jwt.NewJWTManager(
		s.config.JwtAccessTokenSecret,
		time.Duration(s.config.JwtAccessTokenExpirationTime),
	)

	var jwtRefreshTokenManager = jwt.NewJWTManager(
		s.config.JwtRefreshTokenSecret,
		time.Duration(s.config.JwtRefreshTokenExpirationTime),
	)

	var bcrypt = utils.NewBcryptEncoder(bcrypt.DefaultCost)

	// Repositories
	userRepo := userRepository.NewUserRepository(s.db)

	// Service
	var authSrv = authService.NewAuthService(
		&userRepo,
		&jwtAccessTokenManager,
		&jwtRefreshTokenManager,
		&bcrypt,
	)

	// Controller
	var authCtl = authController.NewAuthController(authSrv)

	// Register ProtoBuf
	authPb.RegisterAuthControllerServer(s.grpc, authCtl)
}
