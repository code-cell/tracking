package main

import (
	"log"
	"time"

	yaml "gopkg.in/yaml.v2"
)

type Hour struct {
	Client string
	Day    time.Time
	Hours  float32
}

type sourceHour struct {
	Day   time.Time
	Hours map[string]float32
}

func ParseHours(src string) []*Hour {
	var sourceHours []*sourceHour
	err := yaml.Unmarshal([]byte(src), &sourceHours)
	if err != nil {
		log.Fatal(err)
	}

	hours := make([]*Hour, 0)
	for _, day := range sourceHours {
		for client, hour := range day.Hours {
			hours = append(hours, &Hour{
				Client: client,
				Day:    day.Day,
				Hours:  hour,
			})
		}
	}
	return hours
}
