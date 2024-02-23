package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		showUsage()
		return
	}

	switch os.Args[1] {
	case "set":
		save()
	case "get":
		retrieve()
	case "remove":
		remove()
	case "purge":
		purge()
	case "--help", "-h":
		showUsage()
		return
	default:
		fmt.Println("Error: Unknown Command.\nRun \"keyper --help\" for usage.")
	}
}
