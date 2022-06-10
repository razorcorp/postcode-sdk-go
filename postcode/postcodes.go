package postcode

import (
	"encoding/json"
	"fmt"
	"github.com/praveenprem/postcode-sdk-go/model"
	"github.com/praveenprem/postcode-sdk-go/postcode/internal"
	"net/http"
	"strconv"
	"strings"
)

/**
 * Package name: postcodes
 * Project name: postcode-sdk-go
 * Created by: Praveen Premaratne
 * Created on: 05/09/2021 02:53
 */

type (
	Geocode model.Geocode

	Geocodes struct {
		Geolocations []Geocode `json:"geolocations"`
	}

	Postcodes struct {
		Postcodes []string `json:"postcodes"`
	}
)

//Lookup This uniquely identifies a postcode.
//
//	postcode: Valid postcode to match associated data
//
//Returns a single postcode entity for a given postcode (case, space insensitive).
func Lookup(postcode string) (*model.Postcode, *model.ResponseError) {
	client := internal.Client()
	if err := client.Request(http.MethodGet, fmt.Sprintf("postcodes/%s", postcode), nil); err != nil {
		return nil, internal.RequestBuildError(err)
	}

	response, responseError := client.Do()
	if responseError != nil {
		return nil, responseError
	}

	data := new(model.Postcode)
	if decodeErr := internal.ResponseDecoder(response.Body, data); decodeErr != nil {
		return nil, decodeErr
	}

	return data, nil
}

//BulkLookup Accepts an array of postcodes. Returns a list of matching
//postcodes and respective available data.
//
//	postcodes: List of postcodes
//
//	filter: (not required) A comma separated whitelist of attributes to be returned in the result object(s),
//	e.g. `filter=postcode,longitude,latitude`. `null` responses will continue to return `null`.
//	If no attributes match the result
//
//Accepts up to 100 postcodes
func BulkLookup(postcodes Postcodes, filters []string) ([]model.Postcodes, *model.ResponseError) {
	client := internal.Client()
	if err := postcodes.validate(); err != nil {
		return nil, err
	}

	payload, payloadErr := postcodes.json()
	if payloadErr != nil {
		return nil, internal.PayloadEncodeError(payloadErr)
	}

	if len(filters) > 0 {
		client.Query = append(client.Query, internal.Query{
			Key:   "filter",
			Value: strings.Join(filters, ","),
		})
	}

	if err := client.Request(http.MethodPost, "postcodes", payload); err != nil {
		return nil, internal.RequestBuildError(err)
	}

	response, responseError := client.Do()
	if responseError != nil {
		return nil, responseError
	}

	var data []model.Postcodes
	if decodeErr := internal.ResponseDecoder(response.Body, &data); decodeErr != nil {
		return nil, decodeErr
	}

	return data, nil
}

//ReverseGeocoding Returns the nearest postcodes for a given longitude and latitude.
//
//Optional Query Parameters:
//
//limit= (not required) Limits number of postcodes matches to return. Defaults to 10. Needs to be less than 100.
//
//radius= (not required) Limits number of postcodes matches to return. Defaults to 100m.
//Needs to be less than 2,000m.
//
//widesearch= (not required) Search up to 20km radius, but subject to a maximum of 10 results.
//Since lookups over a wide area can be very expensive, we've created this method to allow you choose to make
//the trade off between search radius and number of results. Defaults to false.
//When enabled, radius and limits over 10 are ignored.
func ReverseGeocoding(geocode Geocode) ([]model.Postcode, *model.ResponseError) {
	if err := geocode.validate(); err != nil {
		return nil, err
	}
	client := internal.Client()
	client.Query = append(client.Query, internal.Query{
		Key:   "lon",
		Value: strconv.FormatFloat(geocode.Longitude, 'f', 20, 64),
	})
	client.Query = append(client.Query, internal.Query{
		Key:   "lat",
		Value: strconv.FormatFloat(geocode.Latitude, 'f', 20, 64),
	})

	if geocode.Limit > 0 {
		client.Query = append(client.Query, internal.Query{
			Key:   "limit",
			Value: strconv.FormatInt(geocode.Limit, 10),
		})
	}
	if geocode.Radius > 0 {
		client.Query = append(client.Query, internal.Query{
			Key:   "radius",
			Value: strconv.FormatInt(geocode.Radius, 10),
		})
	}
	if geocode.WideSearch {
		client.Query = append(client.Query, internal.Query{
			Key:   "widesearch",
			Value: strconv.FormatBool(geocode.WideSearch),
		})
	}

	if err := client.Request(http.MethodGet, "postcodes", nil); err != nil {
		return nil, internal.RequestBuildError(err)
	}

	response, responseError := client.Do()
	if responseError != nil {
		return nil, responseError
	}

	data := new([]model.Postcode)
	if decodeErr := internal.ResponseDecoder(response.Body, data); decodeErr != nil {
		return nil, decodeErr
	}

	return *data, nil
}

