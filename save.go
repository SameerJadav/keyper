package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func save() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: keyper set <project> <key=value> ...")
		return
	}

	project := os.Args[2]
	kvPairs := os.Args[3:]

	envVarFile := getEnvVarFile()
	envVars := loadEnvVars()

	if _, exist := envVars[project]; !exist {
		envVars[project] = make(map[string]string)
	}

	for _, kvPair := range kvPairs {
		kv := strings.Split(kvPair, "=")

		if len(kv) != 2 {
			fmt.Printf("Error: Invalid key=value pair %q\nPlease provide valid key-value pairs in the format \"key=value\".\n", kvPair)
			return
		}

		key, value := kv[0], kv[1]

		envVars[project][key] = value
	}

	jsonData, err := json.Marshal(envVars)
	if err != nil {
		fmt.Println("Error: An error occurred while converting data to JSON format.")
		return
	}

	if err := os.WriteFile(envVarFile, jsonData, 0o644); err != nil {
		fmt.Println("Error: Failed to save the environment variable(s).")
		return
	}
}
