package converter

import (
	"github.com/oganes5796/instagram-clon/internal/domain"
	"github.com/oganes5796/instagram-clon/internal/repository/user/model"
)

func TouserFromRepo(user *domain.User) *model.User {
	return &model.User{
		ID:   user.ID,
		Info: TouserInfoFromRepo(user.Info),
	}
}

func TouserInfoFromRepo(info domain.UserInfo) model.UserInfo {
	return model.UserInfo{
		Email:    info.Email,
		Username: info.Username,
		Password: info.Password,
	}
}
