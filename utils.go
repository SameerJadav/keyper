package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

/*
EnvVars is the structure of keyper.json file

	{
	  "project": {
	    "key": "value"
	  }
	}
*/
type EnvVars map[string]map[string]string

const description = `Keyper is a CLI tool for effortlessly managing all your environment variables.
Save environment variables locally and retrieve them with just one command.
Keyper is simple, useful, and blazingly fast.

Usage
  keyper [command]

Available Commands
  set  Saves project's environment variable locally
       keyper set <project> <key=value> ...
  get  Retrieves project's environment variable
       keyper get <project>

Flags
  --help, -h  help for keyper

Examples
  keyper set my-project DB_HOST=localhost DB_PORT=5432
  keyper get my-project

Learn more about the Keyper at https://github.com/SameerJadav/keyper`

func getEnvVarFile() string {
	envVarFile := filepath.Join(os.Getenv("HOME"), "keyper.json")
	return envVarFile
}

func loadEnvVars() EnvVars {
	envVarFile := getEnvVarFile()

	file, err := os.ReadFile(envVarFile)

	if os.IsNotExist(err) {
		return make(EnvVars)
	} else if err != nil {
		fmt.Println("Error: Unable to open the JSON file containing the environment variables.")
		os.Exit(1)
	}

	var envVars EnvVars

	if err := json.Unmarshal(file, &envVars); err != nil {
		fmt.Println("Error: Failed to decode the JSON file containing the environment variables.")
		os.Exit(1)
	}

	return envVars
}

func showUsage() {
	fmt.Println(description)
}
