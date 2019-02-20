package contest

import (
	"bufio"
	"encoding/csv"
	"log"
	"os"
)

type Contest struct {
	Entries []Entry
	Voters  []string
}

func LoadContest(contestFilePath string, countries []Country) Contest {
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

	voters := getVoters(csvFileContents)
	entries := getEntries(csvFileContents, countries)

	return Contest{Entries: entries, Voters: voters}
}

func getVoters(csv [][]string) []string {
	var voters []string
	if len(csv[0]) >= 6 {
		voters = csv[0][6:len(csv[0])]
	} else {
		log.Fatal("No voters were defined in contest file")
	}
	return voters
}

func getEntries(csv [][]string, countries []Country) []Entry {
	var entries []Entry
	for _, line := range csv[1:] {
		country, found := GetCountry(countries, line[1])
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
	return entries
}
