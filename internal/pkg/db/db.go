package db

//go:generate mockgen --build_flags=--mod=mod -destination=mocks_db/mock_db.go -package=mocks . DB
//go:generate mockgen --build_flags=--mod=mod -destination=mocks_tx/mock_tx.go -package=mocks github.com/jackc/pgx/v5 Tx

import (
	"context"

	"github.com/georgysavva/scany/v2/pgxscan"
	pgx "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type key string

const (
	TxKey key = "tx"
)

type Query struct {
	Name     string
	QueryRaw string
}

type db struct {
	pool *pgxpool.Pool
}
type SQLExecer interface {
	NamedExecer
	QueryExecer
}

type NamedExecer interface {
	GetContext(ctx context.Context, dest interface{}, q Query, args ...interface{}) error
	SelectContext(ctx context.Context, dest interface{}, q Query, args ...interface{}) error
}

type QueryExecer interface {
	ExecContext(ctx context.Context, q Query, args ...interface{}) (pgconn.CommandTag, error)
	QueryContext(ctx context.Context, q Query, args ...interface{}) (pgx.Rows, error)
	QueryRowContext(ctx context.Context, q Query, args ...interface{}) pgx.Row
}

type Pinger interface {
	Ping(ctx context.Context) error
}

type DB interface {
	SQLExecer
	Transactor
	Pinger
	Close()
}

type Transactor interface {
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
}

type TxManager interface {
	ReadCommitted(ctx context.Context, f Handler) error
}

type Handler func(ctx context.Context) error

func (d *db) BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error) {
	return d.pool.BeginTx(ctx, txOptions)
}

func (d *db) GetContext(ctx context.Context, dest interface{}, q Query, args ...interface{}) error {
	return pgxscan.Get(ctx, d.pool, dest, q.QueryRaw, args...)
}

func (d *db) SelectContext(ctx context.Context, dest interface{}, q Query, args ...interface{}) error {
	return pgxscan.Select(ctx, d.pool, dest, q.QueryRaw, args...)
}

func (d *db) ExecContext(ctx context.Context, q Query, args ...interface{}) (pgconn.CommandTag, error) {
	return d.pool.Exec(ctx, q.QueryRaw, args...)
}

func (d *db) QueryContext(ctx context.Context, q Query, args ...interface{}) (pgx.Rows, error) {
	return d.pool.Query(ctx, q.QueryRaw, args...)
}

func (d *db) QueryRowContext(ctx context.Context, q Query, args ...interface{}) pgx.Row {
	return d.pool.QueryRow(ctx, q.QueryRaw, args...)
}

func (d *db) Close() {
	d.pool.Close()
}

func GetContextTx(ctx context.Context, tx pgx.Tx) context.Context {
	return context.WithValue(ctx, TxKey, tx)
}

func (d *db) Ping(ctx context.Context) error {
	return d.pool.Ping(ctx)
}
