package main

import (
    "gopkg.in/yaml.v3"
    "io/ioutil"
    "log"
)

// Example that was followed: https://stackoverflow.com/questions/28682439/go-parse-yaml-file

// Config is an object holds all the data from the config file.
type Config struct {
    TargetEmails        []string `yaml:"targetEmails"`
    DoorOpenWaitSeconds int      `yaml:"doorOpenWaitSeconds"`
    // <=0 for no heartbeat
    HeartbeatSeconds int    `yaml:"heartbeatSeconds"`
    FromEmail        string `yaml:"fromEmail"`
    FromName         string `yaml:"fromName"`
    EnableDoorOpen   bool `yaml:"enableDoorOpen"`
    EnableDoorClosed   bool `yaml:"enableDoorClosed"`
    EnableDoorLeftOpen   bool `yaml:"enableDoorLeftOpen"`
    Mailjet          struct {
        PublicApiKey  string `yaml:"publicApiKey"`
        PrivateApiKey string `yaml:"privateApiKey"`
    } `yaml:"mailjet"`
}

// NewConfiguration creates a new config object from the given file.
func NewConfiguration(configFilePath string) Config {
    config := getConfigObject(getYamlFileContents(configFilePath))

    log.Printf("Configuration:\n%#v\n\n", config)

    return config
}

// getYamlFileContents slurps in the contents of the given file and returns as a byte array.
func getYamlFileContents(configFilePath string) []byte {
    yamlFileContents, err := ioutil.ReadFile(configFilePath)

    if err != nil {
        log.Fatalf("Error opening config file '%v': %v", configFilePath, err)
    }

    return yamlFileContents;
}

// getConfigObject converts the byte array of a YAML config file into a Config object and returns it.
func getConfigObject(configFileContents []byte) Config {
    config := Config{}

    // https://godoc.org/gopkg.in/yaml.v2#UnmarshalStrict
    var err = yaml.Unmarshal(configFileContents, &config)
    if err != nil {
        log.Fatalf("Error parsing config file YAML: %v", err)
    }

    return config
}
