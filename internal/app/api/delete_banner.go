package api

import (
	"context"

	"github.com/nikitads9/banner-service-api/internal/app/repository/banner/postgres"
	"github.com/nikitads9/banner-service-api/internal/middleware/auth"
	desc "github.com/nikitads9/banner-service-api/pkg/banner-api"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

// DeleteBanner implements deleteBanner operation.
//
// Удаление баннера по идентификатору.
//
// DELETE /banner/{id}
func (i *Implementation) DeleteBanner(ctx context.Context, params desc.DeleteBannerParams) (desc.DeleteBannerRes, error) {
	const op = "api.banners.DeleteBanner"

	ctx, span := i.tracer.Start(ctx, op)
	defer span.End()

	scope := auth.ScopeFromContext(ctx)
	if scope != "admin" {
		span.AddEvent("invalid access token scope")
		return &desc.DeleteBannerForbidden{}, nil
	}

	err := i.bannerService.DeleteBanner(ctx, params.ID)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		if err == postgres.ErrNotFound {
			return &desc.DeleteBannerNotFound{}, nil
		}
		return &desc.DeleteBannerInternalServerError{
			Error: err.Error(),
		}, err
	}

	span.AddEvent("banner deleted", trace.WithAttributes(attribute.Int64("id", params.ID)))

	return &desc.DeleteBannerNoContent{}, nil
}
