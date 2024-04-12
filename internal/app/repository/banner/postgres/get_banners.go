package postgres

import (
	"context"
	"errors"
	"log/slog"

	"github.com/jackc/pgx/v5"
	"github.com/nikitads9/banner-service-api/internal/app/model"
	"github.com/nikitads9/banner-service-api/internal/logger/sl"
	"github.com/nikitads9/banner-service-api/internal/pkg/db"
)

func (r *repository) GetBanners(ctx context.Context, mod *model.GetBannersParams) ([]*model.BannerInfo, error) {
	const op = "banner.postgres.GetBanners"

	log := r.log.With(
		slog.String("op", op),
	)
	query := `SELECT banner_id, feature_id, array_agg(tag_id) as tag_ids, content, is_active, updated_at, created_at 
		FROM banners_tags 
		JOIN banners 
			ON banners.id=banners_tags.banner_id 
			AND 
			banner_id IN 
				(SELECT DISTINCT banner_id 
					FROM banners_tags 
					WHERE (feature_id = $1  OR $1 IS NULL) 
					AND 
					(tag_id = $2 OR $2 IS NULL)) 
					GROUP BY banners_tags.banner_id, banners_tags.feature_id, banners.content, banners.is_active, banners.updated_at, banners.created_at LIMIT $3 OFFSET $4;`

	var args = []interface{}{mod.FeatureID, mod.TagID, mod.Limit, mod.Offset}

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}

	var banners []*model.BannerInfo
	err := r.client.DB().SelectContext(ctx, &banners, q, args...)
	if err != nil {
		if errors.As(err, pgNoConnection) {
			log.Error("no connection to database host", sl.Err(err))
			return nil, ErrNoConnection
		}
		if errors.Is(err, pgx.ErrNoRows) {
			log.Error("banners with these criteria not found", sl.Err(err))
			return nil, ErrNotFound
		}
		log.Error("query execution error", sl.Err(err))
		return nil, ErrQuery
	}

	return banners, nil
}
