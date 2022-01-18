package config

import (
	"fmt"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

const (
	CONFIG_FILE_PATH = "config.yml"
)

type Config struct {
}

func Get() Config {
	config, err := parseYamlFile(CONFIG_FILE_PATH)
	if err != nil {
		log.Fatal("Failed to parse config file '%s', %s", CONFIG_FILE_PATH, err.Error())
		return Config{}
	}
	return config
}

func parseYamlFile(file string) (Config, error) {
	buffer, err := ioutil.ReadFile(file)
	if err != nil {
		return Config{}, err
	}

	config := Config{}
	if yaml.Unmarshal(buffer, config) != nil {
		return Config{}, fmt.Errorf("could not unmarshal, %s", err.Error())
	}
	return config, nil
}
