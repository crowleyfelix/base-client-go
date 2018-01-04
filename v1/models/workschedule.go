package models

type WorkSchedule struct {
	WeekDay         int    `json:"weekDay"`
	IsNonWorkingDay bool   `json:"isNonWorkingDay"`
	OpenTime        string `json:"openTime,omitempty"`
	CloseTime       string `json:"closeTime,omitempty"`
}
