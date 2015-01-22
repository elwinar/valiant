package main

import (
	"encoding/json"
	"io/ioutil"
)

// Load a json file into the destination struct
func Load(filename string, destination interface{}) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, destination)
	if err != nil {
		return err
	}

	return nil
}
