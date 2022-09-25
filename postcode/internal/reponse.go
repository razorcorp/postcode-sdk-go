package internal

import (
	"encoding/json"
	"fmt"
	"github.com/razorcorp/postcode-sdk-go/model"
	"io"
	"log"
	"net/http"
)

/**
 * Package name: internal
 * Project name: postcode-sdk-go
 * Created by: Praveen Premaratne
 * Created on: 05/09/2021 03:23
 */

type (
	responseWrapper struct {
		Status int
		Result interface{}
		Error  string `json:"error"`
	}
)

func RequestBuildError(err error) *model.ResponseError {
	return &model.ResponseError{
		Status: http.StatusInternalServerError,
		Error:  fmt.Sprintf("Failed to build request: %s", err.Error()),
	}
}

func PayloadEncodeError(err error) *model.ResponseError {
	return &model.ResponseError{
		Status: http.StatusInternalServerError,
		Error:  fmt.Sprintf("Failed to encode the request body: %s", err.Error()),
	}
}

func ResponseDecodeError(err error) *model.ResponseError {
	return &model.ResponseError{
		Status: http.StatusInternalServerError,
		Error:  fmt.Sprintf("Failed to parse response body: %s", err.Error()),
	}
}

//ResponseDecoder transpose given HTTP response body into the given interface
//
//Parameters:
//
//	body: io.Reader instance of the HTTP response
//
//	model: interface pointer for a Go struct
func ResponseDecoder(body io.Reader, iface interface{}) *model.ResponseError {
	data := responseWrapper{}
	if err := json.NewDecoder(body).Decode(&data); err != nil {
		return ResponseDecodeError(err)
	}

	if data.Status >= 400 {
		return &model.ResponseError{
			Status: data.Status,
			Error:  data.Error,
		}
	}

	jsonObject, marshalErr := json.Marshal(data.Result)
	if marshalErr != nil {
		return ResponseDecodeError(marshalErr)
	}

	if err := json.Unmarshal(jsonObject, iface); err != nil {
		log.Printf("%#v", err)
		return ResponseDecodeError(err)
	}

	return nil
}
