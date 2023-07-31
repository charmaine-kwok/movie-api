package models

type Model interface {
	GetTitle() string
}

type Movie struct {
	Title_zh string `json:"title_zh"`
	Title_en string `json:"title_en"`
	Desc     string `json:"desc"`
	Location string `json:"location"`
	Date     string `json:"date"`
	Rating   string `json:"rating"`
	Pic      string `json:"pic"`
	Wiki_url string `json:"wiki_url"`
}

func (m *Movie) GetTitle() string {
	return m.Title_en
}
