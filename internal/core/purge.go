package core

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/SameerJadav/keyper/internal/utils"
)

func Purge() error {
	purgeCmd := flag.NewFlagSet("purge", flag.ExitOnError)

	var showHelp bool
	purgeCmd.BoolVar(&showHelp, "help", false, "show help")
	purgeCmd.BoolVar(&showHelp, "h", false, "show help (shorthand)")

	if err := purgeCmd.Parse(os.Args[2:]); err != nil {
		return errors.New("unable to understand the command line arguments")
	}

	if showHelp {
		fmt.Println(PURGE_CMD_USAGE)
		return nil
	}

	if purgeCmd.NArg() == 0 {
		return errors.New("please specify a project name\nusage: keyper purge <project> [flags]")
	}

	args := purgeCmd.Args()

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

	if appCmd.NArg() != 0 {
		return errors.New("only one project can be purged at a time\nusage: keyper purge <project>")
	}

	if !force {
		if environment == "" {
			fmt.Printf("Are you sure you want to purge %s? This action cannot be undone. (y/N): ", projectName)
			var response string
			fmt.Scanln(&response)
			if response != "y" && response != "Y" {
				return errors.New("operation cancelled")
			}
		} else {
			fmt.Printf("Are you sure you want to purge %s (%s)? This action cannot be undone. (y/N): ", projectName, environment)
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

	if _, ok := projects[projectName]; !ok {
		return fmt.Errorf("project %s does not exist\nplease check the project name and try again", projectName)
	}

	if environment == "" {
		delete(projects, projectName)
	} else {
		if _, ok := projects[projectName][environment]; !ok {
			return fmt.Errorf("environment %s does not exist for project %s\nplease check the environment name and try again", environment, projectName)
		}

		delete(projects[projectName], environment)

		if len(projects[projectName]) == 0 {
			delete(projects, projectName)
		}
	}

	if err = utils.WriteEnvVarsToFile(projects); err != nil {
		return err
	}

	if environment == "" {
		fmt.Printf("Successfully purged %s\n", projectName)
	} else {
		fmt.Printf("Successfully purged %s (%s)\n", projectName, environment)
	}

	return nil
}
