# Postcode.io SDK for Go

`postcode-sdk-go` is the UK Postcode SDK from https://postcodes.io for the Go programming language

The SDK requires minimum version of `Go 1.16`

## Service Maintenance

For information on service maintenance please refer to the service provider https://postcodes.io.
This is a simple SDK writen to access the postcode service provided by IDDQD Limited under Open Source service postcodes.io.

## Getting started

To get started working with the SDK setup your project for Go modules, and retrieve the SDK dependency with `go get`

### Example project setup

This example shows how you can use the `postcode-sdk-go` SDK to make an API request using the SDK's Postcode Lookup service.

#### Initialize Project
```shell
mkdir ~/postcode-lookup
cd ~/postcode-lookup
go mod init postcode-lookup
```

#### Add SDK Dependency
```shell
go get github.com/praveenprem/postcode-sdk-go
```

#### Write Code
```go
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
```

#### Compile and Execute
```shell
go run .
````

#### Output
```
&model.Postcode{Postcode:"OX1 2JD", OutCode:"OX1", InCode:"2JD", Quality:1, Eastings:451034, Northings:206852,
 Country:"England", NhsHa:"South Central", AdminCounty:"Oxfordshire", AdminDistrict:"Oxford",
  AdminWard:"Carfax & Jericho", Longitude:-1.26201, Latitude:51.758038, ParliamentaryConstituency:"Oxford East",
   EuropeanElectoralRegion:"South East", PrimaryCareTrust:"Oxfordshire", Region:"South East",
    Parish:"Oxford, unparished area", Lsoa:"Oxford 008A", Msoa:"Oxford 008", Ced:"Jericho and Osney",
     Ccg:"NHS Oxfordshire", Nuts:"Oxfordshire CC", Distance:0, Codes:model.Codes{AdminDistrict:"E07000178",
      AdminCounty:"E10000025", AdminWard:"E05013098", Parish:"E43000128", ParliamentaryConstituency:"E14000873",
       Ccg:"E38000136", CcgId:"10Q", CcgCode:"", Ced:"E58001257", Nuts:"TLJ14", Lau2:"E07000178", Lsoa:"E01028521",
        Msoa:"E02005947"}}
```

> More examples available in the [example/postcode/main.go](example/postcode/main.go)
