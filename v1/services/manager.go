package services

import (
	"fmt"

	"github.com/karlkfi/inject"
	"github.com/labstack/gommon/log"
	logisticsdk "github.com/stone-payments/logistic-sdk-go"
	"github.com/stone-payments/logistic-sdk-go/errors"
	"github.com/stone-payments/logistic-sdk-go/http"
)

type manager struct {
	baseURL     string
	credentials Credentials
	processors  []http.Processable
}

func (s *manager) Requester() http.Requestable {
	var r http.Requestable
	inject.ExtractAssignable(logisticsdk.Graph, &r)
	return r
}

func (s *manager) RequestOptions(httpMethod string, url string) http.Request {
	request := http.NewRequest()
	request.SetHeaders(s.headers(httpMethod, url))

	return request
}

func (s *manager) BuildURL(endpoint string, params ...interface{}) string {
	return s.baseURL + fmt.Sprintf(endpoint, params...)
}

func (s *manager) Request(method http.RequestMethod, url string, request http.Request, processors ...http.Processable) (http.Response, error) {
	resp, err := method(url, request)

	if err != nil {
		return nil, errors.NewCallout(err.Error())
	}

	processors = s.allProcessors(processors...)

	for _, processor := range processors {
		request = processor.OnRequest(request)
	}

	if resp, e := method(url, request); e != nil {
		log.Error(fmt.Sprintf("Failed httping %s. Message: %s", url, e.Error()))
		return nil, errors.NewInternalServer(e.Error())

	} else {

		var err error
		for _, processor := range processors {
			resp, err = processor.OnResponse(resp)

			if err != nil {
				break
			}
		}

		return resp, err
	}
}

func (s *manager) allProcessors(processors ...http.Processable) []http.Processable {
	var all []http.Processable
	all = append(all, s.processors...)
	all = append(all, processors...)
	return all
}

func (s *manager) headers(httpMethod string, url string) map[string]string {

	return map[string]string{
		"Authorization":   s.credentials.Signature(httpMethod, url),
		"User-Identifier": s.credentials.User,
		"Content-Type":    "application/json",
	}
}
