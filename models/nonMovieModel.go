package models

import "time"

type NonMovie struct {
	Title    string    `json:"title"`
	Desc     string    `json:"desc"`
	Location string    `json:"location"`
	Date     time.Time `json:"date"`
	Rating   string    `json:"rating"`
	Pic      string    `json:"pic"`
}

func (m *NonMovie) GetTitle() string {
	return m.Title
}
