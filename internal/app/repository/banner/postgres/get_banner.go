package postgres

import (
	"context"

	"github.com/go-faster/jx"
)

func (r *repository) GetBanner(ctx context.Context, featureID int64, tagID int64) (map[string]jx.Raw, error) {
	return map[string]jx.Raw{"title": []byte("\"some_title\""), "text": []byte("\"some_text\""), "url": []byte("\"some_url\"")}, nil
}
