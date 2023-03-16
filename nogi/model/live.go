package model

type Lives struct {
	Count string  `json:"count"`
	Data  []*Live `json:"data"`
}

type Live struct {
	Code    string       `json:"code"`
	Title   string       `json:"title"`
	Img     string       `json:"img"`
	Date    string       `json:"date"`
	Caption string       `json:"caption"`
	Link    string       `json:"link"`
	Cate    string       `json:"cate"`
	Kouen   []*LiveKouen `json:"kouen"`
}

type LiveKouen struct {
	Date      string `json:"date"`
	Youbi     string `json:"youbi"`
	Area      string `json:"area"`
	Place     string `json:"place"`
	OpenTime  string `json:"open_time"`
	StartTime string `json:"start_time"`
	Cate      string `json:"cate"`
}
