package models

import "encoding/json"

type Address struct {
	Unique
	Type                   Enumerated `json:"type"`
	StreetName             string     `json:"streetName"`
	EntranceNumber         int        `json:"entranceNumber"`
	Neighborhood           string     `json:"neighborhood"`
	PostalCode             string     `json:"postalCode"`
	Complement             string     `json:"complement"`
	CommercialCenterName   string     `json:"commercialCenterName"`
	NearestLandmark        string     `json:"nearestLandmark"`
	CityName               string     `json:"cityName"`
	Country                Country    `json:"country"`
	CountrySubdivisionCode string     `json:"countrySubdivisionCode"`
}

func (a *Address) MarshalJSON() ([]byte, error) {
	type Alias Address
	return json.Marshal(&struct {
		*Alias
		TypeID    int `json:"typeId"`
		CountryID int `json:"countryId"`
	}{
		Alias:     (*Alias)(a),
		TypeID:    a.Type.ID,
		CountryID: a.Country.ID,
	})
}
