package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

func get_config() Config {
	conf, err := read_yml_file("../config.yml")
	if err != nil {
		log.Fatal(err)
	}
	return *conf
}

func read_yml_file(filename string) (*Config, error) {
	buffer, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	conf := &Config{}
	err = yaml.Unmarshal(buffer, conf)
	if err != nil {
		return nil, fmt.Errorf("in file %q: %v", filename, err)
	}

	return conf, nil
}
