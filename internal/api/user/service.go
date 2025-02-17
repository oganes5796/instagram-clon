package user

import (
	"github.com/oganes5796/instagram-clon/internal/service"
	desc "github.com/oganes5796/instagram-clon/pkg/user_v1"
)

type Implementation struct {
	desc.UnimplementedAuthV1Server
	userService service.UserService
}

func NewImplementation(userService service.UserService) *Implementation {
	return &Implementation{
		userService: userService,
	}
}
