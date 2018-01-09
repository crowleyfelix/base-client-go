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
	manager http.Manager
}

//NewServiceOrder is constructs ServiceOrder Service
func NewServiceOrder(url string) ServiceOrder {
	return &serviceOrder{
		manager: http.NewManager(url, http.NewInterceptor(nil, processResponse)),
	}
}

func (r *serviceOrder) Get(id string) (*models.ServiceOrder, errors.Error) {
	var serviceOrder models.ServiceOrder
	endPointURL := r.manager.BuildURL(serviceOrderEndpoint, id)

	resp, err := r.manager.Request(http.Requester.Get, endPointURL, http.NewRequest())

	if err == nil {
		err = resp.JSON(&serviceOrder)
	}

	return &serviceOrder, err
}

func (r *serviceOrder) List(filters *models.ServiceOrderFilters) ([]models.ServiceOrder, errors.Error) {
	var serviceOrders []models.ServiceOrder
	endPointURL := r.manager.BuildURL(serviceOrdersEndpoint)

	options := http.NewRequest()
	options.SetParams(filters.ToMap())

	resp, err := r.manager.Request(http.Requester.Get, endPointURL, options)

	if err == nil {
		err = resp.JSON(&serviceOrders)
	}

	return serviceOrders, err
}

func (r *serviceOrder) ByStoneCode(stonecode int) ([]models.ServiceOrder, errors.Error) {
	var serviceOrders []models.ServiceOrder
	endPointURL := r.manager.BuildURL(merchantServiceOrdersEndpoint, stonecode)

	resp, err := r.manager.Request(http.Requester.Get, endPointURL, http.NewRequest())

	if err == nil {
		err = resp.JSON(&serviceOrders)
	}

	return serviceOrders, err
}
