package operations

import (
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

	kvPairs := os.Args[3:]

	envVarFile, err := utils.GetEnvVarsFilePath()
	if err != nil {
		return err
	}

	envVars, err := utils.LoadEnvVars()
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

	if err := utils.WriteEnvVarsToFile(envVars, envVarFile); err != nil {
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
		return fmt.Errorf("only one project's environment variables can be retrieved.\nUsage: keyper get <project>")
	}

	envVars, err := utils.LoadEnvVars()
	if err != nil {
		return err
	}

	projectAsArg := os.Args[2]
	if err := utils.ValidateProjectName(projectAsArg); err != nil {
		return err
	}

	project, exist := envVars[projectAsArg]
	if !exist {
		return fmt.Errorf("project does not exist")
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
		if err := utils.ValidateProjectName(project); err != nil {
			return err
		}

		if _, exist := envVars[project]; !exist {
			return fmt.Errorf("project %q does not exist", project)
		}

		delete(envVars, project)
	}

	if err := utils.WriteEnvVarsToFile(envVars, envVarFile); err != nil {
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

	if _, exist := envVars[project]; !exist {
		return fmt.Errorf("project does not exist")
	}

	for _, key := range keys {
		if _, exist := envVars[project][key]; !exist {
			return fmt.Errorf("no environment variable exists with the key %q for the project %q", key, project)
		}

		delete(envVars[project], key)
	}

	if err := utils.WriteEnvVarsToFile(envVars, envVarFile); err != nil {
		return err
	}

	fmt.Println("Info: successfully removed environment variables")

	return nil
}