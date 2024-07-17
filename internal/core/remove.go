package core

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/SameerJadav/keyper/internal/utils"
)

func Remove() error {
	removeCmd := flag.NewFlagSet("remove", flag.ExitOnError)

	var showHelp bool
	removeCmd.BoolVar(&showHelp, "help", false, "show help")
	removeCmd.BoolVar(&showHelp, "h", false, "show help (shorthand)")

	if err := removeCmd.Parse(os.Args[2:]); err != nil {
		return errors.New("unable to understand the command line arguments")
	}

	if showHelp {
		fmt.Println(REMOVE_CMD_USAGE)
		return nil
	}

	if removeCmd.NArg() == 0 {
		return errors.New("please specify a project name\nusage: keyper remove <project> [flags] [key...]")
	}

	args := removeCmd.Args()

	projectName := strings.TrimSpace(args[0])
	if projectName == "" {
		return errors.New("project name cannot be empty")
	}

	appCmd := flag.NewFlagSet(projectName, flag.ExitOnError)

	var environment string
	appCmd.StringVar(&environment, "environment", "", "Specify the environment (e.g., dev, staging, prod)")
	appCmd.StringVar(&environment, "e", "", "Specify the environment (shorthand)")

	var force bool
	appCmd.BoolVar(&force, "force", false, "force purge without confirmation")
	appCmd.BoolVar(&force, "f", false, "force purge without confirmation (shorthand)")

	if err := appCmd.Parse(os.Args[3:]); err != nil {
		return errors.New("unable to understand the command line arguments")
	}

	keys := appCmd.Args()
	if len(keys) == 0 {
		return errors.New("please specify at least one key to remove")
	}

	if !force {
		if environment == "" {
			fmt.Printf("Are you sure you want to remove the specified keys from project %s? This action cannot be undone. (y/N): ", projectName)
			var response string
			fmt.Scanln(&response)
			if response != "y" && response != "Y" {
				return errors.New("operation cancelled")
			}
		} else {
			fmt.Printf("Are you sure you want to remove the specified keys from project %s (%s)? This action cannot be undone. (y/N): ", projectName, environment)
			var response string
			fmt.Scanln(&response)
			if response != "y" && response != "Y" {
				return errors.New("operation cancelled")
			}
		}
	}

	projects, err := utils.LoadEnvs()
	if err != nil {
		return err
	}

	project, ok := projects[projectName]
	if !ok {
		return fmt.Errorf("project %s does not exist\nplease check the project name and try again", projectName)
	}

	var count int

	if environment == "" {
		for env := range project {
			for _, key := range keys {
				delete(project[env], key)
				count++
			}
		}
	} else {
		if _, ok := project[environment]; !ok {
			return fmt.Errorf("environment %s does not exist for project %s", environment, projectName)
		}

		for _, key := range keys {
			delete(project[environment], key)
			count++
		}
	}

	if err = utils.WriteEnvVarsToFile(projects); err != nil {
		return err
	}

	fmt.Printf("Successfully removed %d environment variable(s) from project %s\n", count, projectName)

	return nil
}
