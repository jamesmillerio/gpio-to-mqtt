package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

//Configuration represents the structure for our .config file to be read.
type Configuration struct {
	PollingIntervalMs int
	Pins              []Pin
	MQTT              MQTT
	path              string
}

//NewSecurityConfiguration returns the settings from
//the settings file ./.jarvis in the root of the project.
func NewSecurityConfiguration() *Configuration {
	return NewSecurityConfigurationFromFile("./.config")
}

//NewSecurityConfigurationFromFile loads the specified settings
//file into a Settings struct and returns it.
func NewSecurityConfigurationFromFile(path string) *Configuration {

	config := new(Configuration)

	config.LoadSettingsFromFile(path)

	config.path = path

	return config

}

//LoadJSONFile loads the specified JSON file into the provided instance.
func (s *Configuration) LoadJSONFile(path string, v interface{}) (interface{}, error) {

	path, _ = filepath.Abs(path)
	file, error := ioutil.ReadFile(path)

	if error != nil {
		return nil, error
	}

	json.Unmarshal(file, v)

	return v, nil
}

//LoadSettingsFromFile loads settings from the specific .json file.
func (s *Configuration) LoadSettingsFromFile(path string) {
	s.LoadJSONFile(path, s)
}

//Save saves the current configuration.
func (s *Configuration) Save() {

	config, err := json.Marshal(s)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = ioutil.WriteFile(s.path, config, 0644)

	if err != nil {
		panic(err)
	}
}
