package integration

import (
	"context"
	"fmt"
	"testing"

	"github.com/meraiku/db_tasks/internal/models"
	"github.com/meraiku/db_tasks/internal/repo/orm"
	"github.com/stretchr/testify/require"
)

func TestORM(t *testing.T) {
	db, err := orm.New()

	require.NoError(t, err)

	ctx := context.TODO()

	dummy := &models.Employee{
		FirstName: "John",
		LastName:  "Doe",
	}

	err = db.Create(ctx, dummy)

	require.NoError(t, err)

	e, err := db.Read(ctx)

	require.NotNil(t, e)

	fmt.Println(e)

	dummy.ID = 5
	dummy.FirstName = "Boris"

	updatedEmployee, err := db.Update(ctx, dummy)

	require.Equal(t, dummy, updatedEmployee)

	err = db.Delete(ctx, dummy.ID)

	require.NoError(t, err)

	e, _ = db.Read(ctx)

	require.NotContains(t, e, *dummy)
	require.Len(t, e, 4)
}
