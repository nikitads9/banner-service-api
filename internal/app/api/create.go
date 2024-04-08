package api

import (
	"context"
	"fmt"

	"github.com/nikitads9/banner-service-api/internal/middleware/auth"
	desc "github.com/nikitads9/banner-service-api/pkg/banner-api"
)

// CreateBanner implements createBanner operation.
//
// Создание нового баннера.
//
// POST /banner
func (i *Implementation) CreateBanner(ctx context.Context, req *desc.CreateBannerRequest) (desc.CreateBannerRes, error) {
	scope := auth.ScopeFromContext(ctx)
	if scope != "admin" {
		return &desc.CreateBannerForbidden{}, nil
	}

	fmt.Println("CreateBanner")
	return &desc.CreateBannerResponse{
		BannerID: desc.OptInt{Value: 1, Set: true},
	}, nil
}

// DeleteBanner implements deleteBanner operation.
//
// Удаление баннера по идентификатору.
//
// DELETE /banner/{id}
func (i *Implementation) DeleteBanner(ctx context.Context, params desc.DeleteBannerParams) (desc.DeleteBannerRes, error) {
	scope := auth.ScopeFromContext(ctx)
	if scope != "admin" {
		return &desc.DeleteBannerForbidden{}, nil
	}
	fmt.Println("DeleteBanner")
	return &desc.DeleteBannerNoContent{}, nil
}

// GetBanner implements getBanner operation.
//
// Получение баннера для пользователя.
//
// GET /user_banner
func (i *Implementation) GetBanner(ctx context.Context, params desc.GetBannerParams) (desc.GetBannerRes, error) {
	scope := auth.ScopeFromContext(ctx)
	if scope != "user" {
		return &desc.GetBannerForbidden{}, nil
	}
	fmt.Println("GetBanner")
	return &desc.GetBannerResponse{"title": []byte("\"some_title\""), "text": []byte("\"some_text\""), "url": []byte("\"some_url\"")}, nil
}

// GetBanners implements getBanners operation.
//
// Получение всех баннеров c фильтрацией по фиче и/или
// тегу.
//
// GET /banner
func (i *Implementation) GetBanners(ctx context.Context, params desc.GetBannersParams) (desc.GetBannersRes, error) {
	scope := auth.ScopeFromContext(ctx)
	if scope != "admin" {
		return &desc.GetBannersForbidden{}, nil
	}
	fmt.Println("GetBanners")
	return &desc.GetBannersResponse{desc.GetBannersResponseItem{}}, nil
}

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
	fmt.Println("SetBanner")
	return &desc.SetBannerOK{}, nil
}
