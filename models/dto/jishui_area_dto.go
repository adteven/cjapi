package dto

type JiShuiAreaDto struct {
	Area           string   `json:"area"`
	Id             int64    `json:"id"`
	JsNumber       string   `json:"jsNumber"`
	Lat            string   `json:"lat"`
	Lng            string   `json:"lng"`
	Name           string   `json:"name"`
	Road           string   `json:"road"`
	ProjectName    string   `json:"projectName"`
	ProjectContent string   `json:"projectContent"`
	ProjectUnit    string   `json:"projectUnit"`
	ProjectPeriod  string   `json:"projectPeriod"`
	Pictures       []string `json:"pictures"`
	Process        string   `json:"process"`
}
