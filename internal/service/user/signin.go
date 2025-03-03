package user

import (
	"context"
	"errors"
	"time"

	"github.com/oganes5796/instagram-clon/internal/domain"
	"github.com/oganes5796/instagram-clon/internal/utils"
)

const (
	refreshTokenSecretKey = "W4/X+LLjehdxptt4YgGFCvMpq5ewptpZZYRHY6A72g0="
	accessTokenSecretKey  = "VqvguGiffXILza1f44TWXowDT4zwf03dtXmqWW4SYyE="

	refreshTokenExpiration = 60 * time.Minute
	accessTokenExpiration  = 5 * time.Minute
)

func (s *serv) GenerateToken(ctx context.Context, username, password string) (string, error) {
	// Лезем в базу или кэш за данными пользователя
	// Сверяем хэши пароля

	refreshToken, err := utils.GenerateToken(domain.UserInfo{
		Username: username,
		Password: password,
	},
		[]byte(refreshTokenSecretKey),
		refreshTokenExpiration,
	)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return refreshToken, nil
}
