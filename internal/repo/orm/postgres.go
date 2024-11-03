package orm

import (
	"database/sql"
	"os"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

const (
	MAX_IDLE_CONNS    = 10
	MAX_CONNS         = 50
	MAX_CONN_LIFETIME = time.Minute * 5
)

type DB struct {
	db *bun.DB
}

func New() (*DB, error) {

	dsn := os.Getenv("POSTGRES_DSN")

	sqlDB := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	db := bun.NewDB(sqlDB, pgdialect.New())

	db.SetMaxIdleConns(MAX_IDLE_CONNS)
	db.SetMaxOpenConns(MAX_CONNS)
	db.SetConnMaxLifetime(MAX_CONN_LIFETIME)

	return &DB{
		db: db,
	}, nil
}
