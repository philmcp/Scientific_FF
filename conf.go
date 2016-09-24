package main

import (
	"encoding/json"
	"fmt"

	"os"
)

type Configuration struct {
	RemoteLoc string
	Database  struct {
		Host     string
		Port     int
		DB       string
		User     string
		Password string
	}
	FFS struct {
		Username string
		Password string
	}
	Buffer      BufferAPI
	Twitter     TwitterAPI
	MinNumTeams int
	MaxWage     float64
	Formation   map[string]int
	Threads     float64
	ValueJump   float64
	MinValue    float64
	Season      int
	Week        int
	DKID        int
	DKName      string
}

// Loaded from conf.json
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
