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
                  foobar: 1`,
			expected: []*Hour{
				{
					Client: "foobar",
					Hours:  1,
					Day:    time.Date(2018, 9, 12, 0, 0, 0, 0, time.UTC),
				},
			},
		},
		{
			input: `---
              - day: 2018-09-12
                hours:
                  foo: 1
                  bar: 2`,
			expected: []*Hour{
				{
					Client: "foo",
					Hours:  1,
					Day:    time.Date(2018, 9, 12, 0, 0, 0, 0, time.UTC),
				},
				{
					Client: "bar",
					Hours:  2,
					Day:    time.Date(2018, 9, 12, 0, 0, 0, 0, time.UTC),
				},
			},
		},
		{
			input: `---
              - day: 2018-09-12
                hours:
                  foo: 1
              - day: 2018-09-13
                hours:
                  bar: 2`,
			expected: []*Hour{
				{
					Client: "foo",
					Hours:  1,
					Day:    time.Date(2018, 9, 12, 0, 0, 0, 0, time.UTC),
				},
				{
					Client: "bar",
					Hours:  2,
					Day:    time.Date(2018, 9, 13, 0, 0, 0, 0, time.UTC),
				},
			},
		},
	}
	for _, test := range tests {
		require.Equal(t, test.expected, ParseHours(test.input))
	}
}
