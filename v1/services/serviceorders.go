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
}

type serviceOrder struct {
	*manager
}

//NewServiceOrder is constructs ServiceOrder Service
func NewServiceOrder(url string, credentials *Credentials) ServiceOrder {

	interceptor := http.NewInterceptor(nil, processResponse)

	return &serviceOrder{
		&manager{
			baseURL:      url,
			credentials:  *credentials,
			interceptors: []http.Interceptor{interceptor},
		},
	}
}

func (r *serviceOrder) Get(key string) (*models.ServiceOrder, errors.Error) {
	var serviceOrder models.ServiceOrder
	endPointURL := r.BuildURL(serviceOrdersEndpoint, key)

	resp, err := r.Request(r.Requester().Get, endPointURL, r.RequestOptions("GET", endPointURL))

	if err == nil {
		err = resp.JSON(&serviceOrder)
	}

	return &serviceOrder, err
}

func (r *serviceOrder) List() ([]models.ServiceOrder, error) {
	var serviceOrders []models.ServiceOrder
	endPointURL := r.BuildURL(serviceOrdersEndpoint)

	resp, err := r.Request(r.Requester().Get, endPointURL, r.RequestOptions("GET", endPointURL))

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
