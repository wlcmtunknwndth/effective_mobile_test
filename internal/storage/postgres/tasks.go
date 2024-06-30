package postgres

import (
	"context"
	"fmt"
	"github.com/wlcmtunknwndth/effective_mobile_test/internal/domain/models"
	"gorm.io/gorm/clause"
)

func (s *Storage) CreateTask(ctx context.Context, task *models.Task) (uint64, error) {
	const op = scope + "CreateTask"

	if res := s.db.Model(&models.Task{}).Clauses(
		clause.Returning{Columns: []clause.Column{{Name: "id"}}}).
		Create(task); res.Error != nil {
		return 0, fmt.Errorf("%s: %w", op, res.Error)
	}

	return task.ID, nil
}

func (s *Storage) GetWorkload(ctx context.Context, passportNumber string) ([]models.Task, error) {
	const op = scope + "GetWorkload"

	var tasks []models.Task
	s.db.Model(&models.Task{}).Order("expires_at desc").Where("user_id = ?")
}
