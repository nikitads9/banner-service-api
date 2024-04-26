//go:build unit

package api

import (
	"context"

	"io"
	"log/slog"
	"testing"

	gofakeit "github.com/brianvoe/gofakeit/v7"
	"github.com/golang/mock/gomock"
	pgxV5 "github.com/jackc/pgx/v5"
	"github.com/nikitads9/banner-service-api/internal/app/model"
	bannerRepoMocks "github.com/nikitads9/banner-service-api/internal/app/repository/banner/mocks"
	"github.com/nikitads9/banner-service-api/internal/app/service/banner"
	"github.com/nikitads9/banner-service-api/internal/middleware/auth"
	dbMocks "github.com/nikitads9/banner-service-api/internal/pkg/db/mocks_db"
	txMocks "github.com/nikitads9/banner-service-api/internal/pkg/db/mocks_tx"
	"github.com/nikitads9/banner-service-api/internal/pkg/db/pg"
	"github.com/nikitads9/banner-service-api/internal/pkg/observability"
	desc "github.com/nikitads9/banner-service-api/pkg/banner-api"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func Test_CreateBanner(t *testing.T) {
	var (
		logger           = slog.New(slog.NewTextHandler(io.Discard, &slog.HandlerOptions{Level: slog.LevelError}))
		mock             = gomock.NewController(t)
		bannerID         = gofakeit.Int64()
		featureID        = gofakeit.Int64()
		tagIds           = []int64{gofakeit.Int64(), gofakeit.Int64(), gofakeit.Int64()}
		ErrAlreadyExists = errors.New("this banner already exists")
		txErr            = errors.Wrap(ErrAlreadyExists, "failed executing code inside transaction")
		txErrText        = txErr.Error()
	)

	content, err := gofakeit.JSON(nil)
	if err != nil {
		t.Fatalf("could not fake json. Error: %s", err)
	}

	validBannerInfo := &model.Banner{
		FeatureID: featureID,
		TagIDs:    tagIds,
		Content:   content,
		IsActive:  false,
	}
	validReq := &desc.CreateBannerRequest{
		FeatureID: featureID,
		TagIds:    tagIds,
		Content:   content,
		IsActive:  false,
	}

	tracer := observability.NewMockTracer()

	bannerRepoMock := bannerRepoMocks.NewMockRepository(mock)
	bannerCacheMock := bannerRepoMocks.NewMockCache(mock)
	dbMock := dbMocks.NewMockDB(mock)
	txMock := txMocks.NewMockTx(mock)
	txManagerMock := pg.NewMockTransactionManager(dbMock)
	api := NewMockImplementation(Implementation{
		bannerService: banner.NewMockBannerService(bannerRepoMock, bannerCacheMock, tracer, logger, txManagerMock),
		logger:        logger,
		tracer:        tracer,
	})

	t.Run("double insert banner error case", func(t *testing.T) {
		ctx := context.WithValue(context.Background(), auth.KeyScope{}, admin)
		txCtx := pg.MakeContextTx(ctx, txMock)

		dbMock.EXPECT().BeginTx(ctx, pgxV5.TxOptions{IsoLevel: pgxV5.ReadCommitted}).Return(txMock, nil)
		bannerRepoMock.EXPECT().CreateBanner(txCtx, validBannerInfo).Return(bannerID, nil).Times(1)
		txMock.EXPECT().Commit(txCtx).Return(nil)
		dbMock.EXPECT().BeginTx(ctx, pgxV5.TxOptions{IsoLevel: pgxV5.ReadCommitted}).Return(txMock, nil)
		bannerRepoMock.EXPECT().CreateBanner(txCtx, &model.Banner{
			TagIDs:    tagIds,
			FeatureID: featureID,
			Content:   content,
			IsActive:  false,
		},
		).Return(int64(0), ErrAlreadyExists).Times(1)
		txMock.EXPECT().Rollback(txCtx).Return(nil)

		res, err := api.CreateBanner(ctx, &desc.CreateBannerRequest{
			FeatureID: featureID,
			TagIds:    tagIds,
			Content:   content,
			IsActive:  false,
		},
		)
		require.Nil(t, err)
		require.Equal(t, &desc.CreateBannerResponse{
			BannerID: bannerID,
		}, res)

		_, err = api.CreateBanner(ctx, &desc.CreateBannerRequest{
			TagIds:    tagIds,
			FeatureID: featureID,
			Content:   content,
			IsActive:  false,
		})
		require.Error(t, err)
		require.Equal(t, txErrText, err.Error())
	})

	t.Run("add banner success case", func(t *testing.T) {
		ctx := context.WithValue(context.Background(), auth.KeyScope{}, admin)
		txCtx := pg.MakeContextTx(ctx, txMock)
		dbMock.EXPECT().BeginTx(ctx, pgxV5.TxOptions{IsoLevel: pgxV5.ReadCommitted}).Return(txMock, nil)
		bannerRepoMock.EXPECT().CreateBanner(txCtx, validBannerInfo).Return(bannerID, nil).Times(1)
		txMock.EXPECT().Commit(txCtx).Return(nil)

		res, err := api.CreateBanner(ctx, validReq)
		require.Nil(t, err)
		require.Equal(t, &desc.CreateBannerResponse{
			BannerID: bannerID,
		}, res)
	})

}
