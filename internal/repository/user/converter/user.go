package converter

import (
	"github.com/oganes5796/instagram-clon/internal/domain"
	"github.com/oganes5796/instagram-clon/internal/repository/user/model"
)

func TouserFromRepo(user *model.User) *domain.User {
	return &domain.User{
		ID:   user.ID,
		Info: TouserInfoFromRepo(user.Info),
	}
}

func TouserInfoFromRepo(info model.UserInfo) domain.UserInfo {
	return domain.UserInfo{
		Email:    info.Email,
		Username: info.Username,
		Password: info.Password,
	}
}
