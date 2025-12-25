package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserPayload struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Role string `json:"role"`
}

func GenerateToken(secret string, userPayload UserPayload) (string, error) {
	claims := Claims{
		UserID: userPayload.UserID,
		Email:  userPayload.Email,
		Role: userPayload.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenValue, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", errors.New("error signing token")
	}
	return tokenValue, nil
}

func VerifyToken(tokenStr, secret string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(
		tokenStr,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		},
	)

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}