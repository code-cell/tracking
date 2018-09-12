package main

import (
	"strings"
	"time"

	blackfriday "gopkg.in/russross/blackfriday.v2"
)

type Hour struct {
	Client string
	Day    time.Time
	Hours  float32
}

func ParseHours(markdown string) []*Hour {
	parser := blackfriday.New()
	ast := parser.Parse([]byte(markdown))

	hours := make([]*Hour, 0)
	ast.Walk(func(node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
		if entering && node.Type == blackfriday.Heading {
			day := mustParseTime(string(node.FirstChild.Literal))

			lines := strings.Split(string(node.Next.FirstChild.Literal), "\n")
			for _, line := range lines {
				parts := strings.Split(line, " ")

				h := mustParseFloat32(parts[0])
				client := parts[1]
				hours = append(hours, &Hour{
					Day:    day,
					Hours:  h,
					Client: client,
				})
			}
		}
		return blackfriday.GoToNext
	})

	return hours
}
