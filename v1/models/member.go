package models

import "encoding/json"

type Member struct {
	Unique
	LegalName               string             `json:"legalName"`
	TradeName               string             `json:"tradeName"`
	LegalPersonality        Enumerated         `json:"legalPersonality"`
	TaxID                   string             `json:"taxId"`
	TaxIDType               Enumerated         `json:"taxIdType"`
	TaxIDRegistrationDate   string             `json:"taxIdRegistrationDate"`
	EstimatedMonthlyBilling int                `json:"estimatedMonthlyBilling"`
	AverageTicket           int                `json:"averageTicket"`
	WebsiteURL              string             `json:"websiteUrl"`
	BirthDate               string             `json:"birthDate"`
	BirthPlace              string             `json:"birthPlace"`
	BirthCountry            Country            `json:"birthCountry"`
	FatherName              string             `json:"fatherName"`
	MotherName              string             `json:"motherName"`
	SpouseName              string             `json:"spouseName"`
	SpouseTaxID             string             `json:"spouseTaxId"`
	StateInscription        string             `json:"stateInscription"`
	MunicipalInscription    string             `json:"municipalInscription"`
	SpouseTaxIDType         Enumerated         `json:"spouseTaxIdType"`
	EconomicActivities      []EconomicActivity `json:"-"`
	BankAccounts            []BankAccount      `json:"-"`
	ExtraData               []ExtraData        `json:"-"`
	Contacts                []Contact          `json:"-"`
	Addresses               []Address          `json:"-"`
	WorkSchedule            []WorkSchedule     `json:"-"`
	Partners                []Partner          `json:"-"`
	Benefits                []Benefit          `json:"-"`
	Tags                    []string           `json:"-"`
}

func (m *Member) MarshalJSON() ([]byte, error) {
	type Alias Member
	return json.Marshal(&struct {
		*Alias
		LegalPersonalityID int                `json:"legalPersonalityId"`
		TaxIDTypeID        int                `json:"taxIdTypeId"`
		BirthCountryID     int                `json:"birthCountryId"`
		SpouseTaxIDTypeID  int                `json:"spouseTaxIdTypeId"`
		EconomicActivities []EconomicActivity `json:"economicActivities,omitempty"`
		BankAccounts       []BankAccount      `json:"bankAccounts,omitempty"`
		ExtraData          []ExtraData        `json:"extraData,omitempty"`
		Contacts           []Contact          `json:"contacts,omitempty"`
		Addresses          []Address          `json:"addresses,omitempty"`
		WorkSchedule       []WorkSchedule     `json:"workSchedule,omitempty"`
		Partners           []Partner          `json:"partners,omitempty"`
		Benefits           []Benefit          `json:"benefits,omitempty"`
		Tags               []string           `json:"tags,omitempty"`
	}{
		Alias:              (*Alias)(m),
		LegalPersonalityID: m.LegalPersonality.ID,
		TaxIDTypeID:        m.TaxIDType.ID,
		SpouseTaxIDTypeID:  m.SpouseTaxIDType.ID,
		BirthCountryID:     m.BirthCountry.ID,
		EconomicActivities: m.EconomicActivities,
		BankAccounts:       m.BankAccounts,
		ExtraData:          m.ExtraData,
		Contacts:           m.Contacts,
		Addresses:          m.Addresses,
		WorkSchedule:       m.WorkSchedule,
		Partners:           m.Partners,
		Benefits:           m.Benefits,
		Tags:               m.Tags,
	})
}
