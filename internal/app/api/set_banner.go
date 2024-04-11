package api

import (
	"context"

	"github.com/go-faster/jx"
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

	//TODO: try implement nullable content object
	var content jx.Raw
	if len(req.GetContent().String()) != 0 {
		content = req.GetContent()
	}

	err := i.bannerService.SetBanner(ctx, convert.ToSetBannerInfo(params.ID, content, req))
	if err != nil {
		return &desc.SetBannerInternalServerError{
			Error: err.Error(),
		}, err
	}

	return &desc.SetBannerOK{}, nil
}
