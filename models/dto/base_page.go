package dto

type BasePage struct {
	Page   int    `json:"page"`
	Size   int    `json:"size"`
	Blurry string `json:"blurry"`
}
