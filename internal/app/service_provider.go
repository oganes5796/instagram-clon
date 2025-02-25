package app

import (
	"context"

	"github.com/oganes5796/instagram-clon/internal/api/user"
	"github.com/oganes5796/instagram-clon/internal/client/db"
	"github.com/oganes5796/instagram-clon/internal/client/db/pg"
	"github.com/oganes5796/instagram-clon/internal/closer"
	"github.com/oganes5796/instagram-clon/internal/config"
	"github.com/oganes5796/instagram-clon/internal/logger"
	"github.com/oganes5796/instagram-clon/internal/repository"
	userReposirory "github.com/oganes5796/instagram-clon/internal/repository/user"
	"github.com/oganes5796/instagram-clon/internal/service"
	userService "github.com/oganes5796/instagram-clon/internal/service/user"

	"go.uber.org/zap"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig

	dbClient       db.Client
	userReposirory repository.UserRepository

	userService service.UserService

	userImpl *user.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			logger.Fatal("failed to load PG config:", zap.Error(err))
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			logger.Fatal("failed to load GRPC config:", zap.Error(err))
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		client, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			logger.Fatal("failed to create DB client:", zap.Error(err))
		}

		err = client.DB().Ping(ctx)
		if err != nil {
			logger.Fatal("failed to ping DB:", zap.Error(err))
		}
		closer.Add(client.Close)

		s.dbClient = client
	}

	return s.dbClient
}

func (s *serviceProvider) UserRepository(ctx context.Context) repository.UserRepository {
	if s.userReposirory == nil {
		s.userReposirory = userReposirory.NewRepository(s.DBClient(ctx))
	}

	return s.userReposirory
}

func (s *serviceProvider) UserService(ctx context.Context) service.UserService {
	if s.userService == nil {
		s.userService = userService.NewService(
			s.UserRepository(ctx),
		)
	}

	return s.userService
}

func (s *serviceProvider) UserImpl(ctx context.Context) *user.Implementation {
	if s.userImpl == nil {
		s.userImpl = user.NewImplementation(s.UserService(ctx))
	}

	return s.userImpl
}
