package Models

import (
	"strings"

	"github.com/microcosm-cc/bluemonday"
)

type Country struct {
	Country string
}

func (country *Country) Sanitize() {
	p := bluemonday.StrictPolicy()
	country.Country = strings.Title(strings.ToLower(p.Sanitize(country.Country)))
}
