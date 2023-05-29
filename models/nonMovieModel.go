package models

import (
	"github.com/globalsign/mgo/bson"
)

type NonMovie struct {
	ID       bson.ObjectId `bson:"_id"`
	Title    string        `json:"title"`
	Desc     string        `json:"desc"`
	Location string        `json:"location"`
	Date     string        `json:"date"`
	Rating   string        `json:"rating"`
	Pic      string        `json:"pic"`
}

func (m *NonMovie) GetTitle() string {
	return m.Title
}
