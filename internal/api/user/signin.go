package user

import (
	"context"

	desc "github.com/oganes5796/instagram-clon/pkg/user_v1"
)

func (i *Implementation) Signin(ctx context.Context, req *desc.SigninRequest) (*desc.SigninResponse, error) {
	_, err := i.userService.Signin(ctx, req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	return &desc.SigninResponse{
		RefreshToken: "work",
	}, nil
}
