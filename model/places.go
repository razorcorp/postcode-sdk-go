package model

/**
 * Package name: places
 * Project name: postcode-sdk-go
 * Created by: Praveen Premaratne
 * Created on: 05/09/2021 02:53
 */

type (
	Place struct {
		Code                string  `json:"code"`
		Name1               string  `json:"name_1"`
		Name1Lang           string  `json:"name_1_lang"`
		Name2               string  `json:"name_2"`
		Name2Lang           string  `json:"name_2_lang"`
		LocalType           string  `json:"local_type"`
		Outcode             string  `json:"outcode"`
		CountyUnitary       string  `json:"county_unitary"`
		CountyUnitaryType   string  `json:"county_unitary_type"`
		DistrictBorough     string  `json:"district_borough"`
		DistrictBoroughType string  `json:"district_borough_type"`
		Region              string  `json:"region"`
		Country             string  `json:"country"`
		Longitude           float64 `json:"longitude"`
		Latitude            float64 `json:"latitude"`
		Eastings            int64   `json:"eastings"`
		Northings           int64   `json:"northings"`
		MinEastings         int64   `json:"min_eastings"`
		MinNorthings        int64   `json:"min_northings"`
		MaxEastings         int64   `json:"max_eastings"`
		MaxNorthings        int64   `json:"max_northings"`
	}
)
