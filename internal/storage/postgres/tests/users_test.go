package tests

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wlcmtunknwndth/effective_mobile_test/internal/config"
	"github.com/wlcmtunknwndth/effective_mobile_test/internal/domain/models"
	"github.com/wlcmtunknwndth/effective_mobile_test/internal/storage/postgres"
	"strconv"
	"testing"
)

var cfg = config.Database{
	DbUser:  "postgres",
	DbPass:  "postgres",
	DbName:  "postgres",
	SslMode: "disable",
	Port:    "5432",
}

var cases = []models.User{
	models.User{
		Model:          models.Model{},
		PassportNumber: "6617 899393",
		PassHash:       []byte{0, 1, 2, 3},
	},
	models.User{
		Model:          models.Model{},
		PassportNumber: "6617 834245",
		PassHash:       []byte{0, 1, 2, 3},
	},
	models.User{
		Model:          models.Model{},
		PassportNumber: "6617 891323",
		PassHash:       []byte{0, 1, 2, 3},
	},
}

func TestCreateUser(t *testing.T) {
	db, err := postgres.New(&cfg)
	require.NoError(t, err)

	for i, tc := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var ctx context.Context
			id, err := db.CreateUser(ctx, &tc)

			require.NoError(t, err)

			assert.Equal(t, id, uint64(i+1))
		})
	}
}

func TestGetUser(t *testing.T) {
	db, err := postgres.New(&cfg)

	require.NoError(t, err)

	for i, tc := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var ctx context.Context
			usr, err := db.GetUser(ctx, uint64(i+1))

			require.NoError(t, err)

			assert.Equal(t, tc.PassportNumber, usr.PassportNumber)
		})
	}
}

func TestUpdateUser(t *testing.T) {
	db, err := postgres.New(&cfg)
	require.NoError(t, err)

	for i, _ := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var ctx context.Context
			passport := gofakeit.IntRange(100000, 1000000)
			err := db.UpdateUser(ctx, &models.User{
				Model:          models.Model{ID: uint64(i + 1)},
				PassportNumber: strconv.Itoa(passport),
				PassHash:       nil,
			})
			require.NoError(t, err)

			usr, err := db.GetUser(ctx, uint64(i+1))
			require.NoError(t, err)

			assert.Equal(t, strconv.Itoa(passport), usr.PassportNumber)
		})
	}
	return
}

func TestGetUserByPassport(t *testing.T) {
	db, err := postgres.New(&cfg)
	require.NoError(t, err)

	for i, _ := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var ctx context.Context
			passport := gofakeit.IntRange(100000, 1000000)
			_, err := db.CreateUser(ctx, &models.User{
				Model:          models.Model{},
				PassportNumber: strconv.Itoa(passport),
				PassHash:       nil,
			})
			require.NoError(t, err)

			usr, err := db.GetUserByPassport(ctx, strconv.Itoa(passport))
			require.NoError(t, err)

			assert.Equal(t, strconv.Itoa(passport), usr.PassportNumber)
		})
	}
	return
}

func TestIsAdmin(t *testing.T) {

	return
}

func TestCleanUp(t *testing.T) {
	err := dropLocalUsersTable(cfg)
	require.NoError(t, err)
}

func dropLocalUsersTable(cfg *config.Database) error {
	db, err := connect(cfg)
	if err != nil {
		return err
	}

	if _, err := db.Exec("DROP TABLE IF EXISTS users"); err != nil {
		return err
	}
	return nil
}

func connect(config *config.Database) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=localhost user=%s password=%s "+
		"dbname=%s port=%s sslmode=%s",
		config.DbUser, config.DbPass, config.DbName,
		config.Port, config.SslMode,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func addSuperUser(cfg *config.Config)
