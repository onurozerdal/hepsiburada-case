package config

import (
	"github.com/BurntSushi/toml"
	"log"
)

type Config struct {
	DataDb DatabaseProperties
	RecommendedDb DatabaseProperties
}

type DatabaseProperties struct {
	Server        string
	Database      string
	Port          int
	User          string
	Password      string
}

func (c *Config) Read() {
	if _, err := toml.DecodeFile("./config.toml", &c); err != nil {
		log.Fatal(err)
	}
}