package main

import (
	"fmt"
	"os"
)

func retrieve() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: keyper get <project>")
		return
	}

	if len(os.Args) > 3 {
		fmt.Println("Error: Only one project's environment variables can be retrieved.\nUsage: keyper get <project>")
		return
	}

	envVars := loadEnvVars()

	project, exist := envVars[os.Args[2]]
	if !exist {
		fmt.Println("Error: The project does not exist.")
		return
	}

	for key, value := range project {
		fmt.Printf("%s=%s\n", key, value)
	}
}
