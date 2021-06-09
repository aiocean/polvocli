package main

import (
	"log"

	"pkg.aiocean.dev/polvocli/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Println(err)
	}
}
