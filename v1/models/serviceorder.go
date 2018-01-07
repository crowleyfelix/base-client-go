package models

type ServiceOrder struct {
	Allowance          bool   `json:"allowance"`
	CancellationReason string `json:"cancellation_reason"`
	Class              string `json:"class_"`
	Client             struct {
		Address struct {
			City           string `json:"city"`
			Complement     string `json:"complement"`
			Country        string `json:"country"`
			District       string `json:"district"`
			DistrictNumber string `json:"district_number"`
			DistrictType   string `json:"district_type"`
			Neighborhood   string `json:"neighborhood"`
			Region         string `json:"region"`
			State          string `json:"state"`
			Zipcode        string `json:"zipcode"`
		} `json:"address"`
		Contact        string `json:"contact"`
		DocumentNumber string `json:"document_number"`
		DocumentType   string `json:"document_type"`
		Name           string `json:"name"`
		OpeningHours   string `json:"opening_hours"`
		PhoneNumber    string `json:"phone_number"`
		StoneCode      string `json:"stone_code"`
	} `json:"client"`
	Contractor       string `json:"contractor"`
	DateAttendance   string `json:"date_attendance"`
	DateModification string `json:"date_modification"`
	Kit              struct {
	} `json:"kit"`
	LimitDate       string `json:"limit_date"`
	LimitDateClient string `json:"limit_date_client"`
	NewChip         struct {
	} `json:"new_chip"`
	NewTerminal struct {
		Model      string `json:"model"`
		Technology string `json:"technology"`
	} `json:"new_terminal"`
	NumberOfRevisits int `json:"number_of_revisits"`
	OldChip          struct {
	} `json:"old_chip"`
	OldTerminal struct {
	} `json:"old_terminal"`
	OpeningDate   string `json:"opening_date"`
	OpeningRemark string `json:"opening_remark"`
	Operator      struct {
	} `json:"operator"`
	OperatorCompany string `json:"operator_company"`
	Provider        string `json:"provider"`
	Service         string `json:"service"`
	ServiceGroup    string `json:"service_group"`
	ServiceNumber   int    `json:"service_number"`
	Status          string `json:"status"`
}
