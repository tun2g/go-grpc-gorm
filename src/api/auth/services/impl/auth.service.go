package auth

// import (
// 	"context"

// 	"golang.org/x/crypto/bcrypt"

// 	auth "app/src/api/auth/dtos"
// 	authDto "app/src/api/auth/dtos"
// 	authService "app/src/api/auth/services"
// 	userModel "app/src/api/user/models"
// 	userRepository "app/src/api/user/repositories"
// 	userRole "app/src/shared/auth/constants"
// 	jwt "app/src/shared/jwt"
// 	utils "app/src/shared/utils"
// )

// type AuthService struct {
// 	userRepository         userRepository.IUserRepository
// 	jwtAccessTokenManager  *jwt.JWTManager
// 	jwtRefreshTokenManager *jwt.JWTManager
// 	bcrypt                 *utils.BcryptEncoder
// }

// func NewAuthService(
// 	userRepository *userRepository.IUserRepository,
// 	jwtAccessTokenManager *jwt.JWTManager,
// 	jwtRefreshTokenManager *jwt.JWTManager,
// 	bcrypt *utils.BcryptEncoder,
// ) authService.IAuthService {
// 	return &AuthService{
// 		userRepository:         *userRepository,
// 		jwtAccessTokenManager:  jwtAccessTokenManager,
// 		jwtRefreshTokenManager: jwtRefreshTokenManager,
// 		bcrypt:                 bcrypt,
// 	}
// }

// func (srv *AuthService) Login(
// 	dto *authDto.LoginReqDto,
// 	ctx *context.Context,
// ) (*userModel.User, *authDto.TokenResDto, error) {
// 	var err error

// 	user, err := srv.userRepository.FindOneBy(&userModel.User{Email: dto.Email})
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	if user == nil {
// 		err = exception.NewBadRequestException(
// 			ctx.GetRequestId(),
// 			[]exception.ErrorDetail{{
// 				Issue: "Email or password is invalid",
// 			}},
// 		)
// 		return nil, nil, err
// 	}

// 	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password))

// 	if err != nil {
// 		err = exception.NewBadRequestException(
// 			ctx.GetRequestId(),
// 			[]exception.ErrorDetail{{
// 				Issue: "Email or password is invalid",
// 			}},
// 		)
// 		return nil, nil, err
// 	}

// 	accessToken, _, err := srv.jwtAccessTokenManager.CreateToken(user)
// 	refreshToken, _, err := srv.jwtRefreshTokenManager.CreateToken(user)
// 	tokens := auth.TokenResDto{
// 		AccessToken:  accessToken,
// 		RefreshToken: refreshToken,
// 	}

// 	return user, &tokens, nil
// }

// func (srv *AuthService) Register(
// 	dto *authDto.RegisterReqDto,
// 	ctx *httpContext.CustomContext,
// ) (*userModel.User, *authDto.TokenResDto, error) {
// 	var err error

// 	user, err := srv.userRepository.FindOneBy(&userModel.User{
// 		Email: dto.Email,
// 	})

// 	if user != nil {
// 		err = exception.NewBadRequestException(
// 			ctx.GetRequestId(),
// 			[]exception.ErrorDetail{{
// 				Issue:   "Email is already in use",
// 				Field:   "email",
// 				IssueId: "exists_email",
// 			}},
// 		)
// 		return nil, nil, err
// 	}

// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	hashedPassword, err := srv.bcrypt.Encrypt(dto.Password)

// 	user, err = srv.userRepository.Create(&userModel.User{
// 		Email:    dto.Email,
// 		Password: hashedPassword,
// 		FullName: dto.FullName,
// 		Role:     userRole.RoleUser.String(),
// 	})

// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	accessToken, _, err := srv.jwtAccessTokenManager.CreateToken(user)
// 	refreshToken, _, err := srv.jwtRefreshTokenManager.CreateToken(user)
// 	tokens := auth.TokenResDto{
// 		AccessToken:  accessToken,
// 		RefreshToken: refreshToken,
// 	}
// 	return user, &tokens, nil
// }

// func (srv *AuthService) RefreshToken(ctx *httpContext.CustomContext) (*authDto.TokenResDto, error) {
// 	user := ctx.GetUser()

// 	if user == nil {
// 		err := exception.NewUnauthorizedException(ctx.GetRequestId())
// 		return nil, err
// 	}

// 	_user, err := srv.userRepository.FindOneBy(&userModel.User{Email: user.Email})

// 	if user == nil {
// 		err := exception.NewUnauthorizedException(ctx.GetRequestId())
// 		return nil, err
// 	}

// 	if err != nil {
// 		err := exception.NewUnauthorizedException(ctx.GetRequestId())
// 		return nil, err
// 	}

// 	accessToken, _, err := srv.jwtAccessTokenManager.CreateToken(_user)
// 	refreshToken, _, err := srv.jwtRefreshTokenManager.CreateToken(_user)
// 	tokens := auth.TokenResDto{
// 		AccessToken:  accessToken,
// 		RefreshToken: refreshToken,
// 	}

// 	return &tokens, nil
// }
