package data

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
                name: Foo Bar Ltd
                projects:
                  - key: p1
                    name: Project1 name
                  - key: p2
                    name: Project2 name
                billing: |-
                  Foo Bar Manager
                  Foo Bar Ltd (12345678)`,
			expected: []*Client{
				{
					Key:  "foobar",
					Name: "Foo Bar Ltd",
					Projects: []*Project{
						{Key: "p1", Name: "Project1 name"},
						{Key: "p2", Name: "Project2 name"},
					},
					BillingInfo: "Foo Bar Manager\nFoo Bar Ltd (12345678)",
				},
			},
		},
		{
			input: `---
              - company: foo
                name: Foo
                projects:
                  - key: p1
                    name: Project1 name
                billing: |-
                  foo info
              - company: bar
                projects:
                  - key: p2
                    name: Project2 name
                name: Bar
                billing: |-
                  bar info`,
			expected: []*Client{
				{
					Key:  "foo",
					Name: "Foo",
					Projects: []*Project{
						{Key: "p1", Name: "Project1 name"},
					},
					BillingInfo: "foo info",
				},
				{
					Key:  "bar",
					Name: "Bar",
					Projects: []*Project{
						{Key: "p2", Name: "Project2 name"},
					},
					BillingInfo: "bar info",
				},
			},
		},
	}
	for _, test := range tests {
		require.ElementsMatch(t, test.expected, ParseClients(test.input))
	}
}
