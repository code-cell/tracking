package data

import (
	"log"

	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v2"
)

type From struct {
	Name        string
	BillingInfo string `yaml:"billing"`
}

func ParseFrom(src string) *From {
	var from *From
	err := yaml.Unmarshal([]byte(src), &from)
	if err != nil {
		log.Fatal(errors.Wrap(err, "error parsing from"))
	}
	return from
}
