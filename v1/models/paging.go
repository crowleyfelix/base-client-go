package models

type Paging struct {
	Page           int `json:"page"`
	RecordsPerPage int `json:"recordsPerPage"`
	TotalPages     int `json:"totalPages"`
	TotalRecords   int `json:"totalRecords"`
}
