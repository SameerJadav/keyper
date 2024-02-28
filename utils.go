package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
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

func getEnvVarFile() string {
	if runtime.GOOS != "linux" && runtime.GOOS != "darwin" && runtime.GOOS != "windows" {
		fmt.Println("Error: OS not supported.")
		os.Exit(1)
	}

	if runtime.GOOS == "windows" {
		return filepath.Join(os.Getenv("USERPROFILE"), "AppData", "Local", "keyper.json")
	}

	xdgConfigHome := os.Getenv("XDG_CONFIG_HOME")
	if xdgConfigHome == "" {
		xdgConfigHome = filepath.Join(os.Getenv("HOME"), ".config")
	}

	envVarFile := filepath.Join(xdgConfigHome, "keyper.json")
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
		fmt.Println("Error: Failed to convert data into JSON format.")
		return
	}

	if err := os.WriteFile(envVarFile, jsonData, 0o644); err != nil {
		fmt.Println("Error: Failed to save the environment variables.")
		return
	}
}