//BulkReverseGeocoding Bulk translates geolocations into Postcodes. Accepts up to 100 geolocations.
//
//Optional Query Parameters:
//
//limit= (not required) Limits number of postcodes matches to return. Defaults to 10. Needs to be less than 100.
//
//radius= (not required) Limits number of postcodes matches to return. Defaults to 100m.
//Needs to be less than 2,000m.
//
//widesearch= (not required) Search up to 20km radius, but subject to a maximum of 10 results.
//Since lookups over a wide area can be very expensive, we've created this method to allow you choose to make
//the trade off between search radius and number of results. Defaults to false.
//When enabled, radius and limits over 10 are ignored.
func BulkReverseGeocoding(geocodes Geocodes, filters []string) ([]model.Geocodes, *model.ResponseError) {
	client := internal.Client()

	if err := geocodes.validate(); err != nil {
		return nil, err
	}

	payload, payloadErr := geocodes.json()
	if payloadErr != nil {
		return nil, internal.PayloadEncodeError(payloadErr)
	}

	if len(filters) > 0 {
		client.Query = append(client.Query, internal.Query{
			Key:   "filter",
			Value: strings.Join(filters, ","),
		})
	}

	if err := client.Request(http.MethodPost, "postcodes", payload); err != nil {
		return nil, internal.RequestBuildError(err)
	}

	response, responseError := client.Do()
	if responseError != nil {
		return nil, responseError
	}

	var data []model.Geocodes
	if decodeErr := internal.ResponseDecoder(response.Body, &data); decodeErr != nil {
		return nil, decodeErr
	}

	return data, nil
}

//Query Submit a postcode query and receive a complete list of postcode matches and all associated
//postcode data.
//
//This is essentially a postcode search which prefix matches and returns postcodes in sorted order
//(case insensitive)
//
//The result set can either be empty or populated with up to 100 postcode entities.
//
//Optional Query Parameters
//
//limit (not required) Limits number of postcodes matches to return. Defaults to 10. Needs to be less than 100.
func Query(postcode string, limit *int64) ([]model.Postcode, *model.ResponseError) {
	if limit != nil && *limit > 100 {
		return nil, &model.ResponseError{
			Status: http.StatusBadRequest,
			Error:  "Maximum limit exceeded! Limit must be less than 100",
		}
	}
	client := internal.Client()

	client.Query = append(client.Query, internal.Query{
		Key:   "q",
		Value: postcode,
	})

	if limit != nil {
		client.Query = append(client.Query, internal.Query{
			Key:   "limit",
			Value: strconv.FormatInt(*limit, 10),
		})
	}

	if err := client.Request(http.MethodGet, "postcodes", nil); err != nil {
		return nil, internal.RequestBuildError(err)
	}

	response, responseError := client.Do()
	if responseError != nil {
		return nil, responseError
	}

	data := new([]model.Postcode)
	if decodeErr := internal.ResponseDecoder(response.Body, &data); decodeErr != nil {
		return nil, decodeErr
	}

	return *data, nil
}

//Validation Convenience method to validate a postcode.
//
//Returns true or false (meaning valid or invalid respectively)
func Validation(postcode string) (bool, *model.ResponseError) {
	client := internal.Client()
	if err := client.Request(http.MethodGet, fmt.Sprintf("postcodes/%s/validate", postcode), nil); err != nil {
		return false, internal.RequestBuildError(err)
	}

	response, responseError := client.Do()
	if responseError != nil {
		return false, responseError
	}

	data := new(bool)
	if decodeErr := internal.ResponseDecoder(response.Body, &data); decodeErr != nil {
		return false, decodeErr
	}

	return *data, nil
}

