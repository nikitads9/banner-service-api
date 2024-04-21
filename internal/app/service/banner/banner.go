package banner

import (
	"log/slog"

	"github.com/nikitads9/banner-service-api/internal/app/repository/banner"
	"github.com/nikitads9/banner-service-api/internal/pkg/db"
	"go.opentelemetry.io/otel/trace"
)

// Service ...
type Service struct {
	postgresRepository banner.Repository
	bannerCache        banner.Cache
	tracer             trace.Tracer
	log                *slog.Logger
	txManager          db.TxManager
}

// NewBannerService ...
func NewBannerService(bannerRepository banner.Repository, bannerCache banner.Cache, tracer trace.Tracer, log *slog.Logger, txManager db.TxManager) *Service {
	return &Service{
		postgresRepository: bannerRepository,
		bannerCache:        bannerCache,
		tracer:             tracer,
		log:                log,
		txManager:          txManager,
	}
}

// NewMockBannerService ...
func NewMockBannerService(deps ...interface{}) *Service {
	is := Service{}
	for _, val := range deps {
		switch s := val.(type) {
		case banner.Repository:
			is.postgresRepository = s
		case banner.Cache:
			is.bannerCache = s
		case trace.Tracer:
			is.tracer = s
		case *slog.Logger:
			is.log = s
		case db.TxManager:
			is.txManager = s
		}
	}
	return &is
}
