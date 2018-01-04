package http

import (
	"github.com/levigross/grequests"
)

//Request exposes request request methods
type Request interface {
	SetHeaders(map[string]string)
	Headers() map[string]string
	SetJSON(interface{})
	JSON() interface{}
}

//NewRequest constructs implementation of Request
func NewRequest() Request {
	return new(request)
}

//request is the struct wrapper of grequests struct
type request struct {
	grequests.RequestOptions
}

func (r *request) SetHeaders(headers map[string]string) {
	r.RequestOptions.Headers = headers
}

func (r *request) Headers() map[string]string {
	return r.RequestOptions.Headers
}

func (r *request) SetJSON(data interface{}) {
	r.RequestOptions.JSON = data
}

func (r *request) JSON() interface{} {
	return r.RequestOptions.JSON
}
