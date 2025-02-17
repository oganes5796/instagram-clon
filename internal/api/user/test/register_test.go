package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	"github.com/oganes5796/instagram-clon/internal/api/user"
	"github.com/oganes5796/instagram-clon/internal/domain"
	"github.com/oganes5796/instagram-clon/internal/service"

	serviceMocks "github.com/oganes5796/instagram-clon/internal/service/mocks"
	desk "github.com/oganes5796/instagram-clon/pkg/user_v1"
)

func TestRegister(t *testing.T) {
	t.Parallel()
	type userServiceMockFunc func(mc *minimock.Controller) service.UserService

	type args struct {
		ctx context.Context
		req *desk.RegisterRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id       = gofakeit.Int64()
		email    = gofakeit.Email()
		username = gofakeit.Username()
		password = gofakeit.Password(true, true, true, true, false, 10)

		serviceErr = fmt.Errorf("service error")

		req = &desk.RegisterRequest{
			Info: &desk.UserInfo{
				Email:    email,
				Username: username,
				Password: password,
			},
		}

		info = &domain.UserInfo{
			Email:    email,
			Username: username,
			Password: password,
		}

		res = &desk.RegisterResponse{
			UserId: id,
		}
	)

	tests := []struct {
		name            string
		args            args
		want            *desk.RegisterResponse
		err             error
		userServiceMock userServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := serviceMocks.NewUserServiceMock(mc)
				mock.RegisterMock.Expect(ctx, info).Return(id, nil)
				return mock
			},
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  serviceErr,
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := serviceMocks.NewUserServiceMock(mc)
				mock.RegisterMock.Expect(ctx, info).Return(0, serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userServiceMock := tt.userServiceMock(mc)
			api := user.NewImplementation(userServiceMock)

			newID, err := api.Register(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, newID)
		})
	}
}
