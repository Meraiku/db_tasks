package orm

import (
	"context"

	"github.com/meraiku/db_tasks/internal/models"
	"github.com/meraiku/db_tasks/internal/repo"
)

func (db *DB) Create(ctx context.Context, employee *models.Employee) error {
	e := repo.FromModelToRepoEmployee(employee)

	_, err := db.db.NewInsert().Model(e).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) Read(ctx context.Context) ([]models.Employee, error) {
	e := make([]repo.Employee, 0)

	if err := db.db.NewSelect().Model(&e).Scan(ctx); err != nil {
		return nil, err
	}

	return repo.MultiToModelFromRepoEmployee(e), nil
}

func (db *DB) Update(ctx context.Context, employee *models.Employee) (*models.Employee, error) {
	e := repo.FromModelToRepoEmployee(employee)

	_, err := db.db.NewUpdate().Model(e).WherePK().Exec(ctx)
	if err != nil {
		return nil, err
	}

	return repo.ToModelFromRepoEmployee(e), nil
}

func (db *DB) Delete(ctx context.Context, id int) error {

	_, err := db.db.NewDelete().Table("employees").Where("id = ?", id).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
