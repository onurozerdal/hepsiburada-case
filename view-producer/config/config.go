package config

import (
	"github.com/BurntSushi/toml"
	"log"
)

type Config struct {
	Kafka KafkaProperties
}

type KafkaProperties struct {
	Bootstrap string
	Group     string
	Offset    string
	Topic     string
}

func (c *Config) Read() {
	if _, err := toml.DecodeFile("./config.toml", &c); err != nil {
		log.Fatal(err)
	}
}
