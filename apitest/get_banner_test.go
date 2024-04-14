package apitest

import (
	"context"
	"database/sql"
	"net/http"
	"testing"

	"github.com/go-faster/jx"
	"github.com/nikitads9/banner-service-api/internal/app/model"
	"github.com/steinfletcher/apitest"
	"github.com/stretchr/testify/require"
)

const (
	activeBanner      = "activeBanner"
	inactiveBanner    = "inactiveBanner"
	cashedBanner      = "cashedBanner"
	otherCashedBanner = "otherCashedBanner"
)

type BannerInfo struct {
	bannerID int64
	banner   *model.Banner
}

func (b *BannerInfo) SetBannerID(id int64) {
	b.bannerID = id
}

func (as *APISuite) TestGetBanner() {
	t := as.Suite.T()
	t.Log("Тестирование метода GetUBanner: GET /user_banner")

	const path = "/user_banner"
	const updatedContent = `{"title": "updated_banner", "width": 30}`

	bannerContentList := map[string]jx.Raw{
		activeBanner:      jx.Raw(`{"title": "banner", "width": 30}`),
		cashedBanner:      jx.Raw(`{"title": "cashed_banner", "width": 30}`),
		otherCashedBanner: jx.Raw(`{"title": "other_cashed_banner", "width": 30}`),
		inactiveBanner:    jx.Raw(`{"title": "disabled_banner", "width": 30}`),
	}

	bannerList := map[string]*BannerInfo{
		activeBanner: {
			banner: &model.Banner{
				FeatureID: 1,
				TagIDs:    []int64{2, 4, 3},
				Content:   []byte(`{"title": "banner", "width": 30}`),
				IsActive:  true,
			},
		},
		cashedBanner: {
			banner: &model.Banner{
				FeatureID: 3,
				TagIDs:    []int64{5, 1, 2},
				Content:   []byte(`{"title": "cashed_banner", "width": 30}`),
				IsActive:  true,
			},
		},
		otherCashedBanner: {
			banner: &model.Banner{
				FeatureID: 4,
				TagIDs:    []int64{5, 1, 2},
				Content:   []byte(`{"title": "other_cashed_banner", "width": 30}`),
				IsActive:  true,
			},
		},
		inactiveBanner: {
			banner: &model.Banner{
				FeatureID: 2,
				TagIDs:    []int64{1, 4, 5},
				Content:   []byte(`{"title": "disabled_banner", "width": 30}`),
				IsActive:  false,
			},
		},
	}

	ctx := context.Background()

	for key, val := range bannerList {
		t.Run("Создание баннеров", func(t *testing.T) {
			id, err := as.bannerService.CreateBanner(ctx, val.banner)
			require.Nil(t, err)
			bannerList[key] = &BannerInfo{
				bannerID: id,
				banner:   val.banner,
			}
		})

	}

	t.Run("Успешное получение видимого баннера", func(t *testing.T) {
		apitest.New().
			Report(apitest.SequenceDiagram()).
			Handler(as.server).
			Get(path).
			Query(FeatureIDParam, "1").Query(TagIDParam, "2").
			Header(TokenHeader, as.userToken).
			Expect(t).
			Body(string(bannerContentList[activeBanner])).
			Status(http.StatusOK).
			End()
	})

	t.Run("Успешное получение обновленного баннера не из кэша", func(t *testing.T) {
		apitest.New().
			Report(apitest.SequenceDiagram()).
			Handler(as.server).
			Get(path).
			Query(FeatureIDParam, "3").Query(TagIDParam, "1").
			Header(TokenHeader, as.userToken).
			Expect(t).
			Body(string(bannerContentList[cashedBanner])).
			Status(http.StatusOK).
			End()

		t.Run("Обновление баннера", func(t *testing.T) {
			err := as.bannerRepository.SetBanner(ctx, &model.SetBannerInfo{
				BannerID: bannerList[cashedBanner].bannerID,
				Content:  model.NullJxRaw{JxRaw: jx.Raw(updatedContent), Valid: true},
				TagIDs:   nil,
				FeatureID: sql.NullInt64{
					Valid: false,
				},
				IsActive: sql.NullBool{
					Valid: false,
				},
			})

			require.Nil(t, err)
		})

		apitest.New().
			Report(apitest.SequenceDiagram()).
			Handler(as.server).
			Get(path).
			Query(FeatureIDParam, "3").Query(TagIDParam, "1").Query(UseLastRevisionParam, "true").
			Header(TokenHeader, as.userToken).
			Expect(t).
			Body(updatedContent).
			Status(http.StatusOK).
			End()
	})
}
