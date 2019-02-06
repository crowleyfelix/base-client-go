package http

import (
	"github.com/crowleyfelix/base-client-go/errors"
	"net/http"
)

//Client exposes request client methods
type Client interface {
	RequestBuilder() RequestBuilder
	Do(request *Request, interceptors ...Interceptor) (*Response, errors.Error)
}

//NewClient constructs a request client
func NewClient(baseURL string, interceptors ...Interceptor) Client {
	return &client{
		client:         new(http.Client),
		requestBuilder: NewRequestBuilder(baseURL),
		interceptors:   interceptors,
	}
}

type client struct {
	client         *http.Client
	requestBuilder RequestBuilder
	interceptors   []Interceptor
}

func (cli *client) RequestBuilder() RequestBuilder {
	return cli.requestBuilder
}

func (cli *client) Do(request *Request, interceptors ...Interceptor) (*Response, errors.Error) {
	interceptors = cli.allInterceptors(interceptors...)

	for _, interceptor := range interceptors {
		interceptor.OnRequest(request)
	}

	r, e := cli.client.Do(request.Request)
	if e != nil {
		return nil, errors.NewCallout(e.Error())

	}

	resp := &Response{r}

	var err errors.Error
	for _, interceptor := range interceptors {
		err = interceptor.OnResponse(resp)

		if err != nil {
			break
		}
	}

	return resp, err
}

func (cli *client) allInterceptors(interceptors ...Interceptor) []Interceptor {
	var all []Interceptor
	all = append(all, Interceptors...)
	all = append(all, cli.interceptors...)
	all = append(all, interceptors...)
	return all
}
