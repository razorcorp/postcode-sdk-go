package model

/**
 * Package name: codes
 * Project name: postcode-sdk-go
 * Created by: Praveen Premaratne
 * Created on: 05/09/2021 03:09
 */

type Codes struct {
	AdminDistrict             string `json:"admin_district,omitempty"`
	AdminCounty               string `json:"admin_county,omitempty"`
	AdminWard                 string `json:"admin_ward,omitempty"`
	Parish                    string `json:"parish,omitempty"`
	ParliamentaryConstituency string `json:"parliamentary_constituency,omitempty"`
	Ccg                       string `json:"ccg,omitempty"`
	CcgId                     string `json:"ccg_id,omitempty"`
	CcgCode                   string `json:"ccg_code,omitempty"`
	Ced                       string `json:"ced,omitempty"`
	Nuts                      string `json:"nuts,omitempty"`
	Lau2                      string `json:"lau2,omitempty"`
	Lsoa                      string `json:"lsoa,omitempty"`
	Msoa                      string `json:"msoa,omitempty"`
}
