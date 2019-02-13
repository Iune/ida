package country

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strings"
)

type Country struct {
	Iso   string
	Forum string
	Names []string
}

func (c Country) Find(name string) bool {
	for _, n := range c.Names {
		if n == name {
			return true
		}
	}
	return false
}

func (c Country) GetPrimaryName() string {
	if len(c.Names) > 0 {
		return c.Names[0]
	}
	return ""
}

func LoadCountries(countryFilePath string) []Country {
	csvFile, err := os.Open(countryFilePath)
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))
	reader.Comma = '\t'
	var countries []Country
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		countries = append(countries, Country{
			Iso:   line[0],
			Forum: line[1],
			Names: strings.Split(line[2], ";"),
		})
	}
	return countries
}
