package cache

import (
	"context"
	"log/slog"
	"time"

	"github.com/go-faster/jx"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"

	"github.com/nikitads9/banner-service-api/internal/app/repository/banner"
	"github.com/nikitads9/banner-service-api/internal/logger/sl"
)

var ErrCacheMiss = errors.New("cache miss")

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
	const op = "cache.banners.Get"

	log := c.log.With(
		slog.String("op", op),
	)

	var content []byte

	err := c.client.Get(ctx, key).Scan(&content)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			err = ErrCacheMiss
		}
		log.Error("get cached content failed", sl.Err(err))
		return nil, errors.Wrapf(err, "get cached content failed with key: %s", key)
	}

	return content, nil
}
func (c *cache) Set(ctx context.Context, key string, content []byte, ttl time.Duration) error {
	const op = "cache.banners.Set"

	log := c.log.With(
		slog.String("op", op),
	)

	if err := c.client.Set(ctx, key, content, ttl).Err(); err != nil {
		log.Error("set cache content failed", sl.Err(err))
		return errors.Wrapf(err, "set cached content failed with key: %s", key)
	}

	return nil
}
