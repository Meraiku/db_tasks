package integration

import (
	"context"
	"testing"

	postgressql "github.com/meraiku/db_tasks/internal/repo/postgres_sql"
	"github.com/stretchr/testify/require"
)

func TestTasks(t *testing.T) {

	db, err := postgressql.New(context.TODO())

	require.NoError(t, err)
	require.NotNil(t, db)

	tt := []struct {
		testName    string
		testFunc    func(ctx context.Context, args ...any) error
		expectError bool
		args        []any
	}{
		{
			testName:    "Выполнение SQL-запросово",
			testFunc:    db.FirstTask,
			expectError: false,
		},
		{
			testName:    "Использование параметров в SQL-запросе",
			testFunc:    db.ThirdTask,
			expectError: false,
			args:        []any{1},
		},
	}

	for _, tc := range tt {
		t.Run(tc.testName, func(t *testing.T) {
			err := tc.testFunc(context.TODO(), tc.args...)
			if tc.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}

}
