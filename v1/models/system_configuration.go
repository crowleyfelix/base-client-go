package models

import "encoding/json"

type SystemConfiguration struct {
	//unique
	Key               string             `json:"key"`
	Value             string             `json:"value"`
	Description       string             `json:"description"`
	CreatedOn         string             `json:"createdOn"`
	LastModifiedOn    string             `json:"lastModifiedOn"`
}

func (s *SystemConfiguration) MarshalJSON() ([]byte, error) {
	type Alias SystemConfiguration
	return json.Marshal(&struct {
		*Alias
		Key               string             `json:"key"`
		Value             string             `json:"value"`
		Description       string             `json:"description"`
		CreatedOn         string             `json:"createdOn"`
		LastModifiedOn    string             `json:"lastModifiedOn"`
	}{
		Alias:     (*Alias)(s),
		Key: s.Key,
		Value: s.Value,
		Description: s.Description,
		CreatedOn: s.CreatedOn,
		LastModifiedOn: s.LastModifiedOn,
	})
}

