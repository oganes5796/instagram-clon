package converter

import (
	"github.com/oganes5796/instagram-clon/internal/domain"
	"github.com/oganes5796/instagram-clon/internal/repository/user/redis/model"
)

func ToUserFromRepo(user *model.User) *domain.User {
	return &domain.User{
		ID: user.ID,
		Info: domain.UserInfo{
			Email:    user.Email,
			Username: user.Username,
			Password: user.Password,
		},
	}
}
