package service

import (
	"context"

	"github.com/oganes5796/instagram-clon/internal/domain"
)

type UserService interface {
	Register(ctx context.Context, user *domain.UserInfo) (int64, error)
}
