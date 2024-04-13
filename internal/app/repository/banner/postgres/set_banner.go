package postgres

import (
	"context"
	"errors"
	"log/slog"
	"strings"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/nikitads9/banner-service-api/internal/app/model"
	t "github.com/nikitads9/banner-service-api/internal/app/repository/banner/table"
	"github.com/nikitads9/banner-service-api/internal/logger/sl"
	"github.com/nikitads9/banner-service-api/internal/pkg/db"
	"go.opentelemetry.io/otel/codes"
)

func (r *repository) SetBanner(ctx context.Context, mod *model.SetBannerInfo) error {
	const op = "banner.postgres.SetBanner"

	log := r.log.With(
		slog.String("op", op),
	)

	ctx, span := r.tracer.Start(ctx, op)
	defer span.End()

	update := sq.Update(t.BannerTable).
		Set(t.UpdatedAt, time.Now()).
		Where(sq.And{
			sq.Eq{t.ID: mod.BannerID},
		}).
		PlaceholderFormat(sq.Dollar)

	if mod.Content.Valid {
		update = update.Set(t.Content, mod.Content.JxRaw)
	}

	if mod.IsActive.Valid {
		update = update.Set(t.IsActive, mod.IsActive.Bool)
	}

	query, args, err := update.ToSql()
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error("failed to build a query", sl.Err(err))
		return ErrQueryBuild
	}

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}

	if err := r.errorHandler(r.client.DB().ExecContext(ctx, q, args...)); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error("update banner content failed", sl.Err(err))
		return err
	}

	return r.SetBannerInfo(ctx, mod)
}

func (r *repository) SetBannerInfo(ctx context.Context, mod *model.SetBannerInfo) error {
	const op = "banner.postgres.SetBannerInfo"

	log := r.log.With(
		slog.String("op", op),
	)

	ctx, span := r.tracer.Start(ctx, op)
	defer span.End()

	switch {
	case mod.FeatureID.Valid && mod.TagIDs == nil:
		update := sq.Update(t.BannerTagTable).
			Set(t.FeatureID, mod.FeatureID).
			Where(sq.And{
				sq.Eq{t.ID: mod.BannerID},
			}).
			PlaceholderFormat(sq.Dollar)

		query, args, err := update.ToSql()
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			log.Error("failed to build a query", sl.Err(err))
			return ErrQueryBuild
		}

		q := db.Query{
			Name:     op,
			QueryRaw: query,
		}

		if err := r.errorHandler(r.client.DB().ExecContext(ctx, q, args...)); err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			log.Error("failed to update feature id", sl.Err(err))
			return err
		}

		return nil

	case mod.TagIDs != nil:
		delete := sq.Delete(t.BannerTagTable).Where(sq.And{
			sq.Eq{t.BannerID: mod.BannerID},
		}).
			Suffix("returning " + t.FeatureID).
			PlaceholderFormat(sq.Dollar)

		query, args, err := delete.ToSql()
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			log.Error("failed to build a query", sl.Err(err))
			return ErrQueryBuild
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
				return ErrNoConnection
			}
			log.Error("query execution error", sl.Err(err))
			return ErrQuery
		}

		defer row.Close()

		var featureID int64

		row.Next()
		err = row.Scan(&featureID)
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			log.Error("failed to scan returning id", sl.Err(err))
			return ErrPgxScan
		}

		if mod.FeatureID.Valid {
			featureID = mod.FeatureID.Int64
		}

		row.Close()
		err = r.LinkBannerTags(ctx, mod.BannerID, featureID, mod.TagIDs)
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			if strings.EqualFold(err.Error(), ErrDuplicate) {
				log.Error("this banner already exists", sl.Err(err))
				return ErrAlreadyExists
			}
			log.Error("failed to link banner and tags", sl.Err(err))
			return err
		}

	}

	return nil
}

func (r *repository) errorHandler(result pgconn.CommandTag, err error) error {
	const op = "banner.postgres.errorHandler"

	log := r.log.With(
		slog.String("op", op),
	)

	if err != nil {
		if errors.As(err, pgNoConnection) {
			log.Error("no connection to database host", sl.Err(err))
			return ErrNoConnection
		}
		log.Error("query execution error", sl.Err(err))
		return ErrQuery
	}

	if result.RowsAffected() == 0 {
		log.Error("unsuccessful update", sl.Err(ErrNoRowsAffected))
		return ErrNotFound
	}

	return nil
}
