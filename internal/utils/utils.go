package utils

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"runtime"
)

type (
	EnvVars     map[string]string
	Environment map[string]EnvVars
	Projects    map[string]Environment
)

func GetConfigFile() (string, error) {
	var configDir string

	switch runtime.GOOS {
	case "darwin", "linux":
		configDir = os.Getenv("XDG_CONFIG_HOME")
		if configDir == "" {
			configDir = os.Getenv("HOME")
			if configDir == "" {
				return "", errors.New("unable to locate your configuration directory\nplease ensure either $XDG_CONFIG_HOME or $HOME is set in your environment")
			}
			configDir = filepath.Join(configDir, ".config", "keyper")
		} else {
			configDir = filepath.Join(configDir, "keyper")
		}
	case "windows":
		configDir = os.Getenv("AppData")
		if configDir == "" {
			return "", errors.New("unable to locate your AppData folder\nplease ensure %AppData% is set in your environment")
		}
		configDir = filepath.Join(configDir, "keyper")
	default:
		return "", errors.New("your operating system is not currently supported by Keyper\nsupported systems are Windows, macOS, and Linux")
	}

	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		if err = os.Mkdir(configDir, 0755); err != nil {
			return "", errors.New("unable to create the configuration directory")
		}
	}

	configFile := filepath.Join(configDir, "keyper.json")

	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		if _, err := os.Create(configFile); err != nil {
			return "", errors.New("unable to create the configuration file")
		}
	}

	return configFile, nil
}

func LoadEnvs() (Projects, error) {
	configFile, err := GetConfigFile()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(configFile)
	if err != nil {
		if os.IsNotExist(err) {
			return make(Projects), nil
		}
		return nil, errors.New("unable to read the configuration file")
	}

	if len(data) == 0 {
		return make(Projects), nil
	}

	var projects Projects

	if err = json.Unmarshal(data, &projects); err != nil {
		return nil, errors.New("the configuration file appears to be corrupted or in an invalid format")
	}

	return projects, nil
}

func WriteEnvVarsToFile(projects Projects) error {
	configFile, err := GetConfigFile()
	if err != nil {
		return err
	}

	data, err := json.Marshal(projects)
	if err != nil {
		return errors.New("an error occurred while preparing your environment variables for saving")
	}

	if err = os.WriteFile(configFile, data, 0644); err != nil {
		return errors.New("unable to save your environment variables to the configuration file")
	}

	return nil
}
