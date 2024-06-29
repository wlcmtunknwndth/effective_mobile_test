package postgres

import (
	"fmt"
	_ "github.com/lib/pq"
	"github.com/wlcmtunknwndth/effective_mobile_test/internal/config"
	"github.com/wlcmtunknwndth/effective_mobile_test/internal/domain/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage struct {
	db *gorm.DB
}

const scope = "internal.storage.postgres."

func New(config *config.Database) (*Storage, error) {
	const op = scope + "New"

	connStr := fmt.Sprintf("postgres://%s:%s@postgres:%s/%s?sslmode=%s",
		config.DbUser, config.DbPass, config.Port,
		config.DbName, config.SslMode,
	)
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if err = db.AutoMigrate(&models.User{}, &models.Task{}, &models.Admin{}); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}