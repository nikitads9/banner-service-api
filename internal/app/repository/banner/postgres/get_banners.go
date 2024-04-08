package postgres

import (
	"context"

	"github.com/nikitads9/banner-service-api/internal/app/model"
)

func (r *repository) GetBanners(ctx context.Context, mod *model.GetBannersParams) ([]*model.BannerInfo, error) {
	return nil, nil
}
