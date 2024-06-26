package postgres

import (
	"context"
	"errors"
	"log/slog"
	"strings"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/nikitads9/banner-service-api/internal/app/model"
	t "github.com/nikitads9/banner-service-api/internal/app/repository/banner/table"
	"github.com/nikitads9/banner-service-api/internal/logger/sl"
	"github.com/nikitads9/banner-service-api/internal/pkg/db"
	"go.opentelemetry.io/otel/codes"
)

func (r *repository) CreateBanner(ctx context.Context, banner *model.Banner) (int64, error) {
	const op = "banner.postgres.CreateBanner"

	log := r.log.With(
		slog.String("op", op),
	)

	ctx, span := r.tracer.Start(ctx, op)
	defer span.End()

	builder := sq.Insert(t.BannerTable).
		Columns(t.Content, t.IsActive, t.CreatedAt).
		Values(banner.Content, banner.IsActive, time.Now())

	query, args, err := builder.PlaceholderFormat(sq.Dollar).Suffix("returning id").ToSql()
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error("failed to build a query", sl.Err(err))
		return 0, errQueryBuild
	}

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}

	row, err := r.client.DB().QueryContext(ctx, q, args...)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		if errors.As(err, pgNoConnection) {
			log.Error("no connection to database host", sl.Err(err))
			return 0, errNoConnection
		}

		log.Error("query execution error", sl.Err(err))
		return 0, errQuery
	}

	defer row.Close()

	var id int64
	row.Next()
	err = row.Scan(&id)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error("failed to scan returning id", sl.Err(err))
		return 0, errPgxScan
	}

	row.Close()

	err = r.LinkBannerTags(ctx, id, banner.FeatureID, banner.TagIDs)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		if strings.EqualFold(err.Error(), errDuplicate) {
			log.Error("this banner already exists", sl.Err(err))
			return 0, ErrAlreadyExists
		}
		log.Error("failed to link banner and tags", sl.Err(err))
		return 0, err
	}

	return id, nil
}

func (r *repository) LinkBannerTags(ctx context.Context, bannerID int64, featureID int64, tagIDs []int64) error {
	countTags := len(tagIDs)
	if countTags == 0 {
		return nil
	}

	rows := make([][]any, 0)

	for _, tagID := range tagIDs {
		rows = append(rows, []any{bannerID, featureID, tagID})
	}

	_, err := r.client.DB().CopyFromContext(
		ctx,
		pgx.Identifier{t.BannerTagTable},
		[]string{t.BannerID, t.FeatureID, t.TagID},
		pgx.CopyFromRows(rows),
	)

	if err != nil {
		return err
	}

	return nil
}
