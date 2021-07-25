package main

import (
	"github.com/Iune/ida/pkg/app"
)

func main() {
	// loadedContest := contest.LoadContest("resources/sample.json")
	// fmt.Println(loadedContest)

	// countries := contest.LoadCountries("resources/countries.json")
	// entries := contest.LoadEntries("resources/entries.txt", countries)
	app.SaveContestDetails("resources/countries.json", "resources/entries.txt", "resources/output")
}
