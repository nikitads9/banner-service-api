package banner

import (
	"errors"
	"log/slog"

	"github.com/nikitads9/banner-service-api/internal/app/repository/banner"
	"github.com/nikitads9/banner-service-api/internal/pkg/db"
)

type Service struct {
	postgresRepository banner.Repository
	bannerCache        banner.Cache
	log                *slog.Logger
	txManager          db.TxManager
}

var (
	ErrNoConnection = errors.New("can't begin transaction, no connection to database")
)

func NewBannerService(bannerRepository banner.Repository, bannerCache banner.Cache, log *slog.Logger, txManager db.TxManager) *Service {
	return &Service{
		postgresRepository: bannerRepository,
		bannerCache:        bannerCache,
		log:                log,
		txManager:          txManager,
	}
}

func NewMockBannerService(deps ...interface{}) *Service {
	is := Service{}
	for _, val := range deps {
		switch s := val.(type) {
		case banner.Repository:
			is.postgresRepository = s
		}
	}
	return &is
}
