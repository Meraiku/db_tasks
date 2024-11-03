package helpers

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/meraiku/db_tasks/internal/repo/postgres_sqlc"
)

func InsertData(ctx context.Context) error {

	pool, err := pgxpool.New(ctx, os.Getenv("POSTGRES_DSN"))
	if err != nil {
		return err
	}
	defer pool.Close()

	queries := postgres_sqlc.New(pool)

	_, err = queries.AddDepartment(ctx, "IT")
	if err != nil {
		return err
	}

	_, err = queries.AddDepartment(ctx, "HR")
	if err != nil {
		return err
	}

	_, err = queries.AddDepartment(ctx, "Sales")
	if err != nil {
		return err
	}

	_, err = queries.AddEmployee(ctx, postgres_sqlc.AddEmployeeParams{
		FirstName:    "John",
		LastName:     "Doe",
		DepartmentID: 1,
	})
	if err != nil {
		return err
	}

	_, err = queries.AddEmployee(ctx, postgres_sqlc.AddEmployeeParams{
		FirstName:    "Jane",
		LastName:     "Doe",
		DepartmentID: 1,
	})
	if err != nil {
		return err
	}

	_, err = queries.AddEmployee(ctx, postgres_sqlc.AddEmployeeParams{
		FirstName:    "Joe",
		LastName:     "Doe",
		DepartmentID: 2,
	})
	if err != nil {
		return err
	}

	_, err = queries.AddEmployee(ctx, postgres_sqlc.AddEmployeeParams{
		FirstName:    "Jill",
		LastName:     "Doe",
		DepartmentID: 3,
	})
	if err != nil {
		return err
	}

	if err := queries.AddProject(ctx, "Project 1"); err != nil {
		return err
	}
	if err := queries.AddProject(ctx, "Project 2"); err != nil {
		return err
	}
	if err := queries.AddProject(ctx, "Project 3"); err != nil {
		return err
	}

	return nil

}

func ResetDB(ctx context.Context) error {

	pool, err := pgxpool.New(ctx, os.Getenv("POSTGRES_DSN"))
	if err != nil {
		return err
	}
	defer pool.Close()

	queries := postgres_sqlc.New(pool)

	if err := queries.ResetDB(ctx); err != nil {
		return err
	}

	return nil
}
