package cache

import (
	"context"
	"log/slog"
	"time"

	"github.com/go-faster/jx"
	"github.com/go-redis/redis"

	"github.com/nikitads9/banner-service-api/internal/app/repository/banner"
)

type cache struct {
	client *redis.Client
	log    *slog.Logger
}

func NewBannerCache(client *redis.Client, log *slog.Logger) banner.Cache {
	return &cache{
		client: client,
		log:    log,
	}
}

func (c *cache) Get(ctx context.Context, key string) (jx.Raw, error) {
	return nil, nil
}
func (c *cache) Set(ctx context.Context, key string, content jx.Raw, ttl time.Duration) error {
	return nil
}
