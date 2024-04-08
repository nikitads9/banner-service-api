package model

import (
	"database/sql"
	"time"

	"github.com/go-faster/jx"
)

type Banner struct {
	FeatureID int64
	TagIDs    []int64
	Content   []byte
	IsActive  bool
}

type BannerInfo struct {
	BannerID  int64
	FeatureID int64
	TagIDs    []int64
	Content   map[string]jx.Raw
	IsActive  bool
	UpdatedAt sql.NullTime
	CreatedAt time.Time
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
