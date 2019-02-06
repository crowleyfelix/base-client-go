package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/crowleyfelix/base-client-go/errors"
)

const (
	//GetMethod HTTTP Get Methjod
	GetMethod RequestMethod = "GET"
	//PostMethod HTTTP Post Methjod
	PostMethod RequestMethod = "POST"
	//PutMethod HTTTP Put Methjod
	PutMethod RequestMethod = "PUT"
	//DeleteMethod HTTTP Delete Methjod
	DeleteMethod RequestMethod = "DELETE"
)

//RequestMethod represents a http request method
type RequestMethod string

//Request represents a http request
type Request struct {
	*http.Request
}

func (req *Request) AddQuery(key, value string) {
	req.URL.RawQuery += key + "=" + value
}

//RequestBuilder exposes methods of a http request method
type RequestBuilder interface {
	SetBaseURL(baseURL string)
	SetBaseHeader(header map[string]string)
	Build(method RequestMethod, endpoint string, body io.Reader, params ...interface{}) (*Request, errors.Error)
	BuildJSON(method RequestMethod, endpoint string, data interface{}, params ...interface{}) (*Request, errors.Error)
}

//NewRequestBuilder creates a new request builder
func NewRequestBuilder(baseURL string) RequestBuilder {
	return &requestBuilder{baseURL: baseURL}
}

type requestBuilder struct {
	baseURL    string
	baseHeader map[string]string
}

func (r *requestBuilder) SetBaseURL(baseURL string) {
	r.baseURL = baseURL
}

func (r *requestBuilder) SetBaseHeader(header map[string]string) {
	r.baseHeader = header
}

func (r *requestBuilder) Build(method RequestMethod, endpoint string, body io.Reader, params ...interface{}) (*Request, errors.Error) {
	req, err := http.NewRequest(string(method), r.buildURL(endpoint, params...), body)

	if err != nil {
		return nil, errors.NewCallout(err.Error())
	}

	for k, v := range r.baseHeader {
		req.Header.Add(k, v)
	}

	return &Request{req}, nil
}

func (r *requestBuilder) BuildJSON(method RequestMethod, endpoint string, data interface{}, params ...interface{}) (*Request, errors.Error) {
	blob, e := json.Marshal(data)

	if e != nil {
		return nil, errors.NewSerializing(e.Error())
	}

	req, err := r.Build(method, endpoint, bytes.NewReader(blob), params...)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	return req, err
}

func (r *requestBuilder) buildURL(endpoint string, params ...interface{}) string {
	return r.baseURL + fmt.Sprintf(endpoint, params...)
}
