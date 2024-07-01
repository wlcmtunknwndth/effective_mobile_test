package tests

import (
	"context"
	"database/sql"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wlcmtunknwndth/effective_mobile_test/internal/domain/models"
	"github.com/wlcmtunknwndth/effective_mobile_test/internal/storage/postgres"
	"strconv"
	"testing"
)

var tasksCreated = 1

var casesTasks = []models.Task{
	{
		Model:       models.Model{},
		UserID:      uint64(gofakeit.IntRange(100, 200)),
		HoursSpent:  gofakeit.Float64Range(2, 200),
		Description: gofakeit.Name(),
		CreatedAt:   gofakeit.PastDate(),
		DoneAt:      gofakeit.Date(),
		ExpiresAt:   gofakeit.FutureDate(),
	},
	{
		Model:       models.Model{},
		UserID:      uint64(gofakeit.IntRange(200, 300)),
		HoursSpent:  gofakeit.Float64Range(2, 200),
		Description: gofakeit.Name(),
		CreatedAt:   gofakeit.PastDate(),
		DoneAt:      gofakeit.Date(),
		ExpiresAt:   gofakeit.FutureDate(),
	},
	{
		Model:       models.Model{},
		UserID:      uint64(gofakeit.IntRange(100, 270)),
		HoursSpent:  gofakeit.Float64Range(2, 200),
		Description: gofakeit.Name(),
		CreatedAt:   gofakeit.PastDate(),
		DoneAt:      gofakeit.Date(),
		ExpiresAt:   gofakeit.FutureDate(),
	},
	{
		Model:       models.Model{},
		UserID:      uint64(gofakeit.IntRange(100, 900)),
		HoursSpent:  gofakeit.Float64Range(2, 200),
		Description: gofakeit.Name(),
		CreatedAt:   gofakeit.PastDate(),
		DoneAt:      gofakeit.Date(),
		ExpiresAt:   gofakeit.FutureDate(),
	},
	{
		Model:       models.Model{},
		UserID:      uint64(gofakeit.IntRange(100, 600)),
		HoursSpent:  gofakeit.Float64Range(2, 200),
		Description: gofakeit.Name(),
		CreatedAt:   gofakeit.PastDate(),
		DoneAt:      gofakeit.Date(),
		ExpiresAt:   gofakeit.FutureDate(),
	},
}

func TestCleanUpTasks(t *testing.T) {
	db, err := connect(&cfg)
	require.NoError(t, err)

	err = dropLocalTasksTable(db)
	require.NoError(t, err)
}

func TestCreateTask(t *testing.T) {
	db, err := postgres.New(&cfg)
	require.NoError(t, err)
	var ctx context.Context
	for i, _ := range casesTasks {
		t.Run(strconv.FormatUint(casesTasks[i].UserID, 10), func(t *testing.T) {
			id, err := db.CreateTask(ctx, &casesTasks[i])
			require.NoError(t, err)

			assert.Equal(t, id, uint64(tasksCreated))
			tasksCreated++
		})
	}
}

var userIdTest uint64 = 1

var userTasks = []models.Task{
	{
		Model:       models.Model{},
		UserID:      userIdTest,
		HoursSpent:  gofakeit.Float64Range(2, 200),
		Description: gofakeit.Name(),
		CreatedAt:   gofakeit.PastDate(),
		DoneAt:      gofakeit.Date(),
		ExpiresAt:   gofakeit.FutureDate(),
	},
	{
		Model:       models.Model{},
		UserID:      userIdTest,
		HoursSpent:  gofakeit.Float64Range(2, 200),
		Description: gofakeit.Name(),
		CreatedAt:   gofakeit.PastDate(),
		DoneAt:      gofakeit.Date(),
		ExpiresAt:   gofakeit.FutureDate(),
	},
	{
		Model:  models.Model{},
		UserID: userIdTest,
		//HoursSpent:  gofakeit.Float64Range(2, 200),
		Description: gofakeit.Name(),
		CreatedAt:   gofakeit.PastDate(),
		//DoneAt:      gofakeit.Date(),
		ExpiresAt: gofakeit.FutureDate(),
	},
}

func TestCreateTaskForUser(t *testing.T) {
	db, err := postgres.New(&cfg)
	require.NoError(t, err)
	var ctx context.Context
	for i, _ := range userTasks {
		t.Run(strconv.FormatUint(userTasks[i].UserID, 10), func(t *testing.T) {
			id, err := db.CreateTask(ctx, &userTasks[i])
			require.NoError(t, err)

			assert.Equal(t, id, uint64(i+9))
		})
	}
}

func TestGetTasksWorkload(t *testing.T) {
	db, err := postgres.New(&cfg)
	require.NoError(t, err)

	var ctx context.Context

	tasks, err := db.GetWorkload(ctx, userIdTest, gofakeit.PastDate(), gofakeit.FutureDate())
	require.NoError(t, err)

	//assert.Equal(t, tasks, userTasks)
	t.Logf("%+v", tasks)
}

func dropLocalTasksTable(db *sql.DB) error {
	if _, err := db.Exec("DROP TABLE IF EXISTS tasks"); err != nil {
		return err
	}
	return nil
}
