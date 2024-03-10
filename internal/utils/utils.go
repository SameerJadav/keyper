package utils

import (
	"encoding/json"
	"errors"
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

func GetEnvVarsFilePath() (string, error) {
	paths := map[string]string{
		"linux":   filepath.Join(os.Getenv("HOME"), ".config", "keyper.json"),
		"darwin":  filepath.Join(os.Getenv("HOME"), ".config", "keyper.json"),
		"windows": filepath.Join(os.Getenv("USERPROFILE"), "AppData", "Local", "keyper.json"),
	}

	path, exist := paths[runtime.GOOS]
	if !exist {
		// return "", errors.New("operating system not supported")
		return "", errors.New("operating system not supported")
	}

	return path, nil
}

func LoadEnvVars() (EnvVars, error) {
	envVarFile, err := GetEnvVarsFilePath()
	if err != nil {
		return nil, err
	}

	file, err := os.ReadFile(envVarFile)

	if os.IsNotExist(err) {
		return make(EnvVars), nil
	} else if err != nil {
		return nil, errors.New("failed to read environment variables file")
	}

	var envVars EnvVars

	if err = json.Unmarshal(file, &envVars); err != nil {
		return nil, errors.New("failed to decode the JSON file that contains the environment variables")
	}

	return envVars, nil
}

func WriteEnvVarsToFile(envVars EnvVars, envVarFile string) error {
	data, err := json.Marshal(envVars)
	if err != nil {
		return errors.New("failed to encode data as JSON")
	}

	if err = os.WriteFile(envVarFile, data, 0o644); err != nil {
		return errors.New("failed to save environment variables")
	}

	return nil
}

func ValidateProjectName(project string) error {
	if project == "" {
		return errors.New("project cannot be an empty string")
	}
	return nil
}
