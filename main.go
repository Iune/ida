package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/alexflint/go-arg"
	"github.com/atotto/clipboard"

	"github.com/iune/ida/contest"
	"github.com/iune/ida/results"
	"github.com/iune/ida/voting"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Argument parsing
	var args struct {
		Countries   string `arg:"positional" help:"Path to tab-separated file with country information"`
		Spreadsheet string `arg:"positional" help:"Path to tab-separated contest file"`
	}
	arg.MustParse(&args)
	if len(args.Countries) == 0 {
		log.Fatal("Input file path for the countries file was empty")
	}
	if len(args.Spreadsheet) == 0 {
		log.Fatal("Input file path for the contest file was empty")
	}

	// Set up logging
	log.SetOutput(os.Stdin)
	log.SetLevel(log.WarnLevel)

	countries := contest.LoadCountries(args.Countries)
	contest := contest.LoadContest(args.Spreadsheet, countries)

	repeat := true
	for repeat {
		// Get voter name
		fmt.Print("Country Name:\n> ")
		input := bufio.NewScanner(os.Stdin)
		voterName := input.Text()
		if err := input.Err(); err != nil {
			log.Fatal(err)
		}

		// Read in votes from standard input
		var lines []string
		input = bufio.NewScanner(os.Stdin)
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

		// Check to see if we should continue the loop
		fmt.Print("Continue? (Y/N)\n> ")
		input = bufio.NewScanner(os.Stdin)
		input.Scan()
		if err := input.Err(); err != nil {
			log.Fatal(err)
		}

		text := input.Text()
		if strings.ToLower(text) == "n" {
			repeat = false
		}
		fmt.Println()
	}
}
