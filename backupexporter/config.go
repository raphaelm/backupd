package main

import (
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"log"
	"os"
)

var config struct {
	Jobs  []JobConfig
	KeyID string
}

type JobConfig struct {
	Name     string
	Module   string
	Argument string
}

func parseConfig(r io.Reader) error {
	content, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(content, &config)
	if err != nil {
		return err
	}
	return nil
}

func readConfig() {
	// Read config file
	f, err := os.Open(confFile)
	if err != nil {
		log.Fatal(err)
	}
	err = parseConfig(f)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}
}
