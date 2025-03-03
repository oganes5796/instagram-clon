package redis

import (
	"context"
	"strconv"

	"github.com/oganes5796/instagram-clon/internal/client/cache"
	"github.com/oganes5796/instagram-clon/internal/domain"
	"github.com/oganes5796/instagram-clon/internal/repository"
	"github.com/oganes5796/instagram-clon/internal/repository/user/redis/model"
)

type repo struct {
	cl cache.RedisClient
}

func NewRepository(cl cache.RedisClient) repository.UserRepository {
	return &repo{cl: cl}
}

func (r *repo) Register(ctx context.Context, user *domain.UserInfo) (int64, error) {
	id := int64(1)

	note := model.User{
		ID:       id,
		Email:    user.Email,
		Username: user.Username,
		Password: user.Password,
	}

	idStr := strconv.FormatInt(id, 10)
	err := r.cl.HashSet(ctx, idStr, note)
	if err != nil {
		return 0, err
	}

	return id, nil
}
