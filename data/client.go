package data

import (
	"log"

	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v2"
)

type Client struct {
	Key         string `yaml:"company"`
	Name        string
	BillingInfo string `yaml:"billing"`
	Projects    []*Project
}

func ParseClients(src string) []*Client {
	var clients []*Client
	err := yaml.Unmarshal([]byte(src), &clients)
	if err != nil {
		log.Fatal(errors.Wrap(err, "error parsing clients"))
	}
	return clients
}

func (client *Client) FindProject(projectKey string) *Project {
	for _, project := range client.Projects {
		if project.Key == projectKey {
			return project
		}
	}
	return nil
}
