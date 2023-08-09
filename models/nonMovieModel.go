package models

import (
	"time"

	"gorm.io/gorm"
)

type NonMovie struct {
	gorm.Model `json:"-"`
	ID         uint      `json:"id" gorm:"primarykey"`
	Title      string    `json:"title" binding:"required"`
	Desc       string    `json:"desc" binding:"required"`
	Location   string    `json:"location" binding:"required"`
	Date       time.Time `json:"date" binding:"required"`
	Rating     string    `json:"rating" binding:"required"`
	Pic        string    `json:"pic" binding:"required"`
	User_id    string    `json:"-"`
}

func (n *NonMovie) GetTitle() string {
	return n.Title
}
