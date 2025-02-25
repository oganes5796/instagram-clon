package user

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/oganes5796/instagram-clon/internal/client/db"
	"github.com/oganes5796/instagram-clon/internal/domain"
	"github.com/oganes5796/instagram-clon/internal/repository"
	"github.com/oganes5796/instagram-clon/internal/repository/user/converter"
	"github.com/oganes5796/instagram-clon/internal/repository/user/model"
)

const (
	tableName = "users"

	idColumn       = "id"
	emailColumn    = "email"
	usernameColumn = "username"
	passwordColumn = "password"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.UserRepository {
	return &repo{db: db}
}

func (r *repo) Register(ctx context.Context, user *domain.UserInfo) (int64, error) {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(emailColumn, usernameColumn, passwordColumn).
		Values(user.Email, user.Username, user.Password).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "user_repository.Regist",
		QueryRaw: query,
	}

	var id int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repo) GetUser(ctx context.Context, username, password string) (*domain.User, error) {
	builder := sq.Select("id").
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{"username": username, "password": password})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "user_repository.Get",
		QueryRaw: query,
	}

	var user model.User
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&user.ID)
	if err != nil {
		return nil, err
	}

	return converter.TouserFromRepo(&user), nil
}
