package vo

type ResultList struct {
	Data   interface{} `json:"data"`
	Total  int         `json:"total"`
	Extend interface{} `json:"extend"`
}
