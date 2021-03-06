package Models

import (
	"errors"
	"strings"

	"github.com/microcosm-cc/bluemonday"
)

type Country struct {
	ID      uint32 //Auto increment, Primary key
	Country string //Unique
}

// Sanitize and Standardize Country and Address inputs.
func (country *Country) Sanitize() {
	p := bluemonday.StrictPolicy()
	country.Country = strings.Title(strings.ToLower(p.Sanitize(country.Country)))
}

// Validates inputs and returns error if any.
func (country *Country) Validate() error {
	if country.Country == "" {
		return errors.New("enter the country name")
	}
	return nil
}
