package src

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

var config map[string]string = map[string]string{}

func ReadConfig() error {
	filepath := os.Getenv("SHOUTRRR_CONFIGURATION_FILE")
	if len(filepath) == 0 {
		filepath = "services.yaml"
	}
	yamlFile, err := ioutil.ReadFile(filepath)
	if err != nil {
		return fmt.Errorf("Error ReadFile configuration file %q:  %v", filepath, err)
	}
	err = yaml.Unmarshal(yamlFile, &config)

	if err != nil {
		return fmt.Errorf("Error while unmarshal configuration file %q:  %v", filepath, err)
	}
	return nil
}

func LoadConfig() error {
	err := ReadConfig()
	return err
}

func GetServiceUrl(service string) (string, bool) {
	value, ok := config[service]
	return value, ok
}
