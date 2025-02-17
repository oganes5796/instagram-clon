package converter

import (
	"github.com/oganes5796/instagram-clon/internal/domain"
	desc "github.com/oganes5796/instagram-clon/pkg/user_v1"
)

func ToUserInfoFromDesc(info *desc.UserInfo) *domain.UserInfo {
	return &domain.UserInfo{
		Email:    info.Email,
		Username: info.Username,
		Password: info.Password,
	}
}
