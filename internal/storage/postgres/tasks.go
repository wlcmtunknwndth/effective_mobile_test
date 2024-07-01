package postgres

import (
	"context"
	"fmt"
	"github.com/wlcmtunknwndth/effective_mobile_test/internal/domain/models"
	"gorm.io/gorm/clause"
	"time"
)

func (s *Storage) CreateTask(ctx context.Context, task *models.Task) (uint64, error) {
	const op = scope + "CreateTask"

	if res := s.db.WithContext(ctx).Model(&models.Task{}).Clauses(
		clause.Returning{Columns: []clause.Column{{Name: "id"}}}).
		Create(&task); res.Error != nil {
		return 0, fmt.Errorf("%s: %w", op, res.Error)
	}

	return task.ID, nil
}

func (s *Storage) GetWorkload(ctx context.Context, userID uint64, start, end time.Time) ([]models.Task, error) {
	const op = scope + "GetWorkload"

	//s.db.Model(&models.User{}).Where(" = ?")
	var tasks []models.Task
	res := s.db.WithContext(ctx).Model(&models.Task{}).Where("user_id = ? AND created_at BETWEEN ? AND ?", userID, start, end).Order("hours_spent desc").Scan(&tasks)
	if res.Error != nil {
		return nil, fmt.Errorf("%s: %w", op, res.Error)
	}

	return tasks, nil
}
