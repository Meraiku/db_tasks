package repo

import "github.com/uptrace/bun"

type Employee struct {
	bun.BaseModel `bun:"table:employees,alias:e"`
	ID            int `bun:",pk,autoincrement"`
	FirstName     string
	LastName      string
}
