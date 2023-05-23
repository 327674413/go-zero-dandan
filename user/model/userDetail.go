package model

type UserDetail struct {
	UserId      string `gorm:"user_id"`
	NameCardImg string `gorm:"name_card_img"`
}
