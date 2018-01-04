package models

import "encoding/json"

type Partner struct {
	NaturalPerson   PartnerNaturalPerson   `json:"naturalPerson"`
	JuridicalPerson PartnerJuridicalPerson `json:"juridicalPerson"`
}

type PartnerNaturalPerson struct {
	Unique
	TaxIDType           Enumerated `json:"taxIdType"`
	TaxID               string     `json:"taxId"`
	Name                string     `json:"name"`
	BirthDate           string     `json:"birthDate"`
	BirthPlace          string     `json:"birthPlace"`
	BirthCountry        Country    `json:"birthCountry"`
	FatherName          string     `json:"fatherName"`
	MotherName          string     `json:"motherName"`
	SpouseName          string     `json:"spouseName"`
	SpouseTaxID         string     `json:"spouseTaxId"`
	SpouseTaxIDType     Enumerated `json:"spouseTaxIdType"`
	OwnershipPercentage float64    `json:"ownershipPercentage"`
	Addresses           []Address  `json:"-"`
}

func (p *PartnerNaturalPerson) MarshalJSON() ([]byte, error) {
	type Alias PartnerNaturalPerson
	return json.Marshal(&struct {
		*Alias
		BirthCountryID    int       `json:"birthCountryId"`
		TaxIDTypeID       int       `json:"taxIdTypeId"`
		SpouseTaxIDTypeID int       `json:"spouseTaxIdTypeId"`
		Addresses         []Address `json:"addresses,omitempty"`
	}{
		Alias:             (*Alias)(p),
		BirthCountryID:    p.BirthCountry.ID,
		TaxIDTypeID:       p.TaxIDType.ID,
		SpouseTaxIDTypeID: p.SpouseTaxIDType.ID,
		Addresses:         p.Addresses,
	})
}

type PartnerJuridicalPerson struct {
	Unique
	LegalName               string     `json:"legalName"`
	TradeName               string     `json:"tradeName"`
	LegalPersonality        Enumerated `json:"legalPersonality"`
	TaxID                   string     `json:"taxId"`
	TaxIDType               Enumerated `json:"taxIdType"`
	TaxIDRegistrationDate   string     `json:"taxIdRegistrationDate"`
	StateInscription        string     `json:"stateInscription"`
	MunicipalInscription    string     `json:"municipalInscription"`
	WebsiteURL              string     `json:"websiteUrl"`
	EstimatedMonthlyBilling int64      `json:"estimatedMonthlyBilling"`
	AverageTicket           int64      `json:"averageTicket"`
	OwnershipPercentage     float64    `json:"ownershipPercentage"`
	Addresses               []Address  `json:"-"`
}

func (p *PartnerJuridicalPerson) MarshalJSON() ([]byte, error) {
	type Alias PartnerJuridicalPerson
	return json.Marshal(&struct {
		*Alias
		LegalPersonalityID int       `json:"legalPersonalityId"`
		TaxIDTypeID        int       `json:"taxIdTypeId"`
		Addresses          []Address `json:"addresses,omitempty"`
	}{
		Alias:              (*Alias)(p),
		LegalPersonalityID: p.LegalPersonality.ID,
		TaxIDTypeID:        p.TaxIDType.ID,
		Addresses:          p.Addresses,
	})
}
