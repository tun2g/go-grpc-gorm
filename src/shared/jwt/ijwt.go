package jwt

import (
	"app/src/api/user/models"
)

type Manager interface {
	CreateToken(user *user.User) (string, *JwtPayload, error)

	VerifyToken(token string) (*JwtPayload, error)
}
