package auth

import (
	authService "app/src/api/auth/services"
	authPb "app/proto/auth"
	authDto "app/src/api/auth/dtos"
	"context"
)

type AuthController struct {
	authPb.UnimplementedAuthControllerServer

	authService authService.IAuthService
}

func NewAuthController(authService authService.IAuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (srv AuthController) SignUp(ctx context.Context, req *authPb.SignUpRequest) (*authPb.SignUpResponse, error){
	user, _, err := srv.authService.Register(
		&authDto.RegisterParamsDto{
			Email: req.Email,
			Password: req.Password,
			FullName: req.FullName,
		},
		&ctx,
	)

	if err!= nil {
		return nil, err
	}

	return &authPb.SignUpResponse{
		Email: user.Email,
		UserId: user.Id,
	}, nil
}

func (srv AuthController) SignIn(ctx context.Context, req *authPb.SignInRequest) (*authPb.SignInResponse, error){
	user, tokens, err := srv.authService.Login(
		&authDto.LoginParamsDto{
			Email: req.Email,
			Password: req.Password,
		},
		&ctx,
	)

	if err!= nil {
		return nil, err
	}

	return &authPb.SignInResponse{
		Tokens: &authPb.Tokens{
			AccessToken: tokens.AccessToken,
			RefreshToken: tokens.RefreshToken,
		},
		Email: user.Email,
		UserId: user.Id,
	}, nil
}

