package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

/*
EnvVars is the structure of scoop.json file

	{
	  "project": {
	    "key": "value"
	  }
	}
*/
type EnvVars map[string]map[string]string

const description = `Scoop is a CLI tool for effortlessly managing all your environment variables.
Save environment variables locally and retrieve them with just one command.
Scoop is simple, useful, and blazingly fast.

Usage
  scoop [command]

Available Commands
  set  Saves project's environment variable locally
       scoop set <project> <key=value> ...
  get  Retrieves project's environment variable
       scoop get <project>

Flags
  --help, -h  help for scoop

Examples
  scoop set my-project DB_HOST=localhost DB_PORT=5432
  scoop get my-project

Learn more about the Scoop at https://github.com/SameerJadav/scoop`

func main() {
	if len(os.Args) == 1 {
		showUsage()
		return
	}

	envVarFile := filepath.Join(os.Getenv("HOME"), "scoop.json")
	envVars := loadEnvVars(envVarFile)

	switch os.Args[1] {
	case "set":
		save(envVars, envVarFile)
	case "get":
		retrieve(envVars)
	case "--help", "-h":
		showUsage()
		return
	default:
		fmt.Println("Error: Unknown Command.\nRun \"scoop --help\" for usage.")
	}
}

func save(envVars EnvVars, envVarFile string) {
	if len(os.Args) < 4 {
		fmt.Println("Usage: scoop set <project> <key=value> ...")
		return
	}

	project := os.Args[2]
	kvPairs := os.Args[3:]

	if _, exist := envVars[project]; !exist {
		envVars[project] = make(map[string]string)
	}

	for _, kvPair := range kvPairs {
		kv := strings.Split(kvPair, "=")

		if len(kv) != 2 {
			fmt.Printf("Invalid key=value pair: %s\n", kvPair)
			return
		}

		key, value := kv[0], kv[1]

		envVars[project][key] = value
	}

	jsonData, err := json.Marshal(envVars)
	if err != nil {
		fmt.Println("Error encoding json.", err)
		return
	}

	if err := os.WriteFile(envVarFile, jsonData, 0o644); err != nil {
		fmt.Println("Error saving env variables.", err)
		return
	}
}

func retrieve(envVars EnvVars) {
	if len(os.Args) < 3 {
		fmt.Println("Usage: scoop get <project>")
		return
	}

	if len(os.Args) > 4 {
		fmt.Println("Only one project's environment variables can be retrieved.")
	}

	project, exist := envVars[os.Args[2]]
	if !exist {
		fmt.Println("The project does not exist.")
		return
	}

	for key, value := range project {
		fmt.Printf("%s=%s\n", key, value)
	}
}

func loadEnvVars(envVarfile string) EnvVars {
	file, err := os.ReadFile(envVarfile)

	if os.IsNotExist(err) {
		return make(EnvVars)
	} else if err != nil {
		fmt.Println("Error opening the file.", err)
		os.Exit(1)
	}

	var envVars EnvVars

	if err := json.Unmarshal(file, &envVars); err != nil {
		fmt.Println("Error decoding json file.", err)
		os.Exit(1)
	}

	return envVars
}

func showUsage() {
	fmt.Println(description)
}
