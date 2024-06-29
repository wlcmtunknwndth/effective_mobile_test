package models

type User struct {
	ID             uint64
	PassportNumber string
	PassHash       []byte
}
