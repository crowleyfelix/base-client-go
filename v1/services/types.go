package services

type responseError struct {
	ErrorList []errorData `json:"error_list"`
	Success   bool        `json:"success"`
}

type errorData struct {
	Code   int    `json:"code"`
	Reason string `json:"reason"`
}
