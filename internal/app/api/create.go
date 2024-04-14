package api

import (
	"context"
	"errors"

	"github.com/nikitads9/banner-service-api/internal/app/convert"
	"github.com/nikitads9/banner-service-api/internal/app/repository/banner/postgres"
	"github.com/nikitads9/banner-service-api/internal/middleware/auth"
	desc "github.com/nikitads9/banner-service-api/pkg/banner-api"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

// CreateBanner implements createBanner operation.
//
// Создание нового баннера.
//
// POST /banner
func (i *Implementation) CreateBanner(ctx context.Context, req *desc.CreateBannerRequest) (desc.CreateBannerRes, error) {
	const op = "api.banners.CreateBanner"

	ctx, span := i.tracer.Start(ctx, op)
	defer span.End()

	scope := auth.ScopeFromContext(ctx)
	if scope != admin {
		span.AddEvent("invalid access token scope")
		return &desc.CreateBannerForbidden{}, nil
	}

	id, err := i.bannerService.CreateBanner(ctx, convert.ToBanner(req))
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		if errors.Is(err, postgres.ErrAlreadyExists) {
			return &desc.CreateBannerBadRequest{
				Error: err.Error(),
			}, nil
		}
		return &desc.CreateBannerInternalServerError{
			Error: err.Error(),
		}, err
	}

	span.AddEvent("banner created", trace.WithAttributes(attribute.Int64("id", id)))

	return &desc.CreateBannerResponse{
		BannerID: id,
	}, nil
}
