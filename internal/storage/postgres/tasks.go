package postgres

import (
	"context"
	"fmt"
	"github.com/wlcmtunknwndth/effective_mobile_test/internal/domain/models"
	"gorm.io/gorm/clause"
)

func (s *Storage) CreateTask(ctx context.Context, task *models.TaskDB) (uint64, error) {
	const op = scope + "CreateTask"

	if res := s.db.WithContext(ctx).Model(&models.TaskDB{}).Clauses(
		clause.Returning{Columns: []clause.Column{{Name: "id"}}}).
		Create(task); res.Error != nil {
		return 0, fmt.Errorf("%s: %w", op, res.Error)
	}

	return task.ID, nil
}

//func (s *Storage) GetWorkload(ctx context.Context, passportNumber string) ([]models.TaskDB, error) {
//	const op = scope + "GetWorkload"
//
//	s.db.Model(&models.UserDB{}).Where(" = ?")
//	var tasks []models.TaskDB
//	s.db.WithContext(ctx).Model(&models.TaskDB{}).Where("user_id = ?").Order("expires_at desc")
//}
