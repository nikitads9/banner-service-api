package db

//go:generate mockgen --build_flags=--mod=mod -destination=mocks_db/db_mock.go -package=mocks . DB
//go:generate mockgen --build_flags=--mod=mod -destination=mocks_tx/tx_mock.go -package=mocks github.com/jackc/pgx/v5 Tx

import (
	"context"

	pgx "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type Query struct {
	Name     string
	QueryRaw string
}

type SQLExecutor interface {
	NamedExecutor
	QueryExecutor
	CopyExecutor
}

type NamedExecutor interface {
	GetContext(ctx context.Context, dest interface{}, q Query, args ...interface{}) error
	SelectContext(ctx context.Context, dest interface{}, q Query, args ...interface{}) error
}

type QueryExecutor interface {
	ExecContext(ctx context.Context, q Query, args ...interface{}) (pgconn.CommandTag, error)
	QueryContext(ctx context.Context, q Query, args ...interface{}) (pgx.Rows, error)
	QueryRowContext(ctx context.Context, q Query, args ...interface{}) pgx.Row
}

type Pinger interface {
	Ping(ctx context.Context) error
}

type DB interface {
	SQLExecutor
	Transactor
	Pinger
	Close()
}

type Transactor interface {
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
}

type CopyExecutor interface {
	CopyFromContext(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error)
}

type TxManager interface {
	ReadCommitted(ctx context.Context, f Handler) error
}

type Handler func(ctx context.Context) error
