// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: reset.sql

package postgres_sqlc

import (
	"context"
)

const resetDB = `-- name: ResetDB :exec
DELETE FROM employees
`

func (q *Queries) ResetDB(ctx context.Context) error {
	_, err := q.db.Exec(ctx, resetDB)
	return err
}
