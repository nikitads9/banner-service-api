package banner

import (
	"context"
	"log/slog"

	"github.com/nikitads9/banner-service-api/internal/app/model"
	"github.com/nikitads9/banner-service-api/internal/logger/sl"
)

func (s *Service) CreateBanner(ctx context.Context, mod *model.Banner) (int64, error) {
	const op = "service.banners.CreateBanner"

	log := s.log.With(
		slog.String("op", op),
	)

	var id int64
	var errTx error

	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		id, errTx = s.postgresRepository.CreateBanner(ctx, mod)
		if errTx != nil {
			log.Error("create banner operation failed", sl.Err(errTx))
			return errTx
		}

		return nil
	})

	if err != nil {
		log.Error("transaction failed", sl.Err(err))
		return 0, err
	}
	return id, nil
}
