package models

import "time"

type Task struct {
	Model
	UserID      uint64 `gorm:"index;type:bigint;autoIncrement"`
	HoursSpent  float64
	Description string `gorm:"type:varchar(2048)"`
	CreatedAt   time.Time
	DoneAt      time.Time
	ExpiresAt   time.Time
}

type TaskAPI struct {
	ID          uint64    `json:"ID,omitempty"`
	UserID      uint64    `json:"user_id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	DoneAt      time.Time `json:"done_at"`
	ExpiresAt   time.Time `json:"expires_at"`
	TimeSpent   time.Time `json:"time_spent"`
}
