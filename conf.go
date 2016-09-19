package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	Database struct {
		Host     string
		Port     int
		DB       string
		User     string
		Password string
	}
	FFPassword string
}

// Load a JSON configuration file which matches the format of the Configuration
// struct. Initiate a new Client using the details porvided.
func LoadConfig(filepath string, conf interface{}) (err error) {
	// Load configuration values from conf.js into the Configuration struct
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Sprintln(os.Stderr, "Could not load %s\n\nError: %s", filepath, err)
		return err
	}

	decoder := json.NewDecoder(file)
	if err = decoder.Decode(conf); err != nil {
		fmt.Sprintln(os.Stderr, "Could not parse %s\n\nError: %s", filepath, err)
		return err
	}

	return nil
}
