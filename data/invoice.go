package data

import (
	"log"
	"time"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type Invoice struct {
	Number string `yaml:"invoice"`
	Client string
	Rate   float32
	From   time.Time
	To     time.Time
}

func ParseInvoices(src string) []*Invoice {
	var invoices []*Invoice
	err := yaml.Unmarshal([]byte(src), &invoices)
	if err != nil {
		log.Fatal(errors.Wrap(err, "error parsing invoices"))
	}
	return invoices
}

func (i *Invoice) Contains(h *Hour) bool {
	if h.Day.Before(i.From) || h.Day.After(i.To) {
		return false
	}
	return true
}
