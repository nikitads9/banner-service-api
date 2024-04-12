package postgres

import (
	"context"
	"errors"
	"log/slog"

	sq "github.com/Masterminds/squirrel"
	t "github.com/nikitads9/banner-service-api/internal/app/repository/banner/table"
	"github.com/nikitads9/banner-service-api/internal/logger/sl"
	"github.com/nikitads9/banner-service-api/internal/pkg/db"
)

func (r *repository) DeleteBanner(ctx context.Context, bannerID int64) error {
	const op = "banner.postgres.DeleteBanner"

	log := r.log.With(
		slog.String("op", op),
	)

	builder := sq.Delete(t.BannerTable).
		Where(
			sq.Eq{t.ID: bannerID},
		).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		log.Error("failed to build a query", sl.Err(err))
		return ErrQueryBuild
	}

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}

	result, err := r.client.DB().ExecContext(ctx, q, args...)
	if err != nil {
		if errors.As(err, pgNoConnection) {
			log.Error("no connection to database host", sl.Err(err))
			return ErrNoConnection
		}
		log.Error("query execution error", sl.Err(err))
		return ErrQuery
	}

	if result.RowsAffected() == 0 {
		log.Error("unsuccessful delete", sl.Err(ErrNoRowsAffected))
		return ErrNotFound
	}

	return nil
}
