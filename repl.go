package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strings"

	pc "github.com/cor0nius/pokedexcli/internal"
)

func cleanInput(text string) []string {
	lowerCase := strings.ToLower(text)
	words := strings.Fields(lowerCase)
	return words
}

func commandExit(cfg *config, cache *pc.Cache) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, cache *pc.Cache) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	comms := getCommands()
	for comm := range comms() {
		fmt.Printf("%s: %s\n", comms()[comm].name, comms()[comm].description)
	}
	return nil
}

func commandMap(cfg *config, cache *pc.Cache) error {
	var locationAreas locationAreaAPI
	var data []byte
	url := cfg.next
	data, ok := cache.Get(url)
	if !ok {
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		cache.Add(url, data)
	}
	err := json.Unmarshal(data, &locationAreas)
	if err != nil {
		return err
	}
	cfg.next = locationAreas.Next
	cfg.previous = locationAreas.Previous
	for i := range locationAreas.Results {
		fmt.Println(locationAreas.Results[i].Name)
	}
	return nil
}

func commandMapb(cfg *config, cache *pc.Cache) error {
	if url := cfg.previous; url != "" {
		var locationAreas locationAreaAPI
		var data []byte
		data, ok := cache.Get(url)
		if !ok {
			resp, err := http.Get(url)
			if err != nil {
				return err
			}
			defer resp.Body.Close()
			data, err = io.ReadAll(resp.Body)
			if err != nil {
				return err
			}
			cache.Add(url, data)
		}
		err := json.Unmarshal(data, &locationAreas)
		if err != nil {
			return err
		}
		cfg.next = locationAreas.Next
		cfg.previous = locationAreas.Previous
		for i := range locationAreas.Results {
			fmt.Println(locationAreas.Results[i].Name)
		}
		return nil
	} else {
		fmt.Println("you're on the first page")
		return nil
	}
}

func commandExplore(cfg *config, cache *pc.Cache) error {
	url := cfg.aux
	var encounters areaEncounters
	var data []byte
	data, ok := cache.Get(url)
	if !ok {
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		cache.Add(url, data)
	}
	err := json.Unmarshal(data, &encounters)
	if err != nil {
		return err
	}
	for encounter := range encounters.PokemonEncounters {
		fmt.Println(encounters.PokemonEncounters[encounter].Pokemon.Name)
	}
	return nil
}

func commandCatch(cfg *config, cache *pc.Cache) error {
	fmt.Printf("Throwing a Pokeball at %s...\n", cfg.name)
	x := rand.Float64()
	url := "https://pokeapi.co/api/v2/pokemon/" + cfg.name
	var pokemon Pokemon
	var data []byte
	data, ok := cache.Get(url)
	if !ok {
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		cache.Add(url, data)
	}
	err := json.Unmarshal(data, &pokemon)
	if err != nil {
		return err
	}
	catchRate := 100 - 0.15*float64(pokemon.BaseExp)
	if x > catchRate/100 {
		fmt.Printf("%s escaped!\n", cfg.name)
	} else {
		cfg.pokedex[pokemon.Name] = pokemon
		fmt.Printf("%s was caught!\n", cfg.name)
	}
	return nil
}

func commandInspect(cfg *config, cache *pc.Cache) error {
	pokemon, ok := cfg.pokedex[cfg.name]
	if !ok {
		fmt.Println("you have not caught that Pokemon")
		return nil
	}
	fmt.Printf("Name: %v\nHeight: %vcm\nWeight: %vkg\nType(s):\n", pokemon.Name, pokemon.Height*10, pokemon.Weight/10)
	for _, typ := range pokemon.Types {
		fmt.Printf("	- %v\n", typ.Type.Name)
	}
	return nil
}
