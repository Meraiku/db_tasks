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
			testFunc:    db.SelectTask,
			expectError: false,
		},
		{
			testName:    "Использование параметров в SQL-запросе",
			testFunc:    db.ParamsTask,
			expectError: false,
			args:        []any{1},
		},
		{
			testName:    "Неправильные параметры в SQL-запросе",
			testFunc:    db.ParamsTask,
			expectError: true,
			args:        []any{-1},
		},
		{
			testName:    "Вставка данных в базу",
			testFunc:    db.InsertTask,
			expectError: false,
			args:        []any{"John", "Doe", 1},
		},
		{
			testName:    "Обновление данных в базе",
			testFunc:    db.UpdateTask,
			expectError: false,
			args:        []any{"John", "Doe", 1},
		},
		{
			testName:    "Транзакция",
			testFunc:    db.TransactionTask,
			expectError: false,
			args:        []any{"IT", "John", "Doe"},
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
