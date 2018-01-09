package services

const (
	version                       = "/v1"
	serviceOrdersEndpoint         = version + "/serviceorders"
	serviceOrderEndpoint          = serviceOrdersEndpoint + "/%v"
	merchantServiceOrdersEndpoint = version + "/merchants/%v/serviceorders"
)
