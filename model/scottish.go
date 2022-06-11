package model

/**
 * Package name: scottish
 * Project name: postcode-sdk-go
 * Created by: Praveen Premaratne
 * Created on: 05/09/2021 02:53
 */

type (
	ScottishCodes struct {
		ScottishParliamentaryConstituency string `json:"scottish_parliamentary_constituency"`
	}

	ScottishPostcodeData struct {
		Postcode                          string        `json:"postcode"`
		ScottishParliamentaryConstituency string        `json:"scottish_parliamentary_constituency"`
		Codes                             ScottishCodes `json:"codes"`
	}
)
