package models

//Terminal represents a client terminal
type Terminal struct {
	Model        string `json:"model"`
	Technology   string `json:"technology"`
	SerialNumber string `json:"serial_number"`
}
