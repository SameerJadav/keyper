package main

import (
	"log"

	"github.com/SameerJadav/keyper/cmd"
)

func main() {
	log.SetFlags(0)
	if err := cmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
