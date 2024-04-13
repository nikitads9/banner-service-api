package api

import (
	"github.com/nikitads9/banner-service-api/internal/app/service/banner"
	desc "github.com/nikitads9/banner-service-api/pkg/banner-api"
	"go.opentelemetry.io/otel/trace"
)

type Implementation struct {
	bannerService *banner.Service
	tracer        trace.Tracer
	desc.UnimplementedHandler
}

func NewImplementation(bannerService *banner.Service, tracer trace.Tracer) *Implementation {
	return &Implementation{
		bannerService,
		tracer,
		desc.UnimplementedHandler{},
	}
}
