package user

import (
	"context"

	"github.com/oganes5796/instagram-clon/internal/logger"
	desc "github.com/oganes5796/instagram-clon/pkg/user_v1"
	"go.uber.org/zap"
)

func (i *Implementation) Signin(ctx context.Context, req *desc.SigninRequest) (*desc.SigninResponse, error) {

	refreshToken, err := i.userService.GenerateToken(ctx, req.GetUsername(), req.GetPassword())
	if err != nil {
		return nil, err
	}

	logger.Info("User signin attempt", zap.String("username", req.GetUsername()))

	return &desc.SigninResponse{
		RefreshToken: refreshToken,
	}, nil
}
