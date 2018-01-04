package models

import "encoding/json"

type BankAccount struct {
	Unique
	Country                 Country    `json:"country"`
	Bank                    Enumerated `json:"bank"`
	BranchCode              string     `json:"branchCode"`
	BranchCodeCheckDigit    string     `json:"branchCodeCheckDigit"`
	AccountNumber           string     `json:"accountNumber"`
	AccountNumberCheckDigit string     `json:"accountNumberCheckDigit"`
	Iban                    string     `json:"iban"`
	Status                  Enumerated `json:"status"`
	AccountType             Enumerated `json:"accountType"`
}

func (b *BankAccount) MarshalJSON() ([]byte, error) {
	type Alias BankAccount
	return json.Marshal(&struct {
		*Alias
		CountryID     int `json:"countryId"`
		BankID        int `json:"bankId"`
		StatusID      int `json:"statusId"`
		AccountTypeID int `json:"accountTypeId"`
	}{
		Alias:         (*Alias)(b),
		CountryID:     b.Country.ID,
		BankID:        b.Bank.ID,
		StatusID:      b.Status.ID,
		AccountTypeID: b.AccountType.ID,
	})
}
