package banner

import (
	"context"
	"log"
	"log/slog"
	"os"

	bannerRepository "github.com/nikitads9/banner-service-api/internal/app/repository/banner"
	bannerService "github.com/nikitads9/banner-service-api/internal/app/service/banner"
	"github.com/nikitads9/banner-service-api/internal/app/service/jwt"

	"github.com/nikitads9/banner-service-api/internal/config"
	"github.com/nikitads9/banner-service-api/internal/logger/sl"
	"github.com/nikitads9/banner-service-api/internal/pkg/db"
	"github.com/nikitads9/banner-service-api/internal/pkg/db/transaction"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

type serviceProvider struct {
	configPath string
	configType string
	config     *config.BannerConfig

	db        db.Client
	txManager db.TxManager

	log *slog.Logger

	bannerRepository bannerRepository.Repository
	bannerService    *bannerService.Service

	jwtService jwt.Service
}

func newServiceProvider(configType string, configPath string) *serviceProvider {
	return &serviceProvider{
		configType: configType,
		configPath: configPath,
	}
}

func (s *serviceProvider) GetDB(ctx context.Context) db.Client {
	if s.db == nil {
		cfg, err := s.GetConfig().GetDBConfig()
		if err != nil {
			s.log.Error("could not get db config: %s", sl.Err(err))
		}
		dbc, err := db.NewClient(ctx, cfg)
		if err != nil {
			s.log.Error("coud not connect to db: %s", sl.Err(err))
		}
		s.db = dbc
	}

	return s.db
}

func (s *serviceProvider) GetConfig() *config.BannerConfig {
	if s.config == nil {
		if s.configType == "env" {
			cfg, err := config.ReadBannerConfigEnv()
			if err != nil {
				log.Fatalf("could not get banner-api config from env: %s", err)
			}
			s.config = cfg
		} else {
			cfg, err := config.ReadBannerConfigFile(s.configPath)
			if err != nil {
				log.Fatalf("could not get banner-api config from file: %s", err)
			}
			s.config = cfg
		}
	}

	return s.config
}

func (s *serviceProvider) GetBannerRepository(ctx context.Context) bannerRepository.Repository {
	if s.bannerRepository == nil {
		s.bannerRepository = bannerRepository.NewBannerRepository(s.GetDB(ctx), s.GetLogger())
		return s.bannerRepository
	}

	return s.bannerRepository
}

func (s *serviceProvider) GetBannerService(ctx context.Context) *bannerService.Service {
	if s.bannerService == nil {
		bannerRepository := s.GetBannerRepository(ctx)
		s.bannerService = bannerService.NewBannerService(bannerRepository, s.GetLogger(), s.TxManager(ctx))
	}

	return s.bannerService
}

func (s *serviceProvider) GetJWTService(ctx context.Context) jwt.Service {
	if s.jwtService == nil {
		s.jwtService = jwt.NewJWTService(s.GetConfig().GetJWTConfig().Secret, s.GetConfig().GetJWTConfig().Expiration, s.GetLogger())
	}

	return s.jwtService
}

func (s *serviceProvider) GetLogger() *slog.Logger {
	if s.log == nil {
		env := s.GetConfig().GetEnv()
		switch env {
		case envLocal:
			s.log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
		case envDev:
			s.log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
		case envProd:
			s.log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
		}

		s.log.With(slog.String("env", env))
	}

	return s.log
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.GetDB(ctx).DB())
	}

	return s.txManager
}
