package auth

import (
	userRepository "app/src/api/user/repositories"
	authPb "app/proto/auth"
	"context"
)

type AuthController struct {
	authPb.UnimplementedAuthControllerServer

	userRepository *userRepository.IUserRepository
}

func NewAuthController(userRepository *userRepository.IUserRepository) AuthController {
	return AuthController{
		userRepository: userRepository,
	}
}

func (srv AuthController) SignUp(ctx context.Context, req *authPb.SignUpRequest) (*authPb.SignUpResponse, error){
	return nil, nil
}

func (srv AuthController) SignIn(ctx context.Context, req *authPb.SignInRequest) (*authPb.SignInResponse, error){
	return nil, nil
}

