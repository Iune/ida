package contest

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Country struct {
	Forum string   `json:"forum"`
	Names []string `json:"names"`
	Flag  string   `json:"flag"`
}

func (c Country) PrimaryName() (primaryName string, found bool) {
	if len(c.Names) > 0 {
		return c.Names[0], true
	}
	return "", false
}

func (c Country) containedIn(text string) bool {
	for _, name := range c.Names {
		if strings.Contains(strings.ToLower(text), strings.ToLower(name)) {
			return true
		}
	}
	return false
}

type countriesFile struct {
	Countries []Country `json:"countries"`
}

func LoadCountries(filePath string) []Country {
	file, _ := os.ReadFile(filePath)
	countriesFile := countriesFile{}
	_ = json.Unmarshal(file, &countriesFile)
	return countriesFile.Countries
}

type Entry struct {
	Country Country `json:"country"`
	Artist  string  `json:"artist"`
	Song    string  `json:"song"`
}

func (e Entry) Flag() string {
	if len(e.Country.Flag) > 0 {
		return fmt.Sprintf("World/%s.png", e.Country.Flag)
	}
	return ""
}

func LoadEntries(filePath string, countries []Country) []Entry {
	file, _ := os.ReadFile(filePath)
	lines := strings.Split(string(file), "\n")
	for i := range lines {
		lines[i] = strings.TrimSpace(lines[i])
	}

	countriesMap := make(map[string]Country)
	for _, country := range countries {
		countriesMap[country.Forum] = country
	}

	entries := []Entry{}
	forumMatch := regexp.MustCompile(`:([A-z]*):`)
	entryMatch := regexp.MustCompile(`\[B\](.*)\[\/B\]`)
	for _, line := range lines {
		forum := strings.ReplaceAll(forumMatch.FindString(line), ":", "")
		country := countriesMap[forum]

		entryLine := strings.Split(strings.ReplaceAll(strings.ReplaceAll(entryMatch.FindString(line), "[B]", ""), "[/B]", ""), " - ")
		if len(entryLine) == 2 {
			artist := entryLine[0]
			song := entryLine[1]
			entries = append(entries, Entry{country, artist, song})
		}
	}

	return entries
}

type Contest struct {
	Entries   []Entry   `json:"entries"`
	Countries []Country `json:"countries"`
	Voters    []Country `json:"voters"`
}

func LoadContest(filePath string) Contest {
	file, _ := os.ReadFile(filePath)
	contest := Contest{}
	_ = json.Unmarshal(file, &contest)
	return contest
}
