package user

import (
	"context"

	"github.com/oganes5796/instagram-clon/internal/domain"
)

func (s *serv) Register(ctx context.Context, user *domain.UserInfo) (int64, error) {
	user.Password = generatePasswordHash(user.Password)
	id, err := s.userRepository.Register(ctx, user)
	if err != nil {
		return 0, err
	}

	return id, nil
}
