package main

import (
	"fmt"
	"os"
	"strings"
)

func save() error {
	if len(os.Args) < 4 {
		fmt.Println("Usage: keyper set <project> <key=value> ...")
		return nil
	}

	project := os.Args[2]
	if project == "" {
		return fmt.Errorf("project cannot be an empty string")
	}

	kvPairs := os.Args[3:]

	envVarFile, err := getEnvVarFile()
	if err != nil {
		return err
	}

	envVars, err := loadEnvVars()
	if err != nil {
		return err
	}

	if _, exist := envVars[project]; !exist {
		envVars[project] = make(map[string]string)
	}

	for _, kvPair := range kvPairs {
		kv := strings.Split(kvPair, "=")

		if len(kv) != 2 {
			return fmt.Errorf("invalid key=value pair %q\nPlease provide valid key-value pairs in the format \"key=value\"", kvPair)
		}

		key, value := kv[0], kv[1]

		if key == "" || value == "" {
			return fmt.Errorf("key and value cannot be an empty string")
		}

		envVars[project][key] = value
	}

	if err := writeEnvVarsToFile(envVars, envVarFile); err != nil {
		return err
	}

	return nil
}
