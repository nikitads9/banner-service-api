package api

import (
	"context"

	"github.com/nikitads9/banner-service-api/internal/app/convert"
	"github.com/nikitads9/banner-service-api/internal/middleware/auth"
	desc "github.com/nikitads9/banner-service-api/pkg/banner-api"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

// GetBanners implements getBanners operation.
//
// Получение всех баннеров c фильтрацией по фиче и/или
// тегу.
//
// GET /banner
func (i *Implementation) GetBanners(ctx context.Context, params desc.GetBannersParams) (desc.GetBannersRes, error) {
	const op = "api.banners.GetBanners"

	ctx, span := i.tracer.Start(ctx, op)
	defer span.End()

	scope := auth.ScopeFromContext(ctx)
	if scope != admin {
		span.AddEvent("invalid access token scope")
		return &desc.GetBannersForbidden{}, nil
	}

	var (
		limit  int64 = 1000
		offset int64
	)

	banners, err := i.bannerService.GetBanners(ctx, convert.ToGetBannersParams(params, limit, offset))
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return &desc.Error{
			Error: err.Error(),
		}, err
	}

	span.AddEvent("filtered banners acquired", trace.WithAttributes(attribute.Int("quantity", len(banners))))

	var response desc.GetBannersResponse = convert.ToGetBannersResponse(banners)
	return &response, nil
}
