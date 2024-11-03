package repo

import "context"

type Repository interface {
	FirstTask(ctx context.Context) ([]Employee, error)
}
