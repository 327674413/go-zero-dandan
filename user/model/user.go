package model

type User struct {
	Account  string `gorm:"account"`
	Password string `gorm:"password"`
	Username string `gorm:"username"`
}
