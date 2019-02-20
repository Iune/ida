package main

import (
	"fmt"

	"github.com/iune/ida/contest"
)

func main() {
	countries := contest.LoadCountries("countries.tsv")
	contest := contest.LoadContest("1991.csv", countries)
	fmt.Println(contest.Voters)
}
