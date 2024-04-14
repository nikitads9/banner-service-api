package model

import (
	"database/sql"
	"time"

	"github.com/go-faster/jx"
)

// Banner model ...
type Banner struct {
	FeatureID int64
	TagIDs    []int64
	Content   jx.Raw
	IsActive  bool
}

// BannerInfo model ...
type BannerInfo struct {
	BannerID  int64        `db:"banner_id"`
	FeatureID int64        `db:"feature_id"`
	TagIDs    []int64      `db:"tag_ids"`
	Content   []byte       `db:"content"`
	IsActive  bool         `db:"is_active"`
	UpdatedAt sql.NullTime `db:"updated_at"`
	CreatedAt time.Time    `db:"created_at"`
}

// SetBannerInfo model ...
type SetBannerInfo struct {
	BannerID  int64
	FeatureID sql.NullInt64
	TagIDs    []int64
	Content   NullJxRaw
	IsActive  sql.NullBool
}

// GetBannersParams model ...
type GetBannersParams struct {
	FeatureID sql.NullInt64
	TagID     sql.NullInt64
	Limit     int64
	Offset    int64
}

// NullJxRaw model ...
type NullJxRaw struct {
	JxRaw jx.Raw
	Valid bool
}
