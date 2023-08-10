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
	ID         uint      `json:"id" gorm:"primarykey" example:"1"`
	Title_zh   string    `json:"title_zh" binding:"required" example:"Pulp Fiction"`
	Title      string    `json:"title" binding:"required" example:"黑色追緝令"`
	Desc       string    `json:"desc" binding:"required" example:"A very good movie."`
	Location   string    `json:"location" binding:"required" example:"K11"`
	Date       time.Time `json:"date" binding:"required" example:"2023-02-15T00:00:00+00"`
	Rating     string    `json:"rating" binding:"required" example:"9.0"`
	Pic        string    `json:"pic" binding:"required" example:"https://upload.wikimedia.org/wikipedia/en/3/3b/Pulp_Fiction_%281994%29_poster.jpg"`
	Wiki_url   string    `json:"wiki_url,omitempty" example:"https://en.wikipedia.org/wiki/Pulp_Fiction"`
	User_id    string    `json:"-"`
}

func (m *Movie) GetTitle() string {
	return m.Title
}
