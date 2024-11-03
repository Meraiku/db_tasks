package main

import (
	"context"
	"log"

	"github.com/joho/godotenv"
	"github.com/meraiku/db_tasks/internal/helpers"
	postgressql "github.com/meraiku/db_tasks/internal/repo/postgres_sql"
	"github.com/meraiku/db_tasks/internal/service/testing"
)

func main() {
	godotenv.Load()
	ctx := context.TODO()

	if err := helpers.ResetDB(ctx); err != nil {
		log.Fatalf("failed to reset DB: %v", err)
	}

	if err := helpers.InsertData(ctx); err != nil {
		log.Fatalf("failed to insert data in DB: %v", err)
	}

	db, err := postgressql.New(ctx)
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}

	ts := testing.New(db)
	if err := ts.Test(ctx); err != nil {
		log.Fatalf("failed to test service: %v", err)
	}

}
