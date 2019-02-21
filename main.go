package main

import (
	"fmt"

	"github.com/iune/ida/contest"
	"github.com/iune/ida/voting"
)

func main() {
	countries := contest.LoadCountries("countries.tsv")
	contest := contest.LoadContest("1991.csv", countries)

	lines := []string{
		":12: :yugoslavia: Yugoslavia: Bebi Doll - Brazil",
		":10: :sweden: Sweden: Carola - Fångad av en stormvind",
		":8: :france: France: Amina - C'est le dernier qui a parlé raison",
		"Hello world",
		":cyprus: Cyprus: Elena Patroklou - SOS",
		":7: UK: Samantha Janus - Message to Your Heart",
	}
	votes := voting.Find(contest, lines)
	fmt.Println(votes)
}
