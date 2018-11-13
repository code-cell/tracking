package data

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseFrom(t *testing.T) {
	tests := []struct {
		input    string
		expected *From
	}{
		{
			input: `---
              name: Foo Bar Ltd
              billing: |-
                Foo Bar Manager
                Foo Bar Ltd (12345678)`,
			expected: &From{
				Name:        "Foo Bar Ltd",
				BillingInfo: "Foo Bar Manager\nFoo Bar Ltd (12345678)",
			},
		},
	}
	for _, test := range tests {
		require.Equal(t, test.expected, ParseFrom(test.input))
	}
}
