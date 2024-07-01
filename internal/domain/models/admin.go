package models

type AdminDB struct {
	UserId  uint64 `gorm:"type:bigint;uniqueIndex;primaryKey"`
	IsAdmin bool   `gorm:"default:false" json:"is_admin"`
}

type AdminAPI struct {
	UserID  uint64 `json:"user_id"`
	IsAdmin bool   `json:"is_admin"`
}
