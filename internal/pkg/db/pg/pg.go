package pg

import (
	"context"
	"log/slog"

	"github.com/georgysavva/scany/v2/pgxscan"
	pgx "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nikitads9/banner-service-api/internal/pkg/db"
)

type key string

const (
	TxKey key = "tx"
)

type pg struct {
	pool   *pgxpool.Pool
	logger *slog.Logger
}

// NewDB ...
func NewDB(dbc *pgxpool.Pool, logger *slog.Logger) db.DB {
	return &pg{
		pool:   dbc,
		logger: logger,
	}
}

func (p *pg) GetContext(ctx context.Context, dest interface{}, q db.Query, args ...interface{}) error {
	return pgxscan.Get(ctx, p.pool, dest, q.QueryRaw, args...)
}

func (p *pg) SelectContext(ctx context.Context, dest interface{}, q db.Query, args ...interface{}) error {
	return pgxscan.Select(ctx, p.pool, dest, q.QueryRaw, args...)
}

func (p *pg) ExecContext(ctx context.Context, q db.Query, args ...interface{}) (pgconn.CommandTag, error) {
	logQuery(ctx, p.logger, q)

	tx, ok := ContextTx(ctx)
	if ok {
		return tx.Exec(ctx, q.QueryRaw, args...)
	}

	return p.pool.Exec(ctx, q.QueryRaw, args...)
}

func (p *pg) QueryContext(ctx context.Context, q db.Query, args ...interface{}) (pgx.Rows, error) {
	logQuery(ctx, p.logger, q)

	tx, ok := ContextTx(ctx)
	if ok {
		return tx.Query(ctx, q.QueryRaw, args...)
	}

	return p.pool.Query(ctx, q.QueryRaw, args...)
}

func (p *pg) QueryRowContext(ctx context.Context, q db.Query, args ...interface{}) pgx.Row {
	logQuery(ctx, p.logger, q)

	tx, ok := ContextTx(ctx)
	if ok {
		return tx.QueryRow(ctx, q.QueryRaw, args...)
	}

	return p.pool.QueryRow(ctx, q.QueryRaw, args...)
}

// CopyFromContext ..
func (p *pg) CopyFromContext(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error) {
	tx, ok := ContextTx(ctx)
	if ok {
		return tx.CopyFrom(
			ctx,
			tableName,
			columnNames,
			rowSrc,
		)
	}

	return p.pool.CopyFrom(
		ctx,
		tableName,
		columnNames,
		rowSrc,
	)
}

func (p *pg) BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error) {
	return p.pool.BeginTx(ctx, txOptions)
}

// Ping ..
func (p *pg) Ping(ctx context.Context) error {
	return p.pool.Ping(ctx)
}

// Close ..
func (p *pg) Close() {
	p.pool.Close()
}

// MakeContextTx add transaction in context
func MakeContextTx(ctx context.Context, tx pgx.Tx) context.Context {
	return context.WithValue(ctx, TxKey, tx)
}

// ContextTx eject transaction from context
func ContextTx(ctx context.Context) (pgx.Tx, bool) {
	tx, ok := ctx.Value(TxKey).(pgx.Tx)
	if !ok {
		return nil, false
	}

	return tx, true
}

func logQuery(_ context.Context, logger *slog.Logger, q db.Query) {
	logger.Debug("logged query", slog.String("sql", q.Name), slog.String("query", q.QueryRaw))
}
