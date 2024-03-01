package main

import (
	"fmt"
	"os"
)

func retrieve() error {
	if len(os.Args) < 3 {
		fmt.Println("Usage: keyper get <project>")
		return nil
	}

	if len(os.Args) > 3 {
		return fmt.Errorf("only one project's environment variables can be retrieved.\nUsage: keyper get <project>")
	}

	envVars, err := loadEnvVars()
	if err != nil {
		return err
	}

	projectAsArg := os.Args[2]
	if projectAsArg == "" {
		return fmt.Errorf("project cannot be an empty string")
	}

	project, exist := envVars[projectAsArg]
	if !exist {
		return fmt.Errorf("the project does not exist")
	}

	for key, value := range project {
		fmt.Printf("%s=%s\n", key, value)
	}

	return nil
}
