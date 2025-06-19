package main

import (
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	lowerCase := strings.ToLower(text)
	words := strings.Fields(lowerCase)
	return words
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	comms := prepareCommands()
	for comm := range comms() {
		fmt.Printf("%s: %s\n", comms()[comm].name, comms()[comm].description)
	}
	return nil
}
