package logisticsdk

import (
	"github.com/karlkfi/inject"
	"github.com/stone-payments/CaduGO/request"
	"github.com/stone-payments/logistic-sdk-go/http"
)

var (
	//Graph holds the dependecies
	Graph     inject.Graph
	requester http.Requestable
)

func init() {
	registerDependencies()
}

func registerDependencies() {
	Graph = inject.NewGraph()
	RegisterRequester(request.NewRequester())
}

//RegisterRequester replaces default http requester
func RegisterRequester(req request.Requestable) {
	Graph.Define(&requester, inject.NewProvider(func() request.Requestable { return req }))
	Graph.Resolve(&requester)
}
