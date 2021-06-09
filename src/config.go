package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

func get_config() string {
	c, err := readConf("../config.yml")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", c)
	fmt.Println(c.Conf.Hits)
	return "abc"
}

func readConf(filename string) (*Config, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	c := &Config{}
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		return nil, fmt.Errorf("in file %q: %v", filename, err)
	}

	return c, nil
}
