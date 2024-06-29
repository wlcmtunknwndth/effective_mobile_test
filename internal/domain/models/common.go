package models

type Model struct {
	ID uint64 `gorm:"primaryKey;uniqueIndex;type:bigint;autoIncrement" json:"id"`
}
