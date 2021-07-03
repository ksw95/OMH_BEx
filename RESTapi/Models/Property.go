package models

import (
	"strings"

	"github.com/microcosm-cc/bluemonday"
)

type Property struct {
	Address    string
	Country    string
	Availabile bool
}

func (property *Property) Sanitize() {
	p := bluemonday.StrictPolicy
	property.Country = strings.Title(strings.ToLower(p.Sanitize(property.Country)))
	property.Address = strings.Title(strings.ToLower(p.Sanitize(property.Address)))
}
