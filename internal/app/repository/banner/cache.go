package banner

import (
	"context"
	"time"

	"github.com/go-faster/jx"
)

type Cache interface {
	Get(ctx context.Context, key string) (jx.Raw, error)
	Set(ctx context.Context, key string, content []byte, ttl time.Duration) error
}
