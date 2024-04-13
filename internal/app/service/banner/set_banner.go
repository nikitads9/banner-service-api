package banner

import (
	"context"
	"log/slog"

	"github.com/nikitads9/banner-service-api/internal/app/model"
	"github.com/nikitads9/banner-service-api/internal/logger/sl"
	"go.opentelemetry.io/otel/codes"
)

func (s *Service) SetBanner(ctx context.Context, mod *model.SetBannerInfo) error {
	const op = "service.banners.SetBanner"

	log := s.log.With(
		slog.String("op", op),
	)

	ctx, span := s.tracer.Start(ctx, op)
	defer span.End()

	var errTx error

	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		errTx = s.postgresRepository.SetBanner(ctx, mod)
		if errTx != nil {
			span.RecordError(errTx)
			span.SetStatus(codes.Error, errTx.Error())
			log.Error("create banner operation failed", sl.Err(errTx))
			return errTx
		}

		return nil
	})

	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error("transaction failed", sl.Err(err))
		return err
	}

	return nil
}
