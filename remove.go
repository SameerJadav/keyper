package main

import (
	"fmt"
	"os"
)

func purge() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: keyper purge <project> ...")
		return
	}

	projects := os.Args[2:]

	envVarFile := getEnvVarFile()
	envVars := loadEnvVars()

	for _, project := range projects {
		if _, exist := envVars[project]; !exist {
			fmt.Printf("Error: The project %q does not exist.\n", project)
			return
		}

		delete(envVars, project)
	}

	writeEnvVarsToFile(envVars, envVarFile)

	fmt.Println("Info: Successfully purged project.")
}

func remove() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: keyper remove <project> <key> ...")
		return
	}

	project := os.Args[2]
	keys := os.Args[3:]

	envVarFile := getEnvVarFile()
	envVars := loadEnvVars()

	if _, exist := envVars[project]; !exist {
		fmt.Println("Error: The project does not exist.")
		return
	}

	for _, key := range keys {
		if _, exist := envVars[project][key]; !exist {
			fmt.Printf("Error: No environment variable exists with the key %q for the project %q\n", key, project)
			return
		}
		delete(envVars[project], key)
	}

	writeEnvVarsToFile(envVars, envVarFile)

	fmt.Println("Info: Successfully removed environment variable from the project.")
}
