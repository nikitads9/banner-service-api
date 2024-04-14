package banner

import (
	"context"

	"github.com/nikitads9/banner-service-api/internal/app/model"
)

// GetBanners это получение всех баннеров c фильтрацией по фиче и/или тегу.
func (s *Service) GetBanners(ctx context.Context, mod *model.GetBannersParams) ([]*model.BannerInfo, error) {
	return s.postgresRepository.GetBanners(ctx, mod)
}
