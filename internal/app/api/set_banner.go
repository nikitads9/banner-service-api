package api

import (
	"context"

	"github.com/nikitads9/banner-service-api/internal/app/convert"
	"github.com/nikitads9/banner-service-api/internal/middleware/auth"
	desc "github.com/nikitads9/banner-service-api/pkg/banner-api"
)

// SetBanner implements setBanner operation.
//
// Обновление содержимого баннера.
//
// PATCH /banner/{id}
func (i *Implementation) SetBanner(ctx context.Context, req *desc.SetBannerRequest, params desc.SetBannerParams) (desc.SetBannerRes, error) {
	scope := auth.ScopeFromContext(ctx)
	if scope != "admin" {
		return &desc.SetBannerForbidden{}, nil
	}

	var content []byte
	var err error
	if req.GetContent().IsSet() {
		content, err = req.Content.MarshalJSON()
		if err != nil {
			return &desc.SetBannerBadRequest{
				Error: err.Error(),
			}, err
		}
	}

	err = i.bannerService.SetBanner(ctx, convert.ToSetBannerInfo(params.ID, content, req))
	if err != nil {
		return &desc.SetBannerInternalServerError{
			Error: err.Error(),
		}, err
	}

	return &desc.SetBannerOK{}, nil
}
