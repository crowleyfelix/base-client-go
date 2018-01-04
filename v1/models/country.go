package models

type Country struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Iso31661Alpha3 string `json:"iso31661Alpha3"`
	Iso31661Alpha2 string `json:"iso31661Alpha2"`
}
