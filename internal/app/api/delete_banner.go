package api

import (
	"context"

	"github.com/nikitads9/banner-service-api/internal/app/repository/banner/postgres"
	"github.com/nikitads9/banner-service-api/internal/middleware/auth"
	desc "github.com/nikitads9/banner-service-api/pkg/banner-api"
)

// DeleteBanner implements deleteBanner operation.
//
// Удаление баннера по идентификатору.
//
// DELETE /banner/{id}
func (i *Implementation) DeleteBanner(ctx context.Context, params desc.DeleteBannerParams) (desc.DeleteBannerRes, error) {
	scope := auth.ScopeFromContext(ctx)
	if scope != "admin" {
		return &desc.DeleteBannerForbidden{}, nil
	}

	err := i.bannerService.DeleteBanner(ctx, params.ID)
	if err != nil {
		if err == postgres.ErrNotFound {
			return &desc.DeleteBannerNotFound{}, nil
		}
		return &desc.DeleteBannerInternalServerError{
			Error: err.Error(),
		}, err
	}

	return &desc.DeleteBannerNoContent{}, nil
}
