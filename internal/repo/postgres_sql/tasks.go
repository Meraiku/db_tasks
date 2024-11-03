package postgressql

import (
	"context"
	"fmt"

	"github.com/meraiku/db_tasks/internal/repo"
)

func (db *DB) FirstTask(ctx context.Context, args ...any) error {
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

func (db *DB) ThirdTask(ctx context.Context, args ...any) error {
	query := `SELECT first_name, last_name
		FROM employees
		WHERE id = $1`

	row := db.db.QueryRowContext(ctx, query, args[0])

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
