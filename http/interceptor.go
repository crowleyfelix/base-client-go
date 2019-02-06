package http

import "github.com/crowleyfelix/base-client-go/errors"

//Interceptor exposes interceptor methods
type Interceptor interface {
	OnRequest(*Request)
	OnResponse(*Response) errors.Error
}

type interceptor struct {
	onRequestCallback  func(*Request)
	onResponseCallback func(*Response) errors.Error
}

//OnRequestCallback define request callback interceptor assignature
type OnRequestCallback func(*Request)

//OnResponseCallback define response callback interceptor assignature
type OnResponseCallback func(*Response) errors.Error

//NewInterceptor constructs interceptor
func NewInterceptor(onRequestCallback OnRequestCallback, onResponseCallback OnResponseCallback) Interceptor {
	return &interceptor{
		onRequestCallback,
		onResponseCallback,
	}
}

func (p *interceptor) OnRequest(request *Request) {
	if p.onRequestCallback != nil {
		p.onRequestCallback(request)
	}

}

func (p *interceptor) OnResponse(response *Response) errors.Error {
	
	if p.onResponseCallback != nil {
		return p.onResponseCallback(response)
	}

	return nil
}
