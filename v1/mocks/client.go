// Code generated by mockery v1.0.0
package mocks

import mock "github.com/stretchr/testify/mock"
import services "github.com/stone-payments/gologistic/v1/services"

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// ServiceOrder provides a mock function with given fields: user
func (_m *Client) ServiceOrder(user string) services.ServiceOrder {
	ret := _m.Called(user)

	var r0 services.ServiceOrder
	if rf, ok := ret.Get(0).(func(string) services.ServiceOrder); ok {
		r0 = rf(user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(services.ServiceOrder)
		}
	}

	return r0
}
