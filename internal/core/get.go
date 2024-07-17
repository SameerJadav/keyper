package core

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/SameerJadav/keyper/internal/utils"
)

func Get() error {
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)

	var showHelp bool
	getCmd.BoolVar(&showHelp, "help", false, "show help")
	getCmd.BoolVar(&showHelp, "h", false, "show help (shorthand)")

	if err := getCmd.Parse(os.Args[2:]); err != nil {
		return errors.New("unable to understand the command line arguments")
	}

	if showHelp {
		fmt.Println(GET_CMD_USAGE)
		return nil
	}

	if getCmd.NArg() == 0 {
		fmt.Println("please specify a project name\nusage: keyper get <project> [flags]")
		return nil
	}

	args := getCmd.Args()

	projectName := strings.TrimSpace(args[0])
	if projectName == "" {
		return errors.New("project name cannot be empty")
	}

	appCmd := flag.NewFlagSet(projectName, flag.ExitOnError)

	var environment string
	appCmd.StringVar(&environment, "environment", "", "Specify the environment (e.g., dev, staging, prod)")
	appCmd.StringVar(&environment, "e", "", "Specify the environment (shorthand)")

	var out string
	appCmd.StringVar(&out, "out", "", "Specify the output file path")
	appCmd.StringVar(&out, "o", "", "Specify the output file path (shorthand)")

	if err := appCmd.Parse(os.Args[3:]); err != nil {
		return errors.New("unable to understand the command line arguments")
	}

	if appCmd.NArg() != 0 {
		return errors.New("only one project's environment variables can be retrieved at a time\nusage: keyper get <project> [flags]")
	}

	projects, err := utils.LoadEnvs()
	if err != nil {
		return err
	}

	if _, ok := projects[projectName]; !ok {
		return fmt.Errorf("project %s does not exist\nplease check the project name and try again", projectName)
	}

	var builder strings.Builder

	if out == "" {
		if environment == "" {
			for env, envMap := range projects[projectName] {
				fmt.Printf("# %s (%s)\n", projectName, env)
				for key, value := range envMap {
					fmt.Printf("%s=%s\n", key, value)
				}
			}
		} else {
			if _, ok := projects[projectName][environment]; !ok {
				return fmt.Errorf("environment %s does not exist for project %s\nplease check the environment name and try again", environment, projectName)
			}

			fmt.Printf("# %s (%s)\n", projectName, environment)
			for key, value := range projects[projectName][environment] {
				fmt.Printf("%s=%s\n", key, value)
			}
		}
	} else {
		if environment == "" {
			for env, envMap := range projects[projectName] {
				builder.WriteString(fmt.Sprintf("# %s (%s)\n", projectName, env))
				for key, value := range envMap {
					builder.WriteString(fmt.Sprintf("%s=%s\n", key, value))
				}
			}

			err = writeToFile(out, builder)
			if err != nil {
				return err
			}

			fmt.Printf("Environment variables written to %s\n", out)
		} else {
			if _, ok := projects[projectName][environment]; !ok {
				return fmt.Errorf("environment %s does not exist for project %s\nplease check the environment name and try again", environment, projectName)
			}

			builder.WriteString(fmt.Sprintf("# %s (%s)\n", projectName, environment))
			for key, value := range projects[projectName][environment] {
				builder.WriteString(fmt.Sprintf("%s=%s\n", key, value))
			}

			err = writeToFile(out, builder)
			if err != nil {
				return err
			}

			fmt.Printf("Environment variables written to %s\n", out)
		}
	}

	return nil
}

func writeToFile(out string, builder strings.Builder) error {
	file, err := os.Create(out)
	if err != nil {
		return fmt.Errorf("unable to create %s", out)
	}
	defer file.Close()

	if _, err = file.WriteString(builder.String()); err != nil {
		return errors.New("unable to write environment variables to the output file")
	}

	if err = file.Sync(); err != nil {
		return errors.New("unable to save changes to the output file")
	}

	return nil
}
