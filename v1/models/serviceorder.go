package models

type ServiceOrder struct {
	Allowance          bool   `json:"allowance"`
	CancellationReason string `json:"cancellation_reason"`
	Class              string `json:"class_"`
	Client             Client `json:"client"`
	Contractor         string `json:"contractor"`
	DateAttendance     string `json:"date_attendance"`
	DateModification   string `json:"date_modification"`
	Kit                struct {
	} `json:"kit"`
	LimitDate       string `json:"limit_date"`
	LimitDateClient string `json:"limit_date_client"`
	NewChip         struct {
	} `json:"new_chip"`
	NewTerminal      Terminal `json:"new_terminal"`
	NumberOfRevisits int      `json:"number_of_revisits"`
	OldChip          struct {
	} `json:"old_chip"`
	OldTerminal     Terminal `json:"old_terminal"`
	OpeningDate     string   `json:"opening_date"`
	OpeningRemark   string   `json:"opening_remark"`
	Operator        Operator `json:"operator"`
	OperatorCompany string   `json:"operator_company"`
	Provider        string   `json:"provider"`
	Service         string   `json:"service"`
	ServiceGroup    string   `json:"service_group"`
	ServiceNumber   int      `json:"service_number"`
	Status          string   `json:"status"`
}

type ServiceOrdersPage struct {
	ServiceOrders []ServiceOrder     `json:"orders_list"`
	Paging        ServiceOrderPaging `json:"page"`
}

type ServiceOrderPaging struct {
	Paging
	TotalOfOrders int `json:"total_of_orders"`
}
