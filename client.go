package main

import (
	"log"

	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v2"
)

type Client struct {
	Key         string `yaml:"company"`
	Name        string
	BillingInfo string `yaml:"billing"`
}

func ParseClients(src string) []*Client {
	var clients []*Client
	err := yaml.Unmarshal([]byte(src), &clients)
	if err != nil {
		log.Fatal(errors.Wrap(err, "error parsing clients"))
	}
	return clients
}
