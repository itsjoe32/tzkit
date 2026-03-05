# tzkit

A lightweight Go package for working with IANA timezones and countries.

## Install

```
go get github.com/itsjoe32/tzkit
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/itsjoe32/tzkit"
)

func main() {
    // Look up a timezone
    tz, ok := tzkit.GetTimezone("America/New_York")
    if ok {
        fmt.Println(tz.SDTAbbreviation) // "EST"
        fmt.Println(tz.HasDST())        // true
    }

    // Look up a country
    country, ok := tzkit.GetCountry("US")
    if ok {
        fmt.Println(country.Name) // "United States"
    }

    // Cross-reference
    tzs := tzkit.TimezonesByCountry("JP")
    countries := tzkit.CountriesByTimezone("Asia/Tokyo")

    // List all
    ids := tzkit.TimezoneIDs()     // sorted
    codes := tzkit.CountryCodes()  // sorted

    fmt.Println(len(ids), len(codes), len(tzs), len(countries))
}
```