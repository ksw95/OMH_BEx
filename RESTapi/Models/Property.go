package Models

import (
	"errors"
	"strings"

	"github.com/microcosm-cc/bluemonday"
)

type Property struct {
	ID          uint32 //Auto increment, Primary key
	Address     string //Unique, not null
	Country     string //Foreign Key
	Description string
	Available   string // not null
}

// Sanitize and Standardize Country and Address inputs.
func (property *Property) Sanitize() {
	p := bluemonday.StrictPolicy()
	property.Description = p.Sanitize(property.Description)
	property.Country = strings.Title(strings.ToLower(p.Sanitize(property.Country)))
	property.Address = strings.Title(strings.ToLower(p.Sanitize(property.Address)))
	property.Available = strings.Title(strings.ToLower(p.Sanitize(property.Available)))
}

// Ensure all inputs are valid as required.
func (property *Property) Validate() error {
	if property.Address == "" {
		return errors.New("enter the address of the property")
	}
	if property.Country == "" {
		return errors.New("enter the country the property is in")
	}
	if property.Available != "Yes" && property.Available != "No" {
		return errors.New("only enter either yes or no for availability")
	}
	return nil
}
