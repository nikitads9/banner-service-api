package postgres

import (
	"errors"
	"log/slog"

	"github.com/nikitads9/banner-service-api/internal/app/repository/banner"
	"github.com/nikitads9/banner-service-api/internal/pkg/db"

	"github.com/jackc/pgx/v5/pgconn"
)

var (
	ErrAlreadyExists  = errors.New("this banner already exists")
	ErrDuplicate      = "ERROR: duplicate key value violates unique constraint \"banners_tags_feature_id_tag_id_key\" (SQLSTATE 23505)"
	ErrNotFound       = errors.New("no banner with these feature id and tag id")
	ErrNoRowsAffected = errors.New("no database entries affected by this operation")

	ErrQuery        = errors.New("failed to execute query")
	ErrQueryBuild   = errors.New("failed to build query")
	ErrPgxScan      = errors.New("failed to read database response")
	ErrNoConnection = errors.New("could not connect to database")

	pgNoConnection = new(*pgconn.ConnectError)
)

type repository struct {
	client db.Client
	log    *slog.Logger
}

func NewBannerRepository(client db.Client, log *slog.Logger) banner.Repository {
	return &repository{
		client: client,
		log:    log,
	}
}
