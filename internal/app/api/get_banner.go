package api

import (
	"context"

	"github.com/nikitads9/banner-service-api/internal/middleware/auth"
	desc "github.com/nikitads9/banner-service-api/pkg/banner-api"
)

// GetBanner implements getBanner operation.
//
// Получение баннера для пользователя.
//
// GET /user_banner
func (i *Implementation) GetBanner(ctx context.Context, params desc.GetBannerParams) (desc.GetBannerRes, error) {
	scope := auth.ScopeFromContext(ctx)
	if scope != "user" {
		return &desc.GetBannerForbidden{}, nil
	}

	content, err := i.bannerService.GetBanner(ctx, params.FeatureID, params.TagID, params.UseLastRevision.Value)
	if err != nil {
		return &desc.GetBannerInternalServerError{
			Error: err.Error(),
		}, err
	}

	var response desc.GetBannerResponse = content
	return &response, nil
}
