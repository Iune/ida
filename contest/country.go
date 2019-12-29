package contest

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

type countries struct {
	Countries []Country `json:"countries"`
}

type Country struct {
	Forum   string   `json:"forum"`
	Names   []string `json:"names"`
}

func (c Country) HasName(name string) bool {
	for _, countryName := range c.Names {
		if strings.Contains(strings.ToLower(name), strings.ToLower(countryName)) {
			return true
		}
	}
	return false
}

func (c Country) GetPrimaryName() string {
	if len(c.Names) > 0 {
		return c.Names[0]
	}
	log.Fatal("A Country must have at least one Name")
	return ""
}

func GetCountry(countries []Country, name string) (country Country, found bool) {
	for _, c := range countries {
		if c.HasName(name) {
			return c, true
		}
	}
	return Country{}, false
}

func LoadCountries(countriesFilePath string) []Country {
	jsonFile, err := os.Open(countriesFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var countries countries
	json.Unmarshal(byteValue, &countries)

	log.WithFields(log.Fields{
		"file": countriesFilePath,
	}).Info("Loaded country details from file")
	return countries.Countries
}