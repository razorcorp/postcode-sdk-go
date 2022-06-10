package model

/**
 * Package name: model
 * Project name: postcode-sdk-go
 * Created by: Praveen Premaratne
 * Created on: 28/05/2022 01:41
 */

type (
	ResponseError struct {
		Status int    `json:"status"`
		Error  string `json:"error"`
	}
)
