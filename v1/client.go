package client

import (
	"github.com/stone-payments/logistic-sdk-go/v1/services"
)

//Client exposes the services to access logistic resources
type Client interface {
	ServiceOrder(user string) services.ServiceOrder
}

type client struct {
	url string
}

//New constructs a new client
func New(url string) Client {
	return &client{url}
}

func (c *client) ServiceOrder(user string) services.ServiceOrder {
	return services.NewServiceOrder(c.url)
}
