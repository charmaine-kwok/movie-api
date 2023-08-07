package models

type User struct {
	Username string `gorm:"unique" json:"username"`
	Password string `json:"password"`
}
