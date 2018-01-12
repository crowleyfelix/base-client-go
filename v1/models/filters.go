package models

//ServiceOrderFilters contains filter fields to service order request
type ServiceOrderFilters struct {
	Page                      *int    `param:"page"`
	Size                      *int    `param:"size"`
	DocumentNumber            *string `param:"document_number"`
	City                      *string `param:"city"`
	StoneCode                 *string `param:"stonecode"`
	Contractor                *string `param:"contractor"`
	OperatorDocumentNumber    *string `param:"operator_document_number"`
	State                     *string `param:"state"`
	ClientGroup               *string `param:"client_group"`
	ServiceGroup              *string `param:"service_group"`
	Model                     *string `param:"model"`
	SerialNumber              *string `param:"serial_number"`
	ServiceNumber             *int    `param:"service_number"`
	Provider                  *string `param:"provider"`
	Reference                 *string `param:"reference"`
	Service                   *string `param:"service"`
	Status                    *string `param:"status"`
	Sync                      *string `param:"sync"`
	StartDateLastModification *string `param:"start_date_last_modification"`
	EndDateLastModification   *string `param:"end_date_last_modification"`
}
