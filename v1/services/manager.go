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
	baseURL      string
	credentials  Credentials
	interceptors []http.Interceptor
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

func (s *manager) Request(method http.RequestMethod, url string, request http.Request, interceptors ...http.Interceptor) (http.Response, errors.Error) {
	interceptors = s.allInterceptors(interceptors...)

	for _, interceptor := range interceptors {
		request = interceptor.OnRequest(request)
	}

	if resp, e := method(url, request); e != nil {
		log.Error(fmt.Sprintf("Failed requesting %s. Message: %s", url, e.Error()))
		return nil, errors.NewCallout(e.Error())

	} else {

		var err errors.Error
		for _, interceptor := range interceptors {
			resp, err = interceptor.OnResponse(resp)

			if err != nil {
				break
			}
		}

		return resp, err
	}
}

func (s *manager) allInterceptors(interceptors ...http.Interceptor) []http.Interceptor {
	var all []http.Interceptor
	inject.FindAssignable(logisticsdk.Graph, &all)
	all = append(all, s.interceptors...)
	all = append(all, interceptors...)
	return all
}

func (s *manager) headers(httpMethod string, url string) map[string]string {

	return map[string]string{
		"Authorization":   s.credentials.Signature(httpMethod, url),
		"User-Identifier": s.credentials.User,
		"Content-Type":    "application/json",
	}
}
