package model

type User struct {
	Username string `gorm:"username"`
	Password string `gorm:"password"`
}
