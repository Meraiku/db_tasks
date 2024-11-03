package postgressql

import (
	"context"
	"fmt"

	"github.com/meraiku/db_tasks/internal/repo"
)

func (db *DB) SelectTask(ctx context.Context, args ...any) error {
	query := `SELECT id, first_name, last_name FROM employees`

	rows, err := db.db.QueryContext(ctx, query)
	if err != nil {
		return err
	}
	defer rows.Close()

	employees := make([]repo.Employee, 0)

	for rows.Next() {
		e := repo.Employee{}

		if err := rows.Scan(
			&e.ID,
			&e.FirstName,
			&e.LastName,
		); err != nil {
			return err
		}

		employees = append(employees, e)
	}

	if len(employees) == 0 {
		return fmt.Errorf("Expect at least one employee")
	}

	return nil
}

func (db *DB) ParamsTask(ctx context.Context, args ...any) error {
	query := `SELECT first_name, last_name
		FROM employees
		WHERE id = $1`

	row := db.db.QueryRowContext(ctx, query, args...)

	e := repo.Employee{}

	if err := row.Scan(
		&e.FirstName,
		&e.LastName,
	); err != nil {
		return err
	}

	if e.FirstName == "" {
		return fmt.Errorf("Expect employee with ID = %d, got nothing", args[0])
	}

	return nil
}

func (db *DB) InsertTask(ctx context.Context, args ...any) error {

	query := `INSERT INTO employees 
	(first_name, last_name, department_id) 
	VALUES ($1, $2, $3)`

	if _, err := db.db.ExecContext(ctx, query, args...); err != nil {
		return err
	}

	return nil
}

func (db *DB) UpdateTask(ctx context.Context, args ...any) error {

	query := `UPDATE employees SET first_name = $1, last_name = $2 WHERE id = $3`

	if _, err := db.db.ExecContext(ctx, query, args...); err != nil {
		return err
	}

	return nil
}

func (db *DB) TransactionTask(ctx context.Context, args ...any) error {

	if len(args) < 3 {
		return fmt.Errorf("Expect at least 3 arguments")
	}
	var err error

	tx, err := db.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	query := `SELECT id FROM departments WHERE title = $1`

	row := tx.QueryRowContext(ctx, query, args[0])

	var id int
	err = row.Scan(&id)
	if err != nil {
		return err
	}

	query = `INSERT INTO employees (first_name, last_name) VALUES ($1, $2)`

	_, err = tx.ExecContext(ctx, query, args[1], args[2])
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
