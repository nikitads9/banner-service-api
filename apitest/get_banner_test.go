//go:build integration

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
	activeBanner        = "activeBanner"
	inactiveBanner      = "inactiveBanner"
	cashedBanner        = "cashedBanner"
	anotherCashedBanner = "anotherCashedBanner"
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
	const updatedContent = `{"title": "updated_banner", "height": 30}`

	bannerContentList := map[string]jx.Raw{
		activeBanner:        jx.Raw(`{"title": "banner", "height": 30}`),
		cashedBanner:        jx.Raw(`{"title": "cashed_banner", "height": 30}`),
		anotherCashedBanner: jx.Raw(`{"title": "another_cashed_banner", "height": 30}`),
		inactiveBanner:      jx.Raw(`{"title": "disabled_banner", "height": 30}`),
	}

	bannerList := map[string]*BannerInfo{
		activeBanner: {
			banner: &model.Banner{
				FeatureID: 1,
				TagIDs:    []int64{2, 4, 3},
				Content:   []byte(`{"title": "banner", "height": 30}`),
				IsActive:  true,
			},
		},
		cashedBanner: {
			banner: &model.Banner{
				FeatureID: 3,
				TagIDs:    []int64{5, 1, 2},
				Content:   []byte(`{"title": "cashed_banner", "height": 30}`),
				IsActive:  true,
			},
		},
		anotherCashedBanner: {
			banner: &model.Banner{
				FeatureID: 4,
				TagIDs:    []int64{5, 1, 2},
				Content:   []byte(`{"title": "another_cashed_banner", "height": 30}`),
				IsActive:  true,
			},
		},
		inactiveBanner: {
			banner: &model.Banner{
				FeatureID: 2,
				TagIDs:    []int64{1, 4, 5},
				Content:   []byte(`{"title": "disabled_banner", "height": 30}`),
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
			Handler(as.server).
			Get(path).
			Query(FeatureIDParam, "1").Query(TagIDParam, "2").
			Header(TokenHeader, as.userToken).
			Expect(t).
			Body(string(bannerContentList[activeBanner])).
			Status(http.StatusOK).
			End()
	})

	t.Run("Успешное получение видимого баннера из кэша", func(t *testing.T) {
		apitest.New().
			Report(apitest.SequenceDiagram()).
			Handler(as.server).
			Get(path).
			Query(FeatureIDParam, "4").Query(TagIDParam, "1").
			Header(TokenHeader, as.userToken).
			Expect(t).
			Body(string(bannerContentList[anotherCashedBanner])).
			Status(http.StatusOK).
			End()

		t.Run("Создание баннеров", func(t *testing.T) {
			err := as.bannerRepository.SetBanner(ctx, &model.SetBannerInfo{
				BannerID: bannerList[anotherCashedBanner].bannerID,
				FeatureID: sql.NullInt64{
					Valid: false,
				},
				TagIDs:  nil,
				Content: model.NullJxRaw{JxRaw: jx.Raw(updatedContent), Valid: true},
				IsActive: sql.NullBool{
					Valid: false,
				},
			})
			require.Nil(t, err)
		})

		apitest.New().
			Handler(as.server).
			Get(path).
			Query(FeatureIDParam, "4").Query(TagIDParam, "1").
			Header(TokenHeader, as.userToken).
			Expect(t).
			Body(string(bannerContentList[anotherCashedBanner])).
			Status(http.StatusOK).
			End()
	})

	t.Run("Успешное получение обновленного баннера в обход кэша", func(t *testing.T) {
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
				FeatureID: sql.NullInt64{
					Valid: false,
				},
				TagIDs:  nil,
				Content: model.NullJxRaw{JxRaw: jx.Raw(updatedContent), Valid: true},
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

	t.Run("Попытка получения баннера неавторизованным пользователем", func(t *testing.T) {
		apitest.New().
			Handler(as.server).
			Get(path).
			Query(FeatureIDParam, "3").Query(TagIDParam, "4").
			Expect(t).
			Status(http.StatusUnauthorized).
			End()
	})

	t.Run("Попытка получения баннера пользователя с неверными правами", func(t *testing.T) {
		apitest.New().
			Handler(as.server).
			Get(path).
			Query(FeatureIDParam, "3").Query(TagIDParam, "4").
			Header(TokenHeader, as.adminToken).
			Expect(t).
			Status(http.StatusForbidden).
			End()
	})

	t.Run("Попытка получения неактивного баннера", func(t *testing.T) {
		apitest.New().
			Handler(as.server).
			Get(path).
			Query(FeatureIDParam, "2").Query(TagIDParam, "4").
			Header(TokenHeader, as.userToken).
			Expect(t).
			Status(http.StatusNotFound).
			End()
	})

	t.Run("Попытка получения несуществующего баннера", func(t *testing.T) {
		apitest.New().
			Handler(as.server).
			Get(path).
			Query(FeatureIDParam, "3").Query(TagIDParam, "4").
			Header(TokenHeader, as.userToken).
			Expect(t).
			Status(http.StatusNotFound).
			End()
	})

	t.Run("Попытка получения баннера с некорректными параметрами запроса", func(t *testing.T) {
		apitest.New().
			Handler(as.server).
			Get(path).
			Query(FeatureIDParam, "rediska").Query(TagIDParam, "4").
			Header(TokenHeader, as.userToken).
			Expect(t).
			Status(http.StatusBadRequest).
			End()
	})
	t.Run("Попытка получения баннера без одного из обязательных параметров", func(t *testing.T) {
		apitest.New().
			Handler(as.server).
			Get(path).
			Query(TagIDParam, "4").
			Header(TokenHeader, as.userToken).
			Expect(t).
			Status(http.StatusBadRequest).
			End()

		apitest.New().
			Handler(as.server).
			Get(path).
			Query(FeatureIDParam, "3").
			Header(TokenHeader, as.userToken).
			Expect(t).
			Status(http.StatusBadRequest).
			End()
	})
}
