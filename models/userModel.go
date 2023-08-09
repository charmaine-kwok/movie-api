package models

import gorm "gorm.io/gorm"

type User struct {
	gorm.Model `json:"-"`
	Username   string `gorm:"unique" json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
}
