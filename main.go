package main

import (
	"github.com/iune/ida/contest"
	"github.com/iune/ida/country"
)

func main() {
	countries := country.LoadCountries("countries.tsv")
	contest := contest.LoadContest("1991.csv", countries)
}
