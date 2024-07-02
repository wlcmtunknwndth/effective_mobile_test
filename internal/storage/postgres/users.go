package postgres

import (
	"context"
	"errors"
	"fmt"
	"github.com/wlcmtunknwndth/effective_mobile_test/internal/domain/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (s *Storage) CreateUser(ctx context.Context, user *models.User) (uint64, error) {
	const op = scope + "CreateUser"

	if res := s.db.WithContext(ctx).Model(&models.User{}).Clauses(clause.Returning{Columns: []clause.Column{{Name: "id"}}}).Create(user); res.Error != nil {
		return 0, fmt.Errorf("%s: %w", op, res.Error)
	}

	return user.ID, nil
}

func (s *Storage) AddUserInfo(ctx context.Context, info *models.UserInfo) error {
	const op = scope + "AddUserInfo"

	if res := s.db.WithContext(ctx).Model(&models.UserInfo{}).Create(info); res.Error != nil {
		return fmt.Errorf("%s: %w", op, res.Error)
	}

	return nil
}

func (s *Storage) DeleteUser(ctx context.Context, id uint64) error {
	const op = scope + "DeleteUser"

	if res := s.db.WithContext(ctx).Delete(&models.User{}, id); res.Error != nil {
		return fmt.Errorf("%s: %w", op, res.Error)
	}

	if res := s.db.WithContext(ctx).Delete(&models.Admin{}, id); res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil
		}
		return fmt.Errorf("%s: %w", op, res.Error)
	}
	return nil
}

func (s *Storage) GetUser(ctx context.Context, id uint64) (*models.User, error) {
	const op = scope + "GetUser"

	var usr models.User
	usr.ID = id
	if res := s.db.WithContext(ctx).First(&usr); res.Error != nil {
		return nil, fmt.Errorf("%s: %w", op, res.Error)
	}

	return &usr, nil
}

func (s *Storage) UpdateUser(ctx context.Context, user *models.User) error {
	const op = scope + "UpdateUser"

	if res := s.db.WithContext(ctx).Model(user).Updates(&user); res.Error != nil {
		return fmt.Errorf("%s: %w", op, res.Error)
	}

	return nil
}

func (s *Storage) GetUserByPassport(ctx context.Context, passportSerie uint16, passportNumber uint32) (*models.User, error) {
	const op = scope + "GetUserByPassport"

	var usr models.User
	if res := s.db.WithContext(ctx).Model(&models.User{}).Where("passport_serie = ? AND passport_number = ?", passportSerie, passportNumber).First(&usr); res.Error != nil {
		return nil, fmt.Errorf("%s: %w", op, res.Error)
	}

	return &usr, nil
}

func (s *Storage) IsAdmin(ctx context.Context, id uint64) (bool, error) {
	const op = scope + "IsAdmin"

	var admin models.Admin
	if res := s.db.WithContext(ctx).Model(&models.Admin{}).Where("user_id = ?", id).First(&admin); res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, fmt.Errorf("%s: %w", op, res.Error)
	}

	return admin.IsAdmin, nil
}
