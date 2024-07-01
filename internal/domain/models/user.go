package models

import (
	"strconv"
	"strings"
)

type User struct {
	Model
	PassportSerie  uint16 `gorm:"type:int;index:passport,unique"`
	PassportNumber uint32 `gorm:"type:int;index:passport,unique"`
	PassHash       []byte `gorm:"type:bytea"`
}

type UserInfo struct {
	UserID     uint64 `gorm:"type:bigint;uniqueIndex;primaryKey"`
	Name       string `gorm:"type:varchar(30)"`
	Surname    string `gorm:"type:varchar(30)"`
	Patronymic string `gorm:"type:varchar(30)"`
	Address    string
}

type CreateUserAPI struct {
	Passport   string `json:"passport"`
	Password   string `json:"password"`
	Name       string `json:"name,omitempty"`
	Surname    string `json:"surname,omitempty"`
	Patronymic string `json:"patronymic,omitempty"`
	Address    string `json:"address,omitempty"`
}

type UserInfoAPI struct {
	UserID     uint64 `json:"user_id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Address    string `json:"address"`
}

type UserAPI struct {
	ID       uint64 `json:"id"`
	Passport string `json:"passport"`
	PassHash []byte `json:"password"`
}

func StringToSerieAndNumber(passportId string) (uint16, uint32) {
	res := strings.Split(passportId, " ")
	if len(res) != 2 {
		return 0, 0
	}
	serie, err := strconv.ParseUint(res[0], 10, 16)
	if err != nil {
		return 0, 0
	}

	number, err := strconv.ParseUint(res[1], 10, 32)
	if err != nil {
		return 0, 0
	}

	return uint16(serie), uint32(number)
}

func ApiToDB(api UserAPI) User {
	serie, number := StringToSerieAndNumber(api.Passport)
	return User{
		Model:          Model{ID: api.ID},
		PassportSerie:  serie,
		PassportNumber: number,
		PassHash:       api.PassHash,
	}
}
