package api

import (
	"context"
	"errors"

	"github.com/nikitads9/banner-service-api/internal/app/convert"
	"github.com/nikitads9/banner-service-api/internal/app/repository/banner/postgres"
	"github.com/nikitads9/banner-service-api/internal/middleware/auth"
	desc "github.com/nikitads9/banner-service-api/pkg/banner-api"
	"go.opentelemetry.io/otel/codes"
)

// SetBanner implements setBanner operation.
//
// Обновление содержимого баннера.
//
// PATCH /banner/{id}
func (i *Implementation) SetBanner(ctx context.Context, req *desc.SetBannerRequest, params desc.SetBannerParams) (desc.SetBannerRes, error) {
	const op = "api.banners.SetBanner"

	ctx, span := i.tracer.Start(ctx, op)
	defer span.End()

	scope := auth.ScopeFromContext(ctx)
	if scope != "admin" {
		span.AddEvent("invalid access token scope")
		return &desc.SetBannerForbidden{}, nil
	}

	//TODO: try implement nullable content object
	/* 	var content jx.Raw
	   	if len(req.GetContent().String()) != 0 {
	   		content = req.GetContent()
	   	} */

	err := i.bannerService.SetBanner(ctx, convert.ToSetBannerInfo(params.ID, req))
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		if errors.Is(err, postgres.ErrNotFound) {
			return &desc.SetBannerBadRequest{
				Error: err.Error(),
			}, nil
		}

		if errors.Is(err, postgres.ErrAlreadyExists) {
			return &desc.SetBannerBadRequest{
				Error: err.Error(),
			}, nil
		}

		return &desc.SetBannerInternalServerError{
			Error: err.Error(),
		}, err
	}

	span.AddEvent("banner updated")

	return &desc.SetBannerOK{}, nil
}
