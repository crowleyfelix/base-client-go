package client

import (
	"github.com/stone-payments/logistic-sdk-go/v1/services"
)

//MembershipClientable exposes the services to access membership resources
type MembershipClientable interface {
	Member(user string) services.MemberService
	SystemConfigurations(user string) services.SystemConfigurationsService
}

type client struct {
	url                  string
	secretKey            string
	clientApplicationKey string
}

//NewMembershipClient constructs a new membership client
func NewMembershipClient(url string, secretKey string, clientApplicationKey string) MembershipClientable {
	return &client{url, secretKey, clientApplicationKey}
}

func (c *client) Member(user string) services.MemberService {
	return services.NewMember(c.url, &services.Credentials{
		SecretKey:            c.secretKey,
		User:                 user,
		ClientApplicationKey: c.clientApplicationKey,
	})
}

func (c *client) SystemConfigurations(user string) services.SystemConfigurationsService {
	return services.NewSystemConfigurations(c.url, &services.Credentials{
		SecretKey:            c.secretKey,
		User:                 user,
		ClientApplicationKey: c.clientApplicationKey,
	})
}
