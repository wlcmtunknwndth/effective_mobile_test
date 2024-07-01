package models

type UserDB struct {
	Model
	PassportSerie  uint16 `gorm:"type:int;uniqueIndex;unique"`
	PassportNumber uint32 `gorm:"type:int;uniqueIndex;unique"`
	PassHash       []byte `gorm:"type:bytea"`
}

type UserInfoDB struct {
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
	ID       uint   `json:"id"`
	Passport string `json:"passport"`
	Password string `json:"password"`
}
