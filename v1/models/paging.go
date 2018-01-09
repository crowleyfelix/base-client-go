package models

type Paging struct {
	Current      int `json:"current"`
	TotalOfPages int `json:"total_of_pages"`
}
