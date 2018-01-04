package models

import "encoding/json"

type Contact struct {
	Unique
	Type            Enumerated `json:"type"`
	Name            string     `json:"name"`
	FriendlyName    string     `json:"friendlyName"`
	TaxID           string     `json:"taxId"`
	TaxIDType       Enumerated `json:"taxIdType"`
	BirthDate       string     `json:"birthDate"`
	Birthplace      string     `json:"birthplace"`
	BirthCountry    Country    `json:"birthCountry"`
	FatherName      string     `json:"fatherName"`
	MotherName      string     `json:"motherName"`
	SpouseTaxID     string     `json:"spouseTaxId"`
	SpouseTaxIDType Enumerated `json:"spouseTaxIdType"`
	SpouseName      string     `json:"spouseName"`
	Phones          []Phone    `json:"-"`
	Emails          []string   `json:"-"`
}

func (c *Contact) MarshalJSON() ([]byte, error) {
	type Alias Contact
	return json.Marshal(&struct {
		*Alias
		TypeID            int      `json:"typeId"`
		TaxIDTypeID       int      `json:"taxIdTypeId"`
		BirthCountryID    int      `json:"birthCountryId"`
		SpouseTaxIDTypeID int      `json:"spouseTaxIdTypeId"`
		Phones            []Phone  `json:"phones"`
		Emails            []string `json:"emails"`
	}{
		Alias:             (*Alias)(c),
		TypeID:            c.Type.ID,
		TaxIDTypeID:       c.TaxIDType.ID,
		BirthCountryID:    c.BirthCountry.ID,
		SpouseTaxIDTypeID: c.SpouseTaxIDType.ID,
		Phones:            c.Phones,
		Emails:            c.Emails,
	})
}
