package model

import (
	"database/sql"
	"time"

	"github.com/go-faster/jx"
)

type Banner struct {
	FeatureID int64
	TagIDs    []int64
	Content   jx.Raw
	IsActive  bool
}

type BannerInfo struct {
	BannerID  int64        `db:"banner_id"`
	FeatureID int64        `db:"feature_id"`
	TagIDs    []int64      `db:"tag_ids"`
	Content   []byte       `db:"content"`
	IsActive  bool         `db:"is_active"`
	UpdatedAt sql.NullTime `db:"updated_at"`
	CreatedAt time.Time    `db:"created_at"`
}

type SetBannerInfo struct {
	BannerID  int64
	FeatureID sql.NullInt64
	TagIDs    []int64
	Content   NullJxRaw
	IsActive  sql.NullBool
}

type GetBannersParams struct {
	FeatureID sql.NullInt64
	TagID     sql.NullInt64
	Limit     int64
	Offset    int64
}

type NullJxRaw struct {
	JxRaw jx.Raw
	Valid bool
}
