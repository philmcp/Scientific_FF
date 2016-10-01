package conf

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// Loaded from conf.json
func LoadConfig(filepath string, conf interface{}) (err error) {
	// Load configuration values from conf.js into the Configuration struct
	log.Println("Loading " + filepath)
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
