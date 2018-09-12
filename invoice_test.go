package tracking

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
			input: removeTrailingSpaces(`
				100001 foobar 12.5 2018-09-12 2018-09-13
			`),
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
			input: removeTrailingSpaces(`
				100001 foo 12.5 2018-09-12 2018-09-13
				100002 bar 23.1 2018-09-08 2018-09-11
			`),
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
		require.Equal(t, test.expected, ParseInvoices(test.input))
	}
}
