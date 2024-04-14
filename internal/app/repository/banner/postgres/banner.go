package postgres

import (
	"errors"
	"log/slog"

	"github.com/nikitads9/banner-service-api/internal/app/repository/banner"
	"github.com/nikitads9/banner-service-api/internal/pkg/db/pg"
	"go.opentelemetry.io/otel/trace"

	"github.com/jackc/pgx/v5/pgconn"
)

var (
	// ErrAlreadyExists ошибка ввода дублирующих баннеров
	ErrAlreadyExists = errors.New("this banner already exists")
	errDuplicate     = "ERROR: duplicate key value violates unique constraint \"banners_tags_feature_id_tag_id_key\" (SQLSTATE 23505)"
	// ErrNotFound ошибка баннер не найден
	ErrNotFound       = errors.New("banner not found")
	errNoRowsAffected = errors.New("no database entries affected by this operation")

	errQuery        = errors.New("failed to execute query")
	errQueryBuild   = errors.New("failed to build query")
	errPgxScan      = errors.New("failed to read database response")
	errNoConnection = errors.New("could not connect to database")

	pgNoConnection = new(*pgconn.ConnectError)
)

type repository struct {
	client pg.Client
	tracer trace.Tracer
	log    *slog.Logger
}

// NewBannerRepository ...
func NewBannerRepository(client pg.Client, tracer trace.Tracer, log *slog.Logger) banner.Repository {
	return &repository{
		client: client,
		tracer: tracer,
		log:    log,
	}
}
