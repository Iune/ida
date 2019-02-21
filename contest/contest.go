package contest

import (
	"bufio"
	"encoding/csv"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

type Contest struct {
	Entries []Entry
	Voters  []string
}

func (c Contest) FindArtist(artist string) (entry Entry, found bool) {
	for _, e := range c.Entries {
		if strings.Contains(strings.ToLower(artist), strings.ToLower(e.Artist)) {
			return e, true
		}
	}
	return Entry{}, false
}

func (c Contest) FindSong(song string) (entry Entry, found bool) {
	for _, e := range c.Entries {
		if strings.Contains(strings.ToLower(song), strings.ToLower(e.Song)) {
			return e, true
		}
	}
	return Entry{}, false
}

func (c Contest) FindCountryByName(name string) (entry Entry, found bool) {
	for _, e := range c.Entries {
		if e.Country.Find(name) {
			return e, true
		}
	}
	return Entry{}, false
}

func (c Contest) FindCountryByForum(forum string) (entry Entry, found bool) {
	for _, e := range c.Entries {
		if strings.Contains(strings.ToLower(forum), strings.ToLower(e.Country.Forum)) {
			return e, true
		}
	}
	return Entry{}, false
}

func (c Contest) GetEntryIndex(entry Entry) (index int, found bool) {
	for idx, e := range c.Entries {
		if e.Equals(entry) {
			return idx, true
		}
	}
	return -1, false
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
