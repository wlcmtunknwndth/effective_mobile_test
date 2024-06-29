package models

import "time"

type Task struct {
	Model
	UserID      uint64    `gorm:"index;type:bigint;autoIncrement" json:"user_id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	DoneAt      time.Time `json:"done_at"`
	ExpiresAt   time.Time `json:"expires_at"`
}
