package banner

import (
	"context"

	"github.com/nikitads9/banner-service-api/internal/app/model"
)

func (s *Service) CreateBanner(ctx context.Context, mod *model.Banner) (int64, error) {
	return s.postgresRepository.CreateBanner(ctx, mod)
}
