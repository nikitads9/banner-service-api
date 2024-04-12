package banner

import (
	"context"
	"log/slog"

	"github.com/nikitads9/banner-service-api/internal/app/model"
	"github.com/nikitads9/banner-service-api/internal/logger/sl"
)

func (s *Service) SetBanner(ctx context.Context, mod *model.SetBannerInfo) error {
	const op = "service.banners.SetBanner"

	log := s.log.With(
		slog.String("op", op),
	)

	var errTx error

	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		errTx = s.postgresRepository.SetBanner(ctx, mod)
		if errTx != nil {
			log.Error("create banner operation failed", sl.Err(errTx))
			return errTx
		}

		return nil
	})

	if err != nil {
		log.Error("transaction failed", sl.Err(err))
		return err
	}

	return nil
}
