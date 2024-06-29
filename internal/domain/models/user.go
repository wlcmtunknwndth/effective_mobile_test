package models

type User struct {
	Model
	PassportNumber string `gorm:"type:varchar(12);uniqueIndex;unique"`
	PassHash       []byte `gorm:"type:bytea"`
}
