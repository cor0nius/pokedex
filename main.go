package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	comms := prepareCommands()
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		if comm, ok := comms()[input[0]]; !ok {
			fmt.Println("Unknown command")
		} else {
			if err := comm.callback(); err != nil {
				fmt.Errorf("Error: %w", err)
			}
		}
	}
}
