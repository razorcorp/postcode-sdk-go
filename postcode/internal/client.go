package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/praveenprem/postcode-sdk-go/model"
	"net/http"
)

/**
 * Package name: internal
 * Project name: postcode-sdk-go
 * Created by: Praveen Premaratne
 * Created on: 05/09/2021 22:48
 */

var API = "https://api.postcodes.io"

type (
	key struct {
		Key   string
		Value string
	}
	Header key
	Query  key
	client struct {
		Url     string
		Headers []Header
		Query   []Query
		req     http.Request
	}

	Http interface {
		Do() (responses []http.Response, error *model.ResponseError)
		Request(method, uri string, payload []byte) error
	}
)

func Client() *client {
	return &client{
		Url: API,
		Headers: []Header{
			{Key: "Content-Type", Value: "application/json"},
			{Key: "Accept", Value: "application/json"},
		},
	}
}

func (c *client) Do() (responses *http.Response, error *model.ResponseError) {
	htClient := new(http.Client)
	resp, err := htClient.Do(&c.req)
	if err != nil {
		return nil, &model.ResponseError{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}
	}

	if resp.StatusCode != 200 {
		errorResponse := new(model.ResponseError)
		if decodeErr := json.NewDecoder(resp.Body).Decode(&errorResponse); decodeErr != nil {
			return nil, &model.ResponseError{
				Status: http.StatusInternalServerError,
				Error:  decodeErr.Error(),
			}
		}

		return nil, errorResponse
	}

	return resp, nil
}

func (c *client) Request(method, uri string, payload []byte) error {
	req, err := http.NewRequest(method, fmt.Sprintf("%s/%s", c.Url, uri), bytes.NewBuffer(payload))
	if err != nil {
		return err
	}

	for _, header := range c.Headers {
		req.Header.Add(header.Key, header.Value)
	}

	params := req.URL.Query()
	for _, query := range c.Query {
		params.Set(query.Key, query.Value)
	}
	req.URL.RawQuery = params.Encode()

	c.req = *req
	return nil
}
