package models

import (
	"time"

	"gorm.io/gorm"
)

type Model interface {
	GetTitle() string
}

type Movie struct {
	gorm.Model `json:"-"`
	ID         uint      `json:"id" gorm:"primarykey"`
	Title_zh   string    `json:"title_zh" binding:"required"`
	Title      string    `json:"title" binding:"required"`
	Desc       string    `json:"desc" binding:"required"`
	Location   string    `json:"location" binding:"required"`
	Date       time.Time `json:"date" binding:"required"`
	Rating     string    `json:"rating" binding:"required"`
	Pic        string    `json:"pic" binding:"required"`
	Wiki_url   string    `json:"wiki_url,omitempty"`
	User_id    string    `json:"-"`
}

func (m *Movie) GetTitle() string {
	return m.Title
}
