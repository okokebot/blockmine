package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type config struct {
	Endpoint string `json:"endpoint"`
	Apikey   string `json:"apikey"`
}

func getConfig() config {
	b, err := ioutil.ReadFile("setting.json")
	if err != nil {
		// fatal("Failed to read config file: %s\n", err)
		fmt.Printf("Failed to read config file: %s\n", err)
	}
	var c config
	err = json.Unmarshal(b, &c)
	if err != nil {
		// fatal("Failed to unmarshal file: %s\n", err)
		fmt.Printf("Failed to unmarshal file: %s\n", err)
	}
	return c
}
