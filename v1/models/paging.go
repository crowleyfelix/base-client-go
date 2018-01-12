package models

//Paging holds base page data
type Paging struct {
	Current      int `json:"current"`
	TotalOfPages int `json:"total_of_pages"`
}
