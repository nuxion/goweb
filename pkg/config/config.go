package config

import (
	"errors"
	"os"

	"github.com/BurntSushi/toml"
)

// Config object
type Config struct {
	General General
	Service []service `toml:"service"`
}

// General config options
type General struct {
	Port string
}

// Service specification
type service struct {
	Hosts []string
	Name  string
	Path  string
	Proto string
	Limit int
	Burst int
}

// LoadTom toml file
func LoadTom(p string) (*Config, error) {
	if _, err := os.Stat(p); os.IsNotExist(err) {
		return nil, errors.New("Config file does not exist")
	} else if err != nil {
		return nil, err
	}

	var conf Config
	if _, err := toml.DecodeFile(p, &conf); err != nil {
		return nil, err
	}

	return &conf, nil
}
