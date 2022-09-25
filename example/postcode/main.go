package main

import (
	"fmt"
	"github.com/razorcorp/postcode-sdk-go"
	"github.com/razorcorp/postcode-sdk-go/postcode"
)

/**
 * Package name: main
 * Project name: postcode-sdk-go
 * Created by: Praveen Premaratne
 * Created on: 05/09/2021 15:50
 */

func executor(title string, f func()) {
	fmt.Printf("===== %s =====\n", title)
	f()
	for i := 0; i < 100; i++ {
		fmt.Print("*")
	}
	fmt.Printf("\n\n")
}

func postcodeLookup() {
	data, lookupError := postcode.Lookup("RG122PE")
	if lookupError != nil {
		panic(fmt.Sprintf("%#v", lookupError))
	}
	fmt.Printf("%#v\n", data)
}

func bulkPostcodeLookup() {
	postcodes := postcode.Postcodes{Postcodes: []string{"RG122PE", "GU479DZ"}}
	data, lookupError := postcode.BulkLookup(postcodes, nil)
	if lookupError != nil {
		panic(fmt.Sprintf("%#v", lookupError))
	}
	fmt.Printf("%#v\n", data)
}

func reverseGeocoding() {
	geocode := &postcode.Geocode{
		Longitude: 0.629834723775309,
		Latitude:  51.7923246977375,
		Radius:    100,
	}
	data, lookupError := postcode.ReverseGeocoding(*geocode)
	if lookupError != nil {
		panic(fmt.Sprintf("%#v", lookupError))
	}
	fmt.Printf("%#v\n", len(data))
	fmt.Printf("%#v\n", data)
}

func bulkReverseGeocoding() {
	geocodes := &postcode.Geocodes{Geolocations: []postcode.Geocode{
		{
			Longitude: -0.740895,
			Latitude:  51.417093,
			Radius:    0,
		},
		{
			Longitude: -0.797388,
			Latitude:  51.343969,
			Radius:    0,
		},
	}}
	data, lookupError := postcode.BulkReverseGeocoding(*geocodes, nil)
	if lookupError != nil {
		panic(fmt.Sprintf("%#v", lookupError))
	}
	fmt.Printf("%#v\n", data)
}

func queryPostcode() {
	limit := int64(3)
	data, lookupError := postcode.Query("RG122P", &limit)
	if lookupError != nil {
		panic(fmt.Sprintf("%#v", lookupError))
	}
	fmt.Printf("%#v\n", len(data))
	fmt.Printf("%#v\n", data)
}

func validatePostcode() {
	data, lookupError := postcode.Validation("RG122PE")
	if lookupError != nil {
		panic(fmt.Sprintf("%#v", lookupError))
	}
	fmt.Printf("%#v\n", data)
}

func nearestPostcode() {
	limit := int64(5)
	radius := int64(200)
	data, lookupError := postcode.NearestPostcode("RG122PE", &limit, &radius)
	if lookupError != nil {
		panic(fmt.Sprintf("%#v", lookupError))
	}
	fmt.Printf("%#v\n", len(data))
	fmt.Printf("%#v\n", data)
}

func autocompletePostcode() {
	limit := int64(5)
	data, lookupError := postcode.Autocomplete("RG122P", &limit)
	if lookupError != nil {
		panic(fmt.Sprintf("%#v", lookupError))
	}
	fmt.Printf("%#v\n", len(data))
	fmt.Printf("%#v\n", data)
}

func randomPostcode() {
	outCode := "RG12"
	data, lookupError := postcode.RandomPostcode(&outCode)
	if lookupError != nil {
		panic(fmt.Sprintf("%#v", lookupError))
	}
	fmt.Printf("%#v\n", data)
}

func outcodeLookup() {
	data, lookupError := postcode.OutcodeLookup("RG42")
	if lookupError != nil {
		panic(fmt.Sprintf("%#v", lookupError))
	}
	fmt.Printf("%#v\n", data)
}

func outcodeReverseGeocoding() {
	geocode := &postcode.Geocode{
		Longitude: -0.740895,
		Latitude:  51.417093,
		Radius:    10000,
	}
	data, lookupError := postcode.OutcodeReverseGeocoding(*geocode)
	if lookupError != nil {
		panic(fmt.Sprintf("%#v", lookupError))
	}
	fmt.Printf("%#v\n", data)
}

func nearestOutcode() {
	limit := int64(5)
	radius := int64(200)
	data, lookupError := postcode.NearestOutcode("RG12", &limit, &radius)
	if lookupError != nil {
		panic(fmt.Sprintf("%#v", lookupError))
	}
	fmt.Printf("%#v\n", len(data))
	fmt.Printf("%#v\n", data)
}

func scottishPostcodeLookup() {
	data, lookupError := postcode.ScottishPostcodeLookup("EH12 8NF")
	if lookupError != nil {
		panic(fmt.Sprintf("%#v", lookupError))
	}
	fmt.Printf("%#v\n", data)
}

func terminatedPostcodeLookup() {
	data, lookupError := postcode.TerminatedPostcodeLookup("E1W 1UU")
	if lookupError != nil {
		panic(fmt.Sprintf("%#v", lookupError))
	}
	fmt.Printf("%#v\n", data)
}

func placeLookup() {
	data, lookupError := postcode.PlaceLookup("osgb4000000074546114")
	if lookupError != nil {
		panic(fmt.Sprintf("%#v", lookupError))
	}
	fmt.Printf("%#v\n", data)
}

func placeQuery() {
	data, lookupError := postcode.PlaceQuery("Bracknell", nil)
	if lookupError != nil {
		panic(fmt.Sprintf("%#v", lookupError))
	}
	fmt.Printf("%#v\n", data)
}

func randomPlace() {
	data, lookupError := postcode.RandomPlace()
	if lookupError != nil {
		panic(fmt.Sprintf("%#v", lookupError))
	}
	fmt.Printf("%#v\n", data)
}

func main() {
	fmt.Printf("Version: %s\n", postcode.VERSION)

	executor("Singe postcode lookup", postcodeLookup)

	executor("Bulk postcode lookup", bulkPostcodeLookup)

	executor("Reverse Geocoding lookup", reverseGeocoding)

	executor("Bulk reverse Geocoding lookup", bulkReverseGeocoding)

	executor("Query postcode", queryPostcode)

	executor("Validate postcode", validatePostcode)

	executor("Nearest postcode", nearestPostcode)

	executor("Autocomplete postcode", autocompletePostcode)

	executor("Random postcode", randomPostcode)

	executor("Outcode lookup", outcodeLookup)

	executor("Outcode Reverse Geocoding lookup", outcodeReverseGeocoding)

	executor("Nearest Outcode", nearestOutcode)

	executor("Scottish postcode lookup", scottishPostcodeLookup)

	executor("Terminated postcode lookup", terminatedPostcodeLookup)

	executor("Place lookup", placeLookup)

	executor("Place query lookup", placeQuery)

	executor("Random place", randomPlace)
}
