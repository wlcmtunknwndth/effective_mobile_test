package models

type Admin struct {
	Model
	IsAdmin bool `gorm:"default:false" json:"is_admin"`
}
