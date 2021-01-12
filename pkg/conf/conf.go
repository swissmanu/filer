package conf

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

// Specification contains all configuration settings for the filer application.
type Specification struct {
	Addr      string `envconfig:"ADDR" default:":8000"`
	UIPath    string `envconfig:"UI_PATH" default:"./ui"`
	InboxPath string `envconfig:"INBOX_PATH" default:"./inbox"`
	DataPath  string `envconfig:"DATA_PATH" default:"./data"`
	RulesPath string `envconfig:"RULES_PATH" default:"./rules.yml"`
}

// NewDefaultConfig reads a Specification from the env.
func NewDefaultConfig() (*Specification, error) {
	config := new(Specification)

	err := envconfig.Process("filer", config)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to parse environment config")
	}

	return config, nil
}
