package operations

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/SameerJadav/keyper/internal/utils"
)

func SetEnvVars() error {
	if len(os.Args) < 4 {
		fmt.Println("Usage: keyper set <project> <key=value> ...")
		return nil
	}

	project := os.Args[2]
	if err := utils.ValidateProjectName(project); err != nil {
		return err
	}

	envVars, err := utils.LoadEnvVars()
	if err != nil {
		return err
	}

	if _, ok := envVars[project]; !ok {
		envVars[project] = make(map[string]string)
	}

	keyValuePairs := os.Args[3:]

	keySet := make(map[string]bool)

	for _, keyValuePair := range keyValuePairs {
		keyValue := strings.Split(keyValuePair, "=")

		if len(keyValue) != 2 {
			return fmt.Errorf("invalid key=value pair %q\nInfo : please provide valid key-value pairs in the format \"key=value\"", keyValuePair)
		}

		key, value := keyValue[0], keyValue[1]

		if key == "" || value == "" {
			return errors.New("key and value cannot be an empty string")
		}

		if _, ok := keySet[key]; ok {
			return fmt.Errorf("key %q is repeated", key)
		}

		envVars[project][key] = value

		keySet[key] = true
	}

	envVarFile, err := utils.GetEnvVarsFilePath()
	if err != nil {
		return err
	}

	if err = utils.WriteEnvVarsToFile(envVars, envVarFile); err != nil {
		return err
	}

	fmt.Println("Info: successfully saved environment variables")

	return nil
}

func GetEnvVars() error {
	if len(os.Args) < 3 {
		fmt.Println("Usage: keyper get <project>")
		return nil
	}

	if len(os.Args) > 3 {
		return errors.New("only one project's environment variables can be retrieved.\nUsage: keyper get <project>")
	}

	envVars, err := utils.LoadEnvVars()
	if err != nil {
		return err
	}

	projectAsArg := os.Args[2]
	if err = utils.ValidateProjectName(projectAsArg); err != nil {
		return err
	}

	project, ok := envVars[projectAsArg]
	if !ok {
		return errors.New("project does not exist")
	}

	for key, value := range project {
		fmt.Printf("%s=%s\n", key, value)
	}

	return nil
}

func RemoveProject() error {
	if len(os.Args) < 3 {
		fmt.Println("Usage: keyper purge <project> ...")
		return nil
	}

	projects := os.Args[2:]

	envVarFile, err := utils.GetEnvVarsFilePath()
	if err != nil {
		return err
	}

	envVars, err := utils.LoadEnvVars()
	if err != nil {
		return err
	}

	for _, project := range projects {
		if err = utils.ValidateProjectName(project); err != nil {
			return err
		}

		if _, ok := envVars[project]; !ok {
			return fmt.Errorf("project %q does not exist", project)
		}

		delete(envVars, project)
	}

	if err = utils.WriteEnvVarsToFile(envVars, envVarFile); err != nil {
		return err
	}

	fmt.Println("Info: successfully purged project")

	return nil
}

func RemoveEnvVars() error {
	if len(os.Args) < 4 {
		fmt.Println("Usage: keyper remove <project> <key> ...")
		return nil
	}

	project := os.Args[2]
	if err := utils.ValidateProjectName(project); err != nil {
		return err
	}

	keys := os.Args[3:]

	envVarFile, err := utils.GetEnvVarsFilePath()
	if err != nil {
		return err
	}

	envVars, err := utils.LoadEnvVars()
	if err != nil {
		return err
	}

	if _, ok := envVars[project]; !ok {
		return errors.New("project does not exist")
	}

	for _, key := range keys {
		if _, ok := envVars[project][key]; !ok {
			return fmt.Errorf("no environment variable exists with the key %q for the project %q", key, project)
		}

		delete(envVars[project], key)
	}

	if err = utils.WriteEnvVarsToFile(envVars, envVarFile); err != nil {
		return err
	}

	fmt.Println("Info: successfully removed environment variables")

	return nil
}

func GetAllEnvVars() error {
	if len(os.Args) > 2 {
		fmt.Println("Usage: keyper list")
		return nil
	}

	envVars, err := utils.LoadEnvVars()
	if err != nil {
		return err
	}

	if len(envVars) == 0 {
		fmt.Println("Info: no projects found")
		return nil
	}

	for project, kvPairs := range envVars {
		fmt.Printf("Project: %s\n", project)

		if len(kvPairs) == 0 {
			fmt.Println("no environment variables found for this project")
		} else {
			for key, value := range kvPairs {
				fmt.Printf("%s=%s\n", key, value)
			}
		}

		fmt.Println()
	}

	return nil
}
