package contest

import (
	"bufio"
	"encoding/csv"
	"log"
	"os"

	"github.com/iune/ida/country"
)

type Contest struct {
	Entries []Entry
	Voters  []string
}

func LoadContest(contestFilePath string, countries []country.Country) Contest {
	csvFile, err := os.Open(contestFilePath)
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))
	reader.FieldsPerRecord = -1
	reader.Comma = '\t'
	csvFileContents, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Get list of voters
	var voters []string
	if len(csvFileContents[0]) >= 6 {
		voters = csvFileContents[0][6:len(csvFileContents[0])]
	} else {
		log.Fatal("No voters were defined in contest file")
	}

	// Get list of entries
	var entries []Entry
	for _, line := range csvFileContents[1:] {
		country, found := country.GetCountry(countries, line[1])
		if !found {
			log.Fatalf("Could not find country %s in countries file", line[1])
		}
		artist := line[3]
		song := line[4]

		entries = append(entries, Entry{
			Country: country,
			Artist:  artist,
			Song:    song,
		})
	}

	return Contest{Entries: entries, Voters: voters}
}
