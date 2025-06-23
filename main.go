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
	pokedex := make(map[string]Pokemon)
	cfg := config{
		"https://pokeapi.co/api/v2/location-area/", "", "", "", pokedex,
	}
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		if comm, ok := comms()[input[0]]; !ok {
			fmt.Println("Unknown command")
		} else {
			if comm.name == "explore" {
				cfg.aux = "https://pokeapi.co/api/v2/location-area/" + input[1]
			}
			if comm.name == "catch" {
				cfg.aux = "https://pokeapi.co/api/v2/pokemon/" + input[1]
				cfg.name = input[1]
			}
			if err := comm.callback(&cfg, cache); err != nil {
				fmt.Printf("Error: %v", err)
			}
		}
	}
}
