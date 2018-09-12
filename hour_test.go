package tracking

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
			input: removeTrailingSpaces(`
				# 2018-09-12
				1 foobar
			`),
			expected: []*Hour{
				{
					Client: "foobar",
					Hours:  1,
					Day:    time.Date(2018, 9, 12, 0, 0, 0, 0, time.UTC),
				},
			},
		},
		{
			input: removeTrailingSpaces(`
				# 2018-09-12
				1 foo
				2 bar
			`),
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
			input: removeTrailingSpaces(`
				# 2018-09-12
				1 foo

				# 2018-09-13
				2 bar
			`),
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
