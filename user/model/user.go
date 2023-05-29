package model

type User struct {
	Account  int    `gorm:"account"`
	Password string `gorm:"password"`
	Username string `gorm:"username"`
}

func (t *User) TableName() string {
	return "dan_user"
}
