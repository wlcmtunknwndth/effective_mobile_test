package models

import "time"

type Task struct {
	Id          uint64
	UserId      uint64
	Description string
	CreatedAt   time.Time
	DoneAt      time.Time
	DeadlineAt  time.Time
}
