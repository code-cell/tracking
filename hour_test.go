package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestParseHour(t *testing.T) {
	tests := []struct {
		input    string
		expected []*Hour
	}{
		{
			input: `---
              - day: 2018-09-12
                hours:
                  foobar.p1: 1`,
			expected: []*Hour{
				{
					Client:  "foobar",
					Project: "p1",
					Hours:   1,
					Day:     time.Date(2018, 9, 12, 0, 0, 0, 0, time.UTC),
				},
			},
		},
		{
			input: `---
              - day: 2018-09-12
                hours:
                  foo.p1: 1
                  bar.p2: 2`,
			expected: []*Hour{
				{
					Client:  "foo",
					Project: "p1",
					Hours:   1,
					Day:     time.Date(2018, 9, 12, 0, 0, 0, 0, time.UTC),
				},
				{
					Client:  "bar",
					Project: "p2",
					Hours:   2,
					Day:     time.Date(2018, 9, 12, 0, 0, 0, 0, time.UTC),
				},
			},
		},
		{
			input: `---
              - day: 2018-09-12
                hours:
                  foo.p1: 1
              - day: 2018-09-13
                hours:
                  bar.p2: 2`,
			expected: []*Hour{
				{
					Client:  "foo",
					Project: "p1",
					Hours:   1,
					Day:     time.Date(2018, 9, 12, 0, 0, 0, 0, time.UTC),
				},
				{
					Client:  "bar",
					Project: "p2",
					Hours:   2,
					Day:     time.Date(2018, 9, 13, 0, 0, 0, 0, time.UTC),
				},
			},
		},
		{
			input: `---
              - day: 2018-09-12
                hours:
                  foo.p1: 1
                  foo.p2: 2`,
			expected: []*Hour{
				{
					Client:  "foo",
					Project: "p1",
					Hours:   1,
					Day:     time.Date(2018, 9, 12, 0, 0, 0, 0, time.UTC),
				},
				{
					Client:  "foo",
					Project: "p2",
					Hours:   2,
					Day:     time.Date(2018, 9, 12, 0, 0, 0, 0, time.UTC),
				},
			},
		},
	}
	for _, test := range tests {
		require.ElementsMatch(t, test.expected, ParseHours(test.input))
	}
}
