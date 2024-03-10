package keyper

import (
	"fmt"
	"log"
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

	log.SetFlags(0)

	switch os.Args[1] {
	case "set":
		if err := operations.SetEnvVars(); err != nil {
			log.Fatalln("Error:", err)
		}
	case "get":
		if err := operations.GetEnvVars(); err != nil {
			log.Fatalln("Error:", err)
		}
	case "remove":
		if err := operations.RemoveEnvVars(); err != nil {
			log.Fatalln("Error:", err)
		}
	case "purge":
		if err := operations.RemoveProject(); err != nil {
			log.Fatalln("Error:", err)
		}
	case "--help", "-h":
		fmt.Println(description)
		return
	default:
		log.Fatalln("Error: unknown Command\nInfo : run \"keyper --help\" for usage")
	}
}
