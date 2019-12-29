package contest

import (
	"strings"
	"github.com/tealeg/xlsx"

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
		if e.Country.HasName(name) {
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
	excelFile, err := xlsx.OpenFile(contestFilePath)
	if err != nil {
		log.Fatal(err)
	}

	sheet := excelFile.Sheets[0]
	voters := getVoters(sheet.Rows[0])
	entries := getEntries(sheet, countries)

	log.WithFields(log.Fields{
		"file": contestFilePath,
	}).Info("Loaded contest details from file")
	return Contest{Entries: entries, Voters: voters}
}

func getVoters(row *xlsx.Row) []string {
	var voters []string
	cells := row.Cells
	if len(cells) >= 6 {
		for _, cell := range row.Cells[6:len(cells)] {
			voters = append(voters, cell.String())
		}
	} else {
		log.Fatal("No voters were defined in contest file")
	}
	return voters
}

func getEntries(sheet *xlsx.Sheet, countries []Country) []Entry {
	var entries []Entry
	for _, row := range sheet.Rows[1:] {
		country, found := GetCountry(countries, row.Cells[1].String())
		if !found {
			log.Fatalf("Could not find country %s in countries file", row.Cells[1])
		}
		
		artist := row.Cells[3].String()
		song := row.Cells[4].String()
		entries = append(entries, Entry{
			Country: country,
			Artist:  artist,
			Song:    song,
		})
	}
	return entries
}
