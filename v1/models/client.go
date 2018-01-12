package models

//Client represents an client model
type Client struct {
	Address        Address `json:"address"`
	Contact        string  `json:"contact"`
	DocumentNumber string  `json:"document_number"`
	DocumentType   string  `json:"document_type"`
	Name           string  `json:"name"`
	OpeningHours   string  `json:"opening_hours"`
	PhoneNumber    string  `json:"phone_number"`
	StoneCode      string  `json:"stone_code"`
}
