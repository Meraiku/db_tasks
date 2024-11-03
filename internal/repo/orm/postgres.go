package orm

import (
	"database/sql"
	"os"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
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

	return &DB{
		db: db,
	}, nil
}
