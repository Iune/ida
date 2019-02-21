package main

import (
	"os"

	"github.com/iune/ida/contest"
	"github.com/iune/ida/results"
	"github.com/iune/ida/voting"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Set up logging
	log.SetOutput(os.Stdout)
	log.SetLevel(log.WarnLevel)

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
	results.Output(contest, votes, "Yugoslavia")
	results.PrintVotes(votes)
}
