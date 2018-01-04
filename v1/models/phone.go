package models

import "encoding/json"

type Phone struct {
	Unique
	PhoneNumber string     `json:"phoneNumber"`
	AreaCode    string     `json:"areaCode"`
	CountryCode string     `json:"countryCode"`
	Extension   string     `json:"extension"`
	Type        Enumerated `json:"type"`
}

func (p *Phone) MarshalJSON() ([]byte, error) {
	type Alias Phone
	return json.Marshal(&struct {
		*Alias
		TypeID int `json:"typeId"`
	}{
		Alias:  (*Alias)(p),
		TypeID: p.Type.ID,
	})
}
