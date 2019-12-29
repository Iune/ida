package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/alexflint/go-arg"
	"github.com/atotto/clipboard"

	"github.com/Iune/ida/contest"
	"github.com/Iune/ida/results"
	"github.com/Iune/ida/voting"

	log "github.com/sirupsen/logrus"
)

type args struct {
	Countries   string `arg:"positional" help:"Path to JSON file with country information"`
	Spreadsheet string `arg:"positional" help:"Path to contest Excel spreadsheet"`
}

func (args) Version() string {
	return "ida 0.2.0"
}

func main() {
	// Argument parsing
	var args args
	arg.MustParse(&args)

	if len(args.Countries) == 0 {
		log.Fatal("Input file path for the countries file was empty")
	}
	if len(args.Spreadsheet) == 0 {
		log.Fatal("Input file path for the contest file was empty")
	}

	// Set up logging
	log.SetOutput(os.Stdin)
	log.SetLevel(log.InfoLevel)

	countries := contest.LoadCountries(args.Countries)
	contest := contest.LoadContest(args.Spreadsheet, countries)

	repeat := true
	for repeat {
		// Get voter name
		fmt.Print("Country Name:\n> ")
		reader := bufio.NewReader(os.Stdin)
		voterName, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		// Read in votes from standard input
		var lines []string
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			lines = append(lines, input.Text())
		}
		if err := input.Err(); err != nil {
			log.Fatal(err)
		}
		fmt.Println()

		// Parse and print votes to standard output and to clipboard
		votes := voting.Find(contest, lines)
		outputString := results.Output(contest, votes, voterName)
		clipboard.WriteAll(outputString)

		results.PrintVotes(votes)
		fmt.Println()

	}
}