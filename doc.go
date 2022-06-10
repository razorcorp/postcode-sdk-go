/*
Package postcode_sdk_go is the UK Postcode SDK from https://postcodes.io for the Go programming language

Getting started

To get started working with the SDK setup your project for Go modules, and retrieve the SDK dependency with `go get`

	go get github.com/praveenprem/postcode-sdk-go

Postcode Lookup

This example shows how you can use the postcode_sdk_go SDK to make an API request using the SDK's Lookup method

	package main

	import (
		"fmt"
		"github.com/praveenprem/postcode-sdk-go/postcode"
	)
	
	func main() {
		data, lookupError := postcode.Lookup("OX12JD")
		if lookupError != nil {
			panic(fmt.Sprintf("%#v", lookupError))
		}
		fmt.Printf("%#v\n", data)
	}
*/

package postcode_sdk_go
