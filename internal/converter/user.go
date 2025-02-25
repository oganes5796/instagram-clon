package converter

import (
	"github.com/oganes5796/instagram-clon/internal/domain"
	desc "github.com/oganes5796/instagram-clon/pkg/user_v1"
)

func ToUserFromService(note *domain.User) *desc.User {

	return &desc.User{
		Id:   note.ID,
		Info: ToUserInfoFromService(note.Info),
	}
}

func ToUserInfoFromService(info domain.UserInfo) *desc.UserInfo {
	return &desc.UserInfo{
		Email:    info.Email,
		Username: info.Username,
		Password: info.Password,
	}
}

func ToUserInfoFromDesc(info *desc.UserInfo) *domain.UserInfo {
	return &domain.UserInfo{
		Email:    info.Email,
		Username: info.Username,
		Password: info.Password,
	}
}
