package services

import (
	"github.com/stone-payments/logistic-sdk-go/errors"
	"github.com/stone-payments/logistic-sdk-go/http"
	"github.com/stone-payments/logistic-sdk-go/v1/models"
)

//ServiceOrder abstracts service orders CRUD
type ServiceOrder interface {
	Get(string) (*models.ServiceOrder, errors.Error)
	ByStoneCode(int) ([]models.ServiceOrder, errors.Error)
	List(*models.ServiceOrderFilters) ([]models.ServiceOrder, errors.Error)
}

type serviceOrder struct {
	*manager
}

//NewServiceOrder is constructs ServiceOrder Service
func NewServiceOrder(url string) ServiceOrder {

	interceptor := http.NewInterceptor(nil, processResponse)

	return &serviceOrder{
		&manager{
			baseURL:      url,
			interceptors: []http.Interceptor{interceptor},
		},
	}
}

func (r *serviceOrder) Get(id string) (*models.ServiceOrder, errors.Error) {
	var serviceOrder models.ServiceOrder
	endPointURL := r.BuildURL(serviceOrdersEndpoint, id)

	resp, err := r.Request(r.Requester().Get, endPointURL, r.RequestOptions("GET", endPointURL))

	if err == nil {
		err = resp.JSON(&serviceOrder)
	}

	return &serviceOrder, err
}

func (r *serviceOrder) List(filters *models.ServiceOrderFilters) ([]models.ServiceOrder, errors.Error) {
	var serviceOrders []models.ServiceOrder
	endPointURL := r.BuildURL(serviceOrdersEndpoint)

	options := r.RequestOptions("GET", endPointURL)
	options.SetParams(filters.ToMap())

	resp, err := r.Request(r.Requester().Get, endPointURL, options)

	if err == nil {
		err = resp.JSON(&serviceOrders)
	}

	return serviceOrders, err
}

func (r *serviceOrder) ByStoneCode(stonecode int) ([]models.ServiceOrder, errors.Error) {
	var serviceOrders []models.ServiceOrder
	endPointURL := r.BuildURL(merchantServiceOrdersEndpoint, stonecode)

	resp, err := r.Request(r.Requester().Get, endPointURL, r.RequestOptions("GET", endPointURL))

	if err == nil {
		err = resp.JSON(&serviceOrders)
	}

	return serviceOrders, err
}
