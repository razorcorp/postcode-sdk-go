package model

/**
 * Package name: terminated
 * Project name: postcode-sdk-go
 * Created by: Praveen Premaratne
 * Created on: 05/09/2021 02:54
 */

type (
	TerminatedPostcode struct {
		Postcode        string  `json:"postcode"`
		YearTerminated  int64   `json:"year_terminated"`
		MonthTerminated int64   `json:"month_terminated"`
		Longitude       float64 `json:"longitude"`
		Latitude        float64 `json:"latitude"`
	}
)
