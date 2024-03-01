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

func getEnvVarFile() (string, error) {
	paths := map[string]string{
		"linux":   filepath.Join(os.Getenv("HOME"), ".config", "keyper.json"),
		"darwin":  filepath.Join(os.Getenv("HOME"), ".config", "keyper.json"),
		"windows": filepath.Join(os.Getenv("USERPROFILE"), "AppData", "Local", "keyper.json"),
	}

	path, exist := paths[runtime.GOOS]
	if !exist {
		return "", fmt.Errorf("operating system not supported")
	}

	return path, nil
}

func loadEnvVars() (EnvVars, error) {
	envVarFile, err := getEnvVarFile()
	if err != nil {
		return nil, err
	}

	file, err := os.ReadFile(envVarFile)

	if os.IsNotExist(err) {
		return make(EnvVars), nil
	} else if err != nil {
		return nil, fmt.Errorf("failed to read environment variables file")
	}

	var envVars EnvVars

	if err := json.Unmarshal(file, &envVars); err != nil {
		return nil, fmt.Errorf("failed to decode the JSON file that contains the environment variables")
	}

	return envVars, nil
}

func writeEnvVarsToFile(envVars EnvVars, envVarFile string) error {
	data, err := json.Marshal(envVars)
	if err != nil {
		return fmt.Errorf("failed to encode data as JSON")
	}

	if err := os.WriteFile(envVarFile, data, 0o644); err != nil {
		return fmt.Errorf("failed to save environment variables")
	}

	return nil
}
