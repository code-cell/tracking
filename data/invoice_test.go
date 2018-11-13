package data

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestParseInvoices(t *testing.T) {
	tests := []struct {
		input    string
		expected []*Invoice
	}{
		{
			input: `---
              - invoice: 100001
                client: foobar
                rate: 12.5
                from: 2018-09-12
                to: 2018-09-13`,
			expected: []*Invoice{
				{
					Number: "100001",
					Client: "foobar",
					Rate:   12.5,
					From:   time.Date(2018, 9, 12, 0, 0, 0, 0, time.UTC),
					To:     time.Date(2018, 9, 13, 0, 0, 0, 0, time.UTC),
				},
			},
		},
		{
			input: `---
              - invoice: 100001
                client: foo
                rate: 12.5
                from: 2018-09-12
                to: 2018-09-13
              - invoice: 100002
                client: bar
                rate: 23.1
                from: 2018-09-08
                to: 2018-09-11`,
			expected: []*Invoice{
				{
					Number: "100001",
					Client: "foo",
					Rate:   12.5,
					From:   time.Date(2018, 9, 12, 0, 0, 0, 0, time.UTC),
					To:     time.Date(2018, 9, 13, 0, 0, 0, 0, time.UTC),
				},
				{
					Number: "100002",
					Client: "bar",
					Rate:   23.1,
					From:   time.Date(2018, 9, 8, 0, 0, 0, 0, time.UTC),
					To:     time.Date(2018, 9, 11, 0, 0, 0, 0, time.UTC),
				},
			},
		},
	}
	for _, test := range tests {
		require.ElementsMatch(t, test.expected, ParseInvoices(test.input))
	}
}
