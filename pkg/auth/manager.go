package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JWTManager struct {
	TokenTTL  time.Duration
	SecretKey string
}

func NewJWTManager(tokenTTL time.Duration, secretKey string) *JWTManager {
	return &JWTManager{
		TokenTTL:  tokenTTL,
		SecretKey: secretKey,
	}
}

func (m *JWTManager) NewToken(userID uint, role string) (string, error) {
	claims := jwt.MapClaims{
		"exp":  time.Now().Add(m.TokenTTL).Unix(),
		"role": role,
		"sub":  userID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(m.SecretKey))
}
