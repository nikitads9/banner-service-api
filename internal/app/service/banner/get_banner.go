package banner

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/go-faster/jx"
	"github.com/nikitads9/banner-service-api/internal/logger/sl"
)

func (s *Service) GetBanner(ctx context.Context, featureID int64, tagID int64, useLastRevision bool) (jx.Raw, error) {
	const op = "service.banners.GetBanner"

	log := s.log.With(
		slog.String("op", op),
	)

	var content jx.Raw

	key := fmt.Sprintf("%d-%d", featureID, tagID)

	if !useLastRevision {
		content, err := s.bannerCache.Get(ctx, key)
		if err == nil {
			return content, nil
		}

		log.Error("could not find content in cache", sl.Err(err))
	}

	content, err := s.postgresRepository.GetBanner(ctx, featureID, tagID)
	if err == nil {
		err = s.bannerCache.Set(ctx, key, content, 5*time.Minute)
		if err != nil {
			log.Error("could not write content to cache", sl.Err(err))
		}
	}

	return content, err
}
