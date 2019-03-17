package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

// https://stackoverflow.com/questions/28682439/go-parse-yaml-file

type Config struct {
	targetEmails []string
	MJ struct {
		publicApiKey string
		privateApiKey string
	}
}

func NewConfiguration(configFilePath string) Config {
	config := getConfigObject(getYamlFileContents(configFilePath))

	fmt.Printf("Configuration:\n%v\n\n", config)

	return config
}

func getYamlFileContents(configFilePath string) []byte {
	yamlFileContents, err := ioutil.ReadFile(configFilePath)

	if err != nil {
		log.Fatalf("Error opening config file '%v': %v", configFilePath, err)
		panic(err)
	}

	return yamlFileContents;
}

func getConfigObject(configFileContents []byte) Config {
	config := Config{}

	err := yaml.Unmarshal(configFileContents, &config)

	if err != nil {
		log.Fatalf("Error parsing config file YAML: %v", err)
	}

	return config
}
