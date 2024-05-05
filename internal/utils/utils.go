package utils

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"runtime"
)

type EnvVars map[string]map[string]string

func GetEnvVarsFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	switch runtime.GOOS {
	case "linux", "darwin":
		return filepath.Join(homeDir, ".config", "keyper.json"), nil
	case "windows":
		return filepath.Join(homeDir, "AppData", "Local", "keyper.json"), nil
	default:
		return "", errors.New("operating system not supported")

	}
}

func LoadEnvVars() (EnvVars, error) {
	envVarFile, err := GetEnvVarsFilePath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(envVarFile)
	if err != nil {
		if os.IsNotExist(err) {
			return make(EnvVars), nil
		}
		return nil, errors.New("failed to read environment variables file")
	}

	var envVars EnvVars

	if err = json.Unmarshal(data, &envVars); err != nil {
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
