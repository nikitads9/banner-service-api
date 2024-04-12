package api

import (
	"context"

	"github.com/nikitads9/banner-service-api/internal/app/convert"
	"github.com/nikitads9/banner-service-api/internal/middleware/auth"
	desc "github.com/nikitads9/banner-service-api/pkg/banner-api"
)

// GetBanners implements getBanners operation.
//
// Получение всех баннеров c фильтрацией по фиче и/или
// тегу.
//
// GET /banner
func (i *Implementation) GetBanners(ctx context.Context, params desc.GetBannersParams) (desc.GetBannersRes, error) {
	scope := auth.ScopeFromContext(ctx)
	if scope != "admin" {
		return &desc.GetBannersForbidden{}, nil
	}

	var (
		limit  int64 = 1000
		offset int64 = 0
	)

	banners, err := i.bannerService.GetBanners(ctx, convert.ToGetBannersParams(params, limit, offset))
	if err != nil {
		return &desc.Error{
			Error: err.Error(),
		}, err
	}

	var response desc.GetBannersResponse = convert.ToGetBannersResponse(banners)
	return &response, nil
}
