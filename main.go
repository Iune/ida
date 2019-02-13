package main

import (
	"fmt"

	"github.com/iune/ida/country"
)

func main() {
	countries := country.LoadCountries("countries.tsv")
	for _, c := range countries {
		fmt.Printf("%s | %s | %s\n", c.Iso, c.Forum, c.GetPrimaryName())
	}
}
