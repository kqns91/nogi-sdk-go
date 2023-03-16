package model

type HandShakes struct {
	Count string       `json:"count"`
	Data  []*HandShake `json:"data"`
}

type HandShake struct {
	Code            string            `json:"code"`
	Cate            string            `json:"cate"`
	Title           string            `json:"title"`
	Subtitle        string            `json:"subtitle"`
	Img             string            `json:"img"`
	Caption         string            `json:"caption"`
	Date            string            `json:"date"`
	Link            string            `json:"link"`
	ReleaseLink     string            `json:"release_link"`
	ArtiCode        string            `json:"arti_code"`
	Kouen           []*HandShakeKouen `json:"kouen"`
	ApplicationTerm string            `json:"application_term"`
}

type HandShakeKouen struct {
	Date      string `json:"date"`
	Youbi     string `json:"youbi"`
	Area      string `json:"area"`
	Place     string `json:"place"`
	OpenTime  string `json:"open_time"`
	StartTime string `json:"start_time"`
}
