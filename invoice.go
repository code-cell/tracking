package tracking

import (
	"strings"
	"time"

	blackfriday "gopkg.in/russross/blackfriday.v2"
)

type Invoice struct {
	Number string
	Client string
	Rate   float32
	From   time.Time
	To     time.Time
}

func ParseInvoices(markdown string) []*Invoice {
	parser := blackfriday.New()
	ast := parser.Parse([]byte(markdown))

	hours := make([]*Invoice, 0)
	ast.Walk(func(node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
		if entering && node.Type == blackfriday.Paragraph {
			lines := strings.Split(string(node.FirstChild.Literal), "\n")
			for _, line := range lines {
				parts := strings.Split(line, " ")

				hours = append(hours, &Invoice{
					Number: parts[0],
					Client: parts[1],
					Rate:   mustParseFloat32(parts[2]),
					From:   mustParseTime(parts[3]),
					To:     mustParseTime(parts[4]),
				})
			}
		}
		return blackfriday.GoToNext
	})

	return hours
}
