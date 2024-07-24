package cmd

import (
	"errors"
	"flag"
	"fmt"

	"github.com/SameerJadav/keyper/internal/core"
)

func Execute() error {
	var showHelp bool
	flag.BoolVar(&showHelp, "help", false, "show help")
	flag.BoolVar(&showHelp, "h", false, "show help (shorthand)")

	flag.Parse()

	usage := `Keyper is a CLI tool for effortlessly managing all your environment variables.
Save environment variables locally and retrieve them with just one command.
Keyper is simple, useful, and blazingly fast.

Usage:
  keyper [flags] [command] 

Commands:
  set       Set environment variables for a project
  get       Retrieve environment variables for a project
  remove    Remove specific environment variables from a project
  purge     Remove all environment variables for a project or environment

Flags:
  -h, --help    Show help information for Keyper or its subcommands

Examples:
  keyper --help
  keyper set myapp -e prod -f prod.env
  keyper get myapp -e prod -o .env
  keyper remove myapp --force -e prod API_KEY SECRET_TOKEN
  keyper purge myapp --force

Use "keyper [command] --help" for more information about a command.
Learn more about the Keyper at https://github.com/SameerJadav/keyper`

	args := flag.Args()

	if len(args) == 0 || showHelp {
		fmt.Println(usage)
		return nil
	}

	switch args[0] {
	case "set":
		if err := core.Set(); err != nil {
			return err
		}
	case "get":
		if err := core.Get(); err != nil {
			return err
		}
	case "purge":
		if err := core.Purge(); err != nil {
			return err
		}
	case "remove":
		if err := core.Remove(); err != nil {
			return err
		}
	default:
		return errors.New("unknown command. run \"keyper --help\" for usage")
	}

	return nil
}
