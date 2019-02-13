package main

import (
	"fmt"
)

func main() {
	countries := LoadCountries("countries.json")
	for _, c := range countries.Countries {
		fmt.Printf("%s | %s | %s\n", c.Iso, c.Forum, c.GetPrimaryName())
	}
}
