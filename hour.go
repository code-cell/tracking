package main

import (
	"log"
	"strings"
	"time"

	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v2"
)

type Hour struct {
	Client  string
	Project string
	Day     time.Time
	Hours   float32
}

type sourceHour struct {
	Day   time.Time
	Hours map[string]float32
}

func ParseHours(src string) []*Hour {
	var sourceHours []*sourceHour
	err := yaml.Unmarshal([]byte(src), &sourceHours)
	if err != nil {
		log.Fatal(errors.Wrap(err, "error parsing hours"))
	}

	hours := make([]*Hour, 0)
	for _, day := range sourceHours {
		for client, hour := range day.Hours {
			parts := strings.Split(client, ".")
			hours = append(hours, &Hour{
				Client:  parts[0],
				Project: parts[1],
				Day:     day.Day,
				Hours:   hour,
			})
		}
	}
	return hours
}
