package model

/**
 * Package name: postcodes
 * Project name: postcode-sdk-go
 * Created by: Praveen Premaratne
 * Created on: 05/09/2021 02:53
 */

type (
	Geocode struct {
		Latitude  float64 `json:"latitude,omitempty"`
		Longitude float64 `json:"longitude,omitempty"`

		//Limit  number of postcodes matches to return.
		//Defaults to 10.
		//Needs to be less than 100.
		Limit int64 `json:"limit,omitempty"`

		//Radius Limits number of postcodes matches to return.
		//Defaults to 100m.
		//Needs to be less than 2,000m.
		Radius int64 `json:"radius,omitempty"`

		//WideSearch Search up to 20km radius, but subject to a maximum of 10 results.
		//When enabled, radius and limits over 10 are ignored.
		WideSearch bool `json:"widesearch,omitempty"`
	}

	Geocodes struct {
		Query    Geocode    `json:"query"`
		Postcode []Postcode `json:"result"`
	}

	Postcodes struct {
		Query    string   `json:"query"`
		Postcode Postcode `json:"result"`
	}

	Postcode struct {
		Postcode                  string  `json:"postcode,omitempty"`
		OutCode                   string  `json:"outcode,omitempty"`
		InCode                    string  `json:"incode,omitempty"`
		Quality                   int64   `json:"quality,omitempty"`
		Eastings                  int64   `json:"eastings,omitempty"`
		Northings                 int64   `json:"northings,omitempty"`
		Country                   string  `json:"country,omitempty"`
		NhsHa                     string  `json:"nhs_ha,omitempty"`
		AdminCounty               string  `json:"admin_county,omitempty"`
		AdminDistrict             string  `json:"admin_district,omitempty"`
		AdminWard                 string  `json:"admin_ward,omitempty"`
		Longitude                 float64 `json:"longitude,omitempty"`
		Latitude                  float64 `json:"latitude,omitempty"`
		ParliamentaryConstituency string  `json:"parliamentary_constituency,omitempty"`
		EuropeanElectoralRegion   string  `json:"european_electoral_region,omitempty"`
		PrimaryCareTrust          string  `json:"primary_care_trust,omitempty"`
		Region                    string  `json:"region,omitempty"`
		Parish                    string  `json:"parish,omitempty"`
		Lsoa                      string  `json:"lsoa,omitempty"`
		Msoa                      string  `json:"msoa,omitempty"`
		Ced                       string  `json:"ced,omitempty"`
		Ccg                       string  `json:"ccg,omitempty"`
		Nuts                      string  `json:"nuts,omitempty"`
		Distance                  float64 `json:"distance,omitempty"`
		Codes                     Codes   `json:"codes,omitempty"`
	}
)
