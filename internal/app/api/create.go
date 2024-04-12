package api

import (
	"context"

	"github.com/nikitads9/banner-service-api/internal/app/convert"
	"github.com/nikitads9/banner-service-api/internal/middleware/auth"
	desc "github.com/nikitads9/banner-service-api/pkg/banner-api"
)

// CreateBanner implements createBanner operation.
//
// Создание нового баннера.
//
// POST /banner
func (i *Implementation) CreateBanner(ctx context.Context, req *desc.CreateBannerRequest) (desc.CreateBannerRes, error) {
	scope := auth.ScopeFromContext(ctx)
	if scope != "admin" {
		return &desc.CreateBannerForbidden{}, nil
	}

	id, err := i.bannerService.CreateBanner(ctx, convert.ToBanner(req))

	if err != nil {
		return &desc.CreateBannerInternalServerError{
			Error: err.Error(),
		}, err
	}

	return &desc.CreateBannerResponse{
		BannerID: id,
	}, nil
}
