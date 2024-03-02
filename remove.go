package main

import (
	"fmt"
	"os"
)

func purge() error {
	if len(os.Args) < 3 {
		fmt.Println("Usage: keyper purge <project> ...")
		return nil
	}

	projects := os.Args[2:]

	envVarFile, err := getEnvVarFile()
	if err != nil {
		return err
	}

	envVars, err := loadEnvVars()
	if err != nil {
		return err
	}

	for _, project := range projects {
		if project == "" {
			return fmt.Errorf("project cannot be an empty string")
		}

		if _, exist := envVars[project]; !exist {
			return fmt.Errorf("project %q does not exist", project)
		}

		delete(envVars, project)
	}

	if err := writeEnvVarsToFile(envVars, envVarFile); err != nil {
		return err
	}

	fmt.Println("Info: successfully purged project.")

	return nil
}

func remove() error {
	if len(os.Args) < 4 {
		fmt.Println("Usage: keyper remove <project> <key> ...")
		return nil
	}

	project := os.Args[2]
	if project == "" {
		return fmt.Errorf("project cannot be an empty string")
	}

	keys := os.Args[3:]

	envVarFile, err := getEnvVarFile()
	if err != nil {
		return err
	}

	envVars, err := loadEnvVars()
	if err != nil {
		return err
	}

	if _, exist := envVars[project]; !exist {
		return fmt.Errorf("project does not exist")
	}

	for _, key := range keys {
		if _, exist := envVars[project][key]; !exist {
			return fmt.Errorf("no environment variable exists with the key %q for the project %q", key, project)
		}

		delete(envVars[project], key)
	}

	if err := writeEnvVarsToFile(envVars, envVarFile); err != nil {
		return err
	}

	fmt.Println("Info: successfully removed environment variables.")

	return nil
}
