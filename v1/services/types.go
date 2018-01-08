package services

import (
	"encoding/json"

	"github.com/stone-payments/logistic-sdk-go/errors"
	"github.com/stone-payments/logistic-sdk-go/v1/models"
)

type response struct {
	Success bool            `json:"success"`
	Data    json.RawMessage `json:"data"`
	Paging  *models.Paging  `json:"paging"`
	Errors  []errorData     `json:"errors"`
}

func (r *response) UnmarshalData(data interface{}) errors.Error {
	err := json.Unmarshal(r.Data, data)

	if err != nil {
		return errors.NewSerializing(err.Error())
	}

	return nil
}

type errorData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
