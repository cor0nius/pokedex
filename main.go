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
	cfg := config{
		"https://pokeapi.co/api/v2/location-area/",
		"",
		"",
	}
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		if comm, ok := comms()[input[0]]; !ok {
			fmt.Println("Unknown command")
		} else {
			if comm.name == "explore" {
				cfg.area = "https://pokeapi.co/api/v2/location-area/" + input[1]
			}
			if err := comm.callback(&cfg, cache); err != nil {
				fmt.Printf("Error: %v", err)
			}
		}
	}
}
