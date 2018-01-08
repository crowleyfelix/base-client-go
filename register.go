package logisticsdk

import (
	"github.com/karlkfi/inject"
	"github.com/stone-payments/logistic-sdk-go/http"
)

var (
	//Graph holds the dependecies
	Graph        inject.Graph
	requester    http.Requestable
	interceptors []http.Interceptor
)

func init() {
	registerDependencies()
}

func registerDependencies() {
	Graph = inject.NewGraph()
	RegisterRequester(http.NewRequester())
}

//RegisterRequester replaces default http requester
func RegisterRequester(req http.Requestable) {
	Graph.Define(&requester, inject.NewProvider(func() http.Requestable { return req }))
	Graph.Resolve(&requester)
}

//RegisterInterceptors add external interceptors for requests
func RegisterInterceptors(itcs ...http.Interceptor) {
	Graph.Define(&interceptors, inject.NewProvider(func() []http.Interceptor { return itcs }))
	Graph.Resolve(&interceptors)
}
