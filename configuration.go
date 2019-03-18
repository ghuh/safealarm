package main

import (
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "log"
)

// https://stackoverflow.com/questions/28682439/go-parse-yaml-file

type Config struct {
    TargetEmails []string `yaml:"targetEmails"`
    DoorOpenWaitSeconds int32 `yaml:"doorOpenWaitSeconds"`
    Mailjet struct {
        PublicApiKey string `yaml:"publicApiKey"`
        PrivateApiKey string `yaml:"privateApiKey"`
    } `yaml:"mailjet"`
}

func NewConfiguration(configFilePath string) Config {
    config := getConfigObject(getYamlFileContents(configFilePath))

    log.Printf("Configuration:\n%#v\n\n", config)

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

    // https://godoc.org/gopkg.in/yaml.v2#UnmarshalStrict
    err := yaml.UnmarshalStrict(configFileContents, &config)

    if err != nil {
        log.Fatalf("Error parsing config file YAML: %v", err)
    }

    return config
}
