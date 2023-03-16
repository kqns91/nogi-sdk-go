package model

type NewsList struct {
	Count string  `json:"count"`
	Data  []*News `json:"data"`
}

type News struct {
	Code     string `json:"code"`
	Title    string `json:"title"`
	Date     string `json:"date"`
	Text     string `json:"text"`
	Cate     string `json:"cate"`
	LinkURL  string `json:"link_url"`
	ArtiCode string `json:"arti_code"`
}
