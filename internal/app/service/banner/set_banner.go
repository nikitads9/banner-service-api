package banner

import (
	"context"

	"github.com/nikitads9/banner-service-api/internal/app/model"
)

func (s *Service) SetBanner(ctx context.Context, mod *model.SetBannerInfo) error {
	return s.postgresRepository.SetBanner(ctx, mod)
}
