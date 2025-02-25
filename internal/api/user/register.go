package user

import (
	"context"

	"github.com/oganes5796/instagram-clon/internal/converter"
	"github.com/oganes5796/instagram-clon/internal/logger"
	desc "github.com/oganes5796/instagram-clon/pkg/user_v1"
	"go.uber.org/zap"
)

func (i *Implementation) Register(ctx context.Context, req *desc.RegisterRequest) (*desc.RegisterResponse, error) {
	id, err := i.userService.Register(ctx, converter.ToUserInfoFromDesc(req.GetInfo()))
	if err != nil {
		return nil, err
	}

	logger.Info("Inserted user with id:", zap.Int64("id", id))

	return &desc.RegisterResponse{
		UserId: id,
	}, nil
}
