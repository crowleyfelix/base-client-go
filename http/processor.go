package http

//Processable exposes processor methods
type Processable interface {
	OnRequest(Request) Request
	OnResponse(Response) (Response, error)
}

type processor struct {
	onRequestCallback  func(Request) Request
	onResponseCallback func(Response) (Response, error)
}

//OnRequestCallback define request callback processor assignature
type OnRequestCallback func(Request) Request

//OnResponseCallback define response callback processor assignature
type OnResponseCallback func(Response) (Response, error)

//NewProcessor constructs processor
func NewProcessor(onRequestCallback OnRequestCallback, onResponseCallback OnResponseCallback) Processable {
	return &processor{
		onRequestCallback,
		onResponseCallback,
	}
}

func (p *processor) OnRequest(request Request) Request {
	if p.onRequestCallback != nil {
		return p.onRequestCallback(request)
	}

	return request
}

func (p *processor) OnResponse(response Response) (Response, error) {
	if p.onResponseCallback != nil {
		return p.onResponseCallback(response)
	}

	return response, nil
}
