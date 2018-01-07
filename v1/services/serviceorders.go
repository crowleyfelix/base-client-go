package services

import (
	"github.com/stone-payments/CaduGO/v1/models"
	"github.com/stone-payments/logistic-sdk-go/errors"
	"github.com/stone-payments/logistic-sdk-go/http"
)

//ServiceOrder abstracts service orders CRUD
type ServiceOrder interface {
	Get(string) (*models.ServiceOrder, errors.Error)
	ByStoneCode(int) ([]models.ServiceOrders, *models.Paging, errors.Error)
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

func (r *serviceOrder) Get(key string) (*models.ServiceOrder, error) {
	var serviceOrder models.ServiceOrder
	endPointURL := r.BuildURL(serviceOrdersEndpoint, key)

	resp, err := r.Request(r.Requester().Get, endPointURL, r.Options("GET", endPointURL))

	if err == nil {
		err = resp.UnmarshalData(&serviceOrder)
	}

	return &serviceOrder, err
}

func (r *serviceOrder) List() ([]models.ServiceOrder, *models.Paging, error) {
	var members []models.ServiceOrder
	endPointURL := r.BuildURL(serviceOrdersEndpoint)

	resp, err := r.Request(r.Requester().Get, endPointURL, r.Options("GET", endPointURL))

	if err != nil {
		return members, nil, err
	}

	err = resp.UnmarshalData(&members)

	return members, resp.Paging, err
}

func (r *serviceOrder) ByStoneCode(stonecode int) ([]models.ServiceOrders, *models.Paging, errors.Error) {
	var serviceOrders []models.ServiceOrder
	endPointURL := r.BuildURL(merchantServiceOrderEndpoint, key)

	resp, err := r.Request(r.Requester().Get, endPointURL, r.Options("GET", endPointURL))

	if err == nil {
		err = resp.UnmarshalData(&serviceOrders)
	}

	return nil, serviceOrders, err
}
