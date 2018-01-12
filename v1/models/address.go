package models

//Address represents an address model
type Address struct {
	City           string `json:"city"`
	Complement     string `json:"complement"`
	Country        string `json:"country"`
	District       string `json:"district"`
	DistrictNumber string `json:"district_number"`
	DistrictType   string `json:"district_type"`
	Neighborhood   string `json:"neighborhood"`
	Region         string `json:"region"`
	State          string `json:"state"`
	Zipcode        string `json:"zipcode"`
}
