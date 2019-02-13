package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Countries struct {
	Countries []Country `json:"countries"`
}

type Country struct {
	Iso   string   `json:"iso"`
	Forum string   `json:"bbcode"`
	Names []string `json:"names"`
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

func LoadCountries(countryFilePath string) Countries {
	jsonData, err := os.Open(countryFilePath)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonData.Close()

	byteValue, _ := ioutil.ReadAll(jsonData)
	var countries Countries
	json.Unmarshal(byteValue, &countries)

	return countries
}
