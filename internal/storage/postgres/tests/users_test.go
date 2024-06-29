package tests

import (
	"context"
	"fmt"
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

func TestCreateUser(t *testing.T) {
	db, err := postgres.New(&cfg)
	if err != nil {
		fmt.Println(err)
		return
	}
	cases := []models.User{
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
	for i, tc := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var ctx context.Context
			id, err := db.CreateUser(ctx, &tc)

			require.NoError(t, err)

			assert.Equal(t, id, uint64(i))
		})
	}
}
