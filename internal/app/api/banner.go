package api

import (
	"github.com/nikitads9/banner-service-api/internal/app/service/banner"
	desc "github.com/nikitads9/banner-service-api/pkg/banner-api"
)

type Implementation struct {
	bannerService *banner.Service
	desc.UnimplementedHandler
}

func NewImplementation(bannerService *banner.Service) *Implementation {
	return &Implementation{
		bannerService,
		desc.UnimplementedHandler{},
	}
}
