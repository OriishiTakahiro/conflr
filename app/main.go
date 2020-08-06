package main

import (
	"conflr/command"
	"log"
)

func main() {
	if err := command.Execute(); err != nil {
		log.Panic(err)
	}
}
