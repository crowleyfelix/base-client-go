package models

type ServiceOrderFilters struct {
	Page                      string `json:"page"`
	Size                      string `json:"size"`
	DocumentNumber            string `json:"document_number"`
	City                      string `json:"city"`
	Stonecode                 string `json:"stonecode"`
	Contractor                string `json:"contractor"`
	OperatorDocumentNumber    string `json:"operator_document_number"`
	State                     string `json:"state"`
	ClientGroup               string `json:"client_group"`
	ClientService             string `json:"client_service"`
	Model                     string `json:"model"`
	SerialNumber              string `json:"serial_number"`
	ServiceOrderNumber        string `json:"service_order_number"`
	Provider                  string `json:"provider"`
	Reference                 string `json:"reference"`
	Service                   string `json:"service"`
	Status                    string `json:"status"`
	Sync                      string `json:"sync"`
	StartDateLastModification string `json:"start_date_last_modification"`
	EndDateLastModification   string `json:"end_date_last_modification"`
}
