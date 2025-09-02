package database

import (
	"auth-api/internal/config"
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Database interface {
	Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
}
type databaseImpl struct {
	conn *pgxpool.Pool
}

func New(ctx context.Context, config config.DBConfig) (Database, error) {
	conn, err := pgxpool.New(ctx, config.Dsn)
	if err != nil {
		return nil, err
	}

	return &databaseImpl{
		conn: conn,
	}, nil
}

func (d *databaseImpl) Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error) {
	return d.conn.Exec(ctx, sql, arguments...)
}

func (d *databaseImpl) Query(ctx context.Context, sql string, arguments ...any) (pgx.Rows, error) {
	return d.conn.Query(ctx, sql, arguments...)
}

func (d *databaseImpl) QueryRow(ctx context.Context, sql string, arguments ...any) pgx.Row {
	return d.conn.QueryRow(ctx, sql, arguments...)
}
