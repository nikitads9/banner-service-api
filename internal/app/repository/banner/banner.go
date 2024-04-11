package banner

import (
	"context"

	"github.com/go-faster/jx"
	"github.com/nikitads9/banner-service-api/internal/app/model"
)

type Repository interface {
	CreateBanner(ctx context.Context, mod *model.Banner) (int64, error)
	GetBanner(ctx context.Context, featureID int64, tagID int64) (jx.Raw, error)
	GetBanners(ctx context.Context, mod *model.GetBannersParams) ([]*model.BannerInfo, error)
	SetBanner(ctx context.Context, mod *model.SetBannerInfo) error
	DeleteBanner(ctx context.Context, bannerID int64) error
}
