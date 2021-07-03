package Models

import (
	"errors"
	"strings"

	"github.com/microcosm-cc/bluemonday"
)

type Property struct {
	ID         uint32 //Auto increment, unique
	Address    string //Primary Key
	Country    string //Foreign Key
	Availabile bool   // not null
}

// Sanitize and Standardize Country and Address inputs
func (property *Property) Sanitize() {
	p := bluemonday.StrictPolicy()
	property.Country = strings.Title(strings.ToLower(p.Sanitize(property.Country)))
	property.Address = strings.Title(strings.ToLower(p.Sanitize(property.Address)))
}

func (property *Property) Validate() error {
	if property.Address == "" {
		return errors.New("enter the address of the property")
	}
	if property.Country == "" {
		return errors.New("enter the country the property is in")
	}
	return nil
}
