package auth

import (
	"context"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"

	authDto "app/src/api/auth/dtos"
	authService "app/src/api/auth/services"
	userModel "app/src/api/user/models"
	constants "app/src/api/auth/constants"
	userRepository "app/src/api/user/repositories"
	exceptions "app/src/shared/exceptions"
	jwt "app/src/shared/jwt"
	"app/src/shared/utils"
)

type AuthService struct {
	userRepository         userRepository.IUserRepository
	jwtAccessTokenManager  *jwt.JWTManager
	jwtRefreshTokenManager *jwt.JWTManager
	bcrypt                 *utils.BcryptEncoder
}

func NewAuthService(
	userRepository *userRepository.IUserRepository,
	jwtAccessTokenManager *jwt.JWTManager,
	jwtRefreshTokenManager *jwt.JWTManager,
	bcrypt *utils.BcryptEncoder,
) authService.IAuthService {
	return &AuthService{
		userRepository:         *userRepository,
		jwtAccessTokenManager:  jwtAccessTokenManager,
		jwtRefreshTokenManager: jwtRefreshTokenManager,
		bcrypt:                 bcrypt,
	}
}

func (srv *AuthService) Login(
	dto *authDto.LoginParamsDto,
	ctx *context.Context,
) (*userModel.User, *authDto.TokenResDto, error) {
	var err error

	user, err := srv.userRepository.FindOneBy(&userModel.User{Email: dto.Email})
	if err != nil {
		return nil, nil, err
	}

	if user == nil {
		err = exceptions.ThrowGrpcError(
			codes.InvalidArgument, 
			"Cannot find user",
			"",	
		)
		return nil, nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password))

	if err != nil {
		err = exceptions.ThrowGrpcError(
			codes.InvalidArgument,
			"Invalid username or password",
			"",
		)
		return nil, nil, err
	}

	accessToken, _, _ := srv.jwtAccessTokenManager.CreateToken(user)
	refreshToken, _, _ := srv.jwtRefreshTokenManager.CreateToken(user)
	tokens := authDto.TokenResDto{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return user, &tokens, nil
}

func (srv *AuthService) Register(
	dto *authDto.RegisterParamsDto,
	ctx *context.Context,
) (*userModel.User, *authDto.TokenResDto, error) {
	var err error

	user, err := srv.userRepository.FindOneBy(&userModel.User{
		Email: dto.Email,
	})

	if user != nil {
		err = exceptions.ThrowGrpcError(
			codes.AlreadyExists,
			"Email is already in use",
			constants.ERR_EXISTED_EMAIL,
		)
		return nil, nil, err
	}

	if err != nil {
		return nil, nil, err
	}

	hashedPassword, _ := srv.bcrypt.Encrypt(dto.Password)

	user, err = srv.userRepository.Create(&userModel.User{
		Email:    dto.Email,
		Password: hashedPassword,
		FullName: dto.FullName,
		Role:     constants.RoleUser.String(),
	})

	if err != nil {
		return nil, nil, err
	}

	accessToken, _, _ := srv.jwtAccessTokenManager.CreateToken(user)
	refreshToken, _, _ := srv.jwtRefreshTokenManager.CreateToken(user)

	tokens := authDto.TokenResDto{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return user, &tokens, nil
}