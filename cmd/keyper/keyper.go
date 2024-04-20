package keyper

import (
	"flag"
	"fmt"
	"os"

	"github.com/SameerJadav/keyper/internal/operations"
)

const description = `Keyper is a CLI tool for effortlessly managing all your environment variables.
Save environment variables locally and retrieve them with just one command.
Keyper is simple, useful, and blazingly fast.

Usage
  keyper [command]

Available Commands
  set     Saves project's environment variables locally
          keyper set <project> <key=value> ...
  get     Retrieves project's environment variables
          keyper get <project>
  remove  Remove specific environment variables from a project
          keyper remove <project> <key> ...
  purge   Remove the entire project and its environment variables
          keyper purge <project> ...
  list    List all projects and their environment variables
          keyper list

Flags
  --help, -h  help for keyper

Examples
  keyper set my-project DB_HOST=localhost DB_PORT=5432
  keyper get my-project

Learn more about the Keyper at https://github.com/SameerJadav/keyper`

func Init() {
	if len(os.Args) == 1 {
		fmt.Println(description)
		return
	}

	helpFlag := false

	flag.BoolVar(&helpFlag, "help", false, "show help")
	flag.BoolVar(&helpFlag, "h", false, "show help (shorthand)")
	flag.Parse()

	if helpFlag {
		fmt.Println(description)
		return
	}

	switch os.Args[1] {
	case "set":
		if err := operations.SetEnvVars(); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	case "get":
		if err := operations.GetEnvVars(); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	case "remove":
		if err := operations.RemoveEnvVars(); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	case "purge":
		if err := operations.RemoveProject(); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	case "list":
		if err := operations.GetAllEnvVars(); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	default:
		fmt.Println("Error: unknown Command\nInfo : run \"keyper --help\" for usage")
		os.Exit(1)
	}
}
