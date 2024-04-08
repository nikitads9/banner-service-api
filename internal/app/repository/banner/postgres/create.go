package postgres

import (
	"context"

	"github.com/nikitads9/banner-service-api/internal/app/model"
)

func (r *repository) CreateBanner(ctx context.Context, mod *model.Banner) (int64, error) {
	return 1, nil
}
