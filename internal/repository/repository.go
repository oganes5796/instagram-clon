package repository

import (
	"context"

	"github.com/oganes5796/instagram-clon/internal/domain"
)

type UserRepository interface {
	Register(ctx context.Context, user *domain.UserInfo) (int64, error)
	// GetUser(ctx context.Context, username, password string) (*domain.User, error)
}
