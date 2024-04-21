package api

import (
	"log/slog"

	"github.com/nikitads9/banner-service-api/internal/app/service/banner"
	desc "github.com/nikitads9/banner-service-api/pkg/banner-api"
	"go.opentelemetry.io/otel/trace"
)

const (
	admin = "admin"
	user  = "user"
)

// Implementation ...
type Implementation struct {
	bannerService *banner.Service
	logger        *slog.Logger
	tracer        trace.Tracer
	desc.UnimplementedHandler
}

// NewImplementation ...
func NewImplementation(bannerService *banner.Service, logger *slog.Logger, tracer trace.Tracer) *Implementation {
	return &Implementation{
		bannerService,
		logger,
		tracer,
		desc.UnimplementedHandler{},
	}
}

func NewMockImplementation(i Implementation) *Implementation {
	return &Implementation{
		i.bannerService,
		i.logger,
		i.tracer,
		desc.UnimplementedHandler{},
	}
}