//NearestPostcode Returns nearest postcodes for a given postcode.
//
//Optional Query Parameters
//
//limit= (not required) Limits number of postcodes matches to return. Defaults to 10. Needs to be less than 100.
//
//radius= (not required) Limits number of postcodes matches to return. Defaults to 100m. Needs to be less than 2,000m.
func NearestPostcode(postcode string, limit, radius *int64) ([]model.Postcode, *model.ResponseError) {
	if limit != nil && *limit > 100 {
		return nil, &model.ResponseError{
			Status: http.StatusBadRequest,
			Error:  "Maximum limit exceeded! Limit must be less than 100",
		}
	}

	if radius != nil && *radius > 2000 {
		return nil, &model.ResponseError{
			Status: http.StatusBadRequest,
			Error:  "Maximum limit exceeded! Limit must be less than 100",
		}
	}

	client := internal.Client()

	if limit != nil {
		client.Query = append(client.Query, internal.Query{
			Key:   "limit",
			Value: strconv.FormatInt(*limit, 10),
		})
	}

	if radius != nil {
		client.Query = append(client.Query, internal.Query{
			Key:   "radius",
			Value: strconv.FormatInt(*radius, 10),
		})
	}

	if err := client.Request(http.MethodGet, fmt.Sprintf("postcodes/%s/nearest", postcode), nil); err != nil {
		return nil, internal.RequestBuildError(err)
	}

	response, responseError := client.Do()
	if responseError != nil {
		return nil, responseError
	}

	data := new([]model.Postcode)
	if decodeErr := internal.ResponseDecoder(response.Body, &data); decodeErr != nil {
		return nil, decodeErr
	}

	return *data, nil
}

// Autocomplete Convenience method to return a list of matching postcodes.
//
//Optional Query Parameters
//	limit= (not required) Limits number of postcodes matches to return. Defaults to 10. Needs to be less than 100.
func Autocomplete(postcode string, limit *int64) ([]string, *model.ResponseError) {

	if limit != nil && *limit > 100 {
		return nil, &model.ResponseError{
			Status: http.StatusBadRequest,
			Error:  "Maximum limit exceeded! Limit must be less than 100",
		}
	}

	client := internal.Client()

	if limit != nil {
		client.Query = append(client.Query, internal.Query{
			Key:   "limit",
			Value: strconv.FormatInt(*limit, 10),
		})
	}

	if err := client.Request(http.MethodGet, fmt.Sprintf("postcodes/%s/autocomplete", postcode), nil); err != nil {
		return nil, internal.RequestBuildError(err)
	}

	response, responseError := client.Do()
	if responseError != nil {
		return nil, responseError
	}

	data := new([]string)
	if decodeErr := internal.ResponseDecoder(response.Body, &data); decodeErr != nil {
		return nil, decodeErr
	}

	return *data, nil
}

//RandomPostcode Returns a random postcode and all available data for that postcode.
//
//Optional Query Parameters
//	outcode= (not required) Filters random postcodes by outcode.
func RandomPostcode(outCode *string) (*model.Postcode, *model.ResponseError) {
	client := internal.Client()
	if outCode != nil {
		client.Query = append(client.Query, internal.Query{
			Key:   "outcode",
			Value: *outCode,
		})
	}
	if err := client.Request(http.MethodGet, "random/postcodes", nil); err != nil {
		return nil, internal.RequestBuildError(err)
	}

	response, responseError := client.Do()
	if responseError != nil {
		return nil, responseError
	}

	data := new(model.Postcode)
	if decodeErr := internal.ResponseDecoder(response.Body, data); decodeErr != nil {
		return nil, decodeErr
	}

	return data, nil
}

//OutcodeLookup Geolocation data for the centroid of the outward code specified.
//The outward code represents the first half of any postcode (separated by a space).
func OutcodeLookup(outCode string) (*model.OutcodeData, *model.ResponseError) {
	client := internal.Client()
	if err := client.Request(http.MethodGet, fmt.Sprintf("outcodes/%s", outCode), nil); err != nil {
		return nil, internal.RequestBuildError(err)
	}

	response, responseError := client.Do()
	if responseError != nil {
		return nil, responseError
	}

	data := new(model.OutcodeData)
	if decodeErr := internal.ResponseDecoder(response.Body, data); decodeErr != nil {
		return nil, decodeErr
	}

	return data, nil
}

