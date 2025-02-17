package user

import (
	"context"

	"github.com/oganes5796/instagram-clon/internal/domain"
)

func (s *serv) Register(ctx context.Context, user *domain.UserInfo) (int64, error) {
	var id int64
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		id, errTx = s.userRepository.Register(ctx, user)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return id, nil
}
