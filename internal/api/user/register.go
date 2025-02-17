package user

import (
	"context"
	"log"

	"github.com/oganes5796/instagram-clon/internal/converter"
	desc "github.com/oganes5796/instagram-clon/pkg/user_v1"
)

func (i *Implementation) Register(ctx context.Context, req *desc.RegisterRequest) (*desc.RegisterResponse, error) {
	id, err := i.userService.Register(ctx, converter.ToUserInfoFromDesc(req.GetInfo()))
	if err != nil {
		return nil, err
	}

	log.Printf("inserted note with id: %d", id)

	return &desc.RegisterResponse{
		UserId: id,
	}, nil
}
