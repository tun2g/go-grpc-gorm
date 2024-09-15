package auth

import (
	dtos "app/src/api/auth/dtos"
	userModel "app/src/api/user/models"
	"context"
)

type IAuthService interface {
	Login(dto *dtos.LoginParamsDto, ctx *context.Context) (*userModel.User, *dtos.TokenResDto, error)
	Register(dto *dtos.RegisterParamsDto, ctx *context.Context) (*userModel.User, *dtos.TokenResDto, error)
}
