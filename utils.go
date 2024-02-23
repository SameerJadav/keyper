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
  set     Saves project's environment variables locally
          keyper set <project> <key=value> ...
  get     Retrieves project's environment variables
          keyper get <project>
  remove  Remove specific environment variables from a project
          keyper remove <project> <key=value> ...
  purge   Remove the entire project and its environment variables
          keyper purge <project> ...


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

func writeEnvVarsToFile(envVars EnvVars, envVarFile string) {
	jsonData, err := json.Marshal(envVars)
	if err != nil {
		fmt.Println("Error: An error occurred while converting data to JSON format.")
		return
	}

	if err := os.WriteFile(envVarFile, jsonData, 0o644); err != nil {
		fmt.Println("Error: Failed to save the environment variables.")
		return
	}
}

func showUsage() {
	fmt.Println(description)
}
