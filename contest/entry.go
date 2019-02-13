package contest

import (
	"github.com/iune/ida/country"
)

type Entry struct {
	Country country.Country
	Artist  string
	Song    string
}
