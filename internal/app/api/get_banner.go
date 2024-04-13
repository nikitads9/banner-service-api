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

// GetBanner implements getBanner operation.
//
// Получение баннера для пользователя.
//
// GET /user_banner
func (i *Implementation) GetBanner(ctx context.Context, params desc.GetBannerParams) (desc.GetBannerRes, error) {
	const op = "api.banners.GetBanner"

	ctx, span := i.tracer.Start(ctx, op)
	defer span.End()

	scope := auth.ScopeFromContext(ctx)
	if scope != "user" {
		span.AddEvent("invalid access token scope")
		return &desc.GetBannerForbidden{}, nil
	}

	content, err := i.bannerService.GetBanner(ctx, params.FeatureID, params.TagID, params.UseLastRevision.Value)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		if err == postgres.ErrNotFound {
			return &desc.GetBannerNotFound{}, nil
		}
		return &desc.GetBannerInternalServerError{
			Error: err.Error(),
		}, err
	}

	span.AddEvent("banner acquired", trace.WithAttributes(attribute.Int64("featureID", params.FeatureID), attribute.Int64("tagID", params.TagID)))

	var response desc.GetBannerResponse = desc.GetBannerResponse(content)
	return &response, nil
}
