package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Client interface {
	Close() error
	DB() *db
}

type client struct {
	db        *db
	closeFunc context.CancelFunc
}

func NewClient(ctx context.Context, config *pgxpool.Config) (Client, error) {
	dbc, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	_, cancel := context.WithCancel(ctx)

	return &client{
		db: &db{
			pool: dbc,
		},
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

func (c *client) DB() *db {
	return c.db
}
