package services

import (
	"github.com/stone-payments/logistic-sdk-go/errors"
	"github.com/stone-payments/logistic-sdk-go/http"
)

func processResponse(resp http.Response) (*response, error) {
	var (
		r   response
		err error
	)

	if resp.Ok() {
		if e := resp.JSON(&r); e != nil {
			err = errors.NewSerializing(e.Error())
		}
	} else {
		err = trackError(resp)
	}

	return &r, err
}

func trackError(resp http.Response) error {
	messages := errorMessages(resp)
	return errors.Build(resp.StatusCode(), messages...)
}

func errorMessages(resp http.Response) []string {
	var messages []string
	var r response
	err := resp.JSON(&r)

	if err == nil {
		for _, e := range r.Errors {
			messages = append(messages, e.Message)
		}
	}

	return messages
}
