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

// CopyFromContext ..
func (d *db) CopyFromContext(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error) {
	tx, ok := ContextTx(ctx)
	if ok {
		return tx.CopyFrom(
			ctx,
			tableName,
			columnNames,
			rowSrc,
		)
	}

	return d.pool.CopyFrom(
		ctx,
		tableName,
		columnNames,
		rowSrc,
	)
}

func (d *db) Close() {
	d.pool.Close()
}

// eject transaction from context
func ContextTx(ctx context.Context) (pgx.Tx, bool) {
	tx, ok := ctx.Value(TxKey).(pgx.Tx)
	if !ok {
		return nil, false
	}

	return tx, true
}

func (d *db) Ping(ctx context.Context) error {
	return d.pool.Ping(ctx)
}
