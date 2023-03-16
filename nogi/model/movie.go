package model

type Movies struct {
	Count string   `json:"count"`
	Data  []*Movie `json:"data"`
}

type Movie struct {
	Code  string `json:"code"`
	Title string `json:"title"`
	Date  string `json:"date"`
	Img   string `json:"img"`
	Cate  string `json:"cate"`
	Link  string `json:"link"`
}
