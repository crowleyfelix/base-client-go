package services

import (
	"github.com/stone-payments/logistic-sdk-go/errors"
	"github.com/stone-payments/logistic-sdk-go/http"
)

func processResponse(resp http.Response) (http.Response, errors.Error) {
	var (
		err errors.Error
	)

	if !resp.Ok() {
		err = trackError(resp)
	}

	return resp, err
}

func trackError(resp http.Response) errors.Error {
	messages := errorMessages(resp)
	return errors.Build(resp.StatusCode(), messages...)
}

func errorMessages(resp http.Response) []string {
	var messages []string
	var r responseError
	err := resp.JSON(&r)

	if err == nil {
		for _, e := range r.ErrorList {
			messages = append(messages, e.Reason)
		}
	}

	return messages
}
