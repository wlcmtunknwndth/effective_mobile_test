package models

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
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
	Address    string `gorm:"type:varchar(128)"`
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

var WrongNumber = errors.New("wrong passport number")
var WrongSerie = errors.New("wrong passport serie")
var WrongFormat = errors.New("wrong format")

func StringToSerieAndNumber(passportId string) (*uint64, *uint64, error) {
	res := strings.Split(passportId, " ")
	if len(res) != 2 {
		return nil, nil, fmt.Errorf("%s: %w", passportId, WrongFormat)
	}

	serie, err := strconv.ParseUint(res[0], 10, 16)
	if err != nil {
		return nil, nil, fmt.Errorf("%s: %w", res[0], WrongSerie)
	}

	number, err := strconv.ParseUint(res[1], 10, 32)
	if err != nil {
		return nil, nil, fmt.Errorf("%s: %w", res[1], WrongNumber)
	}

	return &serie, &number, nil
}

func ApiToDB(api *UserAPI) (*User, error) {
	serie, number, err := StringToSerieAndNumber(api.Passport)
	if err != nil {
		return nil, err
	}
	return &User{
		Model:          Model{ID: api.ID},
		PassportSerie:  uint16(*serie),
		PassportNumber: uint32(*number),
		PassHash:       api.PassHash,
	}, nil
}

func CreateUserToUsersDB(api *CreateUserAPI) (*User, *UserInfo, error) {
	serie, number, err := StringToSerieAndNumber(api.Passport)
	if err != nil {
		return nil, nil, err
	}
	passHash, err := bcrypt.GenerateFromPassword([]byte(api.Password), 5)
	if err != nil {
		return nil, nil, err
	}
	return &User{
			Model:          Model{},
			PassportSerie:  uint16(*serie),
			PassportNumber: uint32(*number),
			PassHash:       passHash,
		}, &UserInfo{
			UserID:     0,
			Name:       api.Name,
			Surname:    api.Surname,
			Patronymic: api.Patronymic,
			Address:    api.Address,
		}, nil
}
