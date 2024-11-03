package testing

import (
	"context"

	"github.com/meraiku/db_tasks/internal/repo"
)

type TestingService struct {
	repo repo.Repository
}

func New(repo repo.Repository) *TestingService {
	return &TestingService{
		repo: repo,
	}
}

func (ts *TestingService) Test(ctx context.Context) error {
	employees, err := ts.repo.FirstTask(ctx)
	if err != nil {
		return err
	}
	for _, e := range employees {
		println(e.FirstName, e.LastName)
	}
	return nil
}
