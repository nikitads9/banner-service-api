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
	BannerID  int64        `db:"id"`
	FeatureID int64        `db:"feature_id"`
	TagIDs    []int64      `db:"tag_id"`
	Content   jx.Raw       `db:"content"`
	IsActive  bool         `db:"is_active"`
	UpdatedAt sql.NullTime `db:"updated_at"`
	CreatedAt time.Time    `db:"created_at"`
}

type SetBannerInfo struct {
	BannerID  int64
	FeatureID sql.NullInt64
	TagIDs    []int64
	Content   []byte
	IsActive  sql.NullBool
}

type GetBannersParams struct {
	FeatureID sql.NullInt64
	TagID     sql.NullInt64
	Limit     int64
	Offset    int64
}
