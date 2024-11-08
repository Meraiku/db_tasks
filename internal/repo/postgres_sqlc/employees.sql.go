// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: employees.sql

package postgres_sqlc

import (
	"context"
)

const addEmployee = `-- name: AddEmployee :one
INSERT INTO employees(first_name, last_name, department_id) 
VALUES($1, $2, $3)
RETURNING id
`

type AddEmployeeParams struct {
	FirstName    string
	LastName     string
	DepartmentID int32
}

func (q *Queries) AddEmployee(ctx context.Context, arg AddEmployeeParams) (int32, error) {
	row := q.db.QueryRow(ctx, addEmployee, arg.FirstName, arg.LastName, arg.DepartmentID)
	var id int32
	err := row.Scan(&id)
	return id, err
}