//OutcodeReverseGeocoding Returns nearest outcodes for a given longitude and latitude.
//
//Optional Query Parameters
//	limit= (not required) Limits number of postcodes matches to return. Defaults to 10. Needs to be less than 100.
//	radius= (not required) Limits number of postcodes matches to return. Defaults to 5,000m.
//Needs to be less than 25,000m.
func OutcodeReverseGeocoding(geocode Geocode) ([]model.OutcodeData, *model.ResponseError) {
	if err := geocode.validate(); err != nil {
		return nil, err
	}
	client := internal.Client()
	client.Query = append(client.Query, internal.Query{
		Key:   "lon",
		Value: strconv.FormatFloat(geocode.Longitude, 'f', 20, 64),
	})
	client.Query = append(client.Query, internal.Query{
		Key:   "lat",
		Value: strconv.FormatFloat(geocode.Latitude, 'f', 20, 64),
	})

	if geocode.Limit > 0 {
		client.Query = append(client.Query, internal.Query{
			Key:   "limit",
			Value: strconv.FormatInt(geocode.Limit, 10),
		})
	}
	if geocode.Radius > 0 {
		client.Query = append(client.Query, internal.Query{
			Key:   "radius",
			Value: strconv.FormatInt(geocode.Radius, 10),
		})
	}

	if err := client.Request(http.MethodGet, "outcodes", nil); err != nil {
		return nil, internal.RequestBuildError(err)
	}

	response, responseError := client.Do()
	if responseError != nil {
		return nil, responseError
	}

	data := new([]model.OutcodeData)
	if decodeErr := internal.ResponseDecoder(response.Body, data); decodeErr != nil {
		return nil, decodeErr
	}

	return *data, nil
}

//NearestOutcode Returns nearest outcodes for a given outcode.
//
//Optional Query Parameters
//	limit= (not required) Limits number of postcodes matches to return. Defaults to 10. Needs to be less than 100.
//	radius= (not required) Limits number of postcodes matches to return. Defaults to 5,000m.
//Needs to be less than 25,000m.
func NearestOutcode(outCode string, limit, radius *int64) ([]model.OutcodeData, *model.ResponseError) {
	client := internal.Client()
	if limit != nil && *limit > 0 {
		client.Query = append(client.Query, internal.Query{
			Key:   "limit",
			Value: strconv.FormatInt(*limit, 10),
		})
	}
	if radius != nil && *radius > 0 {
		client.Query = append(client.Query, internal.Query{
			Key:   "radius",
			Value: strconv.FormatInt(*radius, 10),
		})
	}

	if err := client.Request(http.MethodGet, fmt.Sprintf("outcodes/%s/nearest", outCode), nil); err != nil {
		return nil, internal.RequestBuildError(err)
	}

	response, responseError := client.Do()
	if responseError != nil {
		return nil, responseError
	}

	data := new([]model.OutcodeData)
	if decodeErr := internal.ResponseDecoder(response.Body, data); decodeErr != nil {
		return nil, decodeErr
	}

	return *data, nil
}

//ScottishPostcodeLookup Lookup a Scottish postcode. Returns SPD data associated with postcode.
//At the moment this is just Scottish Parliamentary Constituency.
func ScottishPostcodeLookup(postcode string) (*model.ScottishPostcodeData, *model.ResponseError) {
	client := internal.Client()
	if err := client.Request(http.MethodGet, fmt.Sprintf("scotland/postcodes/%s", postcode), nil); err != nil {
		return nil, internal.RequestBuildError(err)
	}

	response, responseError := client.Do()
	if responseError != nil {
		return nil, responseError
	}

	data := new(model.ScottishPostcodeData)
	if decodeErr := internal.ResponseDecoder(response.Body, data); decodeErr != nil {
		return nil, decodeErr
	}

	return data, nil
}

//TerminatedPostcodeLookup Lookup a terminated postcode. Returns the postcode, year and month of termination.
func TerminatedPostcodeLookup(postcode string) (*model.TerminatedPostcode, *model.ResponseError) {
	client := internal.Client()
	if err := client.Request(http.MethodGet, fmt.Sprintf("terminated_postcodes/%s", postcode), nil); err != nil {
		return nil, internal.RequestBuildError(err)
	}

	response, responseError := client.Do()
	if responseError != nil {
		return nil, responseError
	}

	data := new(model.TerminatedPostcode)
	if decodeErr := internal.ResponseDecoder(response.Body, data); decodeErr != nil {
		return nil, decodeErr
	}

	return data, nil
}

