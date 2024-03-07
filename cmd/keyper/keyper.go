package keyper

import (
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

	switch os.Args[1] {
	case "set":
		if err := operations.SetEnvVars(); err != nil {
			fmt.Println("Error:", err)
			return
		}
	case "get":
		if err := operations.GetEnvVars(); err != nil {
			fmt.Println("Error:", err)
			return
		}
	case "remove":
		if err := operations.RemoveEnvVars(); err != nil {
			fmt.Println("Error:", err)
			return
		}
	case "purge":
		if err := operations.RemoveProject(); err != nil {
			fmt.Println("Error:", err)
			return
		}
	case "--help", "-h":
		fmt.Println(description)
		return
	default:
		fmt.Println("Error: Unknown Command.\nRun \"keyper --help\" for usage.")
	}
}
