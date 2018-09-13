package main

import (
	"log"

	yaml "gopkg.in/yaml.v2"
)

type Client struct {
	Key         string `yaml:"company"`
	BillingInfo string `yaml:"billing"`
}

func ParseClients(src string) []*Client {
	var clients []*Client
	err := yaml.Unmarshal([]byte(src), &clients)
	if err != nil {
		log.Fatal(err)
	}
	return clients
}
