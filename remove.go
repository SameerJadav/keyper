package main

import (
	"fmt"
	"os"
	"strings"
)

func purge() {
	if len(os.Args) < 3 {
		fmt.Println("USAGE: keyper purge <project> ...")
		return
	}

	projects := os.Args[2:]

	envVarFile := getEnvVarFile()
	envVars := loadEnvVars()

	for _, project := range projects {
		if _, exist := envVars[project]; !exist {
			fmt.Printf("Error: The project %q does not exist.\n", project)
			return
		}

		delete(envVars, project)
	}

	writeEnvVarsToFile(envVars, envVarFile)

	fmt.Println("INFO: Successfully purged project.")
}

func remove() {
	if len(os.Args) < 4 {
		fmt.Println("USAGE: keyper remove <project> <key=value> ...")
		return
	}

	project := os.Args[2]
	kvPairs := os.Args[3:]

	envVarFile := getEnvVarFile()
	envVars := loadEnvVars()

	if _, exist := envVars[project]; !exist {
		fmt.Printf("Error: The project %q does not exist.\n", project)
		return
	}

	for _, kvPair := range kvPairs {
		kv := strings.Split(kvPair, "=")

		if len(kv) != 2 {
			fmt.Printf("Error: Invalid key=value pair %q\nPlease provide valid key-value pairs in the format \"key=value\".\n", kvPair)
			return
		}

		key := kv[0]

		if _, exist := envVars[project][key]; !exist {
			fmt.Printf("Error: No environment variable exists with the key %q for the project %q", key, project)
			return
		}

		delete(envVars[project], key)
	}

	writeEnvVarsToFile(envVars, envVarFile)

	fmt.Println("INFO: Successfully removed environment variable from the project.")
}
