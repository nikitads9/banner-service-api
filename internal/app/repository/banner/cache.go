package banner

//go:generate mockgen --build_flags=--mod=mod -destination=./mocks/banner_service_cache.go -package=mocks . Cache

import (
	"context"
	"time"

	"github.com/go-faster/jx"
)

type Cache interface {
	Get(ctx context.Context, key string) (jx.Raw, error)
	Set(ctx context.Context, key string, content []byte, ttl time.Duration) error
}
