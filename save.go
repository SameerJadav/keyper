package main

import (
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

	writeEnvVarsToFile(envVars, envVarFile)
}
