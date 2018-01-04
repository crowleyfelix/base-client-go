package models

import "encoding/json"

type Benefit struct {
	BenefitType                 Enumerated `json:"benefitType"`
	BenefitEstablishmentType    Enumerated `json:"benefitEstablishmentType"`
	NumberOfCashiers            int        `json:"numberOfCashiers"`
	ServiceAreaSi               int        `json:"serviceAreaSi"`
	MeetsHealthyMenu            bool       `json:"meetsHealthyMenu,omitempty"`
	MeetsNutritionalRequirement bool       `json:"meetsNutritionalRequirement,omitempty"`
	NumberOfTables              int        `json:"numberOfTables,omitempty"`
	NumberOfSeats               int        `json:"numberOfSeats,omitempty"`
	MaxNumberOfMeals            int        `json:"maxNumberOfMeals,omitempty"`
}

func (b *Benefit) MarshalJSON() ([]byte, error) {
	type Alias Benefit
	return json.Marshal(&struct {
		*Alias
		BenefitTypeID              int `json:"benefitTypeId"`
		BenefitEstablishmentTypeID int `json:"benefitEstablishmentTypeId"`
	}{
		Alias:                      (*Alias)(b),
		BenefitTypeID:              b.BenefitType.ID,
		BenefitEstablishmentTypeID: b.BenefitEstablishmentType.ID,
	})
}
