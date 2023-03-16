package model

type Schedules struct {
	Count string      `json:"count"`
	Data  []*Schedule `json:"data"`
}

type Schedule struct {
	Code      string     `json:"code"`
	Title     string     `json:"title"`
	Text      string     `json:"text"`
	Date      string     `json:"date"`
	StartTime string     `json:"start_time"`
	EndTime   string     `json:"end_time"`
	Cate      string     `json:"cate"`
	Link      string     `json:"link"`
	ArtiCode  [][]string `json:"arti_code"`
}
