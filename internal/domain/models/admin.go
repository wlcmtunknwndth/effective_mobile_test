package models

type Admin struct {
	UserId  uint64 `gorm:"type:bigint;uniqueIndex;primaryKey"`
	IsAdmin bool   `gorm:"default:false" json:"is_admin"`
}
