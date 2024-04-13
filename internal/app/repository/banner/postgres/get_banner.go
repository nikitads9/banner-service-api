package postgres

import (
	"context"
	"errors"
	"log/slog"

	sq "github.com/Masterminds/squirrel"
	"github.com/go-faster/jx"
	"github.com/jackc/pgx/v5"
	t "github.com/nikitads9/banner-service-api/internal/app/repository/banner/table"
	"github.com/nikitads9/banner-service-api/internal/logger/sl"
	"github.com/nikitads9/banner-service-api/internal/pkg/db"
	"go.opentelemetry.io/otel/codes"
)

func (r *repository) GetBanner(ctx context.Context, featureID int64, tagID int64) (jx.Raw, error) {
	const op = "banner.postgres.GetBanner"

	log := r.log.With(
		slog.String("op", op),
	)

	ctx, span := r.tracer.Start(ctx, op)
	defer span.End()

	builder := sq.Select(t.Content).
		From(t.BannerTable).Join(t.BannerTagTable + " ON " + t.BannerTable + "." + t.ID + "=" + t.BannerTagTable + "." + t.BannerID).
		Where(sq.And{
			sq.Eq{t.FeatureID: featureID},
			sq.Eq{t.TagID: tagID},
			sq.Eq{t.IsActive: true},
		}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error("failed to build a query", sl.Err(err))
		return nil, ErrQueryBuild
	}

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}

	var res []byte

	err = r.client.DB().GetContext(ctx, &res, q, args...)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		if errors.As(err, pgNoConnection) {
			log.Error("no connection to database host", sl.Err(err))
			return nil, ErrNoConnection
		}
		if errors.Is(err, pgx.ErrNoRows) {
			log.Error("banner with these tag and feature not found", sl.Err(err))
			return nil, ErrNotFound
		}
		log.Error("query execution error", err)
		return nil, ErrQuery
	}

	return res, nil
}
