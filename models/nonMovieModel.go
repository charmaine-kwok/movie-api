package models

import (
	"time"

	"gorm.io/gorm"
)

type NonMovie struct {
	gorm.Model `json:"-"`
	ID         uint      `json:"id" gorm:"primarykey" example:"5"`
	Title      string    `json:"title" binding:"required" example:"Westlife The Wild Dreams Tour"`
	Desc       string    `json:"desc" binding:"required" example:"So great to see WESTLIFE live!"`
	Location   string    `json:"location" binding:"required" example:"ASIAWORLD-ARENA"`
	Date       time.Time `json:"date" binding:"required" example:"2023-02-15T00:00:00+00"`
	Rating     string    `json:"rating" binding:"required" example:"9.0"`
	Pic        string    `json:"pic" binding:"required" example:"https://res.klook.com/image/upload/v1670553795/sn2b41ae5zpobabcxya4.jpg"`
	User_id    string    `json:"-"`
}

func (n *NonMovie) GetTitle() string {
	return n.Title
}
