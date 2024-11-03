package postgressql

import (
	"context"
	"database/sql"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

const (
	MAX_IDLE_CONNS = 5
	MAX_CONNS      = 100
)

type DB struct {
	db *sql.DB
}

func New(ctx context.Context) (*DB, error) {
	dsn := os.Getenv("POSTGRES_DSN")

	sqlDb, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err := sqlDb.Ping(); err != nil {
		return nil, err
	}

	sqlDb.SetMaxIdleConns(MAX_IDLE_CONNS)
	sqlDb.SetMaxOpenConns(MAX_CONNS)

	return &DB{
		db: sqlDb,
	}, nil
}
