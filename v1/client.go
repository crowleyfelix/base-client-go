package client

//Client exposes the services to access api resources
type Client interface {
}

type client struct {
	url string
}

//New constructs a new client
func New(url string) Client {
	return &client{url}
}
