package model

/**
 * Package name: model
 * Project name: postcode-sdk-go
 * Created by: Praveen Premaratne
 * Created on: 28/05/2022 18:37
 */

type (
	OutcodeData struct {
		Outcode       string   `json:"outcode,omitempty"`
		Longitude     float64  `json:"longitude,omitempty"`
		Latitude      float64  `json:"latitude,omitempty"`
		Northings     int64    `json:"northings,omitempty"`
		Eastings      int64    `json:"eastings,omitempty"`
		AdminDistrict []string `json:"adminDistrict,omitempty"`
		Parish        []string `json:"parish,omitempty"`
		AdminCounty   []string `json:"adminCounty,omitempty"`
		AdminWard     []string `json:"adminWard,omitempty"`
		Country       []string `json:"country,omitempty"`
	}
)
