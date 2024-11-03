package postgressql

import (
	"context"

	"github.com/meraiku/db_tasks/internal/repo"
)

func (db *DB) FirstTask(ctx context.Context) ([]repo.Employee, error) {
	query := `SELECT id, first_name, last_name FROM employees`

	rows, err := db.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
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
			return nil, err
		}

		employees = append(employees, e)
	}

	return employees, nil
}