//PlaceLookup Find a place by OSGB code (e.g. "osgb4000000074564391").
//Returns all available data if found.
func PlaceLookup(osgbCode string) (*model.Place, *model.ResponseError) {
	client := internal.Client()
	if err := client.Request(http.MethodGet, fmt.Sprintf("places/%s", osgbCode), nil); err != nil {
		return nil, internal.RequestBuildError(err)
	}

	response, responseError := client.Do()
	if responseError != nil {
		return nil, responseError
	}

	data := new(model.Place)
	if decodeErr := internal.ResponseDecoder(response.Body, data); decodeErr != nil {
		return nil, decodeErr
	}

	return data, nil
}

//PlaceQuery Submit a place query and receive a complete list of places matches and associated data.
//
//Optional Query Parameters
//	query: Place name to query and receive a complete list of places matches and associated data.
//	limit: (not required) Limits number of postcodes matches to return. Defaults to 10. Needs to be less than 100.
func PlaceQuery(query string, limit *int64) ([]model.Place, *model.ResponseError) {
	if limit != nil && *limit > 100 {
		return nil, &model.ResponseError{
			Status: http.StatusBadRequest,
			Error:  "Maximum limit exceeded! Limit must be less than 100",
		}
	}

	client := internal.Client()

	client.Query = append(client.Query, internal.Query{
		Key:   "q",
		Value: query,
	})

	if limit != nil && *limit > 0 {
		client.Query = append(client.Query, internal.Query{
			Key:   "limit",
			Value: strconv.FormatInt(*limit, 10),
		})
	}

	if err := client.Request(http.MethodGet, fmt.Sprintf("places"), nil); err != nil {
		return nil, internal.RequestBuildError(err)
	}

	response, responseError := client.Do()
	if responseError != nil {
		return nil, responseError
	}

	data := new([]model.Place)
	if decodeErr := internal.ResponseDecoder(response.Body, data); decodeErr != nil {
		return nil, decodeErr
	}

	return *data, nil
}

//RandomPlace Returns a random place and all associated data
func RandomPlace() (*model.Place, *model.ResponseError) {
	client := internal.Client()
	if err := client.Request(http.MethodGet, "random/places", nil); err != nil {
		return nil, internal.RequestBuildError(err)
	}

	response, responseError := client.Do()
	if responseError != nil {
		return nil, responseError
	}

	data := new(model.Place)
	if decodeErr := internal.ResponseDecoder(response.Body, data); decodeErr != nil {
		return nil, decodeErr
	}

	return data, nil
}

func (p Postcodes) json() ([]byte, error) {
	return json.Marshal(p)
}

func (p Postcodes) validate() *model.ResponseError {
	if len(p.Postcodes) == 0 {
		return &model.ResponseError{
			Status: http.StatusBadRequest,
			Error:  "minimum of 1 postcode required!",
		}
	}
	if len(p.Postcodes) > 100 {
		return &model.ResponseError{
			Status: http.StatusBadRequest,
			Error:  "Maximum postcode limit exceeded! Maximum of 100 postcodes",
		}
	}
	return nil
}

func (g Geocode) json() ([]byte, error) {
	return json.Marshal(g)
}

func (g Geocode) validate() *model.ResponseError {
	if g.Latitude == 0.0 || g.Longitude == 0.0 {
		return &model.ResponseError{
			Status: http.StatusBadRequest,
			Error:  "Latitude and Longitude must be defined",
		}
	}
	return nil
}

func (gs Geocodes) json() ([]byte, error) {
	return json.Marshal(gs)

}

func (gs Geocodes) validate() *model.ResponseError {
	if len(gs.Geolocations) == 0 {
		return &model.ResponseError{
			Status: http.StatusBadRequest,
			Error:  "minimum of 1 geolocations required!",
		}
	}

	if len(gs.Geolocations) > 100 {
		return &model.ResponseError{
			Status: http.StatusBadRequest,
			Error:  "Maximum geolocations limit exceeded! Maximum of 100 geolocations",
		}
	}

	for _, g := range gs.Geolocations {
		if err := g.validate(); err != nil {
			return err
		}
	}

	return nil
}
