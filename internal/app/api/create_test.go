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
	desc "github.com/nikitads9/banner-service-api/pkg/banner-api"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
	"go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/trace/embedded"
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

	tracer := NewMockTracer()

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

type mockTracer struct {
	embedded.Tracer
}

func NewMockTracer() mockTracer {
	return mockTracer{}
}

func (m mockTracer) Start(ctx context.Context, spanName string, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
	return ctx, NewMockSpan()
}

type mockSpan struct {
	embedded.Span
}

func NewMockSpan() mockSpan {
	return mockSpan{}
}

func (ms mockSpan) End(options ...trace.SpanEndOption) {

}

// AddEvent adds an event with the provided name and options.
func (ms mockSpan) AddEvent(name string, options ...trace.EventOption) {

}

// AddLink adds a link.
// Adding links at span creation using WithLinks is preferred to calling AddLink
// later, for contexts that are available during span creation, because head
// sampling decisions can only consider information present during span creation.
func (ms mockSpan) AddLink(link trace.Link) {

}

// IsRecording returns the recording state of the Span. It will return
// true if the Span is active and events can be recorded.
func (ms mockSpan) IsRecording() bool {
	return true
}

// RecordError will record err as an exception span event for this span. An
// additional call to SetStatus is required if the Status of the Span should
// be set to Error, as this method does not change the Span status. If this
// span is not being recorded or err is nil then this method does nothing.
func (ms mockSpan) RecordError(err error, options ...trace.EventOption) {

}

// SpanContext returns the SpanContext of the Span. The returned SpanContext
// is usable even after the End method has been called for the Span.
func (ms mockSpan) SpanContext() trace.SpanContext {
	return trace.NewSpanContext(trace.SpanContextConfig{})
}

// SetStatus sets the status of the Span in the form of a code and a
// description, provided the status hasn't already been set to a higher
// value before (OK > Error > Unset). The description is only included in a
// status when the code is for an error.
func (ms mockSpan) SetStatus(code codes.Code, description string) {

}

// SetName sets the Span name.
func (ms mockSpan) SetName(name string) {

}

// SetAttributes sets kv as attributes of the Span. If a key from kv
// already exists for an attribute of the Span it will be overwritten with
// the value contained in kv.
func (ms mockSpan) SetAttributes(kv ...attribute.KeyValue) {

}

// TracerProvider returns a TracerProvider that can be used to generate
// additional Spans on the same telemetry pipeline as the current Span.
func (ms mockSpan) TracerProvider() trace.TracerProvider {
	mockTraceExporter := tracetest.NewInMemoryExporter()

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.TraceIDRatioBased(0.01)),
		sdktrace.WithSyncer(mockTraceExporter),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(""),
		)),
	)

	otel.SetTracerProvider(tp)
	return tp
}
