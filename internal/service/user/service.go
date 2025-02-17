package user

import (
	"github.com/oganes5796/instagram-clon/internal/repository"
	"github.com/oganes5796/instagram-clon/internal/service"
)

type serv struct {
	userRepository repository.UserRepository
}

func NewService(
	userRepository repository.UserRepository,
) service.UserService {
	return &serv{
		userRepository: userRepository,
	}
}

func NewMockService(deps ...interface{}) service.UserService {
	srv := serv{}

	for _, v := range deps {
		switch s := v.(type) {
		case repository.UserRepository:
			srv.userRepository = s
		}
	}

	return &srv
}
