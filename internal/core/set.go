package core

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/SameerJadav/envparse"
	"github.com/SameerJadav/keyper/internal/utils"
)

func Set() error {
	setCmd := flag.NewFlagSet("set", flag.ExitOnError)

	var showHelp bool
	setCmd.BoolVar(&showHelp, "help", false, "show help")
	setCmd.BoolVar(&showHelp, "h", false, "show help (shorthand)")

	if err := setCmd.Parse(os.Args[2:]); err != nil {
		return errors.New("unable to understand the command line arguments")
	}

	if showHelp {
		fmt.Println(SET_CMD_USAGE)
		return nil
	}

	if setCmd.NArg() == 0 {
		return errors.New("please specify a project name\nusage: keyper set <project> [flags] [key=value...]")
	}

	args := setCmd.Args()

	projectName := strings.TrimSpace(args[0])
	if projectName == "" {
		return errors.New("project name cannot be empty")
	}

	appCmd := flag.NewFlagSet(projectName, flag.ExitOnError)

	var environment string
	defaultEnvironment := "dev"
	appCmd.StringVar(&environment, "environment", defaultEnvironment, "Specify the environment (e.g., dev, staging, prod)")
	appCmd.StringVar(&environment, "e", defaultEnvironment, "Specify the environment (shorthand)")

	var envFile string
	appCmd.StringVar(&envFile, "file", "", "Path to a .env file to load variables from")
	appCmd.StringVar(&envFile, "f", "", "Path to a .env file to load variables from (shorthand)")

	var overwrite bool
	appCmd.BoolVar(&overwrite, "overwrite", false, "overwrite existing variables without warning")

	if err := appCmd.Parse(os.Args[3:]); err != nil {
		return errors.New("unable to understand the command line arguments")
	}

	if envFile == "" && appCmd.NArg() == 0 {
		return errors.New("no environment variables provided\nplease specify variables or provide an .env file\nusage: keyper set <project> [flags] [key=value...]")
	}

	projects, err := utils.LoadEnvs()
	if err != nil {
		return err
	}

	if _, ok := projects[projectName]; !ok {
		projects[projectName] = make(utils.Environment)
	}

	if _, ok := projects[projectName][environment]; !ok {
		projects[projectName][environment] = make(utils.EnvVars)
	}

	var envMap map[string]string

	if envFile == "" {
		envMap, err = envparse.Parse(strings.NewReader(strings.Join(appCmd.Args(), "\n")))
		if err != nil {
			return err
		}
	} else {
		envMap, err = envparse.ParseFile(envFile)
		if err != nil {
			return fmt.Errorf("unable to read %s\nplease check if the file exists and you have permission to read it", envFile)
		}
	}

	for key, value := range envMap {
		if _, exists := projects[projectName][environment][key]; exists && !overwrite {
			return fmt.Errorf("variable %s already exists\nuse --overwrite flag to overwrite", key)
		}
		projects[projectName][environment][key] = value
	}

	if err = utils.WriteEnvVarsToFile(projects); err != nil {
		return err
	}

	fmt.Println("Successfully saved environment variables")

	return nil
}
