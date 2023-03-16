package model

type Comments struct {
	Count string     `json:"count"`
	Data  []*Comment `json:"data"`
}

type Comment struct {
	Kijicode string `json:"kijicode"`
	Code     string `json:"code"`
	Body     string `json:"body"`
	Comment1 string `json:"comment1"`
	Date     string `json:"date"`
}
