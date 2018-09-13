package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseClient(t *testing.T) {
	tests := []struct {
		input    string
		expected []*Client
	}{
		{
			input: `---
              - company: foobar
                billing: |-
                  Foo Bar Manager
                  Foo Bar Ltd (12345678)
                  foo@bar.io
                  12 Foo bar st, Foobarland
                  VAT: FB12345678
                  MOD: 98273311`,
			expected: []*Client{
				{
					Key: "foobar",
					BillingInfo: removeLeadingSpaces(
						`Foo Bar Manager
            Foo Bar Ltd (12345678)
            foo@bar.io
            12 Foo bar st, Foobarland
            VAT: FB12345678
            MOD: 98273311`),
				},
			},
		},
		{
			input: `---
              - company: foo
                billing: |-
                  foo info
              - company: bar
                billing: |-
                  bar info`,
			expected: []*Client{
				{
					Key:         "foo",
					BillingInfo: "foo info",
				},
				{
					Key:         "bar",
					BillingInfo: "bar info",
				},
			},
		},
	}
	for _, test := range tests {
		require.Equal(t, test.expected, ParseClients(test.input))
	}
}
