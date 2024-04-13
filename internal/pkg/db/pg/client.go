package pg

import (
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Client interface {
	Close() error
	DB() *pg
}

type client struct {
	db        *pg
	logger    *slog.Logger
	closeFunc context.CancelFunc
}

func NewClient(ctx context.Context, logger *slog.Logger, config *pgxpool.Config) (Client, error) {
	dbc, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	_, cancel := context.WithCancel(ctx)

	return &client{
		db: &pg{pool: dbc,
			logger: logger},
		logger:    logger,
		closeFunc: cancel,
	}, nil
}

func (c *client) Close() error {
	if c != nil {
		if c.closeFunc != nil {
			c.closeFunc()
		}
	}

	if c.db != nil {
		c.db.pool.Close()
	}

	return nil
}

func (c *client) DB() *pg {
	return c.db
}
