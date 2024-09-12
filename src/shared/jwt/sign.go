package jwt

import (
	"errors"
	"app/src/api/user/models"
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

type JWTManager struct {
	secretKey      string
	expirationTime time.Duration
}

func NewJWTManager(secretKey string, expirationTime time.Duration) JWTManager {
	jwtManager := &JWTManager{
		secretKey:      secretKey,
		expirationTime: expirationTime,
	}
	return *jwtManager
}

func (_jwt *JWTManager) CreateToken(user *user.User) (string, *JwtPayload, error) {
	payload := NewJwtPayload(user, _jwt.expirationTime)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := jwtToken.SignedString([]byte(_jwt.secretKey))
	return token, payload, err
}

func (_jwt *JWTManager) VerifyToken(token string) (*JwtPayload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(_jwt.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &JwtPayload{}, keyFunc)
	if err != nil {
		_err, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(_err.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*JwtPayload)
	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}
