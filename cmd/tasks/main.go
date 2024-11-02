package main

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/meraiku/db_tasks/internal/repo/postgres_sqlc"
)

func main() {
	godotenv.Load()
	ctx := context.TODO()
	pool, err := pgxpool.New(ctx, os.Getenv("POSTGRES_DSN"))
	if err != nil {
		panic(err)
	}
	defer pool.Close()

	queries := postgres_sqlc.New(pool)

	_, err = queries.AddDepartment(ctx, "IT")
	if err != nil {
		panic(err)
	}
	queries.AddDepartment(ctx, "HR")
	queries.AddDepartment(ctx, "Sales")

	_, err = queries.AddEmployee(ctx, postgres_sqlc.AddEmployeeParams{
		FirstName:    "John",
		LastName:     "Doe",
		DepartmentID: 1,
	})
	if err != nil {
		panic(err)
	}

	queries.AddEmployee(ctx, postgres_sqlc.AddEmployeeParams{
		FirstName:    "Jane",
		LastName:     "Doe",
		DepartmentID: 1,
	})

	queries.AddEmployee(ctx, postgres_sqlc.AddEmployeeParams{
		FirstName:    "Joe",
		LastName:     "Doe",
		DepartmentID: 2,
	})

	queries.AddEmployee(ctx, postgres_sqlc.AddEmployeeParams{
		FirstName:    "Jill",
		LastName:     "Doe",
		DepartmentID: 3,
	})

	if err := queries.AddProject(ctx, "Project 1"); err != nil {
		panic(err)
	}
	queries.AddProject(ctx, "Project 2")
	queries.AddProject(ctx, "Project 3")

}
