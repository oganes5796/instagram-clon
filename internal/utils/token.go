package utils

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/oganes5796/instagram-clon/internal/domain"
)

func GenerateToken(info domain.UserInfo, secretKey []byte, duration time.Duration) (string, error) {
	claims := domain.UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(duration).Unix(),
		},
		Username: info.Username,
		Password: info.Password,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secretKey)
}

func VerifyToken(tokenString string, secretKey []byte) (*domain.UserClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&domain.UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, errors.New("unexpected token signing method")
			}
			return secretKey, nil
		},
	)

	if err != nil {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(*domain.UserClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}
