package user

import (
	"context"

	"github.com/oganes5796/instagram-clon/internal/domain"
)

func (s *serv) Signin(ctx context.Context, username, password string) (*domain.User, error) {
	user, err := s.userRepository.GetUser(ctx, username, password)
	if err != nil {
		return nil, err
	}

	return user, nil
}
