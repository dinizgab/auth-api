package database

import (
	"auth-api/internal/config"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Database interface {
}
type databaseImpl struct {
	db *pgxpool.Pool
}

func New(ctx context.Context, config config.DBConfig) (Database, error) {
	conn, err := pgxpool.New(ctx, config.Dsn)
	if err != nil {
		return nil, err
	}

	return &databaseImpl{
		db: conn,
	}, nil
}
