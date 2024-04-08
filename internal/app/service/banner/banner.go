package banner

import (
	"errors"
	"log/slog"

	"github.com/nikitads9/banner-service-api/internal/app/repository/banner"
	"github.com/nikitads9/banner-service-api/internal/pkg/db"

	"github.com/jackc/pgx/v5/pgconn"
)

type Service struct {
	bannerRepository banner.Repository
	log              *slog.Logger
	txManager        db.TxManager
}

var (
	ErrNoConnection = errors.New("can't begin transaction, no connection to database")
	pgNoConnection  = new(*pgconn.ConnectError)
)

func NewBannerService(bannerRepository banner.Repository, log *slog.Logger, txManager db.TxManager) *Service {
	return &Service{
		bannerRepository: bannerRepository,
		log:              log,
		txManager:        txManager,
	}
}
