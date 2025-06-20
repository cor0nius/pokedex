package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	pokecache "github.com/cor0nius/pokedexcli/internal"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	comms := getCommands()
	cache := pokecache.NewCache(5 * time.Second)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		if comm, ok := comms()[input[0]]; !ok {
			fmt.Println("Unknown command")
		} else {
			if err := comm.callback(cache); err != nil {
				fmt.Printf("Error: %v", err)
			}
		}
	}
}
